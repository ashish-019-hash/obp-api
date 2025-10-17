package models

import "time"

type UserAuthContext struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	ConsumerID *string   `json:"consumer_id,omitempty"`
	Key        string    `json:"key"`
	Value      string    `json:"value"`
	CreatedAt  time.Time `json:"created_at"`
}
