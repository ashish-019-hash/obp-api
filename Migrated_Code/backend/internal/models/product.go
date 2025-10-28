package models

import (
	"time"
)

type Product struct {
	ID                      int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	BankID                  string    `json:"bank_id" gorm:"index;size:255;not null"`
	ProductCode             string    `json:"product_code" gorm:"size:50;not null"`
	ParentProductCode       string    `json:"parent_product_code" gorm:"size:50"`
	Name                    string    `json:"name" gorm:"size:125;not null"`
	Category                string    `json:"category" gorm:"size:50"`
	Family                  string    `json:"family" gorm:"size:50"`
	SuperFamily             string    `json:"super_family" gorm:"size:50"`
	MoreInfoURL             string    `json:"more_info_url" gorm:"size:2000"`
	TermsAndConditionsURL   string    `json:"terms_and_conditions_url" gorm:"size:2000"`
	Details                 string    `json:"details" gorm:"size:2000"`
	Description             string    `json:"description" gorm:"size:2000"`
	LicenseID               string    `json:"license_id" gorm:"size:255"`
	LicenseName             string    `json:"license_name" gorm:"size:255"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
	
	Bank                    *Bank     `json:"bank,omitempty" gorm:"foreignKey:BankID;references:BankID"`
}

func NewProduct(bankID, productCode, name string) *Product {
	return &Product{
		BankID:      bankID,
		ProductCode: productCode,
		Name:        name,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (Product) TableName() string {
	return "products"
}
