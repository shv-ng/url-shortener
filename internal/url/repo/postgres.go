package repo

import (
	"database/sql"
	"log"
)

type PostgresRepo struct {
	db *sql.DB
}

func createTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS urldb (
		id SERIAL PRIMARY KEY,
		original_url TEXT NOT NULL,
		short_url TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT NOW()
	);`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func NewPostgresRepo(db *sql.DB) *PostgresRepo {
	createTable(db)
	return &PostgresRepo{
		db: db,
	}
}

func (r *PostgresRepo) Save(original_url, short_url string) error {
	query := `INSERT INTO urldb (original_url, short_url) VALUES ($1, $2)`
	_, err := r.db.Exec(query, original_url, short_url)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *PostgresRepo) GetURL(short_url string) (string, error) {
	query := `SELECT original_url FROM urldb 
						       WHERE short_url = $1;`
	var original_url string
	err := r.db.QueryRow(query, short_url).Scan(&original_url)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return original_url, nil
}
