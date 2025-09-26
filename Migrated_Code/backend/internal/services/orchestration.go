package services

import (
	"github.com/ashish-019-hash/obp-api-backend/internal/repositories"
	"gorm.io/gorm"
)

type OrchestrationService struct {
	currencyService        *CurrencyService
	transactionService     *TransactionService
	counterpartyService    *CounterpartyLimitService
	securityService        *SecurityService
	bankRepo              repositories.BankRepository
	accountRepo           repositories.BankAccountRepository
	customerRepo          repositories.CustomerRepository
	transactionRepo       repositories.TransactionRepository
	transactionRequestRepo repositories.TransactionRequestRepository
	consentRepo           repositories.ConsentRepository
}

func NewOrchestrationService(
	currencyService *CurrencyService,
	transactionService *TransactionService,
	counterpartyService *CounterpartyLimitService,
	securityService *SecurityService,
	bankRepo repositories.BankRepository,
	accountRepo repositories.BankAccountRepository,
	customerRepo repositories.CustomerRepository,
	transactionRepo repositories.TransactionRepository,
	transactionRequestRepo repositories.TransactionRequestRepository,
	consentRepo repositories.ConsentRepository,
) *OrchestrationService {
	return &OrchestrationService{
		currencyService:        currencyService,
		transactionService:     transactionService,
		counterpartyService:    counterpartyService,
		securityService:        securityService,
		bankRepo:              bankRepo,
		accountRepo:           accountRepo,
		customerRepo:          customerRepo,
		transactionRepo:       transactionRepo,
		transactionRequestRepo: transactionRequestRepo,
		consentRepo:           consentRepo,
	}
}

func NewOrchestrationService(db *gorm.DB) *OrchestrationService {
	bankRepo := repositories.NewBankRepository(db)
	accountRepo := repositories.NewBankAccountRepository(db)
	customerRepo := repositories.NewCustomerRepository(db)
	transactionRepo := repositories.NewTransactionRepository(db)
	transactionRequestRepo := repositories.NewTransactionRequestRepository(db)
	consentRepo := repositories.NewConsentRepository(db)

	currencyService := NewCurrencyService()
	transactionService := NewTransactionService(transactionRepo, accountRepo, currencyService)
	counterpartyService := NewCounterpartyLimitService(currencyService)
	securityService := NewSecurityService(consentRepo, currencyService)

	return &OrchestrationService{
		currencyService:        currencyService,
		transactionService:     transactionService,
		counterpartyService:    counterpartyService,
		securityService:        securityService,
		bankRepo:              bankRepo,
		accountRepo:           accountRepo,
		customerRepo:          customerRepo,
		transactionRepo:       transactionRepo,
		transactionRequestRepo: transactionRequestRepo,
		consentRepo:           consentRepo,
	}
}

func (os *OrchestrationService) GetCurrencyService() *CurrencyService {
	return os.currencyService
}

func (os *OrchestrationService) GetTransactionService() *TransactionService {
	return os.transactionService
}

func (os *OrchestrationService) GetCounterpartyService() *CounterpartyLimitService {
	return os.counterpartyService
}

func (os *OrchestrationService) GetSecurityService() *SecurityService {
	return os.securityService
}

func (os *OrchestrationService) GetBankRepo() repositories.BankRepository {
	return os.bankRepo
}

func (os *OrchestrationService) GetAccountRepo() repositories.BankAccountRepository {
	return os.accountRepo
}

func (os *OrchestrationService) GetCustomerRepo() repositories.CustomerRepository {
	return os.customerRepo
}

func (os *OrchestrationService) GetTransactionRepo() repositories.TransactionRepository {
	return os.transactionRepo
}

func (os *OrchestrationService) GetTransactionRequestRepo() repositories.TransactionRequestRepository {
	return os.transactionRequestRepo
}

func (os *OrchestrationService) GetConsentRepo() repositories.ConsentRepository {
	return os.consentRepo
}
