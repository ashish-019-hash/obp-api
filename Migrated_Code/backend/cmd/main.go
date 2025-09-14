package main

import (
	"log"

	"github.com/ashish-019-hash/obp-api-backend/internal/config"
	"github.com/ashish-019-hash/obp-api-backend/internal/repositories"
	"github.com/ashish-019-hash/obp-api-backend/internal/routes"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/ashish-019-hash/obp-api-backend/pkg/db"
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

	router := routes.SetupRoutes(orchestrationService)

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
	log.Println("Server endpoints:")
	log.Println("  GET /health - Health check")
	log.Println("  GET /ping - Ping endpoint")
	log.Println("  GET /api/v1/health - API health check")
	log.Println("v5.1.0 API endpoints:")
	log.Println("  GET /root - API info")
	log.Println("  GET /ui/suggested-session-timeout - Session timeout")
	log.Println("  GET /well-known - OAuth2 well-known URIs")
	log.Println("  POST /banks/{bankId}/agents - Create agent")
	log.Println("  GET /my/consents - Get user consents")
	log.Println("  POST /banks/{bankId}/accounts/{accountId}/views/{viewId}/counterparties/{counterpartyId}/limits - Create counterparty limit")
	log.Println("  ... and 40+ more v5.1.0 endpoints")

	if err := router.Run(port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
