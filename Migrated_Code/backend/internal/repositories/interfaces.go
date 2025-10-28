package repositories

import (
	"github.com/ashish-019-hash/obp-api-backend/internal/models"
)

type BankRepository interface {
	Create(bank *models.Bank) error
	GetByID(id int64) (*models.Bank, error)
	GetByBankID(bankID string) (*models.Bank, error)
	Update(bank *models.Bank) error
	Delete(id int64) error
	List(limit, offset int) ([]*models.Bank, error)
}

type BankAccountRepository interface {
	Create(account *models.BankAccount) error
	GetByID(id int64) (*models.BankAccount, error)
	GetByAccountID(accountID string) (*models.BankAccount, error)
	GetByBankID(bankID string) ([]*models.BankAccount, error)
	UpdateBalance(accountID string, newBalance int64) error
	Update(account *models.BankAccount) error
	Delete(id int64) error
	List(limit, offset int) ([]*models.BankAccount, error)
}

type CustomerRepository interface {
	Create(customer *models.Customer) error
	GetByID(id int64) (*models.Customer, error)
	GetByCustomerID(customerID string) (*models.Customer, error)
	GetByBankID(bankID string) ([]*models.Customer, error)
	UpdateKYCStatus(customerID string, status bool) error
	Update(customer *models.Customer) error
	Delete(id int64) error
	List(limit, offset int) ([]*models.Customer, error)
}

type TransactionRepository interface {
	Create(transaction *models.Transaction) error
	GetByID(id int64) (*models.Transaction, error)
	GetByTransactionID(transactionID string) (*models.Transaction, error)
	GetByAccountID(accountID string) ([]*models.Transaction, error)
	GetByDateRange(accountID string, startDate, endDate string) ([]*models.Transaction, error)
	Update(transaction *models.Transaction) error
	Delete(id int64) error
	List(limit, offset int) ([]*models.Transaction, error)
}

type TransactionRequestRepository interface {
	Create(request *models.TransactionRequest) error
	GetByID(id int64) (*models.TransactionRequest, error)
	GetByTransactionRequestID(requestID string) (*models.TransactionRequest, error)
	UpdateStatus(requestID string, status string) error
	GetByStatus(status string) ([]*models.TransactionRequest, error)
	Update(request *models.TransactionRequest) error
	Delete(id int64) error
	List(limit, offset int) ([]*models.TransactionRequest, error)
}

type ConsentRepository interface {
	Create(consent *models.Consent) error
	GetByID(id int64) (*models.Consent, error)
	GetByConsentID(consentID string) (*models.Consent, error)
	GetByUserID(userID string) ([]*models.Consent, error)
	UpdateStatus(consentID string, status string) error
	Update(consent *models.Consent) error
	Delete(id int64) error
	List(limit, offset int) ([]*models.Consent, error)
}
