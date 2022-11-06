package driver

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type DBModel struct {
	db *sql.DB
}

func NewDB(db *sql.DB) DBModel {
	return DBModel{db}
}

func ConnectDB(dsn string) (*sql.DB, error) {
	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return conn, nil
}
