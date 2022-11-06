package driver

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/moaabb/api-postgres/entities"
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

func (m *DBModel) GetAll() ([]*entities.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	var movies []*entities.Movie

	query := `
		SELECT id, title FROM api.movies
	`

	rows, err := m.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var movie entities.Movie
		err = rows.Scan(
			&movie.Id,
			&movie.Title,
		)
		if err != nil {
			return nil, err
		}
		movies = append(movies, &movie)
	}

	if err = rows.Err(); err != nil {
		if err != nil {
			return nil, err
		}
	}

	return movies, nil
}
