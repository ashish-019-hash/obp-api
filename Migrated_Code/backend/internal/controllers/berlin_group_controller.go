package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"obp-api-backend/internal/services"
)

type BerlinGroupController struct {
	accountService services.AccountService
	balanceService services.BalanceService
	paymentService services.PaymentService
}

func NewBerlinGroupController(
	accountService services.AccountService,
	balanceService services.BalanceService,
	paymentService services.PaymentService,
) *BerlinGroupController {
	return &BerlinGroupController{
		accountService: accountService,
		balanceService: balanceService,
		paymentService: paymentService,
	}
}

func (c *BerlinGroupController) GetAccounts(ctx *gin.Context) {
	accounts, err := c.accountService.GetAccountsForUser(ctx.Request.Context(), "current-user")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	berlinGroupAccounts := make([]map[string]interface{}, len(accounts))
	for i, account := range accounts {
		berlinGroupAccounts[i] = map[string]interface{}{
			"resourceId":      account.AccountId,
			"iban":           account.AccountId,
			"currency":       account.Currency,
			"name":           account.Label,
			"product":        "CurrentAccount",
			"cashAccountType": "CACC",
		}
	}
	
	ctx.JSON(http.StatusOK, gin.H{"accounts": berlinGroupAccounts})
}

func (c *BerlinGroupController) GetAccountBalances(ctx *gin.Context) {
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

func (c *BerlinGroupController) InitiateSEPACreditTransfer(ctx *gin.Context) {
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
		"transactionStatus": "RCVD",
		"paymentId":        paymentID,
	})
}

func (c *BerlinGroupController) GetAccountTransactions(ctx *gin.Context) {
	accountID := ctx.Param("account-id")
	
	transactions, err := c.accountService.GetTransactionsByAccountID(ctx.Request.Context(), accountID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{"transactions": transactions})
}

func (c *BerlinGroupController) GetAccountDetails(ctx *gin.Context) {
	accountID := ctx.Param("account-id")
	
	account, err := c.accountService.GetAccountByID(ctx.Request.Context(), accountID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}
	
	berlinGroupAccount := map[string]interface{}{
		"resourceId":      account.AccountId,
		"iban":           account.AccountId,
		"currency":       account.Currency,
		"name":           account.Label,
		"product":        "CurrentAccount",
		"cashAccountType": "CACC",
	}
	
	ctx.JSON(http.StatusOK, berlinGroupAccount)
}

func (c *BerlinGroupController) CreateAccountInformationConsent(ctx *gin.Context) {
	var consentRequest map[string]interface{}
	if err := ctx.ShouldBindJSON(&consentRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	consentID := "consent-123"
	
	ctx.JSON(http.StatusCreated, gin.H{
		"consentId":     consentID,
		"consentStatus": "received",
	})
}

func (c *BerlinGroupController) GetConsentInformation(ctx *gin.Context) {
	consentID := ctx.Param("consentId")
	
	ctx.JSON(http.StatusOK, gin.H{
		"consentId":     consentID,
		"consentStatus": "valid",
		"validUntil":    "2024-09-16",
	})
}

func (c *BerlinGroupController) DeleteConsent(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, nil)
}

func (c *BerlinGroupController) GetConsentStatus(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"consentStatus": "valid",
	})
}
