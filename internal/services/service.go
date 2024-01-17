package service

import (
	"context"
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

func PersonCreate(ctx context.Context, person models.Person) (uint32, error) {
	log.Debug("service: person create")
	var err error
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

func PersonGet(ctx context.Context, id uint32) (models.Person, error) {
	log.Debug("service: person get")
	return store.PersonGet(ctx, id)
}
