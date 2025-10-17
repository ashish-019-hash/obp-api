package models

import (
	"math/big"
	"time"
)

type StandingOrder struct {
	StandingOrderId string     `json:"standing_order_id"`
	BankId          string     `json:"bank_id"`
	AccountId       string     `json:"account_id"`
	CustomerId      string     `json:"customer_id"`
	UserId          string     `json:"user_id"`
	CounterpartyId  string     `json:"counterparty_id"`
	AmountValue     *big.Float `json:"amount_value"`
	AmountCurrency  string     `json:"amount_currency"`
	WhenFrequency   string     `json:"when_frequency"`
	WhenDetail      string     `json:"when_detail"`
	DateSigned      time.Time  `json:"date_signed"`
	DateCancelled   time.Time  `json:"date_cancelled"`
	DateStarts      time.Time  `json:"date_starts"`
	DateExpires     time.Time  `json:"date_expires"`
	Active          bool       `json:"active"`
}

func NewStandingOrder() *StandingOrder {
	return &StandingOrder{
		AmountValue: big.NewFloat(0.0),
	}
}
