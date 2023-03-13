package repo

import (
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/micro-it-freelance/config"
)

type Adapter struct {
	db *sqlx.DB
}

func NewAdapter() *Adapter {
	source := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%d sslmode=disable",
		config.DB.Name, config.DB.User, config.DB.Password, config.DB.Host, config.DB.Port)

	db, err := sqlx.Connect("pgx", source)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return &Adapter{db}
}
