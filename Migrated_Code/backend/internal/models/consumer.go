package models

import "time"

type Consumer struct {
	ConsumerID         string    `json:"consumer_id"`
	ConsumerKey        string    `json:"consumer_key"`
	ConsumerSecret     string    `json:"-"`
	IsActive           bool      `json:"is_active"`
	Name               string    `json:"name"`
	AppType            string    `json:"app_type,omitempty"`
	Description        string    `json:"description,omitempty"`
	DeveloperEmail     string    `json:"developer_email,omitempty"`
	RedirectURL        string    `json:"redirect_url,omitempty"`
	CreatedByUserID    *string   `json:"created_by_user_id,omitempty"`
	ClientCertificate  string    `json:"client_certificate,omitempty"`
	Company            string    `json:"company,omitempty"`
	LogoURL            string    `json:"logo_url,omitempty"`
	PerSecondCallLimit int       `json:"per_second_call_limit"`
	PerMinuteCallLimit int       `json:"per_minute_call_limit"`
	PerHourCallLimit   int       `json:"per_hour_call_limit"`
	PerDayCallLimit    int       `json:"per_day_call_limit"`
	PerWeekCallLimit   int       `json:"per_week_call_limit"`
	PerMonthCallLimit  int       `json:"per_month_call_limit"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
