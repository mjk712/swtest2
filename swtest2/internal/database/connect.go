package database

import (
	"swtest2/internal/config"

	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"
)

var (
	db *sqlx.DB
)

func Connect() error {
	cfg := config.NewStorageConfig()
	d, err := sqlx.Connect("postgres", cfg.DbString)
	if err != nil {
		return err
	}
	db = d
	if err := goose.Up(db.DB, cfg.MigrationsPath); err != nil {
		panic(err)
	}
	return nil
}

func GetDB() *sqlx.DB {
	return db
}
