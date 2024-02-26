package core

import "time"

type Event struct {
	ID     int       `json:"id" db:"id"`
	Date   time.Time `json:"date" db:"date"`
	UserID string    `json:"user_id" db:"user_id"`
	Title  string    `json:"title" db:"title"`
}
