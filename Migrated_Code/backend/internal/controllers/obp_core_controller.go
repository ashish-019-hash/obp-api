package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"obp-api-backend/internal/services"
	"obp-api-backend/internal/models"
)

type OBPCoreController struct {
	bankService        services.BankService
	accountService     services.AccountService
	transactionService services.TransactionService
	customerService    services.CustomerService
	agentService       services.AgentService
	consentService     services.ConsentService
	balanceService     services.BalanceService
	limitService       services.LimitService
	feeService         services.FeeService
	securityService    services.SecurityService
	validationService  services.ValidationService
	currencyService    services.CurrencyService
	analyticsService   services.AnalyticsService
	rateLimitingService services.RateLimitingService
}

func NewOBPCoreController(
	bankService services.BankService,
	accountService services.AccountService,
	transactionService services.TransactionService,
	customerService services.CustomerService,
	agentService services.AgentService,
	consentService services.ConsentService,
	balanceService services.BalanceService,
	limitService services.LimitService,
	feeService services.FeeService,
	securityService services.SecurityService,
	validationService services.ValidationService,
	currencyService services.CurrencyService,
	analyticsService services.AnalyticsService,
	rateLimitingService services.RateLimitingService,
) *OBPCoreController {
	return &OBPCoreController{
		bankService:         bankService,
		accountService:      accountService,
		transactionService:  transactionService,
		customerService:     customerService,
		agentService:        agentService,
		consentService:      consentService,
		balanceService:      balanceService,
		limitService:        limitService,
		feeService:          feeService,
		securityService:     securityService,
		validationService:   validationService,
		currencyService:     currencyService,
		analyticsService:    analyticsService,
		rateLimitingService: rateLimitingService,
	}
}

func (c *OBPCoreController) GetAPIInfo(ctx *gin.Context) {
	apiInfo := map[string]interface{}{
		"version":        "v5.1.0",
		"version_status": "STABLE",
		"git_commit":     "abc123def456",
		"connector":      "mapped",
	}
	ctx.JSON(http.StatusOK, apiInfo)
}

func (c *OBPCoreController) GetSuggestedSessionTimeout(ctx *gin.Context) {
	timeout := map[string]interface{}{
		"suggested_session_timeout_in_seconds": 3600,
	}
	ctx.JSON(http.StatusOK, timeout)
}

func (c *OBPCoreController) GetBanks(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "50"))
	
	banks, err := c.bankService.GetBanks(ctx.Request.Context(), limit, offset)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{"banks": banks})
}

func (c *OBPCoreController) GetBank(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	
	bank, err := c.bankService.GetBankByID(ctx.Request.Context(), bankID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Bank not found"})
		return
	}
	
	ctx.JSON(http.StatusOK, bank)
}

func (c *OBPCoreController) CreateAgent(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	
	var agent models.Agent
	if err := ctx.ShouldBindJSON(&agent); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	agent.BankId = bankID
	
	if err := c.agentService.CreateAgent(ctx.Request.Context(), &agent); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusCreated, agent)
}

func (c *OBPCoreController) GetAgents(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	
	agents, err := c.agentService.GetAgentsByBankID(ctx.Request.Context(), bankID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{"agents": agents})
}

func (c *OBPCoreController) GetAgent(ctx *gin.Context) {
	agentID := ctx.Param("agentId")
	
	agent, err := c.agentService.GetAgentByID(ctx.Request.Context(), agentID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Agent not found"})
		return
	}
	
	ctx.JSON(http.StatusOK, agent)
}

func (c *OBPCoreController) UpdateAgentStatus(ctx *gin.Context) {
	agentID := ctx.Param("agentId")
	
	var updateData map[string]interface{}
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if err := c.agentService.UpdateAgentStatus(ctx.Request.Context(), agentID, updateData); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{"message": "Agent status updated successfully"})
}

func (c *OBPCoreController) GetAccountsAtBank(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "50"))
	
	accounts, err := c.accountService.GetAccountsAtBank(ctx.Request.Context(), bankID, limit, offset)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{"accounts": accounts})
}

func (c *OBPCoreController) CreateAccount(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	
	var account models.BankAccount
	if err := ctx.ShouldBindJSON(&account); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	account.BankId = bankID
	
	if err := c.accountService.CreateAccount(ctx.Request.Context(), &account); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusCreated, account)
}

func (c *OBPCoreController) GetAccount(ctx *gin.Context) {
	accountID := ctx.Param("accountId")
	
	account, err := c.accountService.GetAccountByID(ctx.Request.Context(), accountID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}
	
	ctx.JSON(http.StatusOK, account)
}

func (c *OBPCoreController) GetAccountBalances(ctx *gin.Context) {
	accountID := ctx.Param("accountId")
	
	currentBalance, err := c.balanceService.CalculateCurrentBalance(ctx.Request.Context(), accountID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	availableBalance, err := c.balanceService.CalculateAvailableBalance(ctx.Request.Context(), accountID, nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	balances := []map[string]interface{}{
		{
			"id":     "balance-current",
			"label":  "Current Balance",
			"amount": map[string]interface{}{
				"currency": "USD",
				"amount":   currentBalance.String(),
			},
		},
		{
			"id":     "balance-available",
			"label":  "Available Balance",
			"amount": map[string]interface{}{
				"currency": "USD",
				"amount":   availableBalance.String(),
			},
		},
	}
	
	ctx.JSON(http.StatusOK, gin.H{"balances": balances})
}

func (c *OBPCoreController) GetTransactions(ctx *gin.Context) {
	accountID := ctx.Param("accountId")
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "50"))
	
	transactions, err := c.transactionService.GetTransactionsByAccountID(ctx.Request.Context(), accountID, limit, offset)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{"transactions": transactions})
}

func (c *OBPCoreController) GetTransaction(ctx *gin.Context) {
	transactionID := ctx.Param("transactionId")
	
	transaction, err := c.transactionService.GetTransactionByID(ctx.Request.Context(), transactionID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}
	
	ctx.JSON(http.StatusOK, transaction)
}

func (c *OBPCoreController) CreateSEPATransactionRequest(ctx *gin.Context) {
	accountID := ctx.Param("accountId")
	
	var request map[string]interface{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	transactionRequest := &models.Transaction{
		ThisAccount:     accountID,
		TransactionType: "SEPA",
	}
	
	if err := c.transactionService.ProcessTransaction(ctx.Request.Context(), transactionRequest); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusCreated, gin.H{
		"id":     "tr-123",
		"type":   "SEPA",
		"status": "INITIATED",
	})
}

func (c *OBPCoreController) GetCustomers(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "50"))
	
	customers, err := c.customerService.GetCustomersByBankID(ctx.Request.Context(), bankID, limit, offset)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{"customers": customers})
}

func (c *OBPCoreController) CreateCustomer(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	
	var customer models.Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	customer.BankId = bankID
	
	if err := c.customerService.CreateCustomer(ctx.Request.Context(), &customer); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusCreated, customer)
}

func (c *OBPCoreController) GetMyConsentsAtBank(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	
	consents, err := c.consentService.GetConsentsByBankID(ctx.Request.Context(), bankID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{"consents": consents})
}

func (c *OBPCoreController) CreateConsent(ctx *gin.Context) {
	var consent models.Consent
	if err := ctx.ShouldBindJSON(&consent); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if err := c.consentService.CreateConsent(ctx.Request.Context(), &consent); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusCreated, consent)
}
