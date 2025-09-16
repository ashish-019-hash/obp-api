package repositories

import (
	"context"
	"math/big"
	"time"

	"obp-api-backend/internal/models"
)

type BankRepository interface {
	Create(ctx context.Context, bank *models.Bank) error
	GetByID(ctx context.Context, bankID string) (*models.Bank, error)
	Update(ctx context.Context, bank *models.Bank) error
	Delete(ctx context.Context, bankID string) error
	List(ctx context.Context, limit, offset int) ([]*models.Bank, error)
}

type BankAccountRepository interface {
	Create(ctx context.Context, account *models.BankAccount) error
	GetByID(ctx context.Context, accountID string) (*models.BankAccount, error)
	GetByBankID(ctx context.Context, bankID string) ([]*models.BankAccount, error)
	Update(ctx context.Context, account *models.BankAccount) error
	Delete(ctx context.Context, accountID string) error
	GetAccountsByRouting(ctx context.Context, scheme, address string) ([]*models.BankAccount, error)
}

type TransactionRepository interface {
	Create(ctx context.Context, transaction *models.Transaction) error
	GetByID(ctx context.Context, transactionID string) (*models.Transaction, error)
	GetByAccountID(ctx context.Context, accountID string, limit, offset int) ([]*models.Transaction, error)
	GetByDateRange(ctx context.Context, accountID string, fromDate, toDate time.Time) ([]*models.Transaction, error)
	GetCreditTransactions(ctx context.Context, accountID string) ([]*models.Transaction, error)
	GetDebitTransactions(ctx context.Context, accountID string) ([]*models.Transaction, error)
	CalculateBalance(ctx context.Context, accountID string) (*big.Float, error)
}

type CounterpartyRepository interface {
	Create(ctx context.Context, counterparty *models.Counterparty) error
	GetByID(ctx context.Context, counterpartyID string) (*models.Counterparty, error)
	Update(ctx context.Context, counterparty *models.Counterparty) error
	Delete(ctx context.Context, counterpartyID string) error
	GetByAccountID(ctx context.Context, accountID string) ([]*models.Counterparty, error)
}

type CounterpartyLimitRepository interface {
	Create(ctx context.Context, limit *models.CounterpartyLimit) error
	GetByCounterpartyID(ctx context.Context, counterpartyID string) (*models.CounterpartyLimit, error)
	Update(ctx context.Context, limit *models.CounterpartyLimit) error
	Delete(ctx context.Context, limitID string) error
	ValidateLimit(ctx context.Context, counterpartyID string, amount *big.Float, currency string) (bool, error)
}

type CustomerRepository interface {
	Create(ctx context.Context, customer *models.Customer) error
	GetByID(ctx context.Context, customerID string) (*models.Customer, error)
	GetByBankID(ctx context.Context, bankID string) ([]*models.Customer, error)
	Update(ctx context.Context, customer *models.Customer) error
	Delete(ctx context.Context, customerID string) error
}

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, userID string) (*models.User, error)
	GetByProvider(ctx context.Context, provider, providerID string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, userID string) error
}

type FXRateRepository interface {
	Create(ctx context.Context, bankID, fromCurrency, toCurrency string, rate *big.Float) error
	GetRate(ctx context.Context, bankID, fromCurrency, toCurrency string) (*big.Float, error)
	GetLatestRate(ctx context.Context, fromCurrency, toCurrency string) (*big.Float, error)
}

type MetricsRepository interface {
	RecordAPICall(ctx context.Context, consumerID, userID, url string, duration int64) error
	GetMetrics(ctx context.Context, fromDate, toDate time.Time, consumerID, userID, url string) (*models.APIMetrics, error)
	GetTopAPIs(ctx context.Context, fromDate, toDate time.Time, limit int) ([]*models.APIRanking, error)
	GetTopConsumers(ctx context.Context, fromDate, toDate time.Time, limit int) ([]*models.ConsumerRanking, error)
}

type RateLimitRepository interface {
	IncrementCounter(ctx context.Context, consumerKey, period, periodKey string, limit int) (int, error)
	GetCounter(ctx context.Context, consumerKey, period, periodKey string) (int, error)
	ResetCounter(ctx context.Context, consumerKey, period, periodKey string) error
}
