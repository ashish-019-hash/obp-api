package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"obp-api-backend/internal/services"
)

type UKOpenBankingController struct {
	accountService services.AccountService
	balanceService services.BalanceService
	paymentService services.PaymentService
}

func NewUKOpenBankingController(
	accountService services.AccountService,
	balanceService services.BalanceService,
	paymentService services.PaymentService,
) *UKOpenBankingController {
	return &UKOpenBankingController{
		accountService: accountService,
		balanceService: balanceService,
		paymentService: paymentService,
	}
}

func (c *UKOpenBankingController) GetAccounts(ctx *gin.Context) {
	accounts, err := c.accountService.GetAccountsForUser(ctx.Request.Context(), "current-user")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	ukAccounts := make([]map[string]interface{}, len(accounts))
	for i, account := range accounts {
		ukAccounts[i] = map[string]interface{}{
			"AccountId":      account.AccountId,
			"Currency":       account.Currency,
			"AccountType":    "Personal",
			"AccountSubType": "CurrentAccount",
			"Nickname":       account.Label,
		}
	}
	
	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"Account": ukAccounts,
		},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) GetAccountBalances(ctx *gin.Context) {
	accountID := ctx.Param("AccountId")
	
	currentBalance, err := c.balanceService.CalculateCurrentBalance(ctx.Request.Context(), accountID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	balances := []map[string]interface{}{
		{
			"AccountId": accountID,
			"Amount": map[string]interface{}{
				"Amount":   currentBalance.String(),
				"Currency": "GBP",
			},
			"CreditDebitIndicator": "Credit",
			"Type":                "ClosingAvailable",
		},
	}
	
	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"Balance": balances,
		},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) GetAccountTransactions(ctx *gin.Context) {
	accountID := ctx.Param("AccountId")
	
	transactions, err := c.accountService.GetTransactionsByAccountID(ctx.Request.Context(), accountID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"Transaction": transactions,
		},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) CreateDomesticPaymentConsents(ctx *gin.Context) {
	var consentRequest map[string]interface{}
	if err := ctx.ShouldBindJSON(&consentRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	consentID := "consent-123"
	
	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"ConsentId": consentID,
			"Status":    "AwaitingAuthorisation",
		},
	}
	ctx.JSON(http.StatusCreated, response)
}

func (c *UKOpenBankingController) CreateDomesticPayments(ctx *gin.Context) {
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
	
	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"DomesticPaymentId": paymentID,
			"Status":           "AcceptedSettlementInProcess",
		},
	}
	ctx.JSON(http.StatusCreated, response)
}

func (c *UKOpenBankingController) GetDomesticPaymentConsent(ctx *gin.Context) {
	consentID := ctx.Param("ConsentId")
	
	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"ConsentId": consentID,
			"Status":    "Authorised",
		},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) CreateFundsConfirmationConsents(ctx *gin.Context) {
	var consentRequest map[string]interface{}
	if err := ctx.ShouldBindJSON(&consentRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	consentID := "funds-consent-123"
	
	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"ConsentId": consentID,
			"Status":    "AwaitingAuthorisation",
		},
	}
	ctx.JSON(http.StatusCreated, response)
}

func (c *UKOpenBankingController) CreateFundsConfirmations(ctx *gin.Context) {
	var fundsRequest map[string]interface{}
	if err := ctx.ShouldBindJSON(&fundsRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"FundsAvailable": true,
		},
	}
	ctx.JSON(http.StatusOK, response)
}
