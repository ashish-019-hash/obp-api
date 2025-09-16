package routes

import (
	"obp-api-backend/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	bankController := controllers.NewBankController()
	accountController := controllers.NewAccountController()

	v1 := r.Group("/obp/v5.1.0")
	{
		v1.GET("/banks", bankController.GetBanks)
		v1.GET("/banks/:bankId", bankController.GetBank)

		v1.GET("/banks/:bankId/accounts", accountController.GetAccounts)
		v1.GET("/banks/:bankId/accounts/:accountId", accountController.GetAccount)
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
			"message": "OBP API Backend is running",
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": "Open Bank Project API",
			"version": "5.1.0",
			"description": "Complete REST API for Open Bank Project",
		})
	})
}
