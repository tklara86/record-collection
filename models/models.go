package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

// User is the user model
type User struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Record is the record model
type Record struct {
	RecordID  int
	Title     string
	Label     string
	Year      string
	Cover     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
