package config

import (
	"github.com/record-collection/models/mysql"
	"log"
)

type Application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Records  *mysql.RecordModel
}
