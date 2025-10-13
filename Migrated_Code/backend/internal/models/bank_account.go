package models

import (
	"time"
)

type BankAccount struct {
	ID              int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	AccountID       string    `json:"account_id" gorm:"uniqueIndex;size:255;not null"`
	BankID          string    `json:"bank_id" gorm:"index;size:255;not null"`
	AccountType     string    `json:"account_type" gorm:"size:255;not null"`
	Balance         int64     `json:"balance" gorm:"not null"` // Balance in smallest currency unit (cents)
	Currency        string    `json:"currency" gorm:"size:10;not null"`
	Name            string    `json:"name" gorm:"size:255;not null"`
	Label           string    `json:"label" gorm:"size:255"`
	Number          string    `json:"number" gorm:"size:255;not null"`
	BranchID        string    `json:"branch_id" gorm:"size:255"`
	AccountHolder   string    `json:"account_holder" gorm:"size:100"`
	LastUpdate      time.Time `json:"last_update"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	
	Bank            *Bank              `json:"bank,omitempty" gorm:"foreignKey:BankID;references:BankID"`
	AccountRoutings []AccountRouting   `json:"account_routings,omitempty" gorm:"foreignKey:AccountID;references:AccountID"`
	Transactions    []Transaction      `json:"transactions,omitempty" gorm:"foreignKey:AccountID;references:AccountID"`
}

func NewBankAccount(accountID, bankID, accountType, currency, name, number string, balance int64) *BankAccount {
	return &BankAccount{
		AccountID:   accountID,
		BankID:      bankID,
		AccountType: accountType,
		Balance:     balance,
		Currency:    currency,
		Name:        name,
		Number:      number,
		LastUpdate:  time.Now(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (BankAccount) TableName() string {
	return "bank_accounts"
}
