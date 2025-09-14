package test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/ashish-019-hash/obp-api-backend/internal/models"
)

func TestBankJSONMarshaling(t *testing.T) {
	bank := models.NewBank("test-bank-id", "Test Bank", "Test Bank Full Name")
	bank.LogoURL = "https://example.com/logo.png"
	bank.WebsiteURL = "https://example.com"
	bank.SwiftBIC = "TESTBIC1"

	jsonData, err := json.Marshal(bank)
	if err != nil {
		t.Fatalf("Failed to marshal Bank to JSON: %v", err)
	}

	var unmarshaled models.Bank
	err = json.Unmarshal(jsonData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal Bank from JSON: %v", err)
	}

	if unmarshaled.BankID != bank.BankID {
		t.Errorf("Expected BankID %s, got %s", bank.BankID, unmarshaled.BankID)
	}
	if unmarshaled.ShortName != bank.ShortName {
		t.Errorf("Expected ShortName %s, got %s", bank.ShortName, unmarshaled.ShortName)
	}
}

func TestBankAccountJSONMarshaling(t *testing.T) {
	account := models.NewBankAccount("test-account-id", "test-bank-id", "CURRENT", "USD", "Test Account", "12345", 100000)
	account.Label = "Primary Account"
	account.AccountHolder = "John Doe"

	jsonData, err := json.Marshal(account)
	if err != nil {
		t.Fatalf("Failed to marshal BankAccount to JSON: %v", err)
	}

	var unmarshaled models.BankAccount
	err = json.Unmarshal(jsonData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal BankAccount from JSON: %v", err)
	}

	if unmarshaled.AccountID != account.AccountID {
		t.Errorf("Expected AccountID %s, got %s", account.AccountID, unmarshaled.AccountID)
	}
	if unmarshaled.Balance != account.Balance {
		t.Errorf("Expected Balance %d, got %d", account.Balance, unmarshaled.Balance)
	}
}

func TestCustomerJSONMarshaling(t *testing.T) {
	customer := models.NewCustomer("test-customer-id", "test-bank-id", "John Doe")
	customer.Email = "john.doe@example.com"
	customer.MobileNumber = "+1234567890"
	customer.KYCStatus = true

	jsonData, err := json.Marshal(customer)
	if err != nil {
		t.Fatalf("Failed to marshal Customer to JSON: %v", err)
	}

	var unmarshaled models.Customer
	err = json.Unmarshal(jsonData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal Customer from JSON: %v", err)
	}

	if unmarshaled.CustomerID != customer.CustomerID {
		t.Errorf("Expected CustomerID %s, got %s", customer.CustomerID, unmarshaled.CustomerID)
	}
	if unmarshaled.KYCStatus != customer.KYCStatus {
		t.Errorf("Expected KYCStatus %t, got %t", customer.KYCStatus, unmarshaled.KYCStatus)
	}
}

func TestTransactionRequestJSONMarshaling(t *testing.T) {
	txnReq := models.NewTransactionRequest("test-txn-req-id", "TRANSFER", "PENDING")
	txnReq.BodyValueAmount = "100.00"
	txnReq.BodyValueCurrency = "USD"
	txnReq.BodyDescription = "Test payment"

	jsonData, err := json.Marshal(txnReq)
	if err != nil {
		t.Fatalf("Failed to marshal TransactionRequest to JSON: %v", err)
	}

	var unmarshaled models.TransactionRequest
	err = json.Unmarshal(jsonData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal TransactionRequest from JSON: %v", err)
	}

	if unmarshaled.TransactionRequestID != txnReq.TransactionRequestID {
		t.Errorf("Expected TransactionRequestID %s, got %s", txnReq.TransactionRequestID, unmarshaled.TransactionRequestID)
	}
	if unmarshaled.Status != txnReq.Status {
		t.Errorf("Expected Status %s, got %s", txnReq.Status, unmarshaled.Status)
	}
}

func TestConsentJSONMarshaling(t *testing.T) {
	consent := models.NewConsent("test-consent-id", "test-user-id", "test-bank-id", "ACTIVE", "PAYMENT")

	jsonData, err := json.Marshal(consent)
	if err != nil {
		t.Fatalf("Failed to marshal Consent to JSON: %v", err)
	}

	var unmarshaled models.Consent
	err = json.Unmarshal(jsonData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal Consent from JSON: %v", err)
	}

	if unmarshaled.ConsentID != consent.ConsentID {
		t.Errorf("Expected ConsentID %s, got %s", consent.ConsentID, unmarshaled.ConsentID)
	}
	if unmarshaled.ConsentType != consent.ConsentType {
		t.Errorf("Expected ConsentType %s, got %s", consent.ConsentType, unmarshaled.ConsentType)
	}
}

func TestUserCustomerLinkJSONMarshaling(t *testing.T) {
	link := models.NewUserCustomerLink("test-link-id", "test-user-id", "test-customer-id")

	jsonData, err := json.Marshal(link)
	if err != nil {
		t.Fatalf("Failed to marshal UserCustomerLink to JSON: %v", err)
	}

	var unmarshaled models.UserCustomerLink
	err = json.Unmarshal(jsonData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal UserCustomerLink from JSON: %v", err)
	}

	if unmarshaled.UserCustomerLinkID != link.UserCustomerLinkID {
		t.Errorf("Expected UserCustomerLinkID %s, got %s", link.UserCustomerLinkID, unmarshaled.UserCustomerLinkID)
	}
	if !unmarshaled.IsActive {
		t.Errorf("Expected IsActive to be true")
	}
}

func TestJSONFieldNaming(t *testing.T) {
	bank := models.NewBank("test-bank-id", "Test Bank", "Test Bank Full Name")
	
	jsonData, err := json.Marshal(bank)
	if err != nil {
		t.Fatalf("Failed to marshal Bank to JSON: %v", err)
	}

	jsonStr := string(jsonData)
	
	expectedFields := []string{
		"\"bank_id\":",
		"\"short_name\":",
		"\"full_name\":",
		"\"created_at\":",
		"\"updated_at\":",
	}
	
	for _, field := range expectedFields {
		if !strings.Contains(jsonStr, field) {
			t.Errorf("Expected JSON to contain field %s, but it was not found in: %s", field, jsonStr)
		}
	}
}
