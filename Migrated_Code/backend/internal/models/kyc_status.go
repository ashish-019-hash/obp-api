package models

import "time"

type KycStatus struct {
	CustomerId     string    `json:"customer_id"`
	CustomerNumber string    `json:"customer_number"`
	Ok             bool      `json:"ok"`
	Date           time.Time `json:"date"`
}

func NewKycStatus() *KycStatus {
	return &KycStatus{}
}
