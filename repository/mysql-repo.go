package repository

import (
	"database/sql"
	"github.com/record-collection/models"
)

type repo struct {
	DB *sql.DB
}


// NewMySQLRepository creates a new repo
func NewMySQLRepository() RecordRepository {
	return &repo{}
}


func (r *repo) OpenDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func (r *repo) Save(record *models.Record) (int, error) {
	stmt := `INSERT INTO records (title, label, year, created_at, updated_at) VALUES (?, ?, ?, UTC_TIMESTAMP(), UTC_TIMESTAMP())`

	result, err := r.DB.Exec(stmt, record.Title, record.Label, record.Year)
	if err != nil {
		return  0, err
	}
	id, err := result.LastInsertId()
	if err != nil {

		return 0, err
	}


	return int(id), nil
}
