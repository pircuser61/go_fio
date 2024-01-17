package postgres

import (
	"context"
	"log/slog"

	//"github.com/georgysavva/scany/pgxscan"
	//"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"

	"github.com/pircuser61/go_fio/internal/models"
	"github.com/pircuser61/go_fio/internal/storage"
)

type PostgresStore struct {
	pool *pgxpool.Pool
	log  *slog.Logger
}

func GetStore(loggerInstance *slog.Logger) storage.Store {
	x := PostgresStore{log: loggerInstance}
	return x
}

func (PostgresStore) Release() {}

func (PostgresStore) PersonCreate(ctx context.Context, _ models.Person) (uint32, error) {
	return 0, nil
}

func (i PostgresStore) PersonGet(ctx context.Context, id uint32) (models.Person, error) {
	i.log.Debug("postgres: get person", slog.Uint64("id", uint64(id)))
	var person models.Person
	return person, nil
}

func (PostgresStore) PersonUpdate(ctx context.Context, _ models.Person) error {
	return nil
}

func (PostgresStore) PersonDelete(ctx context.Context, id uint32) error {
	return nil
}

func (i PostgresStore) PersonList(ctx context.Context) ([]*models.Person, error) {
	i.log.Debug("postgres: get person list")
	var list []*models.Person
	return list, nil
}
