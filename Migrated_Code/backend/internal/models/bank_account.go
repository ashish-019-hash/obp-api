package models

import (
	"math/big"
	"time"
)

type BankAccount struct {
	AccountId   string     `json:"account_id"`
	BankId      string     `json:"bank_id"`
	AccountType string     `json:"account_type"`
	Balance     *big.Float `json:"balance"`
	Currency    string     `json:"currency"`
	Name        string     `json:"name"`
	Number      string     `json:"number"`
	Label       string     `json:"label"`
	LastUpdate  time.Time  `json:"last_update"`
	BranchId    string     `json:"branch_id"`
}

func NewBankAccount() *BankAccount {
	return &BankAccount{
		Balance: big.NewFloat(0.0),
	}
}
