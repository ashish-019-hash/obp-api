package models

import "time"

type KycDocument struct {
	BankId         string    `json:"bank_id"`
	CustomerId     string    `json:"customer_id"`
	IdKycDocument  string    `json:"id_kyc_document"`
	CustomerNumber string    `json:"customer_number"`
	Type           string    `json:"type"`
	Number         string    `json:"number"`
	IssueDate      time.Time `json:"issue_date"`
	IssuePlace     string    `json:"issue_place"`
	ExpiryDate     time.Time `json:"expiry_date"`
}

func NewKycDocument() *KycDocument {
	return &KycDocument{}
}
