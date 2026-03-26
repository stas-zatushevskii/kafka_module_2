package domain

import "time"

type Message struct {
	RecipientID int64     `json:"recipient_id"`
	Message     string    `json:"message"`
	Timestamp   time.Time `json:"timestamp"`
}
