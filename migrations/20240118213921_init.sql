-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS PUBLIC.PERSON (
    ID SERIAL PRIMARY KEY,
    NAME VARCHAR,
    SURNAME VARCHAR,
    PATRONYMIC VARCHAR,
    AGE INTEGER,
    GENDER VARCHAR,
    NATIONALITY VARCHAR
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS PUBLIC.PERSON;

-- +goose StatementEnd