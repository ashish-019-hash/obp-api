package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"obp-api-backend/internal/services"
)

type MxOFAPIController struct {
	accountService services.AccountService
	balanceService services.BalanceService
}

func NewMxOFAPIController(
	accountService services.AccountService,
	balanceService services.BalanceService,
) *MxOFAPIController {
	return &MxOFAPIController{
		accountService: accountService,
		balanceService: balanceService,
	}
}

func (c *MxOFAPIController) GetAccounts(ctx *gin.Context) {
	accounts, err := c.accountService.GetAccountsForUser(ctx.Request.Context(), "current-user")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{"accounts": accounts})
}

func (c *MxOFAPIController) GetAccountBalances(ctx *gin.Context) {
	accountID := ctx.Param("account-id")
	
	currentBalance, err := c.balanceService.CalculateCurrentBalance(ctx.Request.Context(), accountID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	balances := []map[string]interface{}{
		{
			"balanceAmount": map[string]interface{}{
				"currency": "MXN",
				"amount":   currentBalance.String(),
			},
			"balanceType":   "closingBooked",
			"referenceDate": "2023-09-16",
		},
	}
	
	ctx.JSON(http.StatusOK, gin.H{"balances": balances})
}

func (c *MxOFAPIController) GetAccountTransactions(ctx *gin.Context) {
	accountID := ctx.Param("account-id")
	
	transactions, err := c.accountService.GetTransactionsByAccountID(ctx.Request.Context(), accountID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{"transactions": transactions})
}
