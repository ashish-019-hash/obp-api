package models

import (
	"time"
)

type AccountRouting struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	AccountID string    `json:"account_id" gorm:"index;size:255;not null"`
	Scheme    string    `json:"scheme" gorm:"size:100;not null"`
	Address   string    `json:"address" gorm:"size:255;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	
	Account   *BankAccount `json:"account,omitempty" gorm:"foreignKey:AccountID;references:AccountID"`
}

func NewAccountRouting(accountID, scheme, address string) *AccountRouting {
	return &AccountRouting{
		AccountID: accountID,
		Scheme:    scheme,
		Address:   address,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (AccountRouting) TableName() string {
	return "account_routings"
}
