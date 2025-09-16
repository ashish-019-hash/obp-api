package models

import (
	"math/big"
	"time"
)

type Transaction struct {
	Id              string     `json:"id"`
	ThisAccount     string     `json:"this_account"`
	OtherAccount    string     `json:"other_account"`
	TransactionType string     `json:"transaction_type"`
	Amount          *big.Float `json:"amount"`
	Currency        string     `json:"currency"`
	Description     *string    `json:"description,omitempty"`
	StartDate       time.Time  `json:"start_date"`
	FinishDate      time.Time  `json:"finish_date"`
	Balance         *big.Float `json:"balance"`
}

func NewTransaction() *Transaction {
	return &Transaction{
		Amount:  big.NewFloat(0.0),
		Balance: big.NewFloat(0.0),
	}
}
