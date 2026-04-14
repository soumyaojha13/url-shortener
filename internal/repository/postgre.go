package repository

import (
	"database/sql"
)

type PostgresRepo struct {
	DB *sql.DB
}

func (r *PostgresRepo) Save(long, short string) error {
	_, err := r.DB.Exec(
		"INSERT INTO urls (long_url, short_code) VALUES ($1,$2)",
		long, short,
	)
	return err
}

func (r *PostgresRepo) Get(short string) (string, error) {
	var long string
	err := r.DB.QueryRow(
		"SELECT long_url FROM urls WHERE short_code=$1",
		short,
	).Scan(&long)
	return long, err
}
