package models

import "time"

type UserLock struct {
	ID           int       `json:"id"`
	UserID       string    `json:"user_id"`
	TypeOfLock   string    `json:"type_of_lock"`
	LastLockDate time.Time `json:"last_lock_date"`
}
