package routes

import (
	"obp-api-backend/internal/controllers"
	"obp-api-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	obpController *controllers.OBPCoreController,
	obpV3Controller *controllers.OBPv3Controller,
	obpV4Controller *controllers.OBPv4Controller,
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
		v5.GET("/config", obpController.GetConfig)
		v5.GET("/adapter", obpController.GetAdapterInfo)
		v5.GET("/rate-limiting", obpController.GetRateLimitingInfo)
		
		v5.GET("/banks", obpController.GetBanks)
		v5.GET("/banks/:bankId", obpController.GetBank)
		v5.POST("/banks", obpController.CreateBank)
		v5.PUT("/banks/:bankId", obpController.UpdateBank)
		v5.DELETE("/banks/:bankId", obpController.DeleteBank)
		
		v5.GET("/banks/:bankId/accounts", obpController.GetAccounts)
		v5.GET("/banks/:bankId/accounts/:accountId/:viewId/account", obpController.GetAccount)
		v5.POST("/banks/:bankId/accounts", obpController.CreateAccount)
		v5.PUT("/banks/:bankId/accounts/:accountId", obpController.UpdateAccount)
		v5.DELETE("/banks/:bankId/accounts/:accountId", obpController.DeleteAccount)
		
		v5.GET("/banks/:bankId/accounts/:accountId/:viewId/transactions", obpController.GetTransactions)
		v5.GET("/banks/:bankId/accounts/:accountId/:viewId/transactions/:transactionId/transaction", obpController.GetTransaction)
		v5.POST("/banks/:bankId/accounts/:accountId/:viewId/transactions", obpController.CreateTransaction)
		
		v5.GET("/banks/:bankId/customers", obpController.GetCustomers)
		v5.GET("/banks/:bankId/customers/:customerId", obpController.GetCustomer)
		v5.POST("/banks/:bankId/customers", obpController.CreateCustomer)
		v5.PUT("/banks/:bankId/customers/:customerId", obpController.UpdateCustomer)
		v5.DELETE("/banks/:bankId/customers/:customerId", obpController.DeleteCustomer)
		
		v5.GET("/users", obpController.GetUsers)
		v5.GET("/users/:userId", obpController.GetUser)
		v5.POST("/users", obpController.CreateUser)
		v5.PUT("/users/:userId", obpController.UpdateUser)
		v5.DELETE("/users/:userId", obpController.DeleteUser)
		
		v5.GET("/consents", obpController.GetConsents)
		v5.GET("/consents/:consentId", obpController.GetConsent)
		v5.POST("/consents", obpController.CreateConsent)
		v5.PUT("/consents/:consentId", obpController.UpdateConsent)
		v5.DELETE("/consents/:consentId", obpController.DeleteConsent)
		
		v5.GET("/banks/:bankId/accounts/:accountId/:viewId/balances", obpController.GetAccountBalances)
		v5.POST("/banks/:bankId/accounts/:accountId/views/:viewId/target-views", obpController.CreateCustomView)
		v5.PUT("/banks/:bankId/accounts/:accountId/views/:viewId/target-views/:targetViewId", obpController.UpdateCustomView)
		v5.GET("/banks/:bankId/accounts/:accountId/views/:viewId/target-views/:targetViewId", obpController.GetCustomView)
		v5.DELETE("/banks/:bankId/accounts/:accountId/views/:viewId/target-views/:targetViewId", obpController.DeleteCustomView)
		
		v5.POST("/consumer/vrp-consent-requests", obpController.CreateVRPConsentRequest)
		
		v5.POST("/regulated-entities/:regulatedEntityId/attributes", obpController.CreateRegulatedEntityAttribute)
		v5.DELETE("/regulated-entities/:regulatedEntityId/attributes/:regulatedEntityAttributeId", obpController.DeleteRegulatedEntityAttribute)
	}
	
	v4 := router.Group("/obp/v4.0.0")
	{
		v4.GET("/root", obpV4Controller.GetAPIInfo)
		v4.GET("/database/info", obpV4Controller.GetDatabaseInfo)
		v4.GET("/users/current/logout-link", obpV4Controller.GetLogoutLink)
		
		v4.GET("/management/system-dynamic-entities", obpV4Controller.GetSystemDynamicEntities)
		v4.POST("/management/system-dynamic-entities", obpV4Controller.CreateSystemDynamicEntity)
		v4.PUT("/management/system-dynamic-entities/:entityId", obpV4Controller.UpdateSystemDynamicEntity)
		v4.DELETE("/management/system-dynamic-entities/:entityId", obpV4Controller.DeleteSystemDynamicEntity)
		
		v4.GET("/banks/:bankId/dynamic-entities", obpV4Controller.GetBankDynamicEntities)
		v4.POST("/banks/:bankId/dynamic-entities", obpV4Controller.CreateBankDynamicEntity)
		v4.PUT("/banks/:bankId/dynamic-entities/:entityId", obpV4Controller.UpdateBankDynamicEntity)
		v4.DELETE("/banks/:bankId/dynamic-entities/:entityId", obpV4Controller.DeleteBankDynamicEntity)
		
		v4.GET("/my/dynamic-entities", obpV4Controller.GetMyDynamicEntities)
		v4.PUT("/my/dynamic-entities/:entityId", obpV4Controller.UpdateMyDynamicEntity)
		v4.DELETE("/my/dynamic-entities/:entityId", obpV4Controller.DeleteMyDynamicEntity)
		
		v4.POST("/banks/:bankId/accounts/:accountId/:viewId/transaction-request-types/ACCOUNT/transaction-requests", obpV4Controller.CreateAccountTransactionRequest)
		v4.POST("/banks/:bankId/accounts/:accountId/:viewId/transaction-request-types/ACCOUNT_OTP/transaction-requests", obpV4Controller.CreateAccountOTPTransactionRequest)
		v4.POST("/banks/:bankId/accounts/:accountId/:viewId/transaction-request-types/COUNTERPARTY/transaction-requests", obpV4Controller.CreateCounterpartyTransactionRequest)
		v4.POST("/banks/:bankId/accounts/:accountId/:viewId/transaction-request-types/SEPA/transaction-requests", obpV4Controller.CreateSEPATransactionRequest)
		v4.POST("/banks/:bankId/accounts/:accountId/:viewId/transaction-request-types/SIMPLE/transaction-requests", obpV4Controller.CreateSimpleTransactionRequest)
		v4.POST("/banks/:bankId/accounts/:accountId/:viewId/transaction-request-types/FREE_FORM/transaction-requests", obpV4Controller.CreateFreeFormTransactionRequest)
		
		v4.GET("/banks/:bankId/accounts/:accountId/:viewId/transaction-request-types/ACCOUNT/transaction-requests/:transactionRequestId", obpV4Controller.GetAccountTransactionRequest)
		v4.GET("/banks/:bankId/accounts/:accountId/:viewId/transaction-request-types/ACCOUNT_OTP/transaction-requests/:transactionRequestId", obpV4Controller.GetAccountOTPTransactionRequest)
		v4.GET("/banks/:bankId/accounts/:accountId/:viewId/transaction-request-types/COUNTERPARTY/transaction-requests/:transactionRequestId", obpV4Controller.GetCounterpartyTransactionRequest)
		v4.GET("/banks/:bankId/accounts/:accountId/:viewId/transaction-request-types/SEPA/transaction-requests/:transactionRequestId", obpV4Controller.GetSEPATransactionRequest)
		v4.GET("/banks/:bankId/accounts/:accountId/:viewId/transaction-request-types/SIMPLE/transaction-requests/:transactionRequestId", obpV4Controller.GetSimpleTransactionRequest)
		v4.GET("/banks/:bankId/accounts/:accountId/:viewId/transaction-request-types/FREE_FORM/transaction-requests/:transactionRequestId", obpV4Controller.GetFreeFormTransactionRequest)
		
		v4.POST("/banks/:bankId/accounts/:accountId/:viewId/transaction-request-types/ACCOUNT/transaction-requests/:transactionRequestId/challenge", obpV4Controller.AnswerAccountTransactionRequestChallenge)
		v4.POST("/banks/:bankId/accounts/:accountId/:viewId/transaction-request-types/ACCOUNT_OTP/transaction-requests/:transactionRequestId/challenge", obpV4Controller.AnswerAccountOTPTransactionRequestChallenge)
		v4.POST("/banks/:bankId/accounts/:accountId/:viewId/transaction-request-types/COUNTERPARTY/transaction-requests/:transactionRequestId/challenge", obpV4Controller.AnswerCounterpartyTransactionRequestChallenge)
		v4.POST("/banks/:bankId/accounts/:accountId/:viewId/transaction-request-types/SEPA/transaction-requests/:transactionRequestId/challenge", obpV4Controller.AnswerSEPATransactionRequestChallenge)
		v4.POST("/banks/:bankId/accounts/:accountId/:viewId/transaction-request-types/SIMPLE/transaction-requests/:transactionRequestId/challenge", obpV4Controller.AnswerSimpleTransactionRequestChallenge)
		v4.POST("/banks/:bankId/accounts/:accountId/:viewId/transaction-request-types/FREE_FORM/transaction-requests/:transactionRequestId/challenge", obpV4Controller.AnswerFreeFormTransactionRequestChallenge)
		
		v4.GET("/banks/:bankId/accounts/:accountId/:viewId/transaction-request-types", obpV4Controller.GetTransactionRequestTypes)
		v4.GET("/banks/:bankId/accounts/:accountId/:viewId/transaction-requests", obpV4Controller.GetTransactionRequests)
		
		v4.POST("/banks/:bankId/transaction-request-attribute-definitions", obpV4Controller.CreateTransactionRequestAttributeDefinition)
		v4.GET("/banks/:bankId/transaction-request-attribute-definitions", obpV4Controller.GetTransactionRequestAttributeDefinitions)
		v4.PUT("/banks/:bankId/transaction-request-attribute-definitions/:attributeDefinitionId", obpV4Controller.UpdateTransactionRequestAttributeDefinition)
		v4.DELETE("/banks/:bankId/transaction-request-attribute-definitions/:attributeDefinitionId", obpV4Controller.DeleteTransactionRequestAttributeDefinition)
		
		v4.POST("/banks/:bankId/accounts/:accountId/direct-debit-management", obpV4Controller.CreateDirectDebitManagement)
		v4.POST("/banks/:bankId/accounts/:accountId/standing-order-management", obpV4Controller.CreateStandingOrderManagement)
		
		v4.POST("/banks/:bankId/accounts/:accountId/:viewId/account-access/revoke", obpV4Controller.RevokeGrantUserAccessToViews)
		
		v4.POST("/banks/:bankId/accounts/:accountId/:viewId/tags", obpV4Controller.AddTagForViewOnAccount)
		v4.DELETE("/banks/:bankId/accounts/:accountId/:viewId/tags/:tagId", obpV4Controller.DeleteTagForViewOnAccount)
		v4.GET("/banks/:bankId/accounts/:accountId/:viewId/tags", obpV4Controller.GetTagsForViewOnAccount)
		
		v4.POST("/iban-checker", obpV4Controller.IBANChecker)
		v4.GET("/call-context", obpV4Controller.GetCallContext)
		v4.POST("/verify-request-sign-response", obpV4Controller.VerifyRequestSignResponse)
	}
	
	v3 := router.Group("/obp/v3.1.0")
	{
		v3.GET("/root", obpV3Controller.GetAPIInfo)
		v3.GET("/config", obpV3Controller.GetConfig)
		v3.GET("/adapter", obpV3Controller.GetAdapterInfo)
		v3.GET("/rate-limiting", obpV3Controller.GetRateLimitingInfo)
		
		v3.POST("/banks/:bankId/accounts/:accountId/:viewId/webhooks/account", obpV3Controller.CreateAccountWebhook)
		v3.GET("/banks/:bankId/accounts/:accountId/:viewId/webhooks/account", obpV3Controller.GetAccountWebhooks)
		v3.PUT("/banks/:bankId/accounts/:accountId/:viewId/webhooks/account/:webhookId", obpV3Controller.UpdateAccountWebhook)
		v3.DELETE("/banks/:bankId/accounts/:accountId/:viewId/webhooks/account/:webhookId", obpV3Controller.DeleteAccountWebhook)
		
		v3.POST("/banks/:bankId/products", obpV3Controller.CreateProduct)
		v3.GET("/banks/:bankId/products", obpV3Controller.GetProducts)
		v3.GET("/banks/:bankId/products/:productCode", obpV3Controller.GetProduct)
		v3.PUT("/banks/:bankId/products/:productCode", obpV3Controller.UpdateProduct)
		v3.DELETE("/banks/:bankId/products/:productCode", obpV3Controller.DeleteProduct)
		v3.GET("/banks/:bankId/products/:productCode/tree", obpV3Controller.GetProductTree)
		
		v3.POST("/banks/:bankId/customers/:customerId/attributes", obpV3Controller.CreateCustomerAttribute)
		v3.GET("/banks/:bankId/customers/:customerId/attributes", obpV3Controller.GetCustomerAttributes)
		v3.PUT("/banks/:bankId/customers/:customerId/attributes/:attributeId", obpV3Controller.UpdateCustomerAttribute)
		v3.DELETE("/banks/:bankId/customers/:customerId/attributes/:attributeId", obpV3Controller.DeleteCustomerAttribute)
		
		v3.POST("/banks/:bankId/meetings", obpV3Controller.CreateMeeting)
		v3.GET("/banks/:bankId/meetings", obpV3Controller.GetMeetings)
		v3.GET("/banks/:bankId/meetings/:meetingId", obpV3Controller.GetMeeting)
		v3.PUT("/banks/:bankId/meetings/:meetingId", obpV3Controller.UpdateMeeting)
		v3.DELETE("/banks/:bankId/meetings/:meetingId", obpV3Controller.DeleteMeeting)
		
		v3.POST("/banks/:bankId/customers/:customerId/addresses", obpV3Controller.CreateCustomerAddress)
		v3.GET("/banks/:bankId/customers/:customerId/addresses", obpV3Controller.GetCustomerAddresses)
		v3.PUT("/banks/:bankId/customers/:customerId/addresses/:addressId", obpV3Controller.UpdateCustomerAddress)
		v3.DELETE("/banks/:bankId/customers/:customerId/addresses/:addressId", obpV3Controller.DeleteCustomerAddress)
		
		v3.POST("/system-views", obpV3Controller.CreateSystemView)
		v3.GET("/system-views/:viewId", obpV3Controller.GetSystemView)
		v3.PUT("/system-views/:viewId", obpV3Controller.UpdateSystemView)
		v3.DELETE("/system-views/:viewId", obpV3Controller.DeleteSystemView)
		
		v3.POST("/banks/:bankId/accounts/:accountId/:viewId/transactions/:transactionId/metadata/comments", obpV3Controller.CreateTransactionComment)
		v3.GET("/banks/:bankId/accounts/:accountId/:viewId/transactions/:transactionId/metadata/comments", obpV3Controller.GetTransactionComments)
		v3.DELETE("/banks/:bankId/accounts/:accountId/:viewId/transactions/:transactionId/metadata/comments/:commentId", obpV3Controller.DeleteTransactionComment)
		
		v3.POST("/banks/:bankId/accounts/:accountId/:viewId/transactions/:transactionId/metadata/tags", obpV3Controller.CreateTransactionTag)
		v3.GET("/banks/:bankId/accounts/:accountId/:viewId/transactions/:transactionId/metadata/tags", obpV3Controller.GetTransactionTags)
		v3.DELETE("/banks/:bankId/accounts/:accountId/:viewId/transactions/:transactionId/metadata/tags/:tagId", obpV3Controller.DeleteTransactionTag)
		
		v3.POST("/banks/:bankId/accounts/:accountId/:viewId/transactions/:transactionId/metadata/images", obpV3Controller.CreateTransactionImage)
		v3.GET("/banks/:bankId/accounts/:accountId/:viewId/transactions/:transactionId/metadata/images", obpV3Controller.GetTransactionImages)
		v3.DELETE("/banks/:bankId/accounts/:accountId/:viewId/transactions/:transactionId/metadata/images/:imageId", obpV3Controller.DeleteTransactionImage)
		
		v3.POST("/banks/:bankId/accounts/:accountId/:viewId/transactions/:transactionId/metadata/where", obpV3Controller.CreateTransactionWhere)
		v3.GET("/banks/:bankId/accounts/:accountId/:viewId/transactions/:transactionId/metadata/where", obpV3Controller.GetTransactionWhere)
		v3.PUT("/banks/:bankId/accounts/:accountId/:viewId/transactions/:transactionId/metadata/where/:whereId", obpV3Controller.UpdateTransactionWhere)
		v3.DELETE("/banks/:bankId/accounts/:accountId/:viewId/transactions/:transactionId/metadata/where/:whereId", obpV3Controller.DeleteTransactionWhere)
		
		v3.POST("/banks/:bankId/accounts/:accountId/:viewId/other_accounts/:otherAccountId/metadata/more_info", obpV3Controller.CreateOtherAccountMoreInfo)
		v3.GET("/banks/:bankId/accounts/:accountId/:viewId/other_accounts/:otherAccountId/metadata/more_info", obpV3Controller.GetOtherAccountMoreInfo)
		v3.PUT("/banks/:bankId/accounts/:accountId/:viewId/other_accounts/:otherAccountId/metadata/more_info", obpV3Controller.UpdateOtherAccountMoreInfo)
		v3.DELETE("/banks/:bankId/accounts/:accountId/:viewId/other_accounts/:otherAccountId/metadata/more_info", obpV3Controller.DeleteOtherAccountMoreInfo)
	}
	
	berlinGroup := router.Group("/berlin-group/v1.3")
	{
		berlinGroup.GET("/accounts", berlinGroupController.GetAccounts)
		berlinGroup.GET("/accounts/:account-id/balances", berlinGroupController.GetAccountBalances)
		berlinGroup.GET("/accounts/:account-id/transactions", berlinGroupController.GetAccountTransactions)
		berlinGroup.GET("/accounts/:account-id", berlinGroupController.GetAccountDetails)
		berlinGroup.POST("/payments/sepa-credit-transfers", berlinGroupController.InitiateSEPACreditTransfer)
		berlinGroup.POST("/consents", berlinGroupController.CreateConsent)
		berlinGroup.GET("/consents/:consentId", berlinGroupController.GetConsent)
		berlinGroup.DELETE("/consents/:consentId", berlinGroupController.DeleteConsent)
		berlinGroup.GET("/consents/:consentId/status", berlinGroupController.GetConsentStatus)
	}
	
	ukOpenBanking := router.Group("/open-banking/v3.1.0")
	{
		aisp := ukOpenBanking.Group("/aisp")
		{
			aisp.GET("/accounts", ukOpenBankingController.GetAccounts)
			aisp.GET("/accounts/:AccountId", ukOpenBankingController.GetAccount)
			aisp.GET("/accounts/:AccountId/balances", ukOpenBankingController.GetAccountBalances)
			aisp.GET("/accounts/:AccountId/transactions", ukOpenBankingController.GetAccountTransactions)
			aisp.GET("/accounts/:AccountId/statements", ukOpenBankingController.GetAccountStatements)
			aisp.GET("/accounts/:AccountId/statements/:StatementId", ukOpenBankingController.GetAccountStatement)
			aisp.GET("/accounts/:AccountId/statements/:StatementId/file", ukOpenBankingController.GetAccountStatementFile)
			aisp.GET("/accounts/:AccountId/statements/:StatementId/transactions", ukOpenBankingController.GetAccountStatementTransactions)
			aisp.GET("/accounts/:AccountId/standing-orders", ukOpenBankingController.GetAccountStandingOrders)
			aisp.GET("/accounts/:AccountId/direct-debits", ukOpenBankingController.GetAccountDirectDebits)
			aisp.GET("/accounts/:AccountId/beneficiaries", ukOpenBankingController.GetAccountBeneficiaries)
			aisp.GET("/accounts/:AccountId/scheduled-payments", ukOpenBankingController.GetAccountScheduledPayments)
			aisp.GET("/accounts/:AccountId/offers", ukOpenBankingController.GetAccountOffers)
			aisp.GET("/accounts/:AccountId/party", ukOpenBankingController.GetAccountParty)
			aisp.GET("/accounts/:AccountId/product", ukOpenBankingController.GetAccountProduct)
			aisp.GET("/balances", ukOpenBankingController.GetBalances)
			aisp.GET("/beneficiaries", ukOpenBankingController.GetBeneficiaries)
			aisp.GET("/direct-debits", ukOpenBankingController.GetDirectDebits)
			aisp.GET("/offers", ukOpenBankingController.GetOffers)
			aisp.GET("/party", ukOpenBankingController.GetParty)
			aisp.GET("/products", ukOpenBankingController.GetProducts)
			aisp.GET("/scheduled-payments", ukOpenBankingController.GetScheduledPayments)
			aisp.GET("/standing-orders", ukOpenBankingController.GetStandingOrders)
			aisp.GET("/statements", ukOpenBankingController.GetStatements)
			aisp.GET("/transactions", ukOpenBankingController.GetTransactions)
			aisp.GET("/accounts/:AccountId/transactions/:StatementId", ukOpenBankingController.GetAccountTransactionsByStatementId)
			aisp.GET("/transactions/:StatementId", ukOpenBankingController.GetTransactionsByStatementId)
		}
		
		pisp := ukOpenBanking.Group("/pisp")
		{
			pisp.POST("/domestic-payment-consents", ukOpenBankingController.CreateDomesticPaymentConsents)
			pisp.POST("/domestic-payments", ukOpenBankingController.CreateDomesticPayments)
			pisp.GET("/domestic-payment-consents/:ConsentId", ukOpenBankingController.GetDomesticPaymentConsent)
			pisp.GET("/domestic-payments/:DomesticPaymentId", ukOpenBankingController.GetDomesticPayment)
			pisp.POST("/domestic-scheduled-payment-consents", ukOpenBankingController.CreateDomesticScheduledPaymentConsents)
			pisp.POST("/domestic-scheduled-payments", ukOpenBankingController.CreateDomesticScheduledPayments)
			pisp.GET("/domestic-scheduled-payment-consents/:ConsentId", ukOpenBankingController.GetDomesticScheduledPaymentConsent)
			pisp.GET("/domestic-scheduled-payments/:DomesticScheduledPaymentId", ukOpenBankingController.GetDomesticScheduledPayment)
			pisp.POST("/domestic-standing-order-consents", ukOpenBankingController.CreateDomesticStandingOrderConsents)
			pisp.POST("/domestic-standing-orders", ukOpenBankingController.CreateDomesticStandingOrders)
			pisp.GET("/domestic-standing-order-consents/:ConsentId", ukOpenBankingController.GetDomesticStandingOrderConsent)
			pisp.GET("/domestic-standing-orders/:DomesticStandingOrderId", ukOpenBankingController.GetDomesticStandingOrder)
			pisp.POST("/international-payment-consents", ukOpenBankingController.CreateInternationalPaymentConsents)
			pisp.POST("/international-payments", ukOpenBankingController.CreateInternationalPayments)
			pisp.GET("/international-payment-consents/:ConsentId", ukOpenBankingController.GetInternationalPaymentConsent)
			pisp.GET("/international-payments/:InternationalPaymentId", ukOpenBankingController.GetInternationalPayment)
			pisp.POST("/international-scheduled-payment-consents", ukOpenBankingController.CreateInternationalScheduledPaymentConsents)
			pisp.POST("/international-scheduled-payments", ukOpenBankingController.CreateInternationalScheduledPayments)
			pisp.GET("/international-scheduled-payment-consents/:ConsentId", ukOpenBankingController.GetInternationalScheduledPaymentConsent)
			pisp.GET("/international-scheduled-payments/:InternationalScheduledPaymentId", ukOpenBankingController.GetInternationalScheduledPayment)
			pisp.POST("/international-standing-order-consents", ukOpenBankingController.CreateInternationalStandingOrderConsents)
			pisp.POST("/international-standing-orders", ukOpenBankingController.CreateInternationalStandingOrders)
			pisp.GET("/international-standing-order-consents/:ConsentId", ukOpenBankingController.GetInternationalStandingOrderConsent)
			pisp.GET("/international-standing-orders/:InternationalStandingOrderId", ukOpenBankingController.GetInternationalStandingOrder)
			pisp.POST("/file-payment-consents", ukOpenBankingController.CreateFilePaymentConsents)
			pisp.POST("/file-payments", ukOpenBankingController.CreateFilePayments)
			pisp.GET("/file-payment-consents/:ConsentId", ukOpenBankingController.GetFilePaymentConsent)
			pisp.GET("/file-payments/:FilePaymentId", ukOpenBankingController.GetFilePayment)
			pisp.GET("/file-payment-consents/:ConsentId/file", ukOpenBankingController.GetFilePaymentConsentFile)
			pisp.POST("/file-payment-consents/:ConsentId/file", ukOpenBankingController.CreateFilePaymentConsentFile)
		}
		
		cbpii := ukOpenBanking.Group("/cbpii")
		{
			cbpii.POST("/funds-confirmation-consents", ukOpenBankingController.CreateFundsConfirmationConsents)
			cbpii.POST("/funds-confirmations", ukOpenBankingController.CreateFundsConfirmations)
			cbpii.GET("/funds-confirmation-consents/:ConsentId", ukOpenBankingController.GetFundsConfirmationConsent)
			cbpii.DELETE("/funds-confirmation-consents/:ConsentId", ukOpenBankingController.DeleteFundsConfirmationConsent)
		}
		
		eventNotifications := ukOpenBanking.Group("/event-notifications")
		{
			eventNotifications.POST("/", ukOpenBankingController.CreateEventNotification)
			eventNotifications.GET("/", ukOpenBankingController.GetEventNotifications)
		}
		
		vrp := ukOpenBanking.Group("/domestic-vrp-consents")
		{
			vrp.POST("/", ukOpenBankingController.CreateDomesticVRPConsents)
			vrp.GET("/:ConsentId", ukOpenBankingController.GetDomesticVRPConsent)
			vrp.DELETE("/:ConsentId", ukOpenBankingController.DeleteDomesticVRPConsent)
			vrp.POST("/:ConsentId/domestic-vrps", ukOpenBankingController.CreateDomesticVRP)
			vrp.GET("/:ConsentId/domestic-vrps/:DomesticVRPId", ukOpenBankingController.GetDomesticVRP)
		}
	}
	
	australianCDR := router.Group("/cds-au/v1.0.0")
	{
		australianCDR.GET("/banking/accounts", australianCDRController.GetAccounts)
		australianCDR.GET("/banking/accounts/:accountId", australianCDRController.GetAccountDetail)
		australianCDR.GET("/banking/accounts/:accountId/balance", australianCDRController.GetAccountBalance)
		australianCDR.GET("/banking/accounts/:accountId/transactions", australianCDRController.GetAccountTransactions)
		australianCDR.GET("/banking/accounts/:accountId/transactions/:transactionId", australianCDRController.GetAccountTransactionDetail)
		australianCDR.GET("/banking/accounts/:accountId/direct-debits", australianCDRController.GetAccountDirectDebits)
		australianCDR.GET("/banking/accounts/:accountId/scheduled-payments", australianCDRController.GetAccountScheduledPayments)
		australianCDR.GET("/banking/products", australianCDRController.GetProducts)
		australianCDR.GET("/banking/products/:productId", australianCDRController.GetProductDetail)
		australianCDR.GET("/common/customer", australianCDRController.GetCustomer)
		australianCDR.GET("/common/customer/detail", australianCDRController.GetCustomerDetail)
	}
	
	bahrainOBF := router.Group("/bahrain-obf/v1.0.0")
	{
		bahrainOBF.GET("/accounts", bahrainOBFController.GetAccounts)
		bahrainOBF.GET("/accounts/:AccountId", bahrainOBFController.GetAccount)
		bahrainOBF.GET("/accounts/:AccountId/balances", bahrainOBFController.GetAccountBalances)
		bahrainOBF.GET("/accounts/:AccountId/transactions", bahrainOBFController.GetAccountTransactions)
		bahrainOBF.GET("/accounts/:AccountId/statements", bahrainOBFController.GetAccountStatements)
		bahrainOBF.GET("/accounts/:AccountId/statements/:StatementId", bahrainOBFController.GetAccountStatement)
		bahrainOBF.GET("/accounts/:AccountId/standing-orders", bahrainOBFController.GetAccountStandingOrders)
		bahrainOBF.GET("/accounts/:AccountId/direct-debits", bahrainOBFController.GetAccountDirectDebits)
		bahrainOBF.GET("/accounts/:AccountId/beneficiaries", bahrainOBFController.GetAccountBeneficiaries)
		bahrainOBF.GET("/accounts/:AccountId/offers", bahrainOBFController.GetAccountOffers)
		bahrainOBF.GET("/accounts/:AccountId/party", bahrainOBFController.GetAccountParty)
		bahrainOBF.GET("/accounts/:AccountId/product", bahrainOBFController.GetAccountProduct)
		
		bahrainOBF.POST("/domestic-payment-consents", bahrainOBFController.CreateDomesticPaymentConsents)
		bahrainOBF.POST("/domestic-payments", bahrainOBFController.CreateDomesticPayments)
		bahrainOBF.GET("/domestic-payment-consents/:ConsentId", bahrainOBFController.GetDomesticPaymentConsent)
		bahrainOBF.GET("/domestic-payments/:DomesticPaymentId", bahrainOBFController.GetDomesticPayment)
		bahrainOBF.POST("/international-payment-consents", bahrainOBFController.CreateInternationalPaymentConsents)
		bahrainOBF.POST("/international-payments", bahrainOBFController.CreateInternationalPayments)
		bahrainOBF.GET("/international-payment-consents/:ConsentId", bahrainOBFController.GetInternationalPaymentConsent)
		bahrainOBF.GET("/international-payments/:InternationalPaymentId", bahrainOBFController.GetInternationalPayment)
		bahrainOBF.POST("/file-payment-consents", bahrainOBFController.CreateFilePaymentConsents)
		bahrainOBF.POST("/file-payments", bahrainOBFController.CreateFilePayments)
		bahrainOBF.GET("/file-payment-consents/:ConsentId", bahrainOBFController.GetFilePaymentConsent)
		bahrainOBF.GET("/file-payments/:FilePaymentId", bahrainOBFController.GetFilePayment)
		
		bahrainOBF.POST("/domestic-future-dated-payment-consents", bahrainOBFController.CreateDomesticFutureDatedPaymentConsents)
		bahrainOBF.POST("/domestic-future-dated-payments", bahrainOBFController.CreateDomesticFutureDatedPayments)
		bahrainOBF.GET("/domestic-future-dated-payment-consents/:ConsentId", bahrainOBFController.GetDomesticFutureDatedPaymentConsent)
		bahrainOBF.GET("/domestic-future-dated-payments/:DomesticFutureDatedPaymentId", bahrainOBFController.GetDomesticFutureDatedPayment)
		
		bahrainOBF.POST("/account-access-consents", bahrainOBFController.CreateAccountAccessConsents)
		bahrainOBF.GET("/account-access-consents/:ConsentId", bahrainOBFController.GetAccountAccessConsent)
		bahrainOBF.DELETE("/account-access-consents/:ConsentId", bahrainOBFController.DeleteAccountAccessConsent)
	}
	
	polishAPI := router.Group("/polish-api/v2.1.1.1")
	{
		polishAPI.GET("/accounts", polishAPIController.GetAccounts)
		polishAPI.GET("/accounts/:accountId", polishAPIController.GetAccount)
		polishAPI.GET("/accounts/:accountId/balances", polishAPIController.GetAccountBalances)
		polishAPI.GET("/accounts/:accountId/transactions", polishAPIController.GetAccountTransactions)
		polishAPI.POST("/payments/domestic", polishAPIController.CreateDomesticPayment)
		polishAPI.GET("/payments/domestic/:paymentId", polishAPIController.GetDomesticPayment)
		polishAPI.POST("/consents", polishAPIController.CreateConsent)
		polishAPI.GET("/consents/:consentId", polishAPIController.GetConsent)
		polishAPI.DELETE("/consents/:consentId", polishAPIController.DeleteConsent)
		polishAPI.GET("/consents/:consentId/status", polishAPIController.GetConsentStatus)
	}
	
	stetAPI := router.Group("/stet/v1.4")
	{
		stetAPI.GET("/accounts", stetAPIController.GetAccounts)
		stetAPI.GET("/accounts/:accountId", stetAPIController.GetAccount)
		stetAPI.GET("/accounts/:accountId/balances", stetAPIController.GetAccountBalances)
		stetAPI.GET("/accounts/:accountId/transactions", stetAPIController.GetAccountTransactions)
		stetAPI.POST("/payment-requests", stetAPIController.CreatePaymentRequest)
		stetAPI.GET("/payment-requests/:paymentRequestId", stetAPIController.GetPaymentRequest)
		stetAPI.POST("/consents", stetAPIController.CreateConsent)
		stetAPI.GET("/consents/:consentId", stetAPIController.GetConsent)
		stetAPI.DELETE("/consents/:consentId", stetAPIController.DeleteConsent)
		stetAPI.GET("/consents/:consentId/status", stetAPIController.GetConsentStatus)
	}
	
	mxofAPI := router.Group("/mxof/v1.0.0")
	{
		mxofAPI.GET("/accounts", mxofAPIController.GetAccounts)
		mxofAPI.GET("/accounts/:accountId", mxofAPIController.GetAccount)
		mxofAPI.GET("/accounts/:accountId/balances", mxofAPIController.GetAccountBalances)
		mxofAPI.GET("/accounts/:accountId/transactions", mxofAPIController.GetAccountTransactions)
		mxofAPI.POST("/payments", mxofAPIController.CreatePayment)
		mxofAPI.GET("/payments/:paymentId", mxofAPIController.GetPayment)
		mxofAPI.POST("/consents", mxofAPIController.CreateConsent)
		mxofAPI.GET("/consents/:consentId", mxofAPIController.GetConsent)
		mxofAPI.DELETE("/consents/:consentId", mxofAPIController.DeleteConsent)
		mxofAPI.GET("/consents/:consentId/status", mxofAPIController.GetConsentStatus)
	}
	
	return router
}
