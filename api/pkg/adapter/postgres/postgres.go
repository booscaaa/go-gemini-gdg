package postgres

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func Initialize() *sqlx.DB {
	db, err := sqlx.Connect("pgx", "postgres://postgres:postgres@postgres:5432/alexa")

	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(99)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(0)

	return db
}
