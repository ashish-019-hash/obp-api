package services

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"obp-api-backend/internal/models"
	"obp-api-backend/internal/repositories"
)

type AnalyticsService interface {
	CalculateCreditRating(ctx context.Context, customerID string) (*CreditAssessment, error)
	GetTopAPIs(ctx context.Context, fromDate, toDate time.Time, limit int) ([]*models.APIRanking, error)
	GetTopConsumers(ctx context.Context, fromDate, toDate time.Time, limit int) ([]*models.ConsumerRanking, error)
	GetAPIMetrics(ctx context.Context, fromDate, toDate time.Time, consumerID, userID, url string) (*models.APIMetrics, error)
	ProcessStandingOrderAmount(amount *big.Float, currency, frequency string) (*StandingOrderResult, error)
}

type CreditAssessment struct {
	CreditRating string
	CreditLimit  *models.AmountOfMoney
	CreditScore  int
}

type StandingOrderResult struct {
	SmallestUnits int64
	DisplayAmount *big.Float
	OrderID       string
}

type analyticsService struct {
	customerRepo    repositories.CustomerRepository
	metricsRepo     repositories.MetricsRepository
	currencyService CurrencyService
}

func NewAnalyticsService(
	customerRepo repositories.CustomerRepository,
	metricsRepo repositories.MetricsRepository,
	currencyService CurrencyService,
) AnalyticsService {
	return &analyticsService{
		customerRepo:    customerRepo,
		metricsRepo:     metricsRepo,
		currencyService: currencyService,
	}
}

func (s *analyticsService) CalculateCreditRating(ctx context.Context, customerID string) (*CreditAssessment, error) {
	customer, err := s.customerRepo.GetByID(ctx, customerID)
	if err != nil {
		return nil, err
	}

	creditScore := s.calculateBaseScore(customer.EmploymentStatus, customer.HighestEducationAttained)
	creditScore = s.adjustForRelationship(creditScore, customer.RelationshipStatus, customer.Dependents)

	assessment := &CreditAssessment{
		CreditScore: creditScore,
	}

	if creditScore >= 750 {
		assessment.CreditRating = "EXCELLENT"
		assessment.CreditLimit = s.calculateCreditLimit(creditScore, "HIGH", "USD")
	} else if creditScore >= 650 {
		assessment.CreditRating = "GOOD"
		assessment.CreditLimit = s.calculateCreditLimit(creditScore, "MEDIUM", "USD")
	} else if creditScore >= 550 {
		assessment.CreditRating = "FAIR"
		assessment.CreditLimit = s.calculateCreditLimit(creditScore, "LOW", "USD")
	} else {
		assessment.CreditRating = "POOR"
		assessment.CreditLimit = s.calculateCreditLimit(creditScore, "MINIMAL", "USD")
	}

	return assessment, nil
}

func (s *analyticsService) GetTopAPIs(ctx context.Context, fromDate, toDate time.Time, limit int) ([]*models.APIRanking, error) {
	return s.metricsRepo.GetTopAPIs(ctx, fromDate, toDate, limit)
}

func (s *analyticsService) GetTopConsumers(ctx context.Context, fromDate, toDate time.Time, limit int) ([]*models.ConsumerRanking, error) {
	return s.metricsRepo.GetTopConsumers(ctx, fromDate, toDate, limit)
}

func (s *analyticsService) GetAPIMetrics(ctx context.Context, fromDate, toDate time.Time, consumerID, userID, url string) (*models.APIMetrics, error) {
	return s.metricsRepo.GetMetrics(ctx, fromDate, toDate, consumerID, userID, url)
}

func (s *analyticsService) ProcessStandingOrderAmount(amount *big.Float, currency, frequency string) (*StandingOrderResult, error) {
	smallestUnits := s.currencyService.ConvertToSmallestUnits(amount, currency)
	displayAmount := s.currencyService.ConvertFromSmallestUnits(smallestUnits, currency)

	orderID := generateOrderID()

	return &StandingOrderResult{
		SmallestUnits: smallestUnits,
		DisplayAmount: displayAmount,
		OrderID:       orderID,
	}, nil
}

func (s *analyticsService) calculateBaseScore(employmentStatus, education string) int {
	baseScore := 500

	switch employmentStatus {
	case "EMPLOYED_FULL_TIME":
		baseScore += 150
	case "EMPLOYED_PART_TIME":
		baseScore += 100
	case "SELF_EMPLOYED":
		baseScore += 120
	case "UNEMPLOYED":
		baseScore += 0
	case "RETIRED":
		baseScore += 80
	case "STUDENT":
		baseScore += 50
	}

	switch education {
	case "DOCTORATE":
		baseScore += 100
	case "MASTERS":
		baseScore += 80
	case "BACHELORS":
		baseScore += 60
	case "ASSOCIATES":
		baseScore += 40
	case "HIGH_SCHOOL":
		baseScore += 20
	case "SOME_HIGH_SCHOOL":
		baseScore += 0
	}

	return baseScore
}

func (s *analyticsService) adjustForRelationship(score int, relationshipStatus string, dependents int) int {
	switch relationshipStatus {
	case "MARRIED":
		score += 50
	case "SINGLE":
		score += 0
	case "DIVORCED":
		score -= 20
	case "WIDOWED":
		score += 10
	}

	if dependents > 0 {
		score -= dependents * 10
	}

	return score
}

func (s *analyticsService) calculateCreditLimit(score int, tier, currency string) *models.AmountOfMoney {
	var baseLimit float64

	switch tier {
	case "HIGH":
		baseLimit = 50000
	case "MEDIUM":
		baseLimit = 25000
	case "LOW":
		baseLimit = 10000
	case "MINIMAL":
		baseLimit = 2000
	default:
		baseLimit = 5000
	}

	multiplier := float64(score) / 750.0
	if multiplier > 2.0 {
		multiplier = 2.0
	}

	finalLimit := baseLimit * multiplier

	return &models.AmountOfMoney{
		Currency: currency,
		Amount:   big.NewFloat(finalLimit).String(),
	}
}

func generateOrderID() string {
	return fmt.Sprintf("SO_%d", time.Now().UnixNano())
}
