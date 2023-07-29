package database

import (
	"fmt"
	"swtest2/internal/config"

	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"
)

var (
	db *sqlx.DB
)

func Connect() error {
	cfg := config.NewStorageConfig()
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database)
	fmt.Println(url + "gay")
	d, err := sqlx.Connect("postgres", url)
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
