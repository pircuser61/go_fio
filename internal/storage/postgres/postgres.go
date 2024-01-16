package postgres

import (
	"context"

	//"github.com/georgysavva/scany/pgxscan"
	//"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"

	"github.com/pircuser61/go_fio/internal/models"
	"github.com/pircuser61/go_fio/internal/storage"
)

type PostgressStore struct {
	pool *pgxpool.Pool
}

func GetStore() storage.Store {
	var x storage.Store
	return x
}

func (_ PostgressStore) PersonCreate(ctx context.Context, _ models.Person) (int, error) {
	return 0, nil
}

func (_ PostgressStore) PersonGet(ctx context.Context, id int32) (models.Person, error) {
	var person models.Person
	return person, nil
}

func (_ PostgressStore) PersonUpdate(ctx context.Context, _ models.Person) error {
	return nil
}

func (_ PostgressStore) PersonDelete(ctx context.Context, id int32) error {
	return nil
}

func (_ PostgressStore) PersonList(ctx context.Context) ([]*models.Person, error) {
	var list []*models.Person
	return list, nil
}
