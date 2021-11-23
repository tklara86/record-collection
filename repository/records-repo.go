package repository

import "database/sql"

type RecordRepository interface {
	OpenDB(dsn string) (*sql.DB, error)
}

