package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BankController struct{}

func NewBankController() *BankController {
	return &BankController{}
}

func (bc *BankController) GetBanks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"banks":   []interface{}{},
		"message": "Banks endpoint - implementation pending",
	})
}

func (bc *BankController) GetBank(c *gin.Context) {
	bankID := c.Param("bankId")
	c.JSON(http.StatusOK, gin.H{
		"bank_id": bankID,
		"message": "Get bank endpoint - implementation pending",
	})
}

type AccountController struct{}

func NewAccountController() *AccountController {
	return &AccountController{}
}

func (ac *AccountController) GetAccounts(c *gin.Context) {
	bankID := c.Param("bankId")
	c.JSON(http.StatusOK, gin.H{
		"bank_id":  bankID,
		"accounts": []interface{}{},
		"message":  "Accounts endpoint - implementation pending",
	})
}

func (ac *AccountController) GetAccount(c *gin.Context) {
	bankID := c.Param("bankId")
	accountID := c.Param("accountId")
	c.JSON(http.StatusOK, gin.H{
		"bank_id":    bankID,
		"account_id": accountID,
		"message":    "Get account endpoint - implementation pending",
	})
}
