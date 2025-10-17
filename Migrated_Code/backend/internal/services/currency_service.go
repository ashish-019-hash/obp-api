package services

import (
	"context"
	"errors"
	"math/big"

	"obp-api-backend/internal/repositories"
)

type CurrencyService interface {
	GetFallbackExchangeRate(fromCurrency, toCurrency string) (*big.Float, error)
	ConvertCurrency(amount *big.Float, exchangeRate *big.Float) *big.Float
	GetExchangeRate(ctx context.Context, bankID, fromCurrency, toCurrency string) (*big.Float, error)
	ClassifyTransaction(amount *big.Float) (string, string)
	ConvertFromSmallestUnits(smallestUnits int64, currency string) *big.Float
	ConvertToSmallestUnits(amount *big.Float, currency string) int64
}

type currencyService struct {
	fxRateRepo    repositories.FXRateRepository
	fallbackRates map[string]map[string]*big.Float
}

func NewCurrencyService(fxRateRepo repositories.FXRateRepository) CurrencyService {
	return &currencyService{
		fxRateRepo:    fxRateRepo,
		fallbackRates: initializeFallbackRates(),
	}
}

func (s *currencyService) GetFallbackExchangeRate(fromCurrency, toCurrency string) (*big.Float, error) {
	if fromCurrency == toCurrency {
		return big.NewFloat(1.0), nil
	}

	if rates, exists := s.fallbackRates[fromCurrency]; exists {
		if rate, exists := rates[toCurrency]; exists {
			return new(big.Float).Copy(rate), nil
		}
	}

	return nil, errors.New("exchange rate not found")
}

func (s *currencyService) ConvertCurrency(amount *big.Float, exchangeRate *big.Float) *big.Float {
	convertedAmount := new(big.Float).Mul(amount, exchangeRate)
	convertedAmount.SetPrec(64)
	return convertedAmount
}

func (s *currencyService) GetExchangeRate(ctx context.Context, bankID, fromCurrency, toCurrency string) (*big.Float, error) {
	if fromCurrency == toCurrency {
		return big.NewFloat(1.0), nil
	}

	rate, err := s.fxRateRepo.GetRate(ctx, bankID, fromCurrency, toCurrency)
	if err == nil {
		return rate, nil
	}

	rate, err = s.fxRateRepo.GetLatestRate(ctx, fromCurrency, toCurrency)
	if err == nil {
		return rate, nil
	}

	return s.GetFallbackExchangeRate(fromCurrency, toCurrency)
}

func (s *currencyService) ClassifyTransaction(amount *big.Float) (string, string) {
	zero := big.NewFloat(0)

	if amount.Cmp(zero) > 0 {
		return "CREDIT", "INCOMING"
	} else if amount.Cmp(zero) < 0 {
		return "DEBIT", "OUTGOING"
	}

	return "NEUTRAL", "ZERO"
}

func (s *currencyService) ConvertFromSmallestUnits(smallestUnits int64, currency string) *big.Float {
	divisor := getSmallestUnitDivisor(currency)
	amount := big.NewFloat(float64(smallestUnits))
	return amount.Quo(amount, big.NewFloat(float64(divisor)))
}

func (s *currencyService) ConvertToSmallestUnits(amount *big.Float, currency string) int64 {
	multiplier := getSmallestUnitDivisor(currency)
	result := new(big.Float).Mul(amount, big.NewFloat(float64(multiplier)))

	intResult, _ := result.Int64()
	return intResult
}

func initializeFallbackRates() map[string]map[string]*big.Float {
	rates := make(map[string]map[string]*big.Float)

	rates["USD"] = map[string]*big.Float{
		"EUR": big.NewFloat(0.85),
		"GBP": big.NewFloat(0.73),
		"JPY": big.NewFloat(110.0),
		"CHF": big.NewFloat(0.92),
		"CAD": big.NewFloat(1.25),
		"AUD": big.NewFloat(1.35),
	}

	rates["EUR"] = map[string]*big.Float{
		"USD": big.NewFloat(1.18),
		"GBP": big.NewFloat(0.86),
		"JPY": big.NewFloat(129.0),
		"CHF": big.NewFloat(1.08),
		"CAD": big.NewFloat(1.47),
		"AUD": big.NewFloat(1.59),
	}

	rates["GBP"] = map[string]*big.Float{
		"USD": big.NewFloat(1.37),
		"EUR": big.NewFloat(1.16),
		"JPY": big.NewFloat(151.0),
		"CHF": big.NewFloat(1.26),
		"CAD": big.NewFloat(1.71),
		"AUD": big.NewFloat(1.85),
	}

	return rates
}

func getSmallestUnitDivisor(currency string) int {
	switch currency {
	case "JPY", "KRW":
		return 1
	case "BHD", "IQD", "JOD", "KWD", "LYD", "OMR", "TND":
		return 1000
	default:
		return 100
	}
}
