package sql

import (
	"database/sql"
)

type URLRepository struct {
	db *sql.DB
}

func NewURLRepository(db *sql.DB) *URLRepository {
	return &URLRepository{db: db}
}

func (r *URLRepository) CreateURL(longURL string) (int64, error) {
	query := `INSERT INTO urls (long_url) VALUES (?)`

	result, err := r.db.Exec(query, longURL)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Update short_code
func (r *URLRepository) UpdateShortCode(id int64, code string) error {
	query := `UPDATE urls SET short_code = ? WHERE id = ?`

	_, err := r.db.Exec(query, code, id)
	return err
}

func (r *URLRepository) GetByShortCode(code string) (string, error) {
	query := `SELECT long_url FROM urls where short_code = ?`

	var longURL string
	err := r.db.QueryRow(query, code).Scan(&longURL)

	if err != nil {
		return "", err
	}

	return longURL, nil
}
