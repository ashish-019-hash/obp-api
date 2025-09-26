package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ashish-019-hash/obp-api-backend/internal/controllers"
	"github.com/ashish-019-hash/obp-api-backend/internal/middleware"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
)

func SetupV510Routes(router *gin.Engine, orchestrationService *services.OrchestrationService, authMiddleware *middleware.AuthMiddleware) {
	v510Controller := controllers.NewV510Controller(orchestrationService)
	agentController := controllers.NewAgentController(orchestrationService)
	consentController := controllers.NewConsentController(orchestrationService)
	atmController := controllers.NewAtmController(orchestrationService)
	counterpartyController := controllers.NewCounterpartyController(orchestrationService)
	userController := controllers.NewUserController(orchestrationService)
	metricsController := controllers.NewMetricsController(orchestrationService)
	customerController := controllers.NewCustomerController(orchestrationService)
	atmManagementController := controllers.NewAtmManagementController(orchestrationService)
	consumerController := controllers.NewConsumerController(orchestrationService)
	accountAccessController := controllers.NewAccountAccessController(orchestrationService)
	transactionRequestController := controllers.NewTransactionRequestController(orchestrationService)
	balanceController := controllers.NewBalanceController(orchestrationService)
	customViewController := controllers.NewCustomViewController(orchestrationService)
	vrpController := controllers.NewVRPController(orchestrationService)
	regulatedEntityAttributeController := controllers.NewRegulatedEntityAttributeController(orchestrationService)
	webuiController := controllers.NewWebUIController(orchestrationService)
	systemViewController := controllers.NewSystemViewController(orchestrationService)
	apiCollectionController := controllers.NewApiCollectionController(orchestrationService)
	accountController := controllers.NewAccountController(orchestrationService)
	counterpartyLimitController := controllers.NewCounterpartyLimitController(orchestrationService)
	userAttributeController := controllers.NewUserAttributeController(orchestrationService)
	integrityCheckController := controllers.NewIntegrityCheckController(orchestrationService)
	currencyController := controllers.NewCurrencyController(orchestrationService)
	entitlementController := controllers.NewEntitlementController(orchestrationService)
	certificateController := controllers.NewCertificateController(orchestrationService)
	bankController := controllers.NewBankController(orchestrationService)

	v510 := router.Group("/obp/v5.1.0")

	v510.GET("/root", v510Controller.GetRoot)
	v510.GET("/ui/suggested-session-timeout", v510Controller.GetSuggestedSessionTimeout)
	v510.GET("/well-known", v510Controller.GetWellKnown)
	v510.GET("/waiting-for-godot", v510Controller.WaitingForGodot)

	protected := v510.Group("")
	protected.Use(authMiddleware.MultiAuth())
	{
		protected.GET("/tags", v510Controller.GetApiTags)
	}

	protected.POST("/banks", bankController.CreateBank)
	protected.GET("/banks", bankController.GetBanks)
	protected.GET("/banks/:bankId", bankController.GetBankById)

	protected.POST("/users", userController.CreateUser)
	protected.GET("/users", userController.GetUsers)

	protected.GET("/banks/:bankId/accounts/:accountId/views/:viewId", accountController.GetCoreAccountByIdThroughView)

	management := v510.Group("/management")
	management.Use(authMiddleware.MultiAuth())
	management.Use(authMiddleware.RequireEntitlement("CanGetApiCollections"))
	{
		management.GET("/api-collections", apiCollectionController.GetApiCollections)
		management.POST("/api-collections", apiCollectionController.CreateApiCollection)
	}

	my := v510.Group("/my")
	my.Use(authMiddleware.MultiAuth())
	{
		my.GET("/api-collections", apiCollectionController.GetMyApiCollections)
		my.PUT("/api-collections/:collectionId", apiCollectionController.UpdateMyApiCollection)
	}

	protected.GET("/regulated-entities", v510Controller.GetRegulatedEntities)
	protected.GET("/regulated-entities/:regulatedEntityId", v510Controller.GetRegulatedEntityById)
	protected.POST("/regulated-entities", v510Controller.CreateRegulatedEntity)
	protected.DELETE("/regulated-entities/:regulatedEntityId", v510Controller.DeleteRegulatedEntity)

	protected.POST("/banks/:bankId/agents", agentController.CreateAgent)
	protected.PUT("/banks/:bankId/agents/:agentId", agentController.UpdateAgentStatus)
	protected.GET("/banks/:bankId/agents/:agentId", agentController.GetAgent)
	protected.GET("/banks/:bankId/agents", agentController.GetAgents)

	protected.POST("/users/:userId/non-personal/attributes", userAttributeController.CreateNonPersonalUserAttribute)
	protected.DELETE("/users/:userId/non-personal/attributes/:userAttributeId", userAttributeController.DeleteNonPersonalUserAttribute)
	protected.GET("/users/:userId/non-personal/attributes", userAttributeController.GetNonPersonalUserAttributes)
	protected.GET("/users/:userId/banks/:bankId/accounts-held", userAttributeController.GetAccountsHeldByUserAtBank)
	protected.GET("/users/:userId/accounts-held", userAttributeController.GetAccountsHeldByUser)

	protected.GET("/users/provider/:provider/username/:username", userController.GetUserByProviderAndUsername)
	protected.GET("/users/provider/:provider/username/:username/lock-status", userController.GetUserLockStatus)
	protected.PUT("/users/provider/:provider/username/:username/lock-status", userController.UnlockUserByProviderAndUsername)
	protected.POST("/users/provider/:provider/username/:username/locks", userController.LockUserByProviderAndUsername)
	management.PUT("/users/:userId", userController.ValidateUserByUserId)
	protected.POST("/users/provider/:provider/provider-id/:providerId/sync", userController.SyncExternalUser)

	management.GET("/system/integrity/custom-view-names-check", integrityCheckController.CustomViewNamesCheck)
	management.GET("/system/integrity/system-view-names-check", integrityCheckController.SystemViewNamesCheck)
	management.GET("/system/integrity/account-access-unique-index-1-check", integrityCheckController.AccountAccessUniqueIndexCheck)
	management.GET("/system/integrity/banks/:bankId/account-currency-check", integrityCheckController.AccountCurrencyCheck)
	management.GET("/system/integrity/banks/:bankId/orphaned-account-check", integrityCheckController.OrphanedAccountCheck)

	protected.GET("/banks/:bankId/currencies", currencyController.GetCurrenciesAtBank)

	protected.GET("/users/:userId/entitlements-and-permissions", entitlementController.GetEntitlementsAndPermissions)

	protected.GET("/mtls-client-certificate-info", certificateController.GetMtlsClientCertificateInfo)
	my.GET("/mtls/certificate/current", v510Controller.GetMtlsClientCertificateInfo)
	protected.GET("/mtls/client-certificate-info", consentController.GetMtlsClientCertificateInfo)

	protected.POST("/banks/:bankId/atms/:atmId/attributes", atmController.CreateAtmAttribute)
	protected.GET("/banks/:bankId/atms/:atmId/attributes", atmController.GetAtmAttributes)
	protected.GET("/banks/:bankId/atms/:atmId/attributes/:atmAttributeId", atmController.GetAtmAttribute)
	protected.PUT("/banks/:bankId/atms/:atmId/attributes/:atmAttributeId", atmController.UpdateAtmAttribute)
	protected.DELETE("/banks/:bankId/atms/:atmId/attributes/:atmAttributeId", atmController.DeleteAtmAttribute)

	protected.PUT("/banks/:bankId/consents/:consentId/status", consentController.UpdateConsentStatusByConsent)
	protected.PUT("/banks/:bankId/consents/:consentId/account-access", consentController.UpdateConsentAccountAccessByConsentId)
	protected.PUT("/banks/:bankId/consents/:consentId/user-id", consentController.UpdateConsentUserIdByConsentId)
	my.GET("/banks/:bankId/consents", consentController.GetMyConsentsByBank)
	my.GET("/consents", consentController.GetMyConsents)
	protected.GET("/banks/:bankId/consents", consentController.GetConsentsAtBank)
	protected.GET("/consents", consentController.GetConsents)
	protected.GET("/banks/:bankId/consents/:consentId", consentController.GetConsentByConsentId)
	
	consumer := v510.Group("/consumer")
	consumer.Use(authMiddleware.OAuthAuth())
	{
		consumer.GET("/consents/:consentId", consentController.GetConsentByConsentIdViaConsumer)
		consumer.DELETE("/consents/:consentId", consentController.SelfRevokeConsent)
	}
	
	protected.DELETE("/banks/:bankId/consents/:consentId", consentController.RevokeConsentAtBank)
	my.DELETE("/consents/:consentId", consentController.RevokeMyConsent)
	protected.POST("/banks/:bankId/consents", consentController.CreateConsent)
	my.POST("/consents/:scaMethod", consentController.CreateConsentImplicit)
	protected.PUT("/consents/:consentId/status", consentController.UpdateConsentStatus)

	protected.GET("/banks/:bankId/accounts/:accountId/views/:viewId/counterparties", counterpartyController.GetCounterparties)
	protected.POST("/banks/:bankId/accounts/:accountId/views/:viewId/counterparties", counterpartyController.CreateCounterparty)
	protected.GET("/banks/:bankId/accounts/:accountId/views/:viewId/counterparties/:counterpartyId", counterpartyController.GetCounterpartyById)
	protected.PUT("/banks/:bankId/accounts/:accountId/views/:viewId/counterparties/:counterpartyId", counterpartyController.UpdateCounterparty)
	protected.DELETE("/banks/:bankId/accounts/:accountId/views/:viewId/counterparties/:counterpartyId", counterpartyController.DeleteCounterparty)

	protected.POST("/banks/:bankId/accounts/:accountId/views/:viewId/counterparties/:counterpartyId/limits", counterpartyLimitController.CreateCounterpartyLimit)
	protected.PUT("/banks/:bankId/accounts/:accountId/views/:viewId/counterparties/:counterpartyId/limits", counterpartyLimitController.UpdateCounterpartyLimit)
	protected.GET("/banks/:bankId/accounts/:accountId/views/:viewId/counterparties/:counterpartyId/limits", counterpartyLimitController.GetCounterpartyLimit)
	protected.GET("/banks/:bankId/accounts/:accountId/views/:viewId/counterparties/:counterpartyId/limit-status", counterpartyLimitController.GetCounterpartyLimitStatus)
	protected.DELETE("/banks/:bankId/accounts/:accountId/views/:viewId/counterparties/:counterpartyId/limits", counterpartyLimitController.DeleteCounterpartyLimit)


	management.GET("/aggregate-metrics", metricsController.GetAggregateMetrics)
	management.GET("/metrics", metricsController.GetMetrics)

	protected.GET("/users/current/customers/customer_ids", customerController.GetCustomersForUserIdsOnly)
	protected.POST("/banks/:bankId/customers/legal-name", customerController.GetCustomersByLegalName)

	protected.POST("/banks/:bankId/atms", atmManagementController.CreateAtm)
	protected.PUT("/banks/:bankId/atms/:atmId", atmManagementController.UpdateAtm)
	protected.GET("/banks/:bankId/atms", atmManagementController.GetAtms)
	protected.GET("/banks/:bankId/atms/:atmId", atmManagementController.GetAtm)
	protected.DELETE("/banks/:bankId/atms/:atmId", atmManagementController.DeleteAtm)

	protected.POST("/dynamic-registration/consumers", consumerController.CreateConsumerDynamicRegistration)
	management.POST("/consumers", consumerController.CreateConsumer)
	my.POST("/consumers", consumerController.CreateMyConsumer)
	management.PUT("/consumers/:consumerId/consumer/redirect_url", consumerController.UpdateConsumerRedirectURL)
	management.PUT("/consumers/:consumerId/consumer/logo_url", consumerController.UpdateConsumerLogoURL)
	management.PUT("/consumers/:consumerId/consumer/certificate", consumerController.UpdateConsumerCertificate)
	management.PUT("/consumers/:consumerId/consumer/name", consumerController.UpdateConsumerName)
	management.GET("/consumers/:consumerId", consumerController.GetConsumer)
	management.GET("/consumers", consumerController.GetConsumers)


	protected.GET("/user/current/consents/:consentId", consentController.GetConsentByConsentId)
	management.GET("/consents/banks/:bankId", consentController.GetConsentsAtBank)
	management.GET("/consents", consentController.GetConsents)

	protected.POST("/banks/:bankId/accounts/:accountId/views/:viewId/account-access/grant", accountAccessController.GrantUserAccessToViewById)
	protected.POST("/banks/:bankId/accounts/:accountId/views/:viewId/account-access/revoke", accountAccessController.RevokeUserAccessToViewById)
	protected.POST("/banks/:bankId/accounts/:accountId/views/:viewId/user-account-access", accountAccessController.CreateUserWithAccountAccessById)
	protected.GET("/users/:userId/account-access", accountAccessController.GetAccountAccessByUserId)

	management.GET("/transaction-requests/:transactionRequestId", transactionRequestController.GetTransactionRequestById)
	protected.GET("/banks/:bankId/accounts/:accountId/:viewId/transaction-requests", transactionRequestController.GetTransactionRequests)
	management.PUT("/transaction-requests/:transactionRequestId", transactionRequestController.UpdateTransactionRequestStatus)

	protected.GET("/banks/:bankId/accounts/:accountId/views/:viewId/balances", balanceController.GetBankAccountBalances)
	protected.GET("/banks/:bankId/balances", balanceController.GetBankAccountsBalances)
	protected.GET("/banks/:bankId/views/:viewId/balances", balanceController.GetBankAccountsBalancesThroughView)
	protected.POST("/banks/:bankId/accounts/:accountId/balances", balanceController.CreateBankAccountBalance)
	protected.GET("/banks/:bankId/accounts/:accountId/balances/:balanceId", balanceController.GetBankAccountBalanceById)
	protected.GET("/banks/:bankId/accounts/:accountId/balances", balanceController.GetAllBankAccountBalances)
	protected.PUT("/banks/:bankId/accounts/:accountId/balances/:balanceId", balanceController.UpdateBankAccountBalance)
	protected.DELETE("/banks/:bankId/accounts/:accountId/balances/:balanceId", balanceController.DeleteBankAccountBalance)

	protected.POST("/banks/:bankId/accounts/:accountId/views/:viewId/target-views", customViewController.CreateCustomView)
	protected.PUT("/banks/:bankId/accounts/:accountId/views/:viewId/target-views/:targetViewId", customViewController.UpdateCustomView)
	protected.GET("/banks/:bankId/accounts/:accountId/views/:viewId/target-views/:targetViewId", customViewController.GetCustomView)
	protected.DELETE("/banks/:bankId/accounts/:accountId/views/:viewId/target-views/:targetViewId", customViewController.DeleteCustomView)

	consumer.POST("/vrp-consent-requests", vrpController.CreateVRPConsentRequest)

	protected.POST("/regulated-entities/:regulatedEntityId/attributes", regulatedEntityAttributeController.CreateRegulatedEntityAttribute)
	protected.DELETE("/regulated-entities/:regulatedEntityId/attributes/:attributeId", regulatedEntityAttributeController.DeleteRegulatedEntityAttribute)
	protected.GET("/regulated-entities/:regulatedEntityId/attributes/:attributeId", regulatedEntityAttributeController.GetRegulatedEntityAttributeById)
	protected.GET("/regulated-entities/:regulatedEntityId/attributes", regulatedEntityAttributeController.GetAllRegulatedEntityAttributes)
	protected.PUT("/regulated-entities/:regulatedEntityId/attributes/:attributeId", regulatedEntityAttributeController.UpdateRegulatedEntityAttribute)

	protected.GET("/webui-props", webuiController.GetWebUIProps)

	protected.POST("/system-views/:viewId/permissions", systemViewController.AddSystemViewPermission)
	protected.DELETE("/system-views/:viewId/permissions/:permissionName", systemViewController.DeleteSystemViewPermission)
}
