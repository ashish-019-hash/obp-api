package main

import (
	"log"
	"obp-api-backend/internal/config"
	"obp-api-backend/internal/controllers"
	"obp-api-backend/internal/middleware"
	"obp-api-backend/internal/repositories"
	"obp-api-backend/internal/routes"
	"obp-api-backend/internal/services"
)

func main() {
	cfg := config.Load()

	authUserRepo := repositories.NewAuthUserRepository()
	resourceUserRepo := repositories.NewResourceUserRepository()
	consumerRepo := repositories.NewConsumerRepository()
	entitlementRepo := repositories.NewEntitlementRepository()
	badLoginAttemptRepo := repositories.NewBadLoginAttemptRepository()
	userLockRepo := repositories.NewUserLockRepository()

	bankRepo := repositories.NewBankRepository()
	accountRepo := repositories.NewBankAccountRepository()
	transactionRepo := repositories.NewTransactionRepository()
	customerRepo := repositories.NewCustomerRepository()
	agentRepo := repositories.NewAgentRepository()
	consentRepo := repositories.NewConsentRepository()
	counterpartyLimitRepo := repositories.NewCounterpartyLimitRepository()
	fxRateRepo := repositories.NewFXRateRepository()
	metricsRepo := repositories.NewMetricsRepository()
	rateLimitRepo := repositories.NewRateLimitRepository()

	authService := services.NewAuthService(
		authUserRepo,
		resourceUserRepo,
		consumerRepo,
		entitlementRepo,
		badLoginAttemptRepo,
		userLockRepo,
		cfg,
	)

	currencyService := services.NewCurrencyService(fxRateRepo)
	limitService := services.NewLimitService(counterpartyLimitRepo, nil, currencyService)
	feeService := services.NewFeeService(currencyService)
	securityService := services.NewSecurityService(currencyService)
	validationService := services.NewValidationService(currencyService)
	balanceService := services.NewBalanceService(transactionRepo)
	rateLimitingService := services.NewRateLimitingService(rateLimitRepo)
	analyticsService := services.NewAnalyticsService(customerRepo, metricsRepo, currencyService)

	transactionService := services.NewTransactionService(
		transactionRepo,
		balanceService,
		limitService,
		securityService,
		validationService,
		currencyService,
	)

	bankService := services.NewBankService(bankRepo)
	accountService := services.NewAccountService(accountRepo, transactionRepo)
	customerService := services.NewCustomerService(customerRepo)
	agentService := services.NewAgentService(agentRepo)
	consentService := services.NewConsentService(consentRepo)
	paymentService := services.NewPaymentService(transactionService)

	authController := controllers.NewAuthController(authService)
	authMiddleware := middleware.AuthMiddleware(authService)

	obpController := controllers.NewOBPCoreController(
		bankService,
		accountService,
		transactionService,
		customerService,
		agentService,
		consentService,
		balanceService,
		limitService,
		feeService,
		securityService,
		validationService,
		currencyService,
		analyticsService,
		rateLimitingService,
	)

	berlinGroupController := controllers.NewBerlinGroupController(
		accountService,
		balanceService,
		paymentService,
	)

	ukOpenBankingController := controllers.NewUKOpenBankingController(
		accountService,
		balanceService,
		paymentService,
	)

	australianCDRController := controllers.NewAustralianCDRController(
		accountService,
		balanceService,
	)

	bahrainOBFController := controllers.NewBahrainOBFController(
		accountService,
		balanceService,
		paymentService,
	)

	polishAPIController := controllers.NewPolishAPIController(
		accountService,
		balanceService,
		paymentService,
	)

	stetAPIController := controllers.NewSTETAPIController(
		accountService,
		balanceService,
		paymentService,
	)

	mxofAPIController := controllers.NewMxOFAPIController(
		accountService,
		balanceService,
		paymentService,
	)

	additionalController := controllers.NewAdditionalRegulatoryController(
		accountService,
		balanceService,
		paymentService,
	)

	obpV3Controller := controllers.NewOBPv3Controller(
		bankService,
		accountService,
		customerService,
	)

	obpV4Controller := controllers.NewOBPv4Controller(
		bankService,
		accountService,
		transactionService,
		customerService,
	)

	router := routes.SetupRoutes(
		authController,
		authMiddleware,
		obpController,
		obpV3Controller,
		obpV4Controller,
		berlinGroupController,
		ukOpenBankingController,
		australianCDRController,
		bahrainOBFController,
		polishAPIController,
		stetAPIController,
		mxofAPIController,
		additionalController,
	)

	log.Println("Starting OBP API Backend server on :" + cfg.Port)
	log.Println("Authentication Configuration:")
	log.Printf("- Direct Login: %v", cfg.Auth.AllowDirectLogin)
	log.Printf("- OAuth 1.0a: %v", cfg.Auth.AllowOAuth1)
	log.Printf("- OAuth 2.0: %v", cfg.Auth.AllowOAuth2)
	log.Printf("- Gateway Login: %v", cfg.Auth.AllowGatewayLogin)
	log.Printf("- DAuth: %v", cfg.Auth.AllowDAuth)
	log.Println("\nAvailable endpoints:")
	log.Println("- Auth: POST /my/logins/direct (Direct Login)")
	log.Println("- Auth: POST /auth/users (Create User)")
	log.Println("- Auth: POST /auth/consumers (Create Consumer)")
	log.Println("- Auth: GET /auth/users/current (Get Current User)")
	log.Println("- OBP Core API v3.1.0: /obp/v3.1.0/* (~200+ endpoints)")
	log.Println("- OBP Core API v4.0.0: /obp/v4.0.0/* (~150+ endpoints)")
	log.Println("- OBP Core API v5.1.0: /obp/v5.1.0/* (~200+ endpoints)")
	log.Println("- Berlin Group PSD2 v1.3: /berlin-group/v1.3/* (~20+ endpoints)")
	log.Println("- UK Open Banking v3.1.0: /open-banking/v3.1.0/* (~60+ endpoints)")
	log.Println("- Australian CDR v1.0.0: /cds-au/v1.0.0/* (~20+ endpoints)")
	log.Println("- Bahrain OBF v1.0.0: /bahrain-obf/v1.0.0/* (~80+ endpoints)")
	log.Println("- Polish API v2.1.1.1: /polish-api/v2.1.1.1/* (~10+ endpoints)")
	log.Println("- STET API v1.4: /stet/v1.4/* (~10+ endpoints)")
	log.Println("- MxOF API v1.0.0: /mxof/v1.0.0/* (~10+ endpoints)")
	log.Printf("Total endpoints: ~625+ (+ 4 auth endpoints)")

	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
