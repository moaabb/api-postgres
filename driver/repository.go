package driver

import (
	"context"
	"time"

	"github.com/moaabb/api-postgres/entities"
)

// Get All Movies from Database
func (m *DBModel) GetAll() ([]*entities.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	var movies []*entities.Movie

	query := `
		SELECT 
			id, title, description, year, release_date, runtime, rating, mpaa_rating, created_at, updated_at 
		FROM api.movies
	`

	rows, err := m.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var movie entities.Movie
		err = rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Description,
			&movie.Year,
			&movie.ReleaseDate,
			&movie.Runtime,
			&movie.Rating,
			&movie.MPAARating,
			&movie.CreatedAt,
			&movie.UpdatedAt,
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

// Get a movie by it's ID
func (m *DBModel) GetByID(id int) (*entities.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `
		SELECT
			id, title, description, year, release_date, runtime, rating, mpaa_rating, created_at, updated_at
		FROM 
			api.movies
		WHERE
			id = $1
	`

	var movie entities.Movie

	err := m.db.QueryRowContext(ctx, query, id).Scan(
		&movie.ID,
		&movie.Title,
		&movie.Description,
		&movie.Year,
		&movie.ReleaseDate,
		&movie.Runtime,
		&movie.Rating,
		&movie.MPAARating,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &movie, nil
}

// Inserts a movie record to the database
func (m *DBModel) InsertMovie(movie entities.Movie) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	stmt := `
		INSERT INTO api.movies
			(title, description, year, release_date, runtime, rating, mpaa_rating, created_at, updated_at)
		VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id
	`
	var id int

	err := m.db.QueryRowContext(ctx, stmt,
		movie.Title,
		movie.Description,
		movie.Year,
		movie.ReleaseDate,
		movie.Runtime,
		movie.Rating,
		movie.MPAARating,
		movie.CreatedAt,
		movie.UpdatedAt,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

// Updates a movie record from the database
func (m *DBModel) UpdateByID(movie entities.Movie) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	stmt := `
		UPDATE api.movies SET
			title = $2, description = $3, year = $4, release_date = $5, runtime = $6, rating = $7, mpaa_rating = $8, updated_at = $9
		WHERE
			id = $1
	`

	_, err := m.db.ExecContext(ctx, stmt, movie.ID,
		movie.Title,
		movie.Description,
		movie.Year,
		movie.ReleaseDate,
		movie.Runtime,
		movie.Rating,
		movie.MPAARating,
		movie.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil

}

// Deletes a movie record from the database
func (m *DBModel) DeleteMovie(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	stmt := `
		DELETE FROM api.movies
		WHERE id = $1
	`
	_, err := m.db.ExecContext(ctx, stmt, id)

	if err != nil {
		return err
	}

	return nil
}
