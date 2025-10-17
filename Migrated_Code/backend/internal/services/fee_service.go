package services

import (
	"context"
	"math/big"
)

type FeeService interface {
	CalculateProductFee(ctx context.Context, productID string, amount *big.Float, currency string) (*FeeCalculationResult, error)
	CalculateTransactionFee(ctx context.Context, transactionType string, amount *big.Float, currency string) (*FeeCalculationResult, error)
	ApplyFeeStructure(ctx context.Context, feeType string, baseAmount *big.Float, currency string) (*FeeCalculationResult, error)
}

type FeeCalculationResult struct {
	BaseFee       *big.Float
	PercentageFee *big.Float
	TotalFee      *big.Float
	FeeBreakdown  map[string]*big.Float
	Currency      string
}

type feeService struct {
	currencyService CurrencyService
	feeStructures   map[string]FeeStructure
}

type FeeStructure struct {
	BaseFee        *big.Float
	PercentageRate *big.Float
	MinimumFee     *big.Float
	MaximumFee     *big.Float
	Currency       string
}

func NewFeeService(currencyService CurrencyService) FeeService {
	return &feeService{
		currencyService: currencyService,
		feeStructures:   initializeFeeStructures(),
	}
}

func (s *feeService) CalculateProductFee(ctx context.Context, productID string, amount *big.Float, currency string) (*FeeCalculationResult, error) {
	feeStructure, exists := s.feeStructures[productID]
	if !exists {
		feeStructure = s.feeStructures["default"]
	}

	return s.calculateFee(ctx, feeStructure, amount, currency)
}

func (s *feeService) CalculateTransactionFee(ctx context.Context, transactionType string, amount *big.Float, currency string) (*FeeCalculationResult, error) {
	feeKey := "transaction_" + transactionType
	feeStructure, exists := s.feeStructures[feeKey]
	if !exists {
		feeStructure = s.feeStructures["default_transaction"]
	}

	return s.calculateFee(ctx, feeStructure, amount, currency)
}

func (s *feeService) ApplyFeeStructure(ctx context.Context, feeType string, baseAmount *big.Float, currency string) (*FeeCalculationResult, error) {
	feeStructure, exists := s.feeStructures[feeType]
	if !exists {
		feeStructure = s.feeStructures["default"]
	}

	return s.calculateFee(ctx, feeStructure, baseAmount, currency)
}

func (s *feeService) calculateFee(ctx context.Context, structure FeeStructure, amount *big.Float, currency string) (*FeeCalculationResult, error) {
	convertedAmount := amount
	if currency != structure.Currency {
		exchangeRate, err := s.currencyService.GetExchangeRate(ctx, "", currency, structure.Currency)
		if err != nil {
			return nil, err
		}
		convertedAmount = s.currencyService.ConvertCurrency(amount, exchangeRate)
	}

	baseFee := new(big.Float).Copy(structure.BaseFee)

	percentageFee := new(big.Float).Mul(convertedAmount, structure.PercentageRate)
	percentageFee.Quo(percentageFee, big.NewFloat(100))

	totalFee := new(big.Float).Add(baseFee, percentageFee)

	if structure.MinimumFee != nil && totalFee.Cmp(structure.MinimumFee) < 0 {
		totalFee = new(big.Float).Copy(structure.MinimumFee)
	}

	if structure.MaximumFee != nil && totalFee.Cmp(structure.MaximumFee) > 0 {
		totalFee = new(big.Float).Copy(structure.MaximumFee)
	}

	breakdown := map[string]*big.Float{
		"base_fee":       baseFee,
		"percentage_fee": percentageFee,
		"total_fee":      totalFee,
	}

	return &FeeCalculationResult{
		BaseFee:       baseFee,
		PercentageFee: percentageFee,
		TotalFee:      totalFee,
		FeeBreakdown:  breakdown,
		Currency:      structure.Currency,
	}, nil
}

func initializeFeeStructures() map[string]FeeStructure {
	structures := make(map[string]FeeStructure)

	structures["default"] = FeeStructure{
		BaseFee:        big.NewFloat(1.00),
		PercentageRate: big.NewFloat(0.5),
		MinimumFee:     big.NewFloat(0.50),
		MaximumFee:     big.NewFloat(50.00),
		Currency:       "USD",
	}

	structures["default_transaction"] = FeeStructure{
		BaseFee:        big.NewFloat(0.25),
		PercentageRate: big.NewFloat(0.1),
		MinimumFee:     big.NewFloat(0.10),
		MaximumFee:     big.NewFloat(10.00),
		Currency:       "USD",
	}

	structures["wire_transfer"] = FeeStructure{
		BaseFee:        big.NewFloat(15.00),
		PercentageRate: big.NewFloat(0.25),
		MinimumFee:     big.NewFloat(15.00),
		MaximumFee:     big.NewFloat(100.00),
		Currency:       "USD",
	}

	structures["international_transfer"] = FeeStructure{
		BaseFee:        big.NewFloat(25.00),
		PercentageRate: big.NewFloat(0.75),
		MinimumFee:     big.NewFloat(25.00),
		MaximumFee:     big.NewFloat(200.00),
		Currency:       "USD",
	}

	return structures
}
