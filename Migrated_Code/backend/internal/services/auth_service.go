package services

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"obp-api-backend/internal/config"
	"obp-api-backend/internal/models"
	"obp-api-backend/internal/repositories"
	"obp-api-backend/internal/utils"
)

type AuthService interface {
	DirectLogin(ctx context.Context, username, password, consumerKey string) (string, error)
	ValidateDirectLoginToken(ctx context.Context, token string) (*models.ResourceUser, *models.Consumer, error)
	CreateUser(ctx context.Context, username, email, password, firstName, lastName string) (*models.AuthUser, *models.ResourceUser, error)
	CreateConsumer(ctx context.Context, name, appType, description, developerEmail string, createdByUserID *string) (*models.Consumer, error)
	GetUserByID(ctx context.Context, userID string) (*models.ResourceUser, error)
}

type authService struct {
	authUserRepo        repositories.AuthUserRepository
	resourceUserRepo    repositories.ResourceUserRepository
	consumerRepo        repositories.ConsumerRepository
	entitlementRepo     repositories.EntitlementRepository
	badLoginAttemptRepo repositories.BadLoginAttemptRepository
	userLockRepo        repositories.UserLockRepository
	config              *config.Config
}

func NewAuthService(
	authUserRepo repositories.AuthUserRepository,
	resourceUserRepo repositories.ResourceUserRepository,
	consumerRepo repositories.ConsumerRepository,
	entitlementRepo repositories.EntitlementRepository,
	badLoginAttemptRepo repositories.BadLoginAttemptRepository,
	userLockRepo repositories.UserLockRepository,
	cfg *config.Config,
) AuthService {
	return &authService{
		authUserRepo:        authUserRepo,
		resourceUserRepo:    resourceUserRepo,
		consumerRepo:        consumerRepo,
		entitlementRepo:     entitlementRepo,
		badLoginAttemptRepo: badLoginAttemptRepo,
		userLockRepo:        userLockRepo,
		config:              cfg,
	}
}

func (s *authService) DirectLogin(ctx context.Context, username, password, consumerKey string) (string, error) {
	if !s.config.Auth.AllowDirectLogin {
		return "", errors.New("direct login is not enabled")
	}

	consumer, err := s.consumerRepo.GetByConsumerKey(ctx, consumerKey)
	if err != nil || consumer == nil {
		return "", errors.New("invalid consumer")
	}

	if !consumer.IsActive {
		return "", errors.New("consumer is not active")
	}

	authUser, err := s.authUserRepo.GetByUsername(ctx, username)
	if err != nil || authUser == nil {
		s.badLoginAttemptRepo.IncrementAttempts(ctx, username, "local")
		return "", errors.New("invalid username or password")
	}

	resourceUser, err := s.resourceUserRepo.GetByID(ctx, authUser.UserID)
	if err != nil || resourceUser == nil {
		return "", errors.New("user not found")
	}

	isLocked, err := s.userLockRepo.IsLocked(ctx, resourceUser.UserID)
	if err == nil && isLocked {
		return "", errors.New("account is locked")
	}

	if resourceUser.IsDeleted {
		return "", errors.New("account is deleted")
	}

	badAttempt, err := s.badLoginAttemptRepo.GetByUsernameProvider(ctx, username, "local")
	if err == nil && badAttempt != nil {
		if badAttempt.BadAttemptsSinceLastSuccessOrReset >= s.config.Auth.MaxBadLoginAttempts {
			lock := &models.UserLock{
				UserID:       resourceUser.UserID,
				TypeOfLock:   "SECURITY",
				LastLockDate: time.Now(),
			}
			s.userLockRepo.Create(ctx, lock)
			return "", errors.New("account locked due to too many failed login attempts")
		}
	}

	if !utils.CheckPasswordHash(password, authUser.PasswordHash) {
		s.badLoginAttemptRepo.IncrementAttempts(ctx, username, "local")
		return "", errors.New("invalid username or password")
	}

	s.badLoginAttemptRepo.ResetAttempts(ctx, username, "local")

	token, err := utils.GenerateDirectLoginToken(
		resourceUser.UserID,
		consumer.ConsumerID,
		resourceUser.Provider,
		resourceUser.ProviderID,
		s.config.Auth.DirectLoginSecret,
		s.config.Auth.DirectLoginTokenExpiration,
	)

	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}

func (s *authService) ValidateDirectLoginToken(ctx context.Context, tokenString string) (*models.ResourceUser, *models.Consumer, error) {
	claims, err := utils.ValidateDirectLoginToken(tokenString, s.config.Auth.DirectLoginSecret)
	if err != nil {
		return nil, nil, errors.New("invalid token")
	}

	resourceUser, err := s.resourceUserRepo.GetByID(ctx, claims.UserID)
	if err != nil || resourceUser == nil {
		return nil, nil, errors.New("user not found")
	}

	if resourceUser.IsDeleted {
		return nil, nil, errors.New("user account is deleted")
	}

	consumer, err := s.consumerRepo.GetByID(ctx, claims.ConsumerID)
	if err != nil || consumer == nil {
		return nil, nil, errors.New("consumer not found")
	}

	if !consumer.IsActive {
		return nil, nil, errors.New("consumer is not active")
	}

	return resourceUser, consumer, nil
}

func (s *authService) CreateUser(ctx context.Context, username, email, password, firstName, lastName string) (*models.AuthUser, *models.ResourceUser, error) {
	existingAuthUser, _ := s.authUserRepo.GetByUsername(ctx, username)
	if existingAuthUser != nil {
		return nil, nil, errors.New("username already exists")
	}

	existingAuthUser, _ = s.authUserRepo.GetByEmail(ctx, email)
	if existingAuthUser != nil {
		return nil, nil, errors.New("email already exists")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, nil, errors.New("failed to hash password")
	}

	userID := uuid.New().String()
	now := time.Now()

	authUser := &models.AuthUser{
		UserID:        userID,
		Username:      username,
		Email:         email,
		PasswordHash:  hashedPassword,
		EmailVerified: false,
		Validated:     false,
		FirstName:     firstName,
		LastName:      lastName,
		Provider:      "local",
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	err = s.authUserRepo.Create(ctx, authUser)
	if err != nil {
		return nil, nil, err
	}

	resourceUser := &models.ResourceUser{
		UserID:     userID,
		AuthUserID: &userID,
		Provider:   "local",
		ProviderID: username,
		Name:       firstName + " " + lastName,
		Email:      email,
		IsDeleted:  false,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	err = s.resourceUserRepo.Create(ctx, resourceUser)
	if err != nil {
		return nil, nil, err
	}

	return authUser, resourceUser, nil
}

func (s *authService) CreateConsumer(ctx context.Context, name, appType, description, developerEmail string, createdByUserID *string) (*models.Consumer, error) {
	consumerKey, err := utils.GenerateConsumerKey()
	if err != nil {
		return nil, errors.New("failed to generate consumer key")
	}

	consumerSecret, err := utils.GenerateConsumerSecret()
	if err != nil {
		return nil, errors.New("failed to generate consumer secret")
	}

	hashedSecret, err := utils.HashPassword(consumerSecret)
	if err != nil {
		return nil, errors.New("failed to hash consumer secret")
	}

	consumerID := uuid.New().String()
	now := time.Now()

	consumer := &models.Consumer{
		ConsumerID:         consumerID,
		ConsumerKey:        consumerKey,
		ConsumerSecret:     hashedSecret,
		IsActive:           true,
		Name:               name,
		AppType:            appType,
		Description:        description,
		DeveloperEmail:     developerEmail,
		CreatedByUserID:    createdByUserID,
		PerSecondCallLimit: 10,
		PerMinuteCallLimit: 100,
		PerHourCallLimit:   1000,
		PerDayCallLimit:    10000,
		PerWeekCallLimit:   50000,
		PerMonthCallLimit:  200000,
		CreatedAt:          now,
		UpdatedAt:          now,
	}

	err = s.consumerRepo.Create(ctx, consumer)
	if err != nil {
		return nil, err
	}

	consumer.ConsumerSecret = consumerSecret

	return consumer, nil
}

func (s *authService) GetUserByID(ctx context.Context, userID string) (*models.ResourceUser, error) {
	return s.resourceUserRepo.GetByID(ctx, userID)
}
