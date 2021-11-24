package service

import (
	"database/sql"
	"github.com/record-collection/models"
	"github.com/record-collection/repository"
)


type serviceRecord struct {
	DB *sql.DB
}



type RecordService interface {
	Create(record *models.Record) (int, error)
}


var (
	repo repository.RecordRepository
)

// NewRecordService - constructor function
func NewRecordService(r repository.RecordRepository) RecordService {
	repo = r
	return &serviceRecord{}
}



func (r serviceRecord) Create(record *models.Record) (int, error) {
	return repo.Save(record)
}
