package services

import (
	"errors"
	"time"

	"github.com/ashish-019-hash/obp-api-backend/internal/models"
	"github.com/ashish-019-hash/obp-api-backend/internal/repositories"
)

type TransactionService struct {
	transactionRepo repositories.TransactionRepository
	accountRepo     repositories.BankAccountRepository
	currencyService *CurrencyService
}

func NewTransactionService(
	transactionRepo repositories.TransactionRepository,
	accountRepo repositories.BankAccountRepository,
	currencyService *CurrencyService,
) *TransactionService {
	return &TransactionService{
		transactionRepo: transactionRepo,
		accountRepo:     accountRepo,
		currencyService: currencyService,
	}
}

func (s *TransactionService) CreateTransaction(transaction *models.Transaction) error {
	account, err := s.accountRepo.GetByAccountID(transaction.AccountID)
	if err != nil {
		return errors.New("account not found: " + transaction.AccountID)
	}

	if transaction.Amount > 0 {
		transaction.TransactionType = "CREDIT"
	} else {
		transaction.TransactionType = "DEBIT"
	}

	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()

	err = s.transactionRepo.Create(transaction)
	if err != nil {
		return err
	}

	newBalance := account.Balance + transaction.Amount
	err = s.accountRepo.UpdateBalance(account.AccountID, newBalance)
	if err != nil {
		return err
	}

	return nil
}

func (s *TransactionService) GetTransactionsByAccount(accountID string) ([]*models.Transaction, error) {
	return s.transactionRepo.GetByAccountID(accountID)
}

func (s *TransactionService) GetTransactionByID(transactionID string) (*models.Transaction, error) {
	return s.transactionRepo.GetByTransactionID(transactionID)
}

func (s *TransactionService) GetTransactionsByDateRange(accountID, startDate, endDate string) ([]*models.Transaction, error) {
	return s.transactionRepo.GetByDateRange(accountID, startDate, endDate)
}

func (s *TransactionService) ConvertTransactionAmount(transaction *models.Transaction, targetCurrency string) (float64, error) {
	if !s.currencyService.IsValidCurrency(transaction.Currency) {
		return 0, errors.New("invalid source currency: " + transaction.Currency)
	}
	if !s.currencyService.IsValidCurrency(targetCurrency) {
		return 0, errors.New("invalid target currency: " + targetCurrency)
	}

	decimalAmount := s.currencyService.ConvertFromSmallestUnit(transaction.Amount, transaction.Currency)
	
	convertedAmount, err := s.currencyService.ConvertCurrency(decimalAmount, transaction.Currency, targetCurrency)
	if err != nil {
		return 0, err
	}

	return convertedAmount, nil
}

func (s *TransactionService) CalculateAccountBalance(accountID string) (int64, error) {
	transactions, err := s.transactionRepo.GetByAccountID(accountID)
	if err != nil {
		return 0, err
	}

	var balance int64 = 0
	for _, transaction := range transactions {
		balance += transaction.Amount
	}

	return balance, nil
}
