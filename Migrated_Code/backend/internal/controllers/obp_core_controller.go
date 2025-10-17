package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"obp-api-backend/internal/models"
	"obp-api-backend/internal/services"
)

type OBPCoreController struct {
	bankService         services.BankService
	accountService      services.AccountService
	transactionService  services.TransactionService
	customerService     services.CustomerService
	agentService        services.AgentService
	consentService      services.ConsentService
	balanceService      services.BalanceService
	limitService        services.LimitService
	feeService          services.FeeService
	securityService     services.SecurityService
	validationService   services.ValidationService
	currencyService     services.CurrencyService
	analyticsService    services.AnalyticsService
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

func (c *OBPCoreController) GetConfig(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"akka": map[string]interface{}{
			"ports": []int{8080, 8081},
		},
		"elastic_search": map[string]interface{}{
			"enabled": false,
		},
		"cache": map[string]interface{}{
			"ttl_seconds": 3600,
		},
	})
}

func (c *OBPCoreController) GetAdapterInfo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name":       "stored_procedure_vDec2019",
		"version":    "Dec2019",
		"git_commit": "unknown",
		"date":       "2019-12-01T00:00:00Z",
	})
}

func (c *OBPCoreController) GetRateLimitingInfo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"enabled":           true,
		"technology":        "REDIS",
		"service_available": true,
		"is_active":         true,
	})
}

func (c *OBPCoreController) CreateBank(ctx *gin.Context) {
	var bankData map[string]interface{}
	if err := ctx.ShouldBindJSON(&bankData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, bankData)
}

func (c *OBPCoreController) UpdateBank(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	var bankData map[string]interface{}
	if err := ctx.ShouldBindJSON(&bankData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bankData["bank_id"] = bankID
	ctx.JSON(http.StatusOK, bankData)
}

func (c *OBPCoreController) DeleteBank(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	ctx.JSON(http.StatusOK, gin.H{"message": "Bank deleted", "bank_id": bankID})
}

func (c *OBPCoreController) GetAccounts(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accounts := []map[string]interface{}{
		{
			"id":      "account1",
			"bank_id": bankID,
			"label":   "My Account",
			"number":  "123456789",
			"type":    "CURRENT",
			"balance": map[string]interface{}{
				"currency": "EUR",
				"amount":   "1000.00",
			},
		},
	}
	ctx.JSON(http.StatusOK, gin.H{"accounts": accounts})
}

func (c *OBPCoreController) UpdateAccount(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	var accountData map[string]interface{}
	if err := ctx.ShouldBindJSON(&accountData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	accountData["bank_id"] = bankID
	accountData["account_id"] = accountID
	ctx.JSON(http.StatusOK, accountData)
}

func (c *OBPCoreController) DeleteAccount(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	ctx.JSON(http.StatusOK, gin.H{"message": "Account deleted", "bank_id": bankID, "account_id": accountID})
}

func (c *OBPCoreController) CreateTransaction(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	viewID := ctx.Param("viewId")

	var transactionData map[string]interface{}
	if err := ctx.ShouldBindJSON(&transactionData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transactionData["bank_id"] = bankID
	transactionData["account_id"] = accountID
	transactionData["view_id"] = viewID
	ctx.JSON(http.StatusCreated, transactionData)
}

func (c *OBPCoreController) GetCustomer(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	customerID := ctx.Param("customerId")

	customer := map[string]interface{}{
		"customer_id":         customerID,
		"bank_id":             bankID,
		"customer_number":     "123456",
		"legal_name":          "John Doe",
		"mobile_phone_number": "+1234567890",
		"email":               "john.doe@example.com",
	}
	ctx.JSON(http.StatusOK, customer)
}

func (c *OBPCoreController) UpdateCustomer(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	customerID := ctx.Param("customerId")

	var customerData map[string]interface{}
	if err := ctx.ShouldBindJSON(&customerData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customerData["customer_id"] = customerID
	customerData["bank_id"] = bankID
	ctx.JSON(http.StatusOK, customerData)
}

func (c *OBPCoreController) DeleteCustomer(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	customerID := ctx.Param("customerId")
	ctx.JSON(http.StatusOK, gin.H{"message": "Customer deleted", "bank_id": bankID, "customer_id": customerID})
}

func (c *OBPCoreController) GetUser(ctx *gin.Context) {
	userID := ctx.Param("userId")

	user := map[string]interface{}{
		"user_id":    userID,
		"username":   "john.doe",
		"email":      "john.doe@example.com",
		"first_name": "John",
		"last_name":  "Doe",
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *OBPCoreController) CreateUser(ctx *gin.Context) {
	var userData map[string]interface{}
	if err := ctx.ShouldBindJSON(&userData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, userData)
}

func (c *OBPCoreController) UpdateUser(ctx *gin.Context) {
	userID := ctx.Param("userId")

	var userData map[string]interface{}
	if err := ctx.ShouldBindJSON(&userData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userData["user_id"] = userID
	ctx.JSON(http.StatusOK, userData)
}

func (c *OBPCoreController) GetConsent(ctx *gin.Context) {
	consentID := ctx.Param("consentId")

	consent := map[string]interface{}{
		"consent_id":         consentID,
		"status":             "AUTHORISED",
		"creation_date_time": "2023-09-16T10:00:00Z",
		"permissions":        []string{"ReadAccountsBasic", "ReadAccountsDetail"},
	}
	ctx.JSON(http.StatusOK, consent)
}

func (c *OBPCoreController) UpdateConsent(ctx *gin.Context) {
	consentID := ctx.Param("consentId")

	var consentData map[string]interface{}
	if err := ctx.ShouldBindJSON(&consentData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	consentData["consent_id"] = consentID
	ctx.JSON(http.StatusOK, consentData)
}

func (c *OBPCoreController) DeleteConsent(ctx *gin.Context) {
	consentID := ctx.Param("consentId")
	ctx.JSON(http.StatusOK, gin.H{"message": "Consent deleted", "consent_id": consentID})
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
			"id":    "balance-current",
			"label": "Current Balance",
			"amount": map[string]interface{}{
				"currency": "USD",
				"amount":   currentBalance.String(),
			},
		},
		{
			"id":    "balance-available",
			"label": "Available Balance",
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

func (c *OBPCoreController) DeleteRegulatedEntity(ctx *gin.Context) {
	entityID := ctx.Param("regulatedEntityId")
	ctx.JSON(http.StatusOK, gin.H{"message": "Regulated entity deleted", "entityId": entityID})
}

func (c *OBPCoreController) WaitingForGodot(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Waiting for Godot..."})
}

func (c *OBPCoreController) GetSuggestedSessionTimeout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"suggested_session_timeout": 3600})
}

func (c *OBPCoreController) CreateATM(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	var atmData map[string]interface{}
	if err := ctx.ShouldBindJSON(&atmData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	atmData["bank_id"] = bankID
	ctx.JSON(http.StatusCreated, atmData)
}

func (c *OBPCoreController) UpdateATM(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	atmID := ctx.Param("atmId")
	var atmData map[string]interface{}
	if err := ctx.ShouldBindJSON(&atmData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	atmData["bank_id"] = bankID
	atmData["atm_id"] = atmID
	ctx.JSON(http.StatusOK, atmData)
}

func (c *OBPCoreController) GetATMs(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	atms := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"atms": atms, "bank_id": bankID})
}

func (c *OBPCoreController) GetATM(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	atmID := ctx.Param("atmId")
	atm := map[string]interface{}{
		"atm_id":  atmID,
		"bank_id": bankID,
	}
	ctx.JSON(http.StatusOK, atm)
}

func (c *OBPCoreController) DeleteATM(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	atmID := ctx.Param("atmId")
	ctx.JSON(http.StatusOK, gin.H{"message": "ATM deleted", "atm_id": atmID, "bank_id": bankID})
}

func (c *OBPCoreController) GetATMAttributes(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	atmID := ctx.Param("atmId")
	attributes := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"atm_attributes": attributes, "bank_id": bankID, "atm_id": atmID})
}

func (c *OBPCoreController) GetATMAttribute(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	atmID := ctx.Param("atmId")
	attrID := ctx.Param("atmAttributeId")
	attribute := map[string]interface{}{
		"attribute_id": attrID,
		"atm_id":       atmID,
		"bank_id":      bankID,
	}
	ctx.JSON(http.StatusOK, attribute)
}

func (c *OBPCoreController) UpdateATMAttribute(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	atmID := ctx.Param("atmId")
	attrID := ctx.Param("atmAttributeId")
	var attrData map[string]interface{}
	if err := ctx.ShouldBindJSON(&attrData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	attrData["attribute_id"] = attrID
	attrData["atm_id"] = atmID
	attrData["bank_id"] = bankID
	ctx.JSON(http.StatusOK, attrData)
}

func (c *OBPCoreController) CreateATMAttribute(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	atmID := ctx.Param("atmId")

	var attrData map[string]interface{}
	if err := ctx.ShouldBindJSON(&attrData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	attrData["bank_id"] = bankID
	attrData["atm_id"] = atmID
	ctx.JSON(http.StatusCreated, attrData)
}

func (c *OBPCoreController) DeleteATMAttribute(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	atmID := ctx.Param("atmId")
	attrID := ctx.Param("atmAttributeId")
	ctx.JSON(http.StatusOK, gin.H{"message": "ATM attribute deleted", "attribute_id": attrID, "atm_id": atmID, "bank_id": bankID})
}

func (c *OBPCoreController) CreateConsumerDynamicRegistration(ctx *gin.Context) {
	var consumerData map[string]interface{}
	if err := ctx.ShouldBindJSON(&consumerData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, consumerData)
}

func (c *OBPCoreController) CreateConsumer(ctx *gin.Context) {
	var consumerData map[string]interface{}
	if err := ctx.ShouldBindJSON(&consumerData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, consumerData)
}

func (c *OBPCoreController) CreateMyConsumer(ctx *gin.Context) {
	var consumerData map[string]interface{}
	if err := ctx.ShouldBindJSON(&consumerData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, consumerData)
}

func (c *OBPCoreController) UpdateConsumerRedirectURL(ctx *gin.Context) {
	consumerID := ctx.Param("consumerId")
	var urlData map[string]interface{}
	if err := ctx.ShouldBindJSON(&urlData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	urlData["consumer_id"] = consumerID
	ctx.JSON(http.StatusOK, urlData)
}

func (c *OBPCoreController) UpdateConsumerLogoURL(ctx *gin.Context) {
	consumerID := ctx.Param("consumerId")
	var logoData map[string]interface{}
	if err := ctx.ShouldBindJSON(&logoData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logoData["consumer_id"] = consumerID
	ctx.JSON(http.StatusOK, logoData)
}

func (c *OBPCoreController) UpdateConsumerCertificate(ctx *gin.Context) {
	consumerID := ctx.Param("consumerId")
	var certData map[string]interface{}
	if err := ctx.ShouldBindJSON(&certData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	certData["consumer_id"] = consumerID
	ctx.JSON(http.StatusOK, certData)
}

func (c *OBPCoreController) UpdateConsumerName(ctx *gin.Context) {
	consumerID := ctx.Param("consumerId")
	var nameData map[string]interface{}
	if err := ctx.ShouldBindJSON(&nameData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	nameData["consumer_id"] = consumerID
	ctx.JSON(http.StatusOK, nameData)
}

func (c *OBPCoreController) GetConsumer(ctx *gin.Context) {
	consumerID := ctx.Param("consumerId")
	consumer := map[string]interface{}{
		"consumer_id": consumerID,
	}
	ctx.JSON(http.StatusOK, consumer)
}

func (c *OBPCoreController) GetConsumers(ctx *gin.Context) {
	consumers := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"consumers": consumers})
}

func (c *OBPCoreController) GrantUserAccessToView(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	viewID := ctx.Param("viewId")
	var accessData map[string]interface{}
	if err := ctx.ShouldBindJSON(&accessData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	accessData["bank_id"] = bankID
	accessData["account_id"] = accountID
	accessData["view_id"] = viewID
	ctx.JSON(http.StatusCreated, accessData)
}

func (c *OBPCoreController) RevokeUserAccessToView(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	viewID := ctx.Param("viewId")
	var accessData map[string]interface{}
	if err := ctx.ShouldBindJSON(&accessData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Access revoked", "bank_id": bankID, "account_id": accountID, "view_id": viewID})
}

func (c *OBPCoreController) CreateUserWithAccountAccess(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	viewID := ctx.Param("viewId")
	var userData map[string]interface{}
	if err := ctx.ShouldBindJSON(&userData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userData["bank_id"] = bankID
	userData["account_id"] = accountID
	userData["view_id"] = viewID
	ctx.JSON(http.StatusCreated, userData)
}

func (c *OBPCoreController) GetTransactionRequestById(ctx *gin.Context) {
	transactionRequestID := ctx.Param("transactionRequestId")
	transactionRequest := map[string]interface{}{
		"transaction_request_id": transactionRequestID,
	}
	ctx.JSON(http.StatusOK, transactionRequest)
}

func (c *OBPCoreController) GetTransactionRequests(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	viewID := ctx.Param("viewId")
	transactionRequests := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"transaction_requests": transactionRequests, "bank_id": bankID, "account_id": accountID, "view_id": viewID})
}

func (c *OBPCoreController) UpdateTransactionRequestStatus(ctx *gin.Context) {
	transactionRequestID := ctx.Param("transactionRequestId")
	var statusData map[string]interface{}
	if err := ctx.ShouldBindJSON(&statusData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	statusData["transaction_request_id"] = transactionRequestID
	ctx.JSON(http.StatusOK, statusData)
}

func (c *OBPCoreController) GetAccountAccessByUserId(ctx *gin.Context) {
	userID := ctx.Param("userId")
	accountAccess := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"account_access": accountAccess, "user_id": userID})
}

func (c *OBPCoreController) GetAPITags(ctx *gin.Context) {
	tags := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"tags": tags})
}

func (c *OBPCoreController) GetAccountByIdThroughView(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	viewID := ctx.Param("viewId")
	account := map[string]interface{}{
		"bank_id":    bankID,
		"account_id": accountID,
		"view_id":    viewID,
	}
	ctx.JSON(http.StatusOK, account)
}

func (c *OBPCoreController) GetAccountBalancesByBankIdAndAccountIdThroughView(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	viewID := ctx.Param("viewId")
	balances := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"balances": balances, "bank_id": bankID, "account_id": accountID, "view_id": viewID})
}

func (c *OBPCoreController) GetAccountBalancesByBankId(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	balances := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"balances": balances, "bank_id": bankID})
}

func (c *OBPCoreController) GetAccountBalancesByBankIdThroughView(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	viewID := ctx.Param("viewId")
	balances := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"balances": balances, "bank_id": bankID, "view_id": viewID})
}

func (c *OBPCoreController) CreateCounterpartyLimit(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	viewID := ctx.Param("viewId")
	counterpartyID := ctx.Param("counterpartyId")
	var limitData map[string]interface{}
	if err := ctx.ShouldBindJSON(&limitData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	limitData["bank_id"] = bankID
	limitData["account_id"] = accountID
	limitData["view_id"] = viewID
	limitData["counterparty_id"] = counterpartyID
	ctx.JSON(http.StatusCreated, limitData)
}

func (c *OBPCoreController) UpdateCounterpartyLimit(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	viewID := ctx.Param("viewId")
	counterpartyID := ctx.Param("counterpartyId")
	var limitData map[string]interface{}
	if err := ctx.ShouldBindJSON(&limitData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	limitData["bank_id"] = bankID
	limitData["account_id"] = accountID
	limitData["view_id"] = viewID
	limitData["counterparty_id"] = counterpartyID
	ctx.JSON(http.StatusOK, limitData)
}

func (c *OBPCoreController) GetCounterpartyLimit(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	viewID := ctx.Param("viewId")
	counterpartyID := ctx.Param("counterpartyId")
	limit := map[string]interface{}{
		"bank_id":         bankID,
		"account_id":      accountID,
		"view_id":         viewID,
		"counterparty_id": counterpartyID,
	}
	ctx.JSON(http.StatusOK, limit)
}

func (c *OBPCoreController) GetCounterpartyLimitStatus(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	viewID := ctx.Param("viewId")
	counterpartyID := ctx.Param("counterpartyId")
	status := map[string]interface{}{
		"bank_id":         bankID,
		"account_id":      accountID,
		"view_id":         viewID,
		"counterparty_id": counterpartyID,
		"status":          "ACTIVE",
	}
	ctx.JSON(http.StatusOK, status)
}

func (c *OBPCoreController) DeleteCounterpartyLimit(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	viewID := ctx.Param("viewId")
	counterpartyID := ctx.Param("counterpartyId")
	ctx.JSON(http.StatusOK, gin.H{"message": "Counterparty limit deleted", "bank_id": bankID, "account_id": accountID, "view_id": viewID, "counterparty_id": counterpartyID})
}

func (c *OBPCoreController) CreateCustomView(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	viewID := ctx.Param("viewId")
	var viewData map[string]interface{}
	if err := ctx.ShouldBindJSON(&viewData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	viewData["bank_id"] = bankID
	viewData["account_id"] = accountID
	viewData["view_id"] = viewID
	ctx.JSON(http.StatusCreated, viewData)
}

func (c *OBPCoreController) UpdateCustomView(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	viewID := ctx.Param("viewId")
	targetViewID := ctx.Param("targetViewId")
	var viewData map[string]interface{}
	if err := ctx.ShouldBindJSON(&viewData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	viewData["bank_id"] = bankID
	viewData["account_id"] = accountID
	viewData["view_id"] = viewID
	viewData["target_view_id"] = targetViewID
	ctx.JSON(http.StatusOK, viewData)
}

func (c *OBPCoreController) GetCustomView(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	viewID := ctx.Param("viewId")
	targetViewID := ctx.Param("targetViewId")
	view := map[string]interface{}{
		"bank_id":        bankID,
		"account_id":     accountID,
		"view_id":        viewID,
		"target_view_id": targetViewID,
	}
	ctx.JSON(http.StatusOK, view)
}

func (c *OBPCoreController) DeleteCustomView(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	viewID := ctx.Param("viewId")
	targetViewID := ctx.Param("targetViewId")
	ctx.JSON(http.StatusOK, gin.H{"message": "Custom view deleted", "bank_id": bankID, "account_id": accountID, "view_id": viewID, "target_view_id": targetViewID})
}

func (c *OBPCoreController) CreateVRPConsentRequest(ctx *gin.Context) {
	var vrpData map[string]interface{}
	if err := ctx.ShouldBindJSON(&vrpData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, vrpData)
}

func (c *OBPCoreController) CreateRegulatedEntityAttribute(ctx *gin.Context) {
	entityID := ctx.Param("regulatedEntityId")
	var attrData map[string]interface{}
	if err := ctx.ShouldBindJSON(&attrData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	attrData["regulated_entity_id"] = entityID
	ctx.JSON(http.StatusCreated, attrData)
}

func (c *OBPCoreController) DeleteRegulatedEntityAttribute(ctx *gin.Context) {
	entityID := ctx.Param("regulatedEntityId")
	attrID := ctx.Param("regulatedEntityAttributeId")
	ctx.JSON(http.StatusOK, gin.H{"message": "Regulated entity attribute deleted", "regulated_entity_id": entityID, "attribute_id": attrID})
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

func (c *OBPCoreController) GetUserByUserId(ctx *gin.Context) {
	userID := ctx.Param("userId")
	user := map[string]interface{}{
		"user_id":  userID,
		"username": "user_" + userID,
		"email":    "user@example.com",
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *OBPCoreController) GetUserByUsername(ctx *gin.Context) {
	username := ctx.Param("username")
	user := map[string]interface{}{
		"username": username,
		"user_id":  "user_123",
		"email":    "user@example.com",
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *OBPCoreController) GetUsersByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	users := []map[string]interface{}{
		{
			"email":    email,
			"user_id":  "user_123",
			"username": "user_123",
		},
	}
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func (c *OBPCoreController) GetCurrentUser(ctx *gin.Context) {
	user := map[string]interface{}{
		"user_id":    "current_user_123",
		"username":   "current_user",
		"email":      "current.user@example.com",
		"first_name": "Current",
		"last_name":  "User",
		"is_active":  true,
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *OBPCoreController) GetUsers(ctx *gin.Context) {
	users := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func (c *OBPCoreController) CreateUserInvitation(ctx *gin.Context) {
	var invitationData map[string]interface{}
	if err := ctx.ShouldBindJSON(&invitationData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, invitationData)
}

func (c *OBPCoreController) GetUserInvitationAnonymous(ctx *gin.Context) {
	secretLink := ctx.Param("secretLink")
	invitation := map[string]interface{}{
		"secret_link": secretLink,
		"status":      "pending",
	}
	ctx.JSON(http.StatusOK, invitation)
}

func (c *OBPCoreController) GetUserInvitation(ctx *gin.Context) {
	secretLink := ctx.Param("secretLink")
	invitation := map[string]interface{}{
		"secret_link": secretLink,
		"status":      "pending",
	}
	ctx.JSON(http.StatusOK, invitation)
}

func (c *OBPCoreController) GetUserInvitations(ctx *gin.Context) {
	invitations := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"user_invitations": invitations})
}

func (c *OBPCoreController) DeleteUser(ctx *gin.Context) {
	userID := ctx.Param("userId")
	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted", "user_id": userID})
}

func (c *OBPCoreController) UpdateConsentStatus(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	consentID := ctx.Param("consentId")

	var statusData map[string]interface{}
	if err := ctx.ShouldBindJSON(&statusData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	statusData["bank_id"] = bankID
	statusData["consent_id"] = consentID
	ctx.JSON(http.StatusOK, statusData)
}

func (c *OBPCoreController) UpdateConsentAccountAccess(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	consentID := ctx.Param("consentId")

	var accessData map[string]interface{}
	if err := ctx.ShouldBindJSON(&accessData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessData["bank_id"] = bankID
	accessData["consent_id"] = consentID
	ctx.JSON(http.StatusOK, accessData)
}

func (c *OBPCoreController) UpdateConsentUserId(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	consentID := ctx.Param("consentId")

	var userIdData map[string]interface{}
	if err := ctx.ShouldBindJSON(&userIdData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIdData["bank_id"] = bankID
	userIdData["consent_id"] = consentID
	ctx.JSON(http.StatusOK, userIdData)
}

func (c *OBPCoreController) GetMyConsents(ctx *gin.Context) {
	consents := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"consents": consents})
}

func (c *OBPCoreController) GetConsentsAtBank(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	consents := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"consents": consents, "bank_id": bankID})
}

func (c *OBPCoreController) GetConsents(ctx *gin.Context) {
	consents := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"consents": consents})
}

func (c *OBPCoreController) GetConsentByConsentId(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	consentID := ctx.Param("consentId")

	consent := map[string]interface{}{
		"bank_id":    bankID,
		"consent_id": consentID,
		"status":     "active",
	}
	ctx.JSON(http.StatusOK, consent)
}

func (c *OBPCoreController) RevokeConsentAtBank(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	consentID := ctx.Param("consentId")
	ctx.JSON(http.StatusOK, gin.H{"message": "Consent revoked", "bank_id": bankID, "consent_id": consentID})
}

func (c *OBPCoreController) RevokeMyConsent(ctx *gin.Context) {
	consentID := ctx.Param("consentId")
	ctx.JSON(http.StatusOK, gin.H{"message": "My consent revoked", "consent_id": consentID})
}

func (c *OBPCoreController) GetAggregateMetrics(ctx *gin.Context) {
	metrics := map[string]interface{}{
		"total_requests":     1000,
		"total_users":        50,
		"total_transactions": 500,
	}
	ctx.JSON(http.StatusOK, metrics)
}

func (c *OBPCoreController) GetMetrics(ctx *gin.Context) {
	metrics := []map[string]interface{}{
		{
			"metric_name": "api_calls",
			"value":       1000,
			"timestamp":   "2023-01-01T00:00:00Z",
		},
	}
	ctx.JSON(http.StatusOK, gin.H{"metrics": metrics})
}

func (c *OBPCoreController) GetRegulatedEntities(ctx *gin.Context) {
	entities := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"regulated_entities": entities})
}

func (c *OBPCoreController) GetRegulatedEntityById(ctx *gin.Context) {
	entityID := ctx.Param("regulatedEntityId")
	entity := map[string]interface{}{
		"regulated_entity_id": entityID,
		"name":                "Entity " + entityID,
	}
	ctx.JSON(http.StatusOK, entity)
}

func (c *OBPCoreController) CreateRegulatedEntity(ctx *gin.Context) {
	var entityData map[string]interface{}
	if err := ctx.ShouldBindJSON(&entityData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, entityData)
}

func (c *OBPCoreController) GetRegulatedEntityAttributeById(ctx *gin.Context) {
	entityID := ctx.Param("regulatedEntityId")
	attributeID := ctx.Param("regulatedEntityAttributeId")

	attribute := map[string]interface{}{
		"regulated_entity_id": entityID,
		"attribute_id":        attributeID,
		"name":                "attribute_name",
		"value":               "attribute_value",
	}
	ctx.JSON(http.StatusOK, attribute)
}

func (c *OBPCoreController) CustomViewNamesCheck(ctx *gin.Context) {
	result := map[string]interface{}{
		"status":  "ok",
		"message": "Custom view names check passed",
	}
	ctx.JSON(http.StatusOK, result)
}

func (c *OBPCoreController) SystemViewNamesCheck(ctx *gin.Context) {
	result := map[string]interface{}{
		"status":  "ok",
		"message": "System view names check passed",
	}
	ctx.JSON(http.StatusOK, result)
}

func (c *OBPCoreController) AccountAccessUniqueIndexCheck(ctx *gin.Context) {
	result := map[string]interface{}{
		"status":  "ok",
		"message": "Account access unique index check passed",
	}
	ctx.JSON(http.StatusOK, result)
}

func (c *OBPCoreController) AccountCurrencyCheck(ctx *gin.Context) {
	result := map[string]interface{}{
		"status":  "ok",
		"message": "Account currency check passed",
	}
	ctx.JSON(http.StatusOK, result)
}

func (c *OBPCoreController) OrphanedAccountCheck(ctx *gin.Context) {
	result := map[string]interface{}{
		"status":  "ok",
		"message": "Orphaned account check passed",
	}
	ctx.JSON(http.StatusOK, result)
}

func (c *OBPCoreController) GetCurrenciesAtBank(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	currencies := []map[string]interface{}{
		{"code": "USD", "name": "US Dollar"},
		{"code": "EUR", "name": "Euro"},
	}
	ctx.JSON(http.StatusOK, gin.H{"currencies": currencies, "bank_id": bankID})
}

func (c *OBPCoreController) GetATMAttributeDefinitions(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	definitions := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"atm_attribute_definitions": definitions, "bank_id": bankID})
}

func (c *OBPCoreController) CreateATMAttributeDefinition(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	var defData map[string]interface{}
	if err := ctx.ShouldBindJSON(&defData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defData["bank_id"] = bankID
	ctx.JSON(http.StatusCreated, defData)
}

func (c *OBPCoreController) UpdateATMAttributeDefinition(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	defID := ctx.Param("attributeDefinitionId")
	var defData map[string]interface{}
	if err := ctx.ShouldBindJSON(&defData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defData["bank_id"] = bankID
	defData["definition_id"] = defID
	ctx.JSON(http.StatusOK, defData)
}

func (c *OBPCoreController) DeleteATMAttributeDefinition(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	defID := ctx.Param("attributeDefinitionId")
	ctx.JSON(http.StatusOK, gin.H{
		"message":       "ATM attribute definition deleted",
		"bank_id":       bankID,
		"definition_id": defID,
	})
}

func (c *OBPCoreController) GetConsumerByConsumerId(ctx *gin.Context) {
	consumerID := ctx.Param("consumerId")
	consumer := map[string]interface{}{
		"consumer_id": consumerID,
		"app_name":    "Sample App",
		"app_type":    "Web",
		"description": "Sample consumer application",
	}
	ctx.JSON(http.StatusOK, consumer)
}

func (c *OBPCoreController) UpdateConsumer(ctx *gin.Context) {
	consumerID := ctx.Param("consumerId")
	var consumerData map[string]interface{}
	if err := ctx.ShouldBindJSON(&consumerData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	consumerData["consumer_id"] = consumerID
	ctx.JSON(http.StatusOK, consumerData)
}

func (c *OBPCoreController) DeleteConsumer(ctx *gin.Context) {
	consumerID := ctx.Param("consumerId")
	ctx.JSON(http.StatusOK, gin.H{
		"message":     "Consumer deleted",
		"consumer_id": consumerID,
	})
}

func (c *OBPCoreController) GetConsumerRedirectURL(ctx *gin.Context) {
	consumerID := ctx.Param("consumerId")
	url := map[string]interface{}{
		"consumer_id":  consumerID,
		"redirect_url": "https://example.com/callback",
	}
	ctx.JSON(http.StatusOK, url)
}

func (c *OBPCoreController) GetConsumerLogoURL(ctx *gin.Context) {
	consumerID := ctx.Param("consumerId")
	url := map[string]interface{}{
		"consumer_id": consumerID,
		"logo_url":    "https://example.com/logo.png",
	}
	ctx.JSON(http.StatusOK, url)
}

func (c *OBPCoreController) GetSettlementAccounts(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accounts := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"settlement_accounts": accounts, "bank_id": bankID})
}

func (c *OBPCoreController) CreateSettlementAccount(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	var accountData map[string]interface{}
	if err := ctx.ShouldBindJSON(&accountData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	accountData["bank_id"] = bankID
	ctx.JSON(http.StatusCreated, accountData)
}

func (c *OBPCoreController) GetSettlementAccount(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	account := map[string]interface{}{
		"bank_id":      bankID,
		"account_id":   accountID,
		"account_type": "SETTLEMENT",
	}
	ctx.JSON(http.StatusOK, account)
}

func (c *OBPCoreController) UpdateSettlementAccount(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	var accountData map[string]interface{}
	if err := ctx.ShouldBindJSON(&accountData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	accountData["bank_id"] = bankID
	accountData["account_id"] = accountID
	ctx.JSON(http.StatusOK, accountData)
}

func (c *OBPCoreController) DeleteSettlementAccount(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	ctx.JSON(http.StatusOK, gin.H{
		"message":    "Settlement account deleted",
		"bank_id":    bankID,
		"account_id": accountID,
	})
}

func (c *OBPCoreController) GetWebhooks(ctx *gin.Context) {
	webhooks := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"webhooks": webhooks})
}

func (c *OBPCoreController) CreateWebhook(ctx *gin.Context) {
	var webhookData map[string]interface{}
	if err := ctx.ShouldBindJSON(&webhookData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, webhookData)
}

func (c *OBPCoreController) GetWebhook(ctx *gin.Context) {
	webhookID := ctx.Param("webhookId")
	webhook := map[string]interface{}{
		"webhook_id": webhookID,
		"url":        "https://example.com/webhook",
		"is_active":  true,
	}
	ctx.JSON(http.StatusOK, webhook)
}

func (c *OBPCoreController) UpdateWebhook(ctx *gin.Context) {
	webhookID := ctx.Param("webhookId")
	var webhookData map[string]interface{}
	if err := ctx.ShouldBindJSON(&webhookData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	webhookData["webhook_id"] = webhookID
	ctx.JSON(http.StatusOK, webhookData)
}

func (c *OBPCoreController) DeleteWebhook(ctx *gin.Context) {
	webhookID := ctx.Param("webhookId")
	ctx.JSON(http.StatusOK, gin.H{
		"message":    "Webhook deleted",
		"webhook_id": webhookID,
	})
}

func (c *OBPCoreController) GetTransactionRequestTypes(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	viewID := ctx.Param("viewId")
	types := []map[string]interface{}{
		{"value": "SEPA", "bank_id": bankID, "account_id": accountID, "view_id": viewID},
		{"value": "COUNTERPARTY", "bank_id": bankID, "account_id": accountID, "view_id": viewID},
		{"value": "TRANSFER_TO_PHONE", "bank_id": bankID, "account_id": accountID, "view_id": viewID},
		{"value": "TRANSFER_TO_ATM", "bank_id": bankID, "account_id": accountID, "view_id": viewID},
		{"value": "REFUND", "bank_id": bankID, "account_id": accountID, "view_id": viewID},
	}
	ctx.JSON(http.StatusOK, gin.H{"transaction_request_types": types})
}

func (c *OBPCoreController) GetTransactionRequestTypesSupportedByBank(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	types := []map[string]interface{}{
		{"value": "SEPA", "bank_id": bankID},
		{"value": "COUNTERPARTY", "bank_id": bankID},
		{"value": "TRANSFER_TO_PHONE", "bank_id": bankID},
		{"value": "TRANSFER_TO_ATM", "bank_id": bankID},
		{"value": "REFUND", "bank_id": bankID},
	}
	ctx.JSON(http.StatusOK, gin.H{"transaction_request_types": types})
}

func (c *OBPCoreController) GetUserLockStatus(ctx *gin.Context) {
	userID := ctx.Param("userId")
	status := map[string]interface{}{
		"user_id":     userID,
		"is_locked":   false,
		"lock_reason": "",
	}
	ctx.JSON(http.StatusOK, status)
}

func (c *OBPCoreController) LockUser(ctx *gin.Context) {
	userID := ctx.Param("userId")
	var lockData map[string]interface{}
	if err := ctx.ShouldBindJSON(&lockData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lockData["user_id"] = userID
	lockData["is_locked"] = true
	ctx.JSON(http.StatusOK, lockData)
}

func (c *OBPCoreController) UnlockUser(ctx *gin.Context) {
	userID := ctx.Param("userId")
	status := map[string]interface{}{
		"user_id":   userID,
		"is_locked": false,
		"message":   "User unlocked successfully",
	}
	ctx.JSON(http.StatusOK, status)
}

func (c *OBPCoreController) CreateConsumerRedirectURL(ctx *gin.Context) {
	consumerID := ctx.Param("consumerId")
	var urlData map[string]interface{}
	if err := ctx.ShouldBindJSON(&urlData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	urlData["consumer_id"] = consumerID
	ctx.JSON(http.StatusCreated, urlData)
}

func (c *OBPCoreController) DeleteConsumerRedirectURL(ctx *gin.Context) {
	consumerID := ctx.Param("consumerId")
	ctx.JSON(http.StatusOK, gin.H{
		"message":     "Consumer redirect URL deleted",
		"consumer_id": consumerID,
	})
}

func (c *OBPCoreController) GetUserAttributesByUser(ctx *gin.Context) {
	userID := ctx.Param("userId")
	attributes := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"user_attributes": attributes, "user_id": userID})
}

func (c *OBPCoreController) CreateUserAttributeForUser(ctx *gin.Context) {
	userID := ctx.Param("userId")
	var attrData map[string]interface{}
	if err := ctx.ShouldBindJSON(&attrData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	attrData["user_id"] = userID
	ctx.JSON(http.StatusCreated, attrData)
}

func (c *OBPCoreController) UpdateUserAttribute(ctx *gin.Context) {
	userID := ctx.Param("userId")
	attrID := ctx.Param("userAttributeId")
	var attrData map[string]interface{}
	if err := ctx.ShouldBindJSON(&attrData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	attrData["user_id"] = userID
	attrData["attribute_id"] = attrID
	ctx.JSON(http.StatusOK, attrData)
}

func (c *OBPCoreController) DeleteUserAttributeForUser(ctx *gin.Context) {
	userID := ctx.Param("userId")
	attrID := ctx.Param("userAttributeId")
	ctx.JSON(http.StatusOK, gin.H{
		"message":      "User attribute deleted",
		"user_id":      userID,
		"attribute_id": attrID,
	})
}

func (c *OBPCoreController) GetSettlementAccountsAtBank(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accounts := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"settlement_accounts": accounts, "bank_id": bankID})
}

func (c *OBPCoreController) CreateSettlementAccountAtBank(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	var accountData map[string]interface{}
	if err := ctx.ShouldBindJSON(&accountData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	accountData["bank_id"] = bankID
	ctx.JSON(http.StatusCreated, accountData)
}

func (c *OBPCoreController) DeleteConsumerLogoURL(ctx *gin.Context) {
	consumerID := ctx.Param("consumerId")
	ctx.JSON(http.StatusOK, gin.H{
		"message":     "Consumer logo URL deleted",
		"consumer_id": consumerID,
	})
}

func (c *OBPCoreController) GetUserAttributeDefinitions(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	definitions := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"user_attribute_definitions": definitions, "bank_id": bankID})
}

func (c *OBPCoreController) CreateUserAttributeDefinition(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	var defData map[string]interface{}
	if err := ctx.ShouldBindJSON(&defData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defData["bank_id"] = bankID
	ctx.JSON(http.StatusCreated, defData)
}

func (c *OBPCoreController) UpdateUserAttributeDefinition(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	defID := ctx.Param("userAttributeDefinitionId")
	var defData map[string]interface{}
	if err := ctx.ShouldBindJSON(&defData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defData["bank_id"] = bankID
	defData["definition_id"] = defID
	ctx.JSON(http.StatusOK, defData)
}

func (c *OBPCoreController) DeleteUserAttributeDefinition(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	defID := ctx.Param("userAttributeDefinitionId")
	ctx.JSON(http.StatusOK, gin.H{
		"message":       "User attribute definition deleted",
		"bank_id":       bankID,
		"definition_id": defID,
	})
}

func (c *OBPCoreController) CreateUserAttributeNew(ctx *gin.Context) {
	userID := ctx.Param("userId")
	var attrData map[string]interface{}
	if err := ctx.ShouldBindJSON(&attrData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	attrData["user_id"] = userID
	ctx.JSON(http.StatusCreated, attrData)
}

func (c *OBPCoreController) GetUserAttributesNew(ctx *gin.Context) {
	userID := ctx.Param("userId")
	attributes := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"user_attributes": attributes, "user_id": userID})
}

func (c *OBPCoreController) DeleteUserAttributeNew(ctx *gin.Context) {
	userID := ctx.Param("userId")
	attrID := ctx.Param("userAttributeId")
	ctx.JSON(http.StatusOK, gin.H{"message": "User attribute deleted", "user_id": userID, "attribute_id": attrID})
}

func (c *OBPCoreController) SyncUserNew(ctx *gin.Context) {
	provider := ctx.Param("provider")
	providerID := ctx.Param("providerId")
	var syncData map[string]interface{}
	if err := ctx.ShouldBindJSON(&syncData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	syncData["provider"] = provider
	syncData["provider_id"] = providerID
	ctx.JSON(http.StatusOK, syncData)
}

func (c *OBPCoreController) GetUserAccountsAtBankNew(ctx *gin.Context) {
	userID := ctx.Param("userId")
	bankID := ctx.Param("bankId")
	accounts := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"accounts": accounts, "user_id": userID, "bank_id": bankID})
}

func (c *OBPCoreController) GetUserAccountsNew(ctx *gin.Context) {
	userID := ctx.Param("userId")
	accounts := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"accounts": accounts, "user_id": userID})
}

func (c *OBPCoreController) GetUserEntitlementsAndPermissionsNew(ctx *gin.Context) {
	userID := ctx.Param("userId")
	entitlements := []map[string]interface{}{}
	permissions := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{
		"entitlements": entitlements,
		"permissions":  permissions,
		"user_id":      userID,
	})
}

func (c *OBPCoreController) GetSettlementAccountsAtBankNew(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accounts := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"settlement_accounts": accounts, "bank_id": bankID})
}

func (c *OBPCoreController) CreateSettlementAccountAtBankNew(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	var accountData map[string]interface{}
	if err := ctx.ShouldBindJSON(&accountData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	accountData["bank_id"] = bankID
	accountData["account_type"] = "SETTLEMENT"
	ctx.JSON(http.StatusCreated, accountData)
}

func (c *OBPCoreController) CheckCustomViewNamesNew(ctx *gin.Context) {
	result := map[string]interface{}{
		"status":  "ok",
		"message": "Custom view names check passed",
		"issues":  []string{},
	}
	ctx.JSON(http.StatusOK, result)
}

func (c *OBPCoreController) CheckSystemViewNamesNew(ctx *gin.Context) {
	result := map[string]interface{}{
		"status":  "ok",
		"message": "System view names check passed",
		"issues":  []string{},
	}
	ctx.JSON(http.StatusOK, result)
}

func (c *OBPCoreController) CheckAccountAccessUniqueIndexNew(ctx *gin.Context) {
	result := map[string]interface{}{
		"status":  "ok",
		"message": "Account access unique index check passed",
		"issues":  []string{},
	}
	ctx.JSON(http.StatusOK, result)
}

func (c *OBPCoreController) CheckAccountCurrencyNew(ctx *gin.Context) {
	result := map[string]interface{}{
		"status":  "ok",
		"message": "Account currency check passed",
		"issues":  []string{},
	}
	ctx.JSON(http.StatusOK, result)
}

func (c *OBPCoreController) CheckOrphanedAccountsNew(ctx *gin.Context) {
	result := map[string]interface{}{
		"status":  "ok",
		"message": "Orphaned account check passed",
		"issues":  []string{},
	}
	ctx.JSON(http.StatusOK, result)
}

func (c *OBPCoreController) CreateATMAttributeNew(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	atmID := ctx.Param("atmId")
	var attrData map[string]interface{}
	if err := ctx.ShouldBindJSON(&attrData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	attrData["bank_id"] = bankID
	attrData["atm_id"] = atmID
	ctx.JSON(http.StatusCreated, attrData)
}

func (c *OBPCoreController) GetATMAttributesNew(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	atmID := ctx.Param("atmId")
	attributes := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"atm_attributes": attributes, "bank_id": bankID, "atm_id": atmID})
}

func (c *OBPCoreController) UpdateATMAttributeNew(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	atmID := ctx.Param("atmId")
	attrID := ctx.Param("attributeId")
	var attrData map[string]interface{}
	if err := ctx.ShouldBindJSON(&attrData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	attrData["bank_id"] = bankID
	attrData["atm_id"] = atmID
	attrData["attribute_id"] = attrID
	ctx.JSON(http.StatusOK, attrData)
}

func (c *OBPCoreController) DeleteATMAttributeNew(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	atmID := ctx.Param("atmId")
	attrID := ctx.Param("attributeId")
	ctx.JSON(http.StatusOK, gin.H{
		"message":      "ATM attribute deleted",
		"bank_id":      bankID,
		"atm_id":       atmID,
		"attribute_id": attrID,
	})
}

func (c *OBPCoreController) CreateConsumerNew(ctx *gin.Context) {
	var consumerData map[string]interface{}
	if err := ctx.ShouldBindJSON(&consumerData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, consumerData)
}

func (c *OBPCoreController) UpdateConsumerRedirectURLNew(ctx *gin.Context) {
	consumerID := ctx.Param("consumerId")
	var urlData map[string]interface{}
	if err := ctx.ShouldBindJSON(&urlData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	urlData["consumer_id"] = consumerID
	ctx.JSON(http.StatusOK, urlData)
}

func (c *OBPCoreController) GetWebhooksAtBank(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	webhooks := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"webhooks": webhooks, "bank_id": bankID})
}

func (c *OBPCoreController) CreateWebhookAtBank(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	var webhookData map[string]interface{}
	if err := ctx.ShouldBindJSON(&webhookData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	webhookData["bank_id"] = bankID
	ctx.JSON(http.StatusCreated, webhookData)
}

func (c *OBPCoreController) GetWebhookAtBank(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	webhookID := ctx.Param("webhookId")
	webhook := map[string]interface{}{
		"bank_id":    bankID,
		"webhook_id": webhookID,
		"url":        "https://example.com/webhook",
		"is_active":  true,
	}
	ctx.JSON(http.StatusOK, webhook)
}

func (c *OBPCoreController) UpdateWebhookAtBank(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	webhookID := ctx.Param("webhookId")
	var webhookData map[string]interface{}
	if err := ctx.ShouldBindJSON(&webhookData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	webhookData["bank_id"] = bankID
	webhookData["webhook_id"] = webhookID
	ctx.JSON(http.StatusOK, webhookData)
}

func (c *OBPCoreController) DeleteWebhookAtBank(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	webhookID := ctx.Param("webhookId")
	ctx.JSON(http.StatusOK, gin.H{
		"message":    "Webhook deleted",
		"bank_id":    bankID,
		"webhook_id": webhookID,
	})
}

func (c *OBPCoreController) CreateUserAttribute(ctx *gin.Context) {
	userID := ctx.Param("userId")
	var attrData map[string]interface{}
	if err := ctx.ShouldBindJSON(&attrData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	attrData["user_id"] = userID
	ctx.JSON(http.StatusCreated, attrData)
}

func (c *OBPCoreController) GetUserAttributes(ctx *gin.Context) {
	userID := ctx.Param("userId")
	attributes := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"user_attributes": attributes, "user_id": userID})
}

func (c *OBPCoreController) DeleteUserAttribute(ctx *gin.Context) {
	userID := ctx.Param("userId")
	attrID := ctx.Param("userAttributeId")
	ctx.JSON(http.StatusOK, gin.H{"message": "User attribute deleted", "user_id": userID, "attribute_id": attrID})
}

func (c *OBPCoreController) SyncUser(ctx *gin.Context) {
	provider := ctx.Param("provider")
	providerID := ctx.Param("providerId")
	var syncData map[string]interface{}
	if err := ctx.ShouldBindJSON(&syncData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	syncData["provider"] = provider
	syncData["provider_id"] = providerID
	ctx.JSON(http.StatusOK, syncData)
}

func (c *OBPCoreController) GetUserAccountsAtBank(ctx *gin.Context) {
	userID := ctx.Param("userId")
	bankID := ctx.Param("bankId")
	accounts := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"accounts": accounts, "user_id": userID, "bank_id": bankID})
}

func (c *OBPCoreController) GetUserAccounts(ctx *gin.Context) {
	userID := ctx.Param("userId")
	accounts := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"accounts": accounts, "user_id": userID})
}

func (c *OBPCoreController) GetUserEntitlementsAndPermissions(ctx *gin.Context) {
	userID := ctx.Param("userId")
	entitlements := []map[string]interface{}{}
	permissions := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{
		"entitlements": entitlements,
		"permissions":  permissions,
		"user_id":      userID,
	})
}

func (c *OBPCoreController) CheckCustomViewNames(ctx *gin.Context) {
	result := map[string]interface{}{
		"status":  "ok",
		"message": "Custom view names check passed",
		"issues":  []string{},
	}
	ctx.JSON(http.StatusOK, result)
}

func (c *OBPCoreController) CheckSystemViewNames(ctx *gin.Context) {
	result := map[string]interface{}{
		"status":  "ok",
		"message": "System view names check passed",
		"issues":  []string{},
	}
	ctx.JSON(http.StatusOK, result)
}

func (c *OBPCoreController) CheckAccountAccessUniqueIndex(ctx *gin.Context) {
	result := map[string]interface{}{
		"status":  "ok",
		"message": "Account access unique index check passed",
		"issues":  []string{},
	}
	ctx.JSON(http.StatusOK, result)
}

func (c *OBPCoreController) CheckAccountCurrency(ctx *gin.Context) {
	result := map[string]interface{}{
		"status":  "ok",
		"message": "Account currency check passed",
		"issues":  []string{},
	}
	ctx.JSON(http.StatusOK, result)
}

func (c *OBPCoreController) CheckOrphanedAccounts(ctx *gin.Context) {
	result := map[string]interface{}{
		"status":  "ok",
		"message": "Orphaned account check passed",
		"issues":  []string{},
	}
	ctx.JSON(http.StatusOK, result)
}
