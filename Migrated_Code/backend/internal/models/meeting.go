package models

import "time"

type ContactDetails struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type Invitee struct {
	ContactDetails ContactDetails `json:"contact_details"`
	Status         string         `json:"status"`
}

type Meeting struct {
	BankId        string         `json:"bank_id"`
	MeetingId     string         `json:"meeting_id"`
	StaffUser     User           `json:"staff_user"`
	CustomerUser  User           `json:"customer_user"`
	ProviderId    string         `json:"provider_id"`
	PurposeId     string         `json:"purpose_id"`
	When          time.Time      `json:"when"`
	SessionId     string         `json:"session_id"`
	CustomerToken string         `json:"customer_token"`
	StaffToken    string         `json:"staff_token"`
	Creator       ContactDetails `json:"creator"`
	Invitees      []Invitee      `json:"invitees"`
}

func NewMeeting() *Meeting {
	return &Meeting{
		Invitees: make([]Invitee, 0),
	}
}
