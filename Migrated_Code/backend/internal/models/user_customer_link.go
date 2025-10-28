package models

import (
	"time"
)

type UserCustomerLink struct {
	ID                   int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	UserCustomerLinkID   string    `json:"user_customer_link_id" gorm:"uniqueIndex;size:255;not null"`
	UserID               string    `json:"user_id" gorm:"index;size:255;not null"`
	CustomerID           string    `json:"customer_id" gorm:"index;size:255;not null"`
	DateInserted         time.Time `json:"date_inserted"`
	IsActive             bool      `json:"is_active"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	
	User                 *User     `json:"user,omitempty" gorm:"foreignKey:UserID;references:UserID"`
	Customer             *Customer `json:"customer,omitempty" gorm:"foreignKey:CustomerID;references:CustomerID"`
}

func NewUserCustomerLink(userCustomerLinkID, userID, customerID string) *UserCustomerLink {
	now := time.Now()
	return &UserCustomerLink{
		UserCustomerLinkID: userCustomerLinkID,
		UserID:             userID,
		CustomerID:         customerID,
		DateInserted:       now,
		IsActive:           true,
		CreatedAt:          now,
		UpdatedAt:          now,
	}
}

func (UserCustomerLink) TableName() string {
	return "user_customer_links"
}
