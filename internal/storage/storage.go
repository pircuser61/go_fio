package storage

import (
	"context"

	"github.com/pircuser61/go_fio/internal/models"
)

type Store interface {
	PersonCreate(context.Context, models.Person) (uint32, error)
	PersonGet(context.Context, uint32) (models.Person, error)
	PersonUpdate(context.Context, models.Person) error
	PersonDelete(context.Context, uint32) error
	PersonList(context.Context) ([]*models.Person, error)
	Release()
}
