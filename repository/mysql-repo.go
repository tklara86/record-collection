package repository

import "database/sql"

type repo struct {}


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