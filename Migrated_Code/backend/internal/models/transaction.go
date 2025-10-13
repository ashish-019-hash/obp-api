package models

import (
	"time"
)

type Transaction struct {
	ID              int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	TransactionID   string    `json:"transaction_id" gorm:"uniqueIndex;size:255;not null"`
	AccountID       string    `json:"account_id" gorm:"index;size:255;not null"`
	CounterpartyID  *string   `json:"counterparty_id,omitempty" gorm:"size:255"`
	TransactionType string    `json:"transaction_type" gorm:"size:100;not null"`
	Amount          int64     `json:"amount" gorm:"not null"` // Amount in smallest currency unit
	Currency        string    `json:"currency" gorm:"size:10;not null"`
	Description     *string   `json:"description,omitempty" gorm:"size:2000"`
	StartDate       time.Time `json:"start_date"`
	FinishDate      time.Time `json:"finish_date"`
	Balance         int64     `json:"balance" gorm:"not null"` // Balance after transaction
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	
	Account         *BankAccount  `json:"account,omitempty" gorm:"foreignKey:AccountID;references:AccountID"`
	Counterparty    *Counterparty `json:"counterparty,omitempty" gorm:"foreignKey:CounterpartyID;references:CounterpartyID"`
}

func NewTransaction(transactionID, accountID, transactionType, currency string, amount, balance int64) *Transaction {
	now := time.Now()
	return &Transaction{
		TransactionID:   transactionID,
		AccountID:       accountID,
		TransactionType: transactionType,
		Amount:          amount,
		Currency:        currency,
		Balance:         balance,
		StartDate:       now,
		FinishDate:      now,
		CreatedAt:       now,
		UpdatedAt:       now,
	}
}

func (Transaction) TableName() string {
	return "transactions"
}
