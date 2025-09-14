package models

import (
	"time"
)

type TransactionRequest struct {
	ID                              int64      `json:"id" gorm:"primaryKey;autoIncrement"`
	TransactionRequestID            string     `json:"transaction_request_id" gorm:"uniqueIndex;size:255;not null"`
	Type                            string     `json:"type" gorm:"size:32;not null"`
	TransactionIDs                  string     `json:"transaction_ids" gorm:"size:2000"`
	Status                          string     `json:"status" gorm:"size:32;not null"`
	StartDate                       *time.Time `json:"start_date,omitempty"`
	EndDate                         *time.Time `json:"end_date,omitempty"`
	
	ChallengeID                     string     `json:"challenge_id" gorm:"size:64"`
	ChallengeAllowedAttempts        int        `json:"challenge_allowed_attempts"`
	ChallengeChallengeType          string     `json:"challenge_challenge_type" gorm:"size:100"`
	
	ChargeSummary                   string     `json:"charge_summary" gorm:"size:64"`
	ChargeAmount                    string     `json:"charge_amount" gorm:"size:32"`
	ChargeCurrency                  string     `json:"charge_currency" gorm:"size:3"`
	ChargePolicy                    string     `json:"charge_policy" gorm:"size:32"`
	
	BodyValueCurrency               string     `json:"body_value_currency" gorm:"size:3"`
	BodyValueAmount                 string     `json:"body_value_amount" gorm:"size:32"`
	BodyDescription                 string     `json:"body_description" gorm:"size:2000"`
	Details                         string     `json:"details" gorm:"type:text"`
	
	FromBankID                      string     `json:"from_bank_id" gorm:"size:255"`
	FromAccountID                   string     `json:"from_account_id" gorm:"size:255"`
	
	ToBankID                        string     `json:"to_bank_id" gorm:"size:255"`
	ToAccountID                     string     `json:"to_account_id" gorm:"size:255"`
	
	Name                            string     `json:"name" gorm:"size:64"`
	ThisBankID                      string     `json:"this_bank_id" gorm:"size:255"`
	ThisAccountID                   string     `json:"this_account_id" gorm:"size:255"`
	ThisViewID                      string     `json:"this_view_id" gorm:"size:255"`
	CounterpartyID                  string     `json:"counterparty_id" gorm:"size:255"`
	OtherAccountRoutingScheme       string     `json:"other_account_routing_scheme" gorm:"size:32"`
	OtherAccountRoutingAddress      string     `json:"other_account_routing_address" gorm:"size:64"`
	OtherBankRoutingScheme          string     `json:"other_bank_routing_scheme" gorm:"size:32"`
	OtherBankRoutingAddress         string     `json:"other_bank_routing_address" gorm:"size:64"`
	IsBeneficiary                   bool       `json:"is_beneficiary"`
	
	PaymentStartDate                *time.Time `json:"payment_start_date,omitempty"`
	PaymentEndDate                  *time.Time `json:"payment_end_date,omitempty"`
	PaymentExecutionRule            string     `json:"payment_execution_rule" gorm:"size:64"`
	PaymentFrequency                string     `json:"payment_frequency" gorm:"size:64"`
	PaymentDayOfExecution           string     `json:"payment_day_of_execution" gorm:"size:64"`
	
	ConsentReferenceID              string     `json:"consent_reference_id" gorm:"size:64"`
	APIStandard                     string     `json:"api_standard" gorm:"size:50"`
	APIVersion                      string     `json:"api_version" gorm:"size:50"`
	
	CreatedAt                       time.Time  `json:"created_at"`
	UpdatedAt                       time.Time  `json:"updated_at"`
	
	FromAccount                     *BankAccount `json:"from_account,omitempty" gorm:"foreignKey:FromAccountID;references:AccountID"`
	ToAccount                       *BankAccount `json:"to_account,omitempty" gorm:"foreignKey:ToAccountID;references:AccountID"`
	Counterparty                    *Counterparty `json:"counterparty,omitempty" gorm:"foreignKey:CounterpartyID;references:CounterpartyID"`
	Consent                         *Consent     `json:"consent,omitempty" gorm:"foreignKey:ConsentReferenceID;references:ConsentID"`
}

func NewTransactionRequest(transactionRequestID, requestType, status string) *TransactionRequest {
	return &TransactionRequest{
		TransactionRequestID: transactionRequestID,
		Type:                requestType,
		Status:              status,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}
}

func (TransactionRequest) TableName() string {
	return "transaction_requests"
}
