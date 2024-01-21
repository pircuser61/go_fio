package postgres

import (
	"context"
	"log/slog"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"

	"github.com/pircuser61/go_fio/config"
	"github.com/pircuser61/go_fio/internal/models"
	"github.com/pircuser61/go_fio/internal/storage"
)

type PostgresStore struct {
	pool *pgxpool.Pool
	log  *slog.Logger
}

func GetStore(ctx context.Context, loggerInstance *slog.Logger) (storage.Store, error) {
	loggerInstance.Info("DB connecting...")
	pool, err := pgxpool.Connect(ctx, config.GetConnectionString())
	if err != nil {
		return nil, err
	}
	x := PostgresStore{log: loggerInstance, pool: pool}
	x.log.Info("DB connected")
	return &x, nil
}

func (i PostgresStore) Release() {
	i.pool.Close()
}

func (i PostgresStore) PersonCreate(ctx context.Context, person *models.Person) (uint32, error) {
	const queryAdd = "INSERT INTO person (name, surname, patronymic, age, gender, nationality)" +
		" VALUES ($1, $2, $3, $4, $5, $6 )  RETURNING id;"

	i.log.Debug("postgres: create person", slog.String("name", person.Name))

	var newId uint32
	err := i.pool.QueryRow(ctx, queryAdd, person.Name, person.Surname, person.Patronymic,
		person.Age, person.Gender, person.Nationality).Scan(&newId)

	if err == nil {
		i.log.Debug("postgres:  person created", slog.Uint64("id", uint64(newId)))
	}
	return newId, err
}

func (i PostgresStore) PersonGet(ctx context.Context, id uint32) (*models.Person, error) {
	const queryGet = "SELECT * FROM person WHERE id = $1"
	i.log.Debug("postgres: get person", slog.Uint64("id", uint64(id)))
	var person models.Person

	if err := pgxscan.Get(ctx, i.pool, &person, queryGet, id); err != nil {
		if pgxscan.NotFound(err) {
			i.log.Debug("postgres: person not found")
			return nil, storage.ErrNotExists
		}
		i.log.Debug("postgres:  ERROR!", slog.String("msg", err.Error()))
		return nil, err
	}
	return &person, nil
}

func (i PostgresStore) PersonUpdate(ctx context.Context, person *models.Person) error {
	const queryUpdate = "UPDATE person " +
		" SET name = $2, surname = $3, patronymic = $4," +
		" age = $5, gender = $6, nationality = $7" +
		" WHERE id = $1;"
	i.log.Debug("postgres: update person",
		slog.Uint64("Id", uint64(person.Id)),
		slog.String("name", person.Name))
	commandTag, err := i.pool.Exec(ctx, queryUpdate,
		person.Id, person.Name, person.Surname, person.Patronymic,
		person.Age, person.Gender, person.Nationality)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return storage.ErrNotExists
	}
	return nil
}

func (i PostgresStore) PersonDelete(ctx context.Context, id uint32) error {
	i.log.Debug("postgres: delete person", slog.Uint64("id", uint64(id)))
	const queryDelete = "DELETE FROM PERSON WHERE id = $1"
	commandTag, err := i.pool.Exec(ctx, queryDelete, id)
	if err != nil {
		i.log.Debug("postgres:  ERROR!", slog.String("msg", err.Error()))
		return err
	}
	if commandTag.RowsAffected() != 1 {
		i.log.Debug("postgres:  not found!")
		return storage.ErrNotExists
	}
	return nil
}

func (i PostgresStore) PersonList(ctx context.Context, filter *models.Filter) ([]*models.Person, error) {
	i.log.Debug("postgres: list person")
	//const queryList = "SELECT * FROM person"
	if filter.Limit <= 0 || filter.Limit > 20 {
		filter.Limit = 20
	}
	sql := sq.Select("*").From("person").
		Limit(filter.Limit).PlaceholderFormat(sq.Dollar)

	if filter.Offset > 0 {
		sql = sql.Offset(filter.Offset)
	}
	if filter.Age > 0 {
		sql = sql.Where(sq.Eq{"age": filter.Age})
	}

	if filter.Gender != "" {
		sql = sql.Where(sq.Eq{"gender": filter.Gender})
	}

	switch filter.Order {
	case "id":
		sql = sql.OrderBy("id")
	case "-id":
		sql = sql.OrderBy("id DESC")
	case "name":
		sql = sql.OrderBy("name")
	case "-name":
		sql = sql.OrderBy("name DESC")
	case "age":
		sql = sql.OrderBy("age")
	case "-age":
		sql = sql.OrderBy("age DESC")
	case "gender":
		sql = sql.OrderBy("gender")
	case "-gender":
		sql = sql.OrderBy("gender DESC")
	case "nation":
		sql = sql.OrderBy("nationality")
	case "-nation":
		sql = sql.OrderBy("nationality DESC")
	}
	queryList, args, err := sql.ToSql()
	if err != nil {
		return nil, err
	}
	i.log.Debug("postgres: list", slog.String("query", queryList))
	i.log.Debug("postgres: list", slog.Any("args", args))
	var list []*models.Person
	if err := pgxscan.Select(ctx, i.pool, &list, queryList, args...); err != nil {
		i.log.Debug("postgres:  ERROR!", slog.String("msg", err.Error()))
		return nil, err
	}
	return list, nil
}
