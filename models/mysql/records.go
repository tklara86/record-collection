package mysql

import (
	"database/sql"
	"errors"

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


func (r *RecordModel) Get(recordId int) (*models.Record, error) {
	s := &models.Record{}

	err := r.DB.QueryRow("SELECT record_id, title, label, year FROM records WHERE record_id = ?",
		recordId).Scan(&s.RecordID, &s.Title, &s.Label, &s.Year)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (r *RecordModel) GetAll() ([]*models.Record, error) {
	stmt := `SELECT record_id, title, label, year FROM records ORDER BY record_id LIMIT 10`

	rows, err := r.DB.Query(stmt)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var records []*models.Record

	for rows.Next() {
		s := &models.Record{}

		err = rows.Scan(&s.RecordID, &s.Title, &s.Label, &s.Year)
		if err != nil {
			return nil, err
		}

		records = append(records, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return records, nil

}


