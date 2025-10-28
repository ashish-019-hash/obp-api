package services

import (
	"context"
	"errors"
	"math/big"
	"regexp"
	"strings"
	"time"
)

type ValidationService interface {
	ValidateFieldLength(field string, minLength, maxLength int) error
	ValidateNumericRange(value *big.Float, min, max *big.Float) error
	ValidateEnumValue(value string, allowedValues []string) error
	ValidateCurrency(currency string) error
	ValidateIBAN(iban string) error
	ValidateEmail(email string) error
	ValidatePhoneNumber(phone string) error
	ValidateDateRange(startDate, endDate time.Time) error
	ValidateAccountNumber(accountNumber string) error
	ValidateTransactionAmount(amount *big.Float, currency string) error
	ValidateConditionalFields(ctx context.Context, data map[string]interface{}) error
	ValidateCrossFieldDependencies(ctx context.Context, data map[string]interface{}) error
}

type validationService struct {
	currencyService CurrencyService
}

func NewValidationService(currencyService CurrencyService) ValidationService {
	return &validationService{
		currencyService: currencyService,
	}
}

func (s *validationService) ValidateFieldLength(field string, minLength, maxLength int) error {
	length := len(strings.TrimSpace(field))
	
	if length < minLength {
		return errors.New("field length is below minimum required")
	}
	
	if maxLength > 0 && length > maxLength {
		return errors.New("field length exceeds maximum allowed")
	}
	
	return nil
}

func (s *validationService) ValidateNumericRange(value *big.Float, min, max *big.Float) error {
	if min != nil && value.Cmp(min) < 0 {
		return errors.New("value is below minimum allowed")
	}
	
	if max != nil && value.Cmp(max) > 0 {
		return errors.New("value exceeds maximum allowed")
	}
	
	return nil
}

func (s *validationService) ValidateEnumValue(value string, allowedValues []string) error {
	for _, allowed := range allowedValues {
		if strings.EqualFold(value, allowed) {
			return nil
		}
	}
	
	return errors.New("value is not in allowed enumeration")
}

func (s *validationService) ValidateCurrency(currency string) error {
	validCurrencies := []string{
		"USD", "EUR", "GBP", "JPY", "CHF", "CAD", "AUD", "NZD",
		"SEK", "NOK", "DKK", "PLN", "CZK", "HUF", "BGN", "RON",
		"HRK", "RUB", "TRY", "BRL", "MXN", "ARS", "CLP", "COP",
		"PEN", "UYU", "CNY", "HKD", "SGD", "KRW", "THB", "MYR",
		"IDR", "PHP", "VND", "INR", "PKR", "LKR", "BDT", "NPR",
		"ZAR", "EGP", "MAD", "TND", "KES", "UGX", "TZS", "GHS",
		"NGN", "XOF", "XAF", "ILS", "JOD", "KWD", "BHD", "QAR",
		"AED", "SAR", "OMR", "IRR", "AFN", "AMD", "AZN", "GEL",
		"KZT", "KGS", "TJS", "TMT", "UZS", "MNT", "BTN", "MVR",
	}
	
	return s.ValidateEnumValue(currency, validCurrencies)
}

func (s *validationService) ValidateIBAN(iban string) error {
	iban = strings.ReplaceAll(strings.ToUpper(iban), " ", "")
	
	if len(iban) < 15 || len(iban) > 34 {
		return errors.New("IBAN length is invalid")
	}
	
	ibanRegex := regexp.MustCompile(`^[A-Z]{2}[0-9]{2}[A-Z0-9]+$`)
	if !ibanRegex.MatchString(iban) {
		return errors.New("IBAN format is invalid")
	}
	
	return nil
}

func (s *validationService) ValidateEmail(email string) error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	
	if !emailRegex.MatchString(email) {
		return errors.New("email format is invalid")
	}
	
	return nil
}

func (s *validationService) ValidatePhoneNumber(phone string) error {
	phone = strings.ReplaceAll(phone, " ", "")
	phone = strings.ReplaceAll(phone, "-", "")
	phone = strings.ReplaceAll(phone, "(", "")
	phone = strings.ReplaceAll(phone, ")", "")
	
	phoneRegex := regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)
	
	if !phoneRegex.MatchString(phone) {
		return errors.New("phone number format is invalid")
	}
	
	return nil
}

func (s *validationService) ValidateDateRange(startDate, endDate time.Time) error {
	if startDate.After(endDate) {
		return errors.New("start date cannot be after end date")
	}
	
	now := time.Now()
	if startDate.After(now.AddDate(10, 0, 0)) {
		return errors.New("start date is too far in the future")
	}
	
	if endDate.Before(now.AddDate(-100, 0, 0)) {
		return errors.New("end date is too far in the past")
	}
	
	return nil
}

func (s *validationService) ValidateAccountNumber(accountNumber string) error {
	if err := s.ValidateFieldLength(accountNumber, 1, 50); err != nil {
		return err
	}
	
	accountRegex := regexp.MustCompile(`^[A-Za-z0-9\-_]+$`)
	if !accountRegex.MatchString(accountNumber) {
		return errors.New("account number contains invalid characters")
	}
	
	return nil
}

func (s *validationService) ValidateTransactionAmount(amount *big.Float, currency string) error {
	if err := s.ValidateCurrency(currency); err != nil {
		return err
	}
	
	zero := big.NewFloat(0)
	if amount.Cmp(zero) <= 0 {
		return errors.New("transaction amount must be positive")
	}
	
	maxAmount := big.NewFloat(1000000000)
	if amount.Cmp(maxAmount) > 0 {
		return errors.New("transaction amount exceeds maximum allowed")
	}
	
	return nil
}

func (s *validationService) ValidateConditionalFields(ctx context.Context, data map[string]interface{}) error {
	return nil
}

func (s *validationService) ValidateCrossFieldDependencies(ctx context.Context, data map[string]interface{}) error {
	return nil
}
