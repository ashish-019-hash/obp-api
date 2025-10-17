package services

import (
	"obp-api-backend/internal/repositories"
)

type Services struct {
	Currency   CurrencyService
	Balance    BalanceService
	Limit      LimitService
	RateLimit  RateLimitingService
	Analytics  AnalyticsService
	Validation ValidationService
}

func NewServices(repos *Repositories) *Services {
	currencyService := NewCurrencyService(repos.FXRate)
	balanceService := NewBalanceService(repos.Transaction)
	limitService := NewLimitService(repos.CounterpartyLimit, balanceService, currencyService)
	rateLimitService := NewRateLimitingService(repos.RateLimit)
	analyticsService := NewAnalyticsService(repos.Customer, repos.Metrics, currencyService)
	validationService := NewValidationService(currencyService)

	return &Services{
		Currency:   currencyService,
		Balance:    balanceService,
		Limit:      limitService,
		RateLimit:  rateLimitService,
		Analytics:  analyticsService,
		Validation: validationService,
	}
}

type Repositories struct {
	Bank              repositories.BankRepository
	BankAccount       repositories.BankAccountRepository
	Transaction       repositories.TransactionRepository
	Counterparty      repositories.CounterpartyRepository
	CounterpartyLimit repositories.CounterpartyLimitRepository
	Customer          repositories.CustomerRepository
	User              repositories.UserRepository
	FXRate            repositories.FXRateRepository
	Metrics           repositories.MetricsRepository
	RateLimit         repositories.RateLimitRepository
}

func NewRepositories() *Repositories {
	return &Repositories{
		Bank:              repositories.NewBankRepository(),
		BankAccount:       repositories.NewBankAccountRepository(),
		Transaction:       repositories.NewTransactionRepository(),
		Counterparty:      repositories.NewCounterpartyRepository(),
		CounterpartyLimit: repositories.NewCounterpartyLimitRepository(),
		Customer:          repositories.NewCustomerRepository(),
		User:              repositories.NewUserRepository(),
		FXRate:            repositories.NewFXRateRepository(),
		Metrics:           repositories.NewMetricsRepository(),
		RateLimit:         repositories.NewRateLimitRepository(),
	}
}
