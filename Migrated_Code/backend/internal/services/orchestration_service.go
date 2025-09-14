package services

import (
	"errors"
	"strconv"
	"time"

	"github.com/ashish-019-hash/obp-api-backend/internal/models"
	"github.com/ashish-019-hash/obp-api-backend/internal/repositories"
)

type OrchestrationService struct {
	currencyService         *CurrencyService
	transactionService      *TransactionService
	counterpartyService     *CounterpartyLimitService
	securityService         *SecurityService

	bankRepo               repositories.BankRepository
	accountRepo            repositories.BankAccountRepository
	customerRepo           repositories.CustomerRepository
	transactionRepo        repositories.TransactionRepository
	transactionRequestRepo repositories.TransactionRequestRepository
	consentRepo            repositories.ConsentRepository
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
		currencyService:         currencyService,
		transactionService:      transactionService,
		counterpartyService:     counterpartyService,
		securityService:         securityService,
		bankRepo:               bankRepo,
		accountRepo:            accountRepo,
		customerRepo:           customerRepo,
		transactionRepo:        transactionRepo,
		transactionRequestRepo: transactionRequestRepo,
		consentRepo:            consentRepo,
	}
}

func (s *OrchestrationService) ProcessPaymentRequest(request *models.TransactionRequest, userID string) error {
	if request.TransactionRequestID == "" || request.Type == "" {
		return errors.New("invalid transaction request: missing required fields")
	}

	if !s.currencyService.IsValidCurrency(request.BodyValueCurrency) {
		return errors.New("unsupported currency: " + request.BodyValueCurrency)
	}

	amount, err := strconv.ParseFloat(request.BodyValueAmount, 64)
	if err != nil {
		return errors.New("invalid amount format: " + request.BodyValueAmount)
	}
	amountCents := s.currencyService.ConvertToSmallestUnit(amount, request.BodyValueCurrency)

	authorized, err := s.securityService.CheckTransactionAuthorization(userID, amountCents, request.BodyValueCurrency)
	if err != nil {
		return err
	}
	if !authorized {
		return errors.New("transaction not authorized")
	}

	if request.BodyToCounterpartyID != nil {
		allTransactions, err := s.transactionRepo.List(1000, 0)
		if err != nil {
			return err
		}

		counterpartyTransactions := s.counterpartyService.GetCounterpartyTransactionHistory(*request.BodyToCounterpartyID, allTransactions)
		limits := s.counterpartyService.GetDefaultLimits(request.BodyValueCurrency)

		err = s.counterpartyService.ValidateTransactionAgainstLimits(
			amountCents,
			request.BodyValueCurrency,
			*request.BodyToCounterpartyID,
			counterpartyTransactions,
			limits,
		)
		if err != nil {
			return err
		}
	}

	request.Status = "PENDING"
	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()

	err = s.transactionRequestRepo.Create(request)
	if err != nil {
		return err
	}

	if request.Type == "TRANSFER" || request.Type == "PAYMENT" {
		transaction := &models.Transaction{
			TransactionID:   generateTransactionID(),
			AccountID:       request.BodyFromAccountID,
			Amount:          amountCents,
			Currency:        request.BodyValueCurrency,
			Description:     &request.BodyDescription,
			CounterpartyID:  request.BodyToCounterpartyID,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}

		err = s.transactionService.CreateTransaction(transaction)
		if err != nil {
			s.transactionRequestRepo.UpdateStatus(request.TransactionRequestID, "FAILED")
			return err
		}

		err = s.transactionRequestRepo.UpdateStatus(request.TransactionRequestID, "COMPLETED")
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *OrchestrationService) GetAccountSummary(accountID string, targetCurrency string) (*AccountSummary, error) {
	account, err := s.accountRepo.GetByAccountID(accountID)
	if err != nil {
		return nil, err
	}

	transactions, err := s.transactionService.GetTransactionsByAccount(accountID)
	if err != nil {
		return nil, err
	}

	balanceDecimal := s.currencyService.ConvertFromSmallestUnit(account.Balance, account.Currency)
	convertedBalance, err := s.currencyService.ConvertCurrency(balanceDecimal, account.Currency, targetCurrency)
	if err != nil {
		return nil, err
	}

	var convertedTransactions []TransactionSummary
	for _, tx := range transactions {
		txDecimal := s.currencyService.ConvertFromSmallestUnit(tx.Amount, tx.Currency)
		convertedAmount, err := s.currencyService.ConvertCurrency(txDecimal, tx.Currency, targetCurrency)
		if err != nil {
			convertedAmount = 0
		}

		description := ""
		if tx.Description != nil {
			description = *tx.Description
		}
		
		convertedTransactions = append(convertedTransactions, TransactionSummary{
			TransactionID: tx.TransactionID,
			Amount:        convertedAmount,
			Currency:      targetCurrency,
			Type:          tx.TransactionType,
			Description:   description,
			Date:          tx.CreatedAt,
		})
	}

	return &AccountSummary{
		AccountID:            account.AccountID,
		AccountName:          account.Name,
		Balance:              convertedBalance,
		Currency:             targetCurrency,
		OriginalBalance:      balanceDecimal,
		OriginalCurrency:     account.Currency,
		RecentTransactions:   convertedTransactions,
		LastUpdate:           account.LastUpdate,
	}, nil
}

func (s *OrchestrationService) ValidateCustomerKYC(customerID string) error {
	customer, err := s.customerRepo.GetByCustomerID(customerID)
	if err != nil {
		return err
	}

	if customer.LegalName == "" || customer.DateOfBirth == nil {
		return errors.New("insufficient customer information for KYC validation")
	}

	err = s.customerRepo.UpdateKYCStatus(customerID, true)
	if err != nil {
		return err
	}

	return nil
}

type AccountSummary struct {
	AccountID            string               `json:"account_id"`
	AccountName          string               `json:"account_name"`
	Balance              float64              `json:"balance"`
	Currency             string               `json:"currency"`
	OriginalBalance      float64              `json:"original_balance"`
	OriginalCurrency     string               `json:"original_currency"`
	RecentTransactions   []TransactionSummary `json:"recent_transactions"`
	LastUpdate           time.Time            `json:"last_update"`
}

type TransactionSummary struct {
	TransactionID string    `json:"transaction_id"`
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	Type          string    `json:"type"`
	Description   string    `json:"description"`
	Date          time.Time `json:"date"`
}

func generateTransactionID() string {
	return "txn_" + time.Now().Format("20060102150405")
}
