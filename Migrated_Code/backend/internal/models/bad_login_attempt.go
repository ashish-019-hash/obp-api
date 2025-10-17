package models

import "time"

type BadLoginAttempt struct {
	ID                                 int       `json:"id"`
	Provider                           string    `json:"provider"`
	Username                           string    `json:"username"`
	LastFailureDate                    time.Time `json:"last_failure_date"`
	BadAttemptsSinceLastSuccessOrReset int       `json:"bad_attempts_since_last_success_or_reset"`
}
