package main

import (
	"log"

	"github.com/ashish-019-hash/obp-api-backend/internal/config"
	"github.com/ashish-019-hash/obp-api-backend/internal/controllers"
	"github.com/ashish-019-hash/obp-api-backend/internal/middleware"
	"github.com/ashish-019-hash/obp-api-backend/internal/repositories"
	"github.com/ashish-019-hash/obp-api-backend/internal/routes"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/ashish-019-hash/obp-api-backend/pkg/db"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	if err := db.InitDB(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.CloseDB()

	bankRepo := repositories.NewBankRepository(db.GetDB())
	accountRepo := repositories.NewBankAccountRepository(db.GetDB())
	customerRepo := repositories.NewCustomerRepository(db.GetDB())
	transactionRepo := repositories.NewTransactionRepository(db.GetDB())
	transactionRequestRepo := repositories.NewTransactionRequestRepository(db.GetDB())
	consentRepo := repositories.NewConsentRepository(db.GetDB())

	currencyService := services.NewCurrencyService()
	transactionService := services.NewTransactionService(transactionRepo, accountRepo, currencyService)
	counterpartyService := services.NewCounterpartyLimitService(currencyService)
	securityService := services.NewSecurityService(consentRepo, currencyService)
	orchestrationService := services.NewOrchestrationService(
		currencyService,
		transactionService,
		counterpartyService,
		securityService,
		bankRepo,
		accountRepo,
		customerRepo,
		transactionRepo,
		transactionRequestRepo,
		consentRepo,
	)

	authRepo := repositories.NewAuthRepository(db.GetDB())
	configService := services.NewConfigService(db.GetDB())
	authService := services.NewAuthenticationService(db.GetDB(), authRepo, cfg.JWT.Secret, configService)
	rateLimiter := services.NewRateLimiter(configService, db.GetDB())
	randomService := services.NewSecureRandomService()
	sessionService := services.NewSessionService(db.GetDB(), configService, randomService)
	dauthService := services.NewDAuthService(db.GetDB(), configService, randomService, cfg.JWT.Secret)
	gatewayService := services.NewGatewayLoginService(db.GetDB(), configService, randomService, cfg.JWT.Secret)
	authController := controllers.NewAuthController(authService)
	authMiddleware := middleware.NewAuthMiddleware(authService, rateLimiter, dauthService, gatewayService, cfg.JWT.Secret)

	if err := configService.InitializeDefaultConfigs(); err != nil {
		log.Printf("Warning: Failed to initialize default configurations: %v", err)
	}

	log.Println("Advanced authentication features enabled:")
	log.Println("  - X.509 Certificate Processing with PSD2 role extraction")
	log.Println("  - JWKS-based OAuth 2.0/OIDC token validation")
	log.Println("  - Berlin Group PSD2 consent management")
	log.Println("  - Multi-factor authentication (TOTP, SMS, Backup codes)")
	log.Println("  - Certificate-based client authentication")
	log.Println("  - Hydra ORY integration support")
	log.Println("  - Advanced security features and compliance")

	router := gin.Default()

	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.LoggerMiddleware())

	routes.SetupRoutes(router, orchestrationService, authMiddleware)
	routes.SetupAuthRoutes(router, authController, authMiddleware)

	if err := services.SeedAuthenticationData(db.GetDB(), authRepo); err != nil {
		log.Printf("Warning: Failed to seed authentication data: %v", err)
	}

	sessionService.StartSessionCleanupRoutine()

	port := ":" + cfg.Port
	log.Printf("Starting OBP-API Backend Server on port %s", cfg.Port)
	log.Println("Based on Open Bank Project API analysis from phase-01-output")
	log.Println("Database: SQLite in-memory with GORM")
	log.Println("Services initialized:")
	log.Println("  - Currency Service (14 supported currencies)")
	log.Println("  - Transaction Service (credit/debit classification)")
	log.Println("  - Counterparty Limit Service (6-dimensional validation)")
	log.Println("  - Security Service (challenge thresholds)")
	log.Println("  - Orchestration Service (workflow coordination)")
	log.Println("  - Authentication Service (JWT, OAuth, DirectLogin)")
	log.Println("Authentication endpoints:")
	log.Println("  POST /auth/direct-login - DirectLogin authentication")
	log.Println("  POST /auth/consumers - Consumer registration")
	log.Println("  POST /auth/users - User registration")
	log.Println("  POST /oauth/initiate - OAuth request token")
	log.Println("  POST /oauth/token - OAuth access token")
	log.Println("  GET /oauth/authorize - OAuth authorization")
	log.Println("  GET /my/user - Current user info (protected)")
	log.Println("Server endpoints:")
	log.Println("  GET /health - Health check")
	log.Println("  GET /ping - Ping endpoint")
	log.Println("  GET /api/v1/health - API health check")
	log.Println("v5.1.0 API endpoints (protected):")
	log.Println("  GET /obp/v5.1.0/root - API info (public)")
	log.Println("  GET /obp/v5.1.0/well-known - OAuth2 well-known URIs (public)")
	log.Println("  GET /obp/v5.1.0/banks - Get banks (protected)")
	log.Println("  POST /obp/v5.1.0/banks - Create bank (protected)")
	log.Println("  GET /obp/v5.1.0/my/consents - Get user consents (protected)")
	log.Println("  ... and 40+ more protected v5.1.0 endpoints")

	if err := router.Run(port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
