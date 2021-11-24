package models

import "time"

// Record is the record model
type Record struct {
	RecordID int
	Title string
	Label string
	Year string
	CreatedAt time.Time
	UpdatedAt time.Time
}