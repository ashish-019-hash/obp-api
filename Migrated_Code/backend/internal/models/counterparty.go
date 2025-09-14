package models

import (
	"time"
)

type Counterparty struct {
	ID                                    int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	CounterpartyID                        string    `json:"counterparty_id" gorm:"uniqueIndex;size:255;not null"`
	CreatedByUserID                       string    `json:"created_by_user_id" gorm:"size:255;not null"`
	Name                                  string    `json:"name" gorm:"size:255;not null"`
	Description                           string    `json:"description" gorm:"size:2000"`
	ThisBankID                            string    `json:"this_bank_id" gorm:"size:255;not null"`
	ThisAccountID                         string    `json:"this_account_id" gorm:"size:255;not null"`
	ThisViewID                            string    `json:"this_view_id" gorm:"size:255;not null"`
	OtherAccountRoutingScheme             string    `json:"other_account_routing_scheme" gorm:"size:100"`
	OtherAccountRoutingAddress            string    `json:"other_account_routing_address" gorm:"size:255"`
	OtherAccountSecondaryRoutingScheme    string    `json:"other_account_secondary_routing_scheme" gorm:"size:100"`
	OtherAccountSecondaryRoutingAddress   string    `json:"other_account_secondary_routing_address" gorm:"size:255"`
	OtherBankRoutingScheme                string    `json:"other_bank_routing_scheme" gorm:"size:100"`
	OtherBankRoutingAddress               string    `json:"other_bank_routing_address" gorm:"size:255"`
	OtherBranchRoutingScheme              string    `json:"other_branch_routing_scheme" gorm:"size:100"`
	OtherBranchRoutingAddress             string    `json:"other_branch_routing_address" gorm:"size:255"`
	IsBeneficiary                         bool      `json:"is_beneficiary"`
	Currency                              string    `json:"currency" gorm:"size:10"`
	CreatedAt                             time.Time `json:"created_at"`
	UpdatedAt                             time.Time `json:"updated_at"`
	
	CreatedByUser                         *User        `json:"created_by_user,omitempty" gorm:"foreignKey:CreatedByUserID;references:UserID"`
	ThisAccount                           *BankAccount `json:"this_account,omitempty" gorm:"foreignKey:ThisAccountID;references:AccountID"`
	Transactions                          []Transaction `json:"transactions,omitempty" gorm:"foreignKey:CounterpartyID;references:CounterpartyID"`
}

func NewCounterparty(counterpartyID, createdByUserID, name, thisBankID, thisAccountID, thisViewID string) *Counterparty {
	return &Counterparty{
		CounterpartyID:    counterpartyID,
		CreatedByUserID:   createdByUserID,
		Name:              name,
		ThisBankID:        thisBankID,
		ThisAccountID:     thisAccountID,
		ThisViewID:        thisViewID,
		IsBeneficiary:     false,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
}

func (Counterparty) TableName() string {
	return "counterparties"
}
