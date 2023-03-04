package repo

import (
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type DBConfig struct {
	Database string
	User     string
	Password string
	Host     string
	Port     int
}

type Adapter struct {
	db *sqlx.DB
}

func NewAdapter(c *DBConfig) *Adapter {
	source := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%d sslmode=disable", c.Database, c.User, c.Password, c.Host, c.Port)
	db, err := sqlx.Connect("pgx", source)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return &Adapter{db}
}
