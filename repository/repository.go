package repository

import "github.com/record-collection/models"

type DatabaseRepo interface {
	Insert(r models.Record) (int, error)
}
