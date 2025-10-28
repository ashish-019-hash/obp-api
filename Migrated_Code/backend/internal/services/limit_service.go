package services

import (
	"context"
	"errors"
	"math/big"

	"obp-api-backend/internal/repositories"
)

type LimitService interface {
	ValidateCounterpartyLimit(ctx context.Context, counterpartyID string, amount *big.Float, currency string) error
	ValidateTransactionLimit(ctx context.Context, accountID string, amount *big.Float, currency string, period string) error
	CheckPaymentCoverage(ctx context.Context, accountID string, paymentAmount *big.Float, paymentCurrency, accountCurrency string) (*PaymentCoverageResult, error)
}

type PaymentCoverageResult struct {
	CoverageStatus     string
	CoverageConfidence string
	AvailableAmount    *big.Float
	ShortfallAmount    *big.Float
}

type limitService struct {
	counterpartyLimitRepo repositories.CounterpartyLimitRepository
	balanceService        BalanceService
	currencyService       CurrencyService
}

func NewLimitService(
	counterpartyLimitRepo repositories.CounterpartyLimitRepository,
	balanceService BalanceService,
	currencyService CurrencyService,
) LimitService {
	return &limitService{
		counterpartyLimitRepo: counterpartyLimitRepo,
		balanceService:        balanceService,
		currencyService:       currencyService,
	}
}

func (s *limitService) ValidateCounterpartyLimit(ctx context.Context, counterpartyID string, amount *big.Float, currency string) error {
	limit, err := s.counterpartyLimitRepo.GetByCounterpartyID(ctx, counterpartyID)
	if err != nil {
		return err
	}
	
	if currency != limit.Currency {
		exchangeRate, err := s.currencyService.GetExchangeRate(ctx, "", currency, limit.Currency)
		if err != nil {
			return err
		}
		amount = s.currencyService.ConvertCurrency(amount, exchangeRate)
	}
	
	if amount.Cmp(limit.MaxSingleAmount) > 0 {
		return errors.New("amount exceeds counterparty single transaction limit")
	}
	
	return nil
}

func (s *limitService) ValidateTransactionLimit(ctx context.Context, accountID string, amount *big.Float, currency string, period string) error {
	return nil
}

func (s *limitService) CheckPaymentCoverage(ctx context.Context, accountID string, paymentAmount *big.Float, paymentCurrency, accountCurrency string) (*PaymentCoverageResult, error) {
	accountBalance, err := s.balanceService.CalculateCurrentBalance(ctx, accountID)
	if err != nil {
		return nil, err
	}
	
	convertedPaymentAmount := paymentAmount
	if paymentCurrency != accountCurrency {
		exchangeRate, err := s.currencyService.GetExchangeRate(ctx, "", paymentCurrency, accountCurrency)
		if err != nil {
			return nil, err
		}
		convertedPaymentAmount = s.currencyService.ConvertCurrency(paymentAmount, exchangeRate)
	}
	
	reservedAmount := big.NewFloat(0)
	minimumBalance := big.NewFloat(0)
	
	availableBalance := new(big.Float).Sub(accountBalance, reservedAmount)
	availableBalance.Sub(availableBalance, minimumBalance)
	
	result := &PaymentCoverageResult{}
	
	if availableBalance.Cmp(convertedPaymentAmount) >= 0 {
		result.CoverageStatus = "COVERED"
		result.CoverageConfidence = "HIGH"
		result.AvailableAmount = new(big.Float).Copy(availableBalance)
		result.ShortfallAmount = big.NewFloat(0)
	} else {
		partialThreshold := new(big.Float).Mul(convertedPaymentAmount, big.NewFloat(0.9))
		if availableBalance.Cmp(partialThreshold) >= 0 {
			result.CoverageStatus = "PARTIALLY_COVERED"
			result.CoverageConfidence = "MEDIUM"
		} else {
			result.CoverageStatus = "NOT_COVERED"
			result.CoverageConfidence = "LOW"
		}
		
		result.AvailableAmount = new(big.Float).Copy(availableBalance)
		result.ShortfallAmount = new(big.Float).Sub(convertedPaymentAmount, availableBalance)
		
		if result.ShortfallAmount.Cmp(big.NewFloat(0)) < 0 {
			result.ShortfallAmount = big.NewFloat(0)
		}
	}
	
	return result, nil
}
