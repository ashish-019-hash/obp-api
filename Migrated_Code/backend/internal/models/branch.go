package models

import "time"

type Address struct {
	Line1       string `json:"line_1"`
	Line2       string `json:"line_2"`
	Line3       string `json:"line_3"`
	City        string `json:"city"`
	County      string `json:"county"`
	State       string `json:"state"`
	PostCode    string `json:"post_code"`
	CountryCode string `json:"country_code"`
}

type Location struct {
	Latitude  float64    `json:"latitude"`
	Longitude float64    `json:"longitude"`
	Date      *time.Time `json:"date,omitempty"`
	User      *BasicUser `json:"user,omitempty"`
}

type BasicUser struct {
	UserId   string `json:"user_id"`
	Provider string `json:"provider"`
	Name     string `json:"name"`
}

type Lobby struct {
	Hours string `json:"hours"`
}

type DriveUp struct {
	Hours string `json:"hours"`
}

type Routing struct {
	Scheme  string `json:"scheme"`
	Address string `json:"address"`
}

type Branch struct {
	BranchId            string     `json:"branch_id"`
	BankId              string     `json:"bank_id"`
	Name                string     `json:"name"`
	Address             Address    `json:"address"`
	Location            Location   `json:"location"`
	Lobby               *Lobby     `json:"lobby,omitempty"`
	DriveUp             *DriveUp   `json:"drive_up,omitempty"`
	IsAccessible        *bool      `json:"is_accessible,omitempty"`
	AccessibleFeatures  *string    `json:"accessible_features,omitempty"`
	BranchType          *string    `json:"branch_type,omitempty"`
	MoreInfo            *string    `json:"more_info,omitempty"`
	PhoneNumber         *string    `json:"phone_number,omitempty"`
	BranchRouting       *Routing   `json:"branch_routing,omitempty"`
}

func NewBranch() *Branch {
	return &Branch{}
}
