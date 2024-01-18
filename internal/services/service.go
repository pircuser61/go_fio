package service

import (
	"context"
	"errors"
	"log/slog"

	"github.com/pircuser61/go_fio/internal/api/fio_api"
	"github.com/pircuser61/go_fio/internal/models"
	"github.com/pircuser61/go_fio/internal/storage"
)

var api fio_api.Api
var store storage.Store
var log *slog.Logger

func SetApp(apiInstance fio_api.Api, dbInstance storage.Store, loggerInstance *slog.Logger) {
	api, store, log = apiInstance, dbInstance, loggerInstance
}

func PersonList(ctx context.Context) ([]*models.Person, error) {
	log.Debug("service: person list")
	return store.PersonList(ctx)
}

var (
	ErrValidationName    = errors.New("empty name")
	ErrValidationSurname = errors.New("empty surname")
)

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
	if person.Age, err = api.GetAge(person.Name); err != nil {
		return 0, err
	}
	if person.Gender, err = api.GetGender(person.Name); err != nil {
		return 0, err
	}
	if person.Nationality, err = api.GetNationality(person.Name); err != nil {
		return 0, err
	}

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
