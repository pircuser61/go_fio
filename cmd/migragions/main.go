package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	goose "github.com/pressly/goose/v3"

	config "github.com/pircuser61/go_fio/config"
)

func main() {
	db, err := sql.Open("postgres", config.GetConnectionString())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	/*
		slog.Info("goose down")
		err = goose.DownTo(db, "./../../migrations/", 0)
		if err != nil {
			fmt.Println(err)
		}
	*/
	err = goose.Up(db, "./../../migrations/")
	if err != nil {
		fmt.Println(err)
	}
}
