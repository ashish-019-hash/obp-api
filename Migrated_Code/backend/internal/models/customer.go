package models

import (
	"time"
)

type Customer struct {
	ID                         int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	CustomerID                 string    `json:"customer_id" gorm:"uniqueIndex;size:255;not null"`
	BankID                     string    `json:"bank_id" gorm:"index;size:255;not null"`
	Number                     string    `json:"number" gorm:"size:50"`
	LegalName                  string    `json:"legal_name" gorm:"size:100;not null"`
	MobileNumber               string    `json:"mobile_number" gorm:"size:100"`
	Email                      string    `json:"email" gorm:"size:100"`
	FaceImageURL               string    `json:"face_image_url" gorm:"size:255"`
	FaceImageDate              *time.Time `json:"face_image_date,omitempty"`
	DateOfBirth                *time.Time `json:"date_of_birth,omitempty"`
	RelationshipStatus         string    `json:"relationship_status" gorm:"size:25"`
	Dependents                 int       `json:"dependents"`
	HighestEducationAttained   string    `json:"highest_education_attained" gorm:"size:100"`
	EmploymentStatus           string    `json:"employment_status" gorm:"size:100"`
	KYCStatus                  bool      `json:"kyc_status"`
	LastOKDate                 *time.Time `json:"last_ok_date,omitempty"`
	Title                      string    `json:"title" gorm:"size:10"`
	BranchID                   string    `json:"branch_id" gorm:"size:255"`
	NameSuffix                 string    `json:"name_suffix" gorm:"size:10"`
	CreatedAt                  time.Time `json:"created_at"`
	UpdatedAt                  time.Time `json:"updated_at"`
	
	Bank                       *Bank                    `json:"bank,omitempty" gorm:"foreignKey:BankID;references:BankID"`
	UserCustomerLinks          []UserCustomerLink       `json:"user_customer_links,omitempty" gorm:"foreignKey:CustomerID;references:CustomerID"`
	CustomerAccountLinks       []CustomerAccountLink    `json:"customer_account_links,omitempty" gorm:"foreignKey:CustomerID;references:CustomerID"`
}

func NewCustomer(customerID, bankID, legalName string) *Customer {
	return &Customer{
		CustomerID: customerID,
		BankID:     bankID,
		LegalName:  legalName,
		KYCStatus:  false,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

func (Customer) TableName() string {
	return "customers"
}
