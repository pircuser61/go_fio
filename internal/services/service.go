package service

import (
	"context"
	"errors"
	"log/slog"

	"github.com/pircuser61/go_fio/internal/api/fio_api"
	"github.com/pircuser61/go_fio/internal/models"
	"github.com/pircuser61/go_fio/internal/storage"
)

var (
	api                  fio_api.Api
	store                storage.Store
	log                  *slog.Logger
	ErrValidationName    = errors.New("empty name")
	ErrValidationSurname = errors.New("empty surname")
)

func SetApp(apiInstance fio_api.Api, dbInstance storage.Store, loggerInstance *slog.Logger) {
	api, store, log = apiInstance, dbInstance, loggerInstance
}

func PersonList(ctx context.Context, filter *models.Filter) ([]*models.Person, error) {
	log.Debug("service: person list")
	return store.PersonList(ctx, filter)
}

func PersonValidate(person *models.Person) error {
	if person.Name == "" {
		return ErrValidationName
	}
	if person.Surname == "" {
		return ErrValidationSurname
	}
	return nil
}

func PersonCreate(ctx context.Context, person *models.Person) (uint32, error) {

	log.Debug("service: person create")
	err := PersonValidate(person)
	if err != nil {
		return 0, err
	}

	type intResult struct {
		val int
		err error
	}
	type stringResult struct {
		val string
		err error
	}

	/*
		как вариант использовать waitGroup
			var wg sync.WaitGroup
			wg.Add(3)
			wg.Wait()
		и общий канал для ошибок,
		в конце проверять наличи записи в ошибках через
		   select {
			case err, ok <- err_chanel
		    default: ... }

		тогда можно писать сразу в person
		но потребуется синхронизировать доступ
	*/

	chan_age := make(chan intResult)
	chan_gender := make(chan stringResult)
	chan_nation := make(chan stringResult)

	go func() {
		log.Debug("service: go api age")
		var res intResult
		res.val, res.err = api.GetAge(person.Name)
		log.Debug("service: api returns", slog.Int("age", res.val))
		chan_age <- res
	}()

	go func() {
		log.Debug("service: go api gender")
		var res stringResult
		res.val, res.err = api.GetGender(person.Name)
		log.Debug("service: api returns", slog.String("gender", res.val))
		chan_gender <- res
	}()

	go func() {
		log.Debug("service: go api nation")
		var res stringResult
		res.val, res.err = api.GetNationality(person.Name)
		log.Debug("service: api returns", slog.String("nationality", res.val))
		chan_nation <- res
	}()

	age := <-chan_age
	if age.err != nil {
		return 0, age.err
	}
	gender := <-chan_gender
	if gender.err != nil {
		return 0, gender.err
	}
	nation := <-chan_nation
	if nation.err != nil {
		return 0, err
	}
	person.Age = age.val
	person.Gender = gender.val
	person.Nationality = nation.val

	id, err := store.PersonCreate(ctx, person)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func PersonUpdate(ctx context.Context, person *models.Person) error {
	log.Debug("service: person update")
	err := PersonValidate(person)
	if err != nil {
		return err
	}
	if person.Id <= 0 {
		return errors.New("id must be positive number")
	}
	return store.PersonUpdate(ctx, person)
}

func PersonGet(ctx context.Context, id uint32) (*models.Person, error) {
	log.Debug("service: person get")
	if id <= 0 {
		return nil, errors.New("id must be positive number")
	}
	return store.PersonGet(ctx, id)
}

func PersonDelete(ctx context.Context, id uint32) error {
	log.Debug("service: person delete")
	if id <= 0 {
		return errors.New("id must be positive number")
	}
	return store.PersonDelete(ctx, id)
}
