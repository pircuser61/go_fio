package storage

import (
	"context"

	"github.com/pircuser61/go_fio/internal/models"
)

type Store interface {
	PersonCreate(context.Context, models.Person) (int32, error)
	PersonGet(context.Context, int32) (models.Person, error)
	PersonUpdate(context.Context, models.Person) error
	PersonDelete(context.Context, int32) error
	PersonList(context.Context) ([]*models.Person, error)
	Release()
}
