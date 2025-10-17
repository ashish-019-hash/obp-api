package models

import "time"

type Scope struct {
	ScopeID    string    `json:"scope_id"`
	ConsumerID string    `json:"consumer_id"`
	RoleName   string    `json:"role_name"`
	BankID     string    `json:"bank_id,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
}
