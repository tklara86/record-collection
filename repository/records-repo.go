package repository

import (
	"database/sql"
	"github.com/record-collection/models"
)

type RecordRepository interface {
	OpenDB(dsn string) (*sql.DB, error)
	Save(record *models.Record) (int, error)
}


