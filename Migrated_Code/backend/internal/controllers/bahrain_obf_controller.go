package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"obp-api-backend/internal/services"
)

type BahrainOBFController struct {
	accountService services.AccountService
	balanceService services.BalanceService
	paymentService services.PaymentService
}

func NewBahrainOBFController(
	accountService services.AccountService,
	balanceService services.BalanceService,
	paymentService services.PaymentService,
) *BahrainOBFController {
	return &BahrainOBFController{
		accountService: accountService,
		balanceService: balanceService,
		paymentService: paymentService,
	}
}

func (c *BahrainOBFController) GetAccounts(ctx *gin.Context) {
	accounts, err := c.accountService.GetAccountsForUser(ctx.Request.Context(), "current-user")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"Account": accounts,
		},
		"Links": map[string]interface{}{
			"Self": "/bahrain-obf/v1.0.0/accounts",
		},
		"Meta": map[string]interface{}{
			"TotalPages": 1,
		},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetAccount(ctx *gin.Context) {
	accountID := ctx.Param("AccountId")
	
	account, err := c.accountService.GetAccountByID(ctx.Request.Context(), accountID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}
	
	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"Account": account,
		},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetAccountBalances(ctx *gin.Context) {
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
				"Currency": "BHD",
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

func (c *BahrainOBFController) CreateDomesticPaymentConsent(ctx *gin.Context) {
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

func (c *BahrainOBFController) CreateDomesticPayment(ctx *gin.Context) {
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
