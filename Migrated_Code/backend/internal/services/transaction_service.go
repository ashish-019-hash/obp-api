package services

import (
	"context"
	"errors"
	"math/big"

	"obp-api-backend/internal/models"
	"obp-api-backend/internal/repositories"
)

type TransactionService interface {
	ProcessTransaction(ctx context.Context, transaction *models.Transaction) error
	ClassifyTransaction(ctx context.Context, transaction *models.Transaction) (*TransactionClassification, error)
	ValidateTransaction(ctx context.Context, transaction *models.Transaction) error
	ApplyBusinessRules(ctx context.Context, transaction *models.Transaction) error
	CalculateNewBalance(ctx context.Context, accountID string, transactionAmount *big.Float) (*big.Float, error)
}

type TransactionClassification struct {
	Type        string
	Direction   string
	Category    string
	RiskLevel   string
	RequiresApproval bool
}

type transactionService struct {
	transactionRepo repositories.TransactionRepository
	balanceService  BalanceService
	limitService    LimitService
	securityService SecurityService
	validationService ValidationService
	currencyService CurrencyService
}

func NewTransactionService(
	transactionRepo repositories.TransactionRepository,
	balanceService BalanceService,
	limitService LimitService,
	securityService SecurityService,
	validationService ValidationService,
	currencyService CurrencyService,
) TransactionService {
	return &transactionService{
		transactionRepo:   transactionRepo,
		balanceService:    balanceService,
		limitService:      limitService,
		securityService:   securityService,
		validationService: validationService,
		currencyService:   currencyService,
	}
}

func (s *transactionService) ProcessTransaction(ctx context.Context, transaction *models.Transaction) error {
	if err := s.ValidateTransaction(ctx, transaction); err != nil {
		return err
	}
	
	if err := s.ApplyBusinessRules(ctx, transaction); err != nil {
		return err
	}
	
	classification, err := s.ClassifyTransaction(ctx, transaction)
	if err != nil {
		return err
	}
	
	if classification.RequiresApproval {
		return errors.New("transaction requires manual approval")
	}
	
	newBalance, err := s.CalculateNewBalance(ctx, transaction.ThisAccount, transaction.Amount)
	if err != nil {
		return err
	}
	
	transaction.Balance = newBalance
	
	return s.transactionRepo.Create(ctx, transaction)
}

func (s *transactionService) ClassifyTransaction(ctx context.Context, transaction *models.Transaction) (*TransactionClassification, error) {
	transactionType, direction := s.currencyService.ClassifyTransaction(transaction.Amount)
	
	classification := &TransactionClassification{
		Type:      transactionType,
		Direction: direction,
		Category:  s.categorizeTransaction(transaction),
		RiskLevel: s.assessRiskLevel(ctx, transaction),
	}
	
	classification.RequiresApproval = s.requiresApproval(ctx, transaction, classification)
	
	return classification, nil
}

func (s *transactionService) ValidateTransaction(ctx context.Context, transaction *models.Transaction) error {
	if err := s.validationService.ValidateTransactionAmount(transaction.Amount, transaction.Currency); err != nil {
		return err
	}
	
	if transaction.OtherAccount != "" {
		if err := s.limitService.ValidateCounterpartyLimit(ctx, transaction.OtherAccount, transaction.Amount, transaction.Currency); err != nil {
			return err
		}
	}
	
	return nil
}

func (s *transactionService) ApplyBusinessRules(ctx context.Context, transaction *models.Transaction) error {
	requiresChallenge, err := s.securityService.CheckChallengeThreshold(ctx, transaction.Amount, transaction.Currency)
	if err != nil {
		return err
	}
	
	if requiresChallenge {
		return errors.New("transaction exceeds challenge threshold - additional authentication required")
	}
	
	return nil
}

func (s *transactionService) CalculateNewBalance(ctx context.Context, accountID string, transactionAmount *big.Float) (*big.Float, error) {
	currentBalance, err := s.balanceService.CalculateCurrentBalance(ctx, accountID)
	if err != nil {
		return nil, err
	}
	
	newBalance := new(big.Float).Add(currentBalance, transactionAmount)
	return newBalance, nil
}

func (s *transactionService) categorizeTransaction(transaction *models.Transaction) string {
	if transaction.Description == nil || *transaction.Description == "" {
		return "GENERAL"
	}
	
	description := *transaction.Description
	
	if contains(description, "ATM") {
		return "ATM_WITHDRAWAL"
	} else if contains(description, "TRANSFER") {
		return "TRANSFER"
	} else if contains(description, "PAYMENT") {
		return "PAYMENT"
	} else if contains(description, "DEPOSIT") {
		return "DEPOSIT"
	} else if contains(description, "FEE") {
		return "FEE"
	}
	
	return "OTHER"
}

func (s *transactionService) assessRiskLevel(ctx context.Context, transaction *models.Transaction) string {
	amount := transaction.Amount
	
	highRiskThreshold := big.NewFloat(10000)
	mediumRiskThreshold := big.NewFloat(1000)
	
	if amount.Cmp(highRiskThreshold) > 0 {
		return "HIGH"
	} else if amount.Cmp(mediumRiskThreshold) > 0 {
		return "MEDIUM"
	}
	
	return "LOW"
}

func (s *transactionService) requiresApproval(ctx context.Context, transaction *models.Transaction, classification *TransactionClassification) bool {
	if classification.RiskLevel == "HIGH" {
		return true
	}
	
	if classification.Category == "INTERNATIONAL_TRANSFER" {
		return true
	}
	
	return false
}

func parseAmount(amountStr string) *big.Float {
	amount, ok := new(big.Float).SetString(amountStr)
	if !ok {
		return big.NewFloat(0)
	}
	return amount
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[:len(substr)] == substr
}
