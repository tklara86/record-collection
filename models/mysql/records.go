package mysql

import (
	"database/sql"
	"github.com/record-collection/models"
)


type RecordModel struct {
	DB *sql.DB

}

func (r *RecordModel) Insert(record *models.Record) (int, error) {
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