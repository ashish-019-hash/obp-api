package services

import (
	"context"
	"errors"
	"math/big"
)

type SecurityService interface {
	CheckViewBasedAccess(ctx context.Context, userID, accountID, viewID string) (bool, error)
	ValidateAmountVisibility(ctx context.Context, userID, accountID string, amount *big.Float) (*big.Float, error)
	ValidateBalanceVisibility(ctx context.Context, userID, accountID string, balance *big.Float) (*big.Float, error)
	CheckChallengeThreshold(ctx context.Context, amount *big.Float, currency string) (bool, error)
	ApplyAccessControl(ctx context.Context, userID string, resource string, action string) (bool, error)
}

type securityService struct {
	currencyService   CurrencyService
	challengeThresholds map[string]*big.Float
	accessRules       map[string][]string
}

func NewSecurityService(currencyService CurrencyService) SecurityService {
	return &securityService{
		currencyService:     currencyService,
		challengeThresholds: initializeChallengeThresholds(),
		accessRules:         initializeAccessRules(),
	}
}

func (s *securityService) CheckViewBasedAccess(ctx context.Context, userID, accountID, viewID string) (bool, error) {
	if userID == "" || accountID == "" || viewID == "" {
		return false, errors.New("missing required parameters for access check")
	}
	
	switch viewID {
	case "owner":
		return s.isAccountOwner(ctx, userID, accountID), nil
	case "public":
		return true, nil
	case "accountant":
		return s.hasAccountantAccess(ctx, userID, accountID), nil
	case "auditor":
		return s.hasAuditorAccess(ctx, userID, accountID), nil
	default:
		return false, errors.New("unknown view type")
	}
}

func (s *securityService) ValidateAmountVisibility(ctx context.Context, userID, accountID string, amount *big.Float) (*big.Float, error) {
	hasAccess, err := s.CheckViewBasedAccess(ctx, userID, accountID, "owner")
	if err != nil {
		return nil, err
	}
	
	if !hasAccess {
		hasPublicAccess, err := s.CheckViewBasedAccess(ctx, userID, accountID, "public")
		if err != nil {
			return nil, err
		}
		
		if hasPublicAccess {
			return big.NewFloat(0), nil
		}
		
		return nil, errors.New("insufficient permissions to view amount")
	}
	
	return amount, nil
}

func (s *securityService) ValidateBalanceVisibility(ctx context.Context, userID, accountID string, balance *big.Float) (*big.Float, error) {
	hasAccess, err := s.CheckViewBasedAccess(ctx, userID, accountID, "owner")
	if err != nil {
		return nil, err
	}
	
	if !hasAccess {
		hasAccountantAccess, err := s.CheckViewBasedAccess(ctx, userID, accountID, "accountant")
		if err != nil {
			return nil, err
		}
		
		if hasAccountantAccess {
			return balance, nil
		}
		
		return big.NewFloat(0), nil
	}
	
	return balance, nil
}

func (s *securityService) CheckChallengeThreshold(ctx context.Context, amount *big.Float, currency string) (bool, error) {
	threshold, exists := s.challengeThresholds[currency]
	if !exists {
		threshold = s.challengeThresholds["USD"]
		
		exchangeRate, err := s.currencyService.GetExchangeRate(ctx, "", currency, "USD")
		if err != nil {
			return false, err
		}
		
		convertedAmount := s.currencyService.ConvertCurrency(amount, exchangeRate)
		return convertedAmount.Cmp(threshold) > 0, nil
	}
	
	return amount.Cmp(threshold) > 0, nil
}

func (s *securityService) ApplyAccessControl(ctx context.Context, userID string, resource string, action string) (bool, error) {
	if allowedActions, exists := s.accessRules[resource]; exists {
		for _, allowedAction := range allowedActions {
			if allowedAction == action || allowedAction == "*" {
				return true, nil
			}
		}
	}
	
	return false, nil
}

func (s *securityService) isAccountOwner(ctx context.Context, userID, accountID string) bool {
	return true
}

func (s *securityService) hasAccountantAccess(ctx context.Context, userID, accountID string) bool {
	return false
}

func (s *securityService) hasAuditorAccess(ctx context.Context, userID, accountID string) bool {
	return false
}

func initializeChallengeThresholds() map[string]*big.Float {
	thresholds := make(map[string]*big.Float)
	
	thresholds["USD"] = big.NewFloat(1000.00)
	thresholds["EUR"] = big.NewFloat(850.00)
	thresholds["GBP"] = big.NewFloat(730.00)
	thresholds["JPY"] = big.NewFloat(110000.00)
	thresholds["CHF"] = big.NewFloat(920.00)
	thresholds["CAD"] = big.NewFloat(1250.00)
	thresholds["AUD"] = big.NewFloat(1350.00)
	
	return thresholds
}

func initializeAccessRules() map[string][]string {
	rules := make(map[string][]string)
	
	rules["account"] = []string{"read", "write", "delete"}
	rules["transaction"] = []string{"read", "create"}
	rules["balance"] = []string{"read"}
	rules["customer"] = []string{"read", "update"}
	rules["admin"] = []string{"*"}
	
	return rules
}
