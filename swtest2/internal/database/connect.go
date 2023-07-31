package database

import (
	"fmt"
	"swtest2/internal/config"

	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"
)

func Connect() (*sqlx.DB, error) {
	cfg := config.NewStorageConfig()
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database)
	fmt.Println(url)
	base, err := sqlx.Connect("postgres", url)
	if err != nil {
		return nil, err
	}
	if err := goose.Up(base.DB, cfg.MigrationsPath); err != nil {
		panic(err)
	}
	return base, nil
}

type RepoDB struct {
	db *sqlx.DB
}

func NewRepoDB(db *sqlx.DB) RepoDB {
	return RepoDB{db}
}
