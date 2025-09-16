package routes

import (
	"github.com/gin-gonic/gin"
	"obp-api-backend/internal/controllers"
	"obp-api-backend/internal/middleware"
)

func SetupRoutes(
	obpController *controllers.OBPCoreController,
	berlinGroupController *controllers.BerlinGroupController,
	ukOpenBankingController *controllers.UKOpenBankingController,
	australianCDRController *controllers.AustralianCDRController,
	bahrainOBFController *controllers.BahrainOBFController,
	polishAPIController *controllers.PolishAPIController,
	stetAPIController *controllers.STETAPIController,
	mxofAPIController *controllers.MxOFAPIController,
	additionalController *controllers.AdditionalRegulatoryController,
) *gin.Engine {
	router := gin.Default()
	
	router.Use(middleware.CORS())
	router.Use(middleware.Logger())
	router.Use(middleware.ErrorHandler())
	
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
			"message": "OBP API Backend is running",
		})
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": "Open Bank Project API",
			"version": "5.1.0",
			"description": "Complete REST API for Open Bank Project",
		})
	})
	
	v5 := router.Group("/obp/v5.1.0")
	{
		v5.GET("/root", obpController.GetAPIInfo)
		v5.GET("/ui/suggested-session-timeout", obpController.GetSuggestedSessionTimeout)
		v5.GET("/banks", obpController.GetBanks)
		v5.GET("/banks/:bankId", obpController.GetBank)
		
		bankGroup := v5.Group("/banks/:bankId")
		{
			bankGroup.POST("/agents", obpController.CreateAgent)
			bankGroup.GET("/agents", obpController.GetAgents)
			bankGroup.GET("/agents/:agentId", obpController.GetAgent)
			bankGroup.PUT("/agents/:agentId", obpController.UpdateAgentStatus)
			
			bankGroup.GET("/accounts", obpController.GetAccountsAtBank)
			bankGroup.POST("/accounts", obpController.CreateAccount)
			bankGroup.GET("/accounts/:accountId", obpController.GetAccount)
			bankGroup.GET("/accounts/:accountId/balances", obpController.GetAccountBalances)
			
			accountViewGroup := bankGroup.Group("/accounts/:accountId/:viewId")
			{
				accountViewGroup.GET("/transactions", obpController.GetTransactions)
				accountViewGroup.GET("/transactions/:transactionId", obpController.GetTransaction)
				accountViewGroup.POST("/transaction-request-types/SEPA/transaction-requests", obpController.CreateSEPATransactionRequest)
			}
			
			bankGroup.GET("/customers", obpController.GetCustomers)
			bankGroup.POST("/customers", obpController.CreateCustomer)
			
			bankGroup.GET("/my/consents", obpController.GetMyConsentsAtBank)
		}
		
		v5.POST("/consumer/consents", obpController.CreateConsent)
	}
	
	berlinGroup := router.Group("/berlin-group/v1.3")
	{
		berlinGroup.GET("/accounts", berlinGroupController.GetAccounts)
		berlinGroup.GET("/accounts/:account-id/balances", berlinGroupController.GetAccountBalances)
		berlinGroup.GET("/accounts/:account-id/transactions", berlinGroupController.GetAccountTransactions)
		berlinGroup.GET("/accounts/:account-id", berlinGroupController.GetAccountDetails)
		berlinGroup.POST("/payments/sepa-credit-transfers", berlinGroupController.InitiateSEPACreditTransfer)
		berlinGroup.POST("/consents", berlinGroupController.CreateAccountInformationConsent)
		berlinGroup.GET("/consents/:consentId", berlinGroupController.GetConsentInformation)
		berlinGroup.DELETE("/consents/:consentId", berlinGroupController.DeleteConsent)
		berlinGroup.GET("/consents/:consentId/status", berlinGroupController.GetConsentStatus)
	}
	
	ukOpenBanking := router.Group("/open-banking/v3.1.0")
	{
		aisp := ukOpenBanking.Group("/aisp")
		{
			aisp.GET("/accounts", ukOpenBankingController.GetAccounts)
			aisp.GET("/accounts/:AccountId/balances", ukOpenBankingController.GetAccountBalances)
			aisp.GET("/accounts/:AccountId/transactions", ukOpenBankingController.GetAccountTransactions)
		}
		
		pisp := ukOpenBanking.Group("/pisp")
		{
			pisp.POST("/domestic-payment-consents", ukOpenBankingController.CreateDomesticPaymentConsents)
			pisp.POST("/domestic-payments", ukOpenBankingController.CreateDomesticPayments)
			pisp.GET("/domestic-payment-consents/:ConsentId", ukOpenBankingController.GetDomesticPaymentConsent)
		}
		
		cbpii := ukOpenBanking.Group("/cbpii")
		{
			cbpii.POST("/funds-confirmation-consents", ukOpenBankingController.CreateFundsConfirmationConsents)
			cbpii.POST("/funds-confirmations", ukOpenBankingController.CreateFundsConfirmations)
		}
	}
	
	australianCDR := router.Group("/cds-au/v1.0.0")
	{
		banking := australianCDR.Group("/banking")
		{
			banking.GET("/accounts", australianCDRController.GetAccounts)
			banking.GET("/accounts/:accountId", australianCDRController.GetAccountDetail)
			banking.GET("/accounts/:accountId/balance", australianCDRController.GetAccountBalance)
			banking.GET("/accounts/:accountId/transactions", australianCDRController.GetAccountTransactions)
		}
		
		common := australianCDR.Group("/common")
		{
			common.GET("/customer", australianCDRController.GetCustomer)
			common.GET("/customer/detail", australianCDRController.GetCustomerDetail)
		}
	}
	
	bahrainOBF := router.Group("/bahrain-obf/v1.0.0")
	{
		bahrainOBF.GET("/accounts", bahrainOBFController.GetAccounts)
		bahrainOBF.GET("/accounts/:AccountId", bahrainOBFController.GetAccount)
		bahrainOBF.GET("/accounts/:AccountId/balances", bahrainOBFController.GetAccountBalances)
		bahrainOBF.POST("/domestic-payment-consents", bahrainOBFController.CreateDomesticPaymentConsent)
		bahrainOBF.POST("/domestic-payments", bahrainOBFController.CreateDomesticPayment)
	}
	
	polishAPI := router.Group("/polish-api/v2.1.1.1")
	{
		polishAPI.GET("/accounts", polishAPIController.GetAccounts)
		polishAPI.GET("/accounts/:account-id/balances", polishAPIController.GetAccountBalances)
		polishAPI.POST("/payments/domestic-credit-transfers", polishAPIController.InitiateDomesticCreditTransfer)
	}
	
	stetAPI := router.Group("/stet/v1.4")
	{
		stetAPI.GET("/accounts", stetAPIController.GetAccounts)
		stetAPI.GET("/accounts/:account-id/balances", stetAPIController.GetAccountBalances)
		stetAPI.POST("/payment-requests", stetAPIController.CreatePaymentRequest)
	}
	
	mxofAPI := router.Group("/mxof/v1.0.0")
	{
		mxofAPI.GET("/accounts", mxofAPIController.GetAccounts)
		mxofAPI.GET("/accounts/:account-id/balances", mxofAPIController.GetAccountBalances)
		mxofAPI.GET("/accounts/:account-id/transactions", mxofAPIController.GetAccountTransactions)
	}
	
	router.GET("/api/health", additionalController.GetHealthCheck)
	router.GET("/api/versions", additionalController.GetAPIVersions)
	
	return router
}
