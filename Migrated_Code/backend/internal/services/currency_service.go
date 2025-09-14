package services

import (
	"errors"
	"math"
	"math/big"
)

type CurrencyService struct {
	fallbackRates map[string]float64
}

func NewCurrencyService() *CurrencyService {
	fallbackRates := map[string]float64{
		"USD": 1.0,    // Base currency
		"EUR": 0.85,   // Euro
		"GBP": 0.73,   // British Pound
		"JPY": 110.0,  // Japanese Yen
		"CHF": 0.92,   // Swiss Franc
		"CAD": 1.25,   // Canadian Dollar
		"AUD": 1.35,   // Australian Dollar
		"CNY": 6.45,   // Chinese Yuan
		"INR": 74.5,   // Indian Rupee
		"KRW": 1180.0, // South Korean Won
		"SGD": 1.35,   // Singapore Dollar
		"HKD": 7.8,    // Hong Kong Dollar
		"SEK": 8.6,    // Swedish Krona
		"NOK": 8.5,    // Norwegian Krone
	}

	return &CurrencyService{
		fallbackRates: fallbackRates,
	}
}

func (s *CurrencyService) ConvertCurrency(amount float64, fromCurrency, toCurrency string) (float64, error) {
	if fromCurrency == toCurrency {
		return amount, nil
	}

	
	fromRate, fromExists := s.fallbackRates[fromCurrency]
	toRate, toExists := s.fallbackRates[toCurrency]

	if !fromExists {
		return 0, errors.New("unsupported source currency: " + fromCurrency)
	}
	if !toExists {
		return 0, errors.New("unsupported target currency: " + toCurrency)
	}

	usdAmount := amount / fromRate
	targetAmount := usdAmount * toRate

	return s.roundHalfUp(targetAmount, 2), nil
}

func (s *CurrencyService) roundHalfUp(value float64, decimals int) float64 {
	multiplier := math.Pow(10, float64(decimals))
	
	bf := big.NewFloat(value * multiplier)
	bf.Add(bf, big.NewFloat(0.5))
	
	result, _ := bf.Int64()
	return float64(result) / multiplier
}

func (s *CurrencyService) ConvertFromSmallestUnit(amount int64, currency string) float64 {
	switch currency {
	case "JPY", "KRW": // Currencies without subdivisions
		return float64(amount)
	default:
		return float64(amount) / 100.0
	}
}

func (s *CurrencyService) ConvertToSmallestUnit(amount float64, currency string) int64 {
	switch currency {
	case "JPY", "KRW": // Currencies without subdivisions
		return int64(math.Round(amount))
	default:
		return int64(math.Round(amount * 100))
	}
}

func (s *CurrencyService) GetSupportedCurrencies() []string {
	currencies := make([]string, 0, len(s.fallbackRates))
	for currency := range s.fallbackRates {
		currencies = append(currencies, currency)
	}
	return currencies
}

func (s *CurrencyService) IsValidCurrency(currency string) bool {
	_, exists := s.fallbackRates[currency]
	return exists
}
