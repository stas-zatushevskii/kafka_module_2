package domain

import "time"

type Message struct {
	UserID      int64
	RecipientID int64
	Message     string
	Timestamp   time.Time
}

func NewMessage(userID int64, recipientID int64, message string) *Message {
	return &Message{
		UserID:      userID,
		RecipientID: recipientID,
		Message:     message,
		Timestamp:   time.Now(),
	}
}
