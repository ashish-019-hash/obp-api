package services

import (
	"context"
	"obp-api-backend/internal/models"
	"obp-api-backend/internal/repositories"
)

type AccountService interface {
	GetAccountsForUser(ctx context.Context, userID string) ([]*models.BankAccount, error)
	GetAccountByID(ctx context.Context, accountID string) (*models.BankAccount, error)
	CreateAccount(ctx context.Context, account *models.BankAccount) error
	GetAccountsAtBank(ctx context.Context, bankID string, limit, offset int) ([]*models.BankAccount, error)
	GetTransactionsByAccountID(ctx context.Context, accountID string) ([]*models.Transaction, error)
}

type accountService struct {
	accountRepo     repositories.BankAccountRepository
	transactionRepo repositories.TransactionRepository
}

func NewAccountService(
	accountRepo repositories.BankAccountRepository,
	transactionRepo repositories.TransactionRepository,
) AccountService {
	return &accountService{
		accountRepo:     accountRepo,
		transactionRepo: transactionRepo,
	}
}

func (s *accountService) GetAccountsForUser(ctx context.Context, userID string) ([]*models.BankAccount, error) {
	return s.accountRepo.GetByBankID(ctx, userID, 50, 0)
}

func (s *accountService) GetAccountByID(ctx context.Context, accountID string) (*models.BankAccount, error) {
	return s.accountRepo.GetByID(ctx, accountID)
}

func (s *accountService) CreateAccount(ctx context.Context, account *models.BankAccount) error {
	return s.accountRepo.Create(ctx, account)
}

func (s *accountService) GetAccountsAtBank(ctx context.Context, bankID string, limit, offset int) ([]*models.BankAccount, error) {
	return s.accountRepo.GetByBankID(ctx, bankID, limit, offset)
}

func (s *accountService) GetTransactionsByAccountID(ctx context.Context, accountID string) ([]*models.Transaction, error) {
	return s.transactionRepo.GetByAccountID(ctx, accountID, 50, 0)
}
