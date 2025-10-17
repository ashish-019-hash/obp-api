package models

import "math/big"

type ProductFee struct {
	BankId       string     `json:"bank_id"`
	ProductCode  string     `json:"product_code"`
	ProductFeeId string     `json:"product_fee_id"`
	Name         string     `json:"name"`
	IsActive     bool       `json:"is_active"`
	MoreInfo     string     `json:"more_info"`
	Currency     string     `json:"currency"`
	Amount       *big.Float `json:"amount"`
	Frequency    string     `json:"frequency"`
	Type         string     `json:"type"`
}

func NewProductFee() *ProductFee {
	return &ProductFee{
		Amount: big.NewFloat(0.0),
	}
}
