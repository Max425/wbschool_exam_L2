package core

import "time"

type Event struct {
	ID         int       `json:"id" db:"id"`
	Date       time.Time `json:"-" db:"date"`
	StringDate string    `json:"date" db:"-"`
	UserID     string    `json:"user_id" db:"user_id"`
	Title      string    `json:"title" db:"title"`
}
