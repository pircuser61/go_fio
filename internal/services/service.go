package service

import (
	"context"

	"github.com/pircuser61/go_fio/internal/api/fio_api"
	"github.com/pircuser61/go_fio/internal/models"
	"github.com/pircuser61/go_fio/internal/storage"
)

type App struct {
	store storage.Store
	api   fio_api.Api
}

func (i App) PersonCreate(ctx context.Context, person models.Person) error {
	var err error
	if person.Age, err = i.api.GetAge(person.Name); err != nil {
		return err
	}
	if _, err = i.store.PersonCreate(ctx, person); err != nil {
		return err
	}
	return nil
}
