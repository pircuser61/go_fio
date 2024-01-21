package storage

import (
	"context"
	"errors"

	"github.com/pircuser61/go_fio/internal/models"
)

var (
	ErrNotExists = errors.New("obj does not exist")
	ErrExists    = errors.New("obj exist")
	ErrTimeout   = errors.New("Timeout")
)

type Store interface {
	PersonCreate(context.Context, *models.Person) (uint32, error)
	PersonGet(context.Context, uint32) (*models.Person, error)
	PersonUpdate(context.Context, *models.Person) error
	PersonDelete(context.Context, uint32) error
	PersonList(context.Context, *models.Filter) ([]*models.Person, error)
	Release()
}
