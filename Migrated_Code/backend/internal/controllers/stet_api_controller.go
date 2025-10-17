package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"obp-api-backend/internal/services"
)

type STETAPIController struct {
	accountService services.AccountService
	balanceService services.BalanceService
	paymentService services.PaymentService
}

func NewSTETAPIController(
	accountService services.AccountService,
	balanceService services.BalanceService,
	paymentService services.PaymentService,
) *STETAPIController {
	return &STETAPIController{
		accountService: accountService,
		balanceService: balanceService,
		paymentService: paymentService,
	}
}

func (c *STETAPIController) GetAccounts(ctx *gin.Context) {
	accounts, err := c.accountService.GetAccountsForUser(ctx.Request.Context(), "current-user")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"accounts": accounts})
}

func (c *STETAPIController) GetAccountBalances(ctx *gin.Context) {
	accountID := ctx.Param("account-id")

	currentBalance, err := c.balanceService.CalculateCurrentBalance(ctx.Request.Context(), accountID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	balances := []map[string]interface{}{
		{
			"balanceAmount": map[string]interface{}{
				"currency": "EUR",
				"amount":   currentBalance.String(),
			},
			"balanceType":   "closingBooked",
			"referenceDate": "2023-09-16",
		},
	}

	ctx.JSON(http.StatusOK, gin.H{"balances": balances})
}

func (c *STETAPIController) CreatePaymentRequest(ctx *gin.Context) {
	var paymentRequest map[string]interface{}
	if err := ctx.ShouldBindJSON(&paymentRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	paymentID, err := c.paymentService.InitiatePayment(ctx.Request.Context(), paymentRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"paymentRequestResourceId": paymentID,
		"transactionStatus":        "RCVD",
	})
}
func (c *STETAPIController) GetAccount(ctx *gin.Context) {
	accountID := ctx.Param("accountId")

	account := map[string]interface{}{
		"resourceId": accountID,
		"currency":   "EUR",
		"name":       "Sample Account",
	}
	ctx.JSON(http.StatusOK, account)
}

func (c *STETAPIController) GetAccountTransactions(ctx *gin.Context) {
	accountID := ctx.Param("accountId")

	transactions := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{
		"account": map[string]interface{}{
			"resourceId": accountID,
		},
		"transactions": transactions,
	})
}

func (c *STETAPIController) GetPaymentRequest(ctx *gin.Context) {
	paymentID := ctx.Param("paymentId")

	payment := map[string]interface{}{
		"paymentRequestResourceId": paymentID,
		"transactionStatus":        "ACSC",
	}
	ctx.JSON(http.StatusOK, payment)
}

func (c *STETAPIController) CreateConsent(ctx *gin.Context) {
	var consentData map[string]interface{}
	if err := ctx.ShouldBindJSON(&consentData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	consentID := "consent-123"
	ctx.JSON(http.StatusCreated, gin.H{
		"consentId":     consentID,
		"consentStatus": "received",
	})
}

func (c *STETAPIController) GetConsent(ctx *gin.Context) {
	consentID := ctx.Param("consentId")

	consent := map[string]interface{}{
		"consentId":     consentID,
		"consentStatus": "valid",
	}
	ctx.JSON(http.StatusOK, consent)
}

func (c *STETAPIController) DeleteConsent(ctx *gin.Context) {
	consentID := ctx.Param("consentId")

	ctx.JSON(http.StatusOK, gin.H{
		"message":   "Consent deleted",
		"consentId": consentID,
	})
}

func (c *STETAPIController) GetConsentStatus(ctx *gin.Context) {
	consentID := ctx.Param("consentId")

	status := map[string]interface{}{
		"consentId":     consentID,
		"consentStatus": "valid",
	}
	ctx.JSON(http.StatusOK, status)
}
