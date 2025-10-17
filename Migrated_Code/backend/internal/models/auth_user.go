package models

import "time"

type AuthUser struct {
	UserID        string    `json:"user_id"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	PasswordHash  string    `json:"-"`
	EmailVerified bool      `json:"email_verified"`
	Validated     bool      `json:"validated"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Provider      string    `json:"provider"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
