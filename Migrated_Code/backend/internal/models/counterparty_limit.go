package models

import "math/big"

type CounterpartyLimit struct {
	CounterpartyLimitId              string     `json:"counterparty_limit_id"`
	BankId                           string     `json:"bank_id"`
	AccountId                        string     `json:"account_id"`
	ViewId                           string     `json:"view_id"`
	CounterpartyId                   string     `json:"counterparty_id"`
	Currency                         string     `json:"currency"`
	MaxSingleAmount                  *big.Float `json:"max_single_amount"`
	MaxMonthlyAmount                 *big.Float `json:"max_monthly_amount"`
	MaxNumberOfMonthlyTransactions   int        `json:"max_number_of_monthly_transactions"`
	MaxYearlyAmount                  *big.Float `json:"max_yearly_amount"`
	MaxNumberOfYearlyTransactions    int        `json:"max_number_of_yearly_transactions"`
	MaxTotalAmount                   *big.Float `json:"max_total_amount"`
	MaxNumberOfTransactions          int        `json:"max_number_of_transactions"`
}

func NewCounterpartyLimit() *CounterpartyLimit {
	return &CounterpartyLimit{
		MaxSingleAmount:  big.NewFloat(0.0),
		MaxMonthlyAmount: big.NewFloat(0.0),
		MaxYearlyAmount:  big.NewFloat(0.0),
		MaxTotalAmount:   big.NewFloat(0.0),
	}
}
