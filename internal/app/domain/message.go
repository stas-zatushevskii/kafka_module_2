package domain

import "time"

type Message struct {
	UserID      int64     `json:"user_id"`
	RecipientID int64     `json:"recipient_id"`
	Message     string    `json:"message"`
	Timestamp   time.Time `json:"timestamp"`
}

func NewMessage(userID int64, recipientID int64, message string) *Message {
	return &Message{
		UserID:      userID,
		RecipientID: recipientID,
		Message:     message,
		Timestamp:   time.Now(),
	}
}
