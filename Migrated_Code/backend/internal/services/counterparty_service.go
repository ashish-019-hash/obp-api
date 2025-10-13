package services

import (
	"errors"
	"time"

	"github.com/ashish-019-hash/obp-api-backend/internal/models"
)

type CounterpartyLimitService struct {
	currencyService *CurrencyService
}

func NewCounterpartyLimitService(currencyService *CurrencyService) *CounterpartyLimitService {
	return &CounterpartyLimitService{
		currencyService: currencyService,
	}
}

type CounterpartyLimits struct {
	SingleTransactionAmount int64 // Maximum amount per transaction
	MonthlyTransactionAmount int64 // Maximum total amount per month
	YearlyTransactionAmount  int64 // Maximum total amount per year
	SingleTransactionCount   int   // Maximum transactions per single operation
	MonthlyTransactionCount  int   // Maximum transactions per month
	YearlyTransactionCount   int   // Maximum transactions per year
	Currency                 string
}

func (s *CounterpartyLimitService) GetDefaultLimits(currency string) *CounterpartyLimits {
	baseLimits := &CounterpartyLimits{
		SingleTransactionAmount:  1000000, // $10,000 in cents
		MonthlyTransactionAmount: 5000000, // $50,000 in cents
		YearlyTransactionAmount:  60000000, // $600,000 in cents
		SingleTransactionCount:   1,
		MonthlyTransactionCount:  100,
		YearlyTransactionCount:   1200,
		Currency:                 "USD",
	}

	if currency != "USD" {
		singleAmount := s.currencyService.ConvertFromSmallestUnit(baseLimits.SingleTransactionAmount, "USD")
		monthlyAmount := s.currencyService.ConvertFromSmallestUnit(baseLimits.MonthlyTransactionAmount, "USD")
		yearlyAmount := s.currencyService.ConvertFromSmallestUnit(baseLimits.YearlyTransactionAmount, "USD")

		convertedSingle, _ := s.currencyService.ConvertCurrency(singleAmount, "USD", currency)
		convertedMonthly, _ := s.currencyService.ConvertCurrency(monthlyAmount, "USD", currency)
		convertedYearly, _ := s.currencyService.ConvertCurrency(yearlyAmount, "USD", currency)

		baseLimits.SingleTransactionAmount = s.currencyService.ConvertToSmallestUnit(convertedSingle, currency)
		baseLimits.MonthlyTransactionAmount = s.currencyService.ConvertToSmallestUnit(convertedMonthly, currency)
		baseLimits.YearlyTransactionAmount = s.currencyService.ConvertToSmallestUnit(convertedYearly, currency)
		baseLimits.Currency = currency
	}

	return baseLimits
}

func (s *CounterpartyLimitService) ValidateTransactionAgainstLimits(
	amount int64,
	currency string,
	counterpartyID string,
	existingTransactions []*models.Transaction,
	limits *CounterpartyLimits,
) error {
	transactionAmount := amount
	if currency != limits.Currency {
		decimalAmount := s.currencyService.ConvertFromSmallestUnit(amount, currency)
		convertedAmount, err := s.currencyService.ConvertCurrency(decimalAmount, currency, limits.Currency)
		if err != nil {
			return err
		}
		transactionAmount = s.currencyService.ConvertToSmallestUnit(convertedAmount, limits.Currency)
	}

	if transactionAmount > limits.SingleTransactionAmount {
		return errors.New("transaction amount exceeds single transaction limit")
	}

	now := time.Now()
	currentMonth := now.Month()
	currentYear := now.Year()

	var monthlyAmount int64 = 0
	var yearlyAmount int64 = 0
	var monthlyCount int = 0
	var yearlyCount int = 0

	for _, tx := range existingTransactions {
		txTime := tx.CreatedAt
		
		txAmount := tx.Amount
		if tx.Currency != limits.Currency {
			decimalTxAmount := s.currencyService.ConvertFromSmallestUnit(tx.Amount, tx.Currency)
			convertedTxAmount, err := s.currencyService.ConvertCurrency(decimalTxAmount, tx.Currency, limits.Currency)
			if err != nil {
				continue // Skip if conversion fails
			}
			txAmount = s.currencyService.ConvertToSmallestUnit(convertedTxAmount, limits.Currency)
		}

		if txTime.Year() == currentYear {
			yearlyAmount += txAmount
			yearlyCount++

			if txTime.Month() == currentMonth {
				monthlyAmount += txAmount
				monthlyCount++
			}
		}
	}

	if monthlyAmount+transactionAmount > limits.MonthlyTransactionAmount {
		return errors.New("transaction would exceed monthly amount limit")
	}

	if yearlyAmount+transactionAmount > limits.YearlyTransactionAmount {
		return errors.New("transaction would exceed yearly amount limit")
	}

	if limits.SingleTransactionCount < 1 {
		return errors.New("single transaction count limit exceeded")
	}

	if monthlyCount+1 > limits.MonthlyTransactionCount {
		return errors.New("transaction would exceed monthly count limit")
	}

	if yearlyCount+1 > limits.YearlyTransactionCount {
		return errors.New("transaction would exceed yearly count limit")
	}

	return nil
}

func (s *CounterpartyLimitService) GetCounterpartyTransactionHistory(
	counterpartyID string,
	transactions []*models.Transaction,
) []*models.Transaction {
	var counterpartyTransactions []*models.Transaction
	
	for _, tx := range transactions {
		if tx.CounterpartyID != nil && *tx.CounterpartyID == counterpartyID {
			counterpartyTransactions = append(counterpartyTransactions, tx)
		}
	}
	
	return counterpartyTransactions
}
