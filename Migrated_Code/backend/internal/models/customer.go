package models

import "time"

type CustomerFaceImage struct {
	Url  string    `json:"url"`
	Date time.Time `json:"date"`
}

type Customer struct {
	CustomerId               string            `json:"customer_id"`
	BankId                   string            `json:"bank_id"`
	Number                   string            `json:"number"`
	LegalName                string            `json:"legal_name"`
	MobileNumber             string            `json:"mobile_number"`
	Email                    string            `json:"email"`
	FaceImage                CustomerFaceImage `json:"face_image"`
	DateOfBirth              time.Time         `json:"date_of_birth"`
	RelationshipStatus       string            `json:"relationship_status"`
	Dependents               int               `json:"dependents"`
	DobOfDependents          []time.Time       `json:"dob_of_dependents"`
	HighestEducationAttained string            `json:"highest_education_attained"`
	EmploymentStatus         string            `json:"employment_status"`
	KycStatus                bool              `json:"kyc_status"`
	LastOkDate               time.Time         `json:"last_ok_date"`
}

func NewCustomer() *Customer {
	return &Customer{
		DobOfDependents: make([]time.Time, 0),
	}
}
