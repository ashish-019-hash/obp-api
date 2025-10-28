package services

import (
	"context"
	"math/big"

	"obp-api-backend/internal/repositories"
)

type BalanceService interface {
	CalculateCurrentBalance(ctx context.Context, accountID string) (*big.Float, error)
	CalculateAvailableBalance(ctx context.Context, accountID string, heldAmount *big.Float) (*big.Float, error)
	CalculateCreditBalance(ctx context.Context, accountID string) (*big.Float, error)
	CalculateDebitBalance(ctx context.Context, accountID string) (*big.Float, error)
	CalculateMultiAccountBalance(ctx context.Context, accountIDs []string, balanceType string) (*big.Float, error)
}

type balanceService struct {
	transactionRepo repositories.TransactionRepository
}

func NewBalanceService(transactionRepo repositories.TransactionRepository) BalanceService {
	return &balanceService{
		transactionRepo: transactionRepo,
	}
}

func (s *balanceService) CalculateCurrentBalance(ctx context.Context, accountID string) (*big.Float, error) {
	creditTransactions, err := s.transactionRepo.GetCreditTransactions(ctx, accountID)
	if err != nil {
		return nil, err
	}
	
	debitTransactions, err := s.transactionRepo.GetDebitTransactions(ctx, accountID)
	if err != nil {
		return nil, err
	}
	
	creditSum := big.NewFloat(0)
	for _, tx := range creditTransactions {
		creditSum.Add(creditSum, tx.Amount)
	}
	
	debitSum := big.NewFloat(0)
	for _, tx := range debitTransactions {
		amount := new(big.Float).Set(tx.Amount)
		amount.Abs(amount)
		debitSum.Add(debitSum, amount)
	}
	
	currentBalance := new(big.Float).Sub(creditSum, debitSum)
	return currentBalance, nil
}

func (s *balanceService) CalculateAvailableBalance(ctx context.Context, accountID string, heldAmount *big.Float) (*big.Float, error) {
	currentBalance, err := s.CalculateCurrentBalance(ctx, accountID)
	if err != nil {
		return nil, err
	}
	
	availableBalance := new(big.Float).Sub(currentBalance, heldAmount)
	
	zero := big.NewFloat(0)
	if availableBalance.Cmp(zero) < 0 {
		availableBalance = big.NewFloat(0)
	}
	
	return availableBalance, nil
}

func (s *balanceService) CalculateCreditBalance(ctx context.Context, accountID string) (*big.Float, error) {
	creditTransactions, err := s.transactionRepo.GetCreditTransactions(ctx, accountID)
	if err != nil {
		return nil, err
	}
	
	creditBalance := big.NewFloat(0)
	for _, tx := range creditTransactions {
		creditBalance.Add(creditBalance, tx.Amount)
	}
	
	return creditBalance, nil
}

func (s *balanceService) CalculateDebitBalance(ctx context.Context, accountID string) (*big.Float, error) {
	debitTransactions, err := s.transactionRepo.GetDebitTransactions(ctx, accountID)
	if err != nil {
		return nil, err
	}
	
	debitBalance := big.NewFloat(0)
	for _, tx := range debitTransactions {
		amount := new(big.Float).Set(tx.Amount)
		amount.Abs(amount)
		debitBalance.Add(debitBalance, amount)
	}
	
	return debitBalance, nil
}

func (s *balanceService) CalculateMultiAccountBalance(ctx context.Context, accountIDs []string, balanceType string) (*big.Float, error) {
	totalBalance := big.NewFloat(0)
	
	for _, accountID := range accountIDs {
		var accountBalance *big.Float
		var err error
		
		switch balanceType {
		case "current":
			accountBalance, err = s.CalculateCurrentBalance(ctx, accountID)
		case "available":
			accountBalance, err = s.CalculateAvailableBalance(ctx, accountID, big.NewFloat(0))
		case "credit":
			accountBalance, err = s.CalculateCreditBalance(ctx, accountID)
		case "debit":
			accountBalance, err = s.CalculateDebitBalance(ctx, accountID)
		default:
			accountBalance, err = s.CalculateCurrentBalance(ctx, accountID)
		}
		
		if err != nil {
			return nil, err
		}
		
		totalBalance.Add(totalBalance, accountBalance)
	}
	
	return totalBalance, nil
}
