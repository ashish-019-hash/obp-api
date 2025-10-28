package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"obp-api-backend/internal/services"
)

type AustralianCDRController struct {
	accountService services.AccountService
	balanceService services.BalanceService
}

func NewAustralianCDRController(
	accountService services.AccountService,
	balanceService services.BalanceService,
) *AustralianCDRController {
	return &AustralianCDRController{
		accountService: accountService,
		balanceService: balanceService,
	}
}

func (c *AustralianCDRController) GetAccounts(ctx *gin.Context) {
	accounts, err := c.accountService.GetAccountsForUser(ctx.Request.Context(), "current-user")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	cdrAccounts := make([]map[string]interface{}, len(accounts))
	for i, account := range accounts {
		cdrAccounts[i] = map[string]interface{}{
			"accountId":       account.AccountId,
			"displayName":     account.Label,
			"nickname":        account.Label,
			"maskedNumber":    "xxxx1234",
			"openStatus":      "OPEN",
			"isOwned":         true,
			"productCategory": "TRANS_AND_SAVINGS_ACCOUNTS",
			"productName":     "Complete Freedom",
		}
	}
	
	response := map[string]interface{}{
		"data": map[string]interface{}{
			"accounts": cdrAccounts,
		},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *AustralianCDRController) GetAccountDetail(ctx *gin.Context) {
	accountID := ctx.Param("accountId")
	
	account, err := c.accountService.GetAccountByID(ctx.Request.Context(), accountID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}
	
	cdrAccount := map[string]interface{}{
		"accountId":       account.AccountId,
		"displayName":     account.Label,
		"nickname":        account.Label,
		"maskedNumber":    "xxxx1234",
		"openStatus":      "OPEN",
		"isOwned":         true,
		"productCategory": "TRANS_AND_SAVINGS_ACCOUNTS",
		"productName":     "Complete Freedom",
	}
	
	response := map[string]interface{}{
		"data": cdrAccount,
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *AustralianCDRController) GetAccountBalance(ctx *gin.Context) {
	accountID := ctx.Param("accountId")
	
	currentBalance, err := c.balanceService.CalculateCurrentBalance(ctx.Request.Context(), accountID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	response := map[string]interface{}{
		"data": map[string]interface{}{
			"accountId":        accountID,
			"currentBalance":   currentBalance.String(),
			"availableBalance": currentBalance.String(),
		},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *AustralianCDRController) GetAccountTransactions(ctx *gin.Context) {
	accountID := ctx.Param("accountId")
	
	transactions, err := c.accountService.GetTransactionsByAccountID(ctx.Request.Context(), accountID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	response := map[string]interface{}{
		"data": map[string]interface{}{
			"transactions": transactions,
		},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *AustralianCDRController) GetCustomer(ctx *gin.Context) {
	response := map[string]interface{}{
		"data": map[string]interface{}{
			"customerUType": "person",
			"person": map[string]interface{}{
				"lastUpdateTime": "2023-09-16T10:30:00Z",
				"firstName":      "John",
				"lastName":       "Smith",
			},
		},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *AustralianCDRController) GetCustomerDetail(ctx *gin.Context) {
	response := map[string]interface{}{
		"data": map[string]interface{}{
			"customerUType": "person",
			"person": map[string]interface{}{
				"lastUpdateTime": "2023-09-16T10:30:00Z",
				"firstName":      "John",
				"lastName":       "Smith",
				"middleNames":    []string{},
				"prefix":         "Mr",
				"suffix":         "",
			},
		},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *AustralianCDRController) GetAccountTransactionDetail(ctx *gin.Context) {
	accountID := ctx.Param("accountId")
	transactionID := ctx.Param("transactionId")
	
	response := map[string]interface{}{
		"data": map[string]interface{}{
			"accountId":     accountID,
			"transactionId": transactionID,
			"status":        "POSTED",
		},
		"links": map[string]interface{}{
			"self": "/cds-au/v1/banking/accounts/" + accountID + "/transactions/" + transactionID,
		},
		"meta": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *AustralianCDRController) GetAccountDirectDebits(ctx *gin.Context) {
	accountID := ctx.Param("accountId")
	
	directDebits := []map[string]interface{}{}
	
	response := map[string]interface{}{
		"data": map[string]interface{}{
			"directDebits": directDebits,
		},
		"links": map[string]interface{}{
			"self": "/cds-au/v1/banking/accounts/" + accountID + "/direct-debits",
		},
		"meta": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *AustralianCDRController) GetAccountScheduledPayments(ctx *gin.Context) {
	accountID := ctx.Param("accountId")
	
	scheduledPayments := []map[string]interface{}{}
	
	response := map[string]interface{}{
		"data": map[string]interface{}{
			"scheduledPayments": scheduledPayments,
		},
		"links": map[string]interface{}{
			"self": "/cds-au/v1/banking/accounts/" + accountID + "/payments/scheduled",
		},
		"meta": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *AustralianCDRController) GetProducts(ctx *gin.Context) {
	products := []map[string]interface{}{}
	
	response := map[string]interface{}{
		"data": map[string]interface{}{
			"products": products,
		},
		"links": map[string]interface{}{
			"self": "/cds-au/v1/banking/products",
		},
		"meta": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *AustralianCDRController) GetProductDetail(ctx *gin.Context) {
	productID := ctx.Param("productId")
	
	response := map[string]interface{}{
		"data": map[string]interface{}{
			"productId": productID,
			"name":      "Sample Product",
		},
		"links": map[string]interface{}{
			"self": "/cds-au/v1/banking/products/" + productID,
		},
		"meta": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}
