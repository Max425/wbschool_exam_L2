package core

import "time"

type Event struct {
	Date   time.Time `json:"date"`
	UserID string    `json:"user_id"`
	ID     string    `json:"event_id"`
	Title  string    `json:"title"`
}
