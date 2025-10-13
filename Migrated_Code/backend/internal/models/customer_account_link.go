package models

import (
	"time"
)

type CustomerAccountLink struct {
	ID         int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	CustomerID string    `json:"customer_id" gorm:"index;size:255;not null"`
	BankID     string    `json:"bank_id" gorm:"index;size:255;not null"`
	AccountID  string    `json:"account_id" gorm:"index;size:255;not null"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	
	Customer   *Customer    `json:"customer,omitempty" gorm:"foreignKey:CustomerID;references:CustomerID"`
	Account    *BankAccount `json:"account,omitempty" gorm:"foreignKey:AccountID;references:AccountID"`
}

func NewCustomerAccountLink(customerID, bankID, accountID string) *CustomerAccountLink {
	return &CustomerAccountLink{
		CustomerID: customerID,
		BankID:     bankID,
		AccountID:  accountID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

func (CustomerAccountLink) TableName() string {
	return "customer_account_links"
}
