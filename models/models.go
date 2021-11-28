package models

import (
	"errors"
	"time"
)


var ErrNoRecord = errors.New("models: no matching record found")

// Record is the record model
type Record struct {
	RecordID int
	Title string
	Label string
	Year string
	Cover string
	CreatedAt time.Time
	UpdatedAt time.Time
}