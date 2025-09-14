package models

import (
	"time"
)

type Bank struct {
	ID                   int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	BankID               string `json:"bank_id" gorm:"uniqueIndex;size:255;not null"`
	ShortName            string `json:"short_name" gorm:"size:100;not null"`
	FullName             string `json:"full_name" gorm:"size:255;not null"`
	LogoURL              string `json:"logo_url" gorm:"size:255"`
	WebsiteURL           string `json:"website_url" gorm:"size:255"`
	SwiftBIC             string `json:"swift_bic" gorm:"size:255"`
	NationalIdentifier   string `json:"national_identifier" gorm:"size:255"`
	BankRoutingScheme    string `json:"bank_routing_scheme" gorm:"size:255"`
	BankRoutingAddress   string `json:"bank_routing_address" gorm:"size:255"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

func NewBank(bankID, shortName, fullName string) *Bank {
	return &Bank{
		BankID:    bankID,
		ShortName: shortName,
		FullName:  fullName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (Bank) TableName() string {
	return "banks"
}
