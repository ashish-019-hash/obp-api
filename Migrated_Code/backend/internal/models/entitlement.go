package models

import "time"

type Entitlement struct {
	EntitlementID   string    `json:"entitlement_id"`
	UserID          string    `json:"user_id"`
	RoleName        string    `json:"role_name"`
	BankID          string    `json:"bank_id,omitempty"`
	CreatedByUserID *string   `json:"created_by_user_id,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
}

func NewEntitlement() *Entitlement {
	return &Entitlement{}
}
