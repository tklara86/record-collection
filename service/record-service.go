package service

import "github.com/record-collection/repository"

type RecordService interface {
	// TODO: add method specific for Record Service
}


type serviceRecord struct {}

var (
	repo repository.RecordRepository
)

// NewRecordService - constructor function
 func NewRecordService(r repository.RecordRepository) RecordService {
	 repo = r
	 return &serviceRecord{}
}