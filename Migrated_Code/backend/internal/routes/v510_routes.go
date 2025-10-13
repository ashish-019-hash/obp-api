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
	protected.GET("/banks/:BANK_ID", bankController.GetBankById)

	protected.POST("/users", userController.CreateUser)
	protected.GET("/users", userController.GetUsers)

	protected.GET("/banks/:BANK_ID/accounts/:ACCOUNT_ID/views/:VIEW_ID", accountController.GetCoreAccountByIdThroughView)

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
		my.PUT("/api-collections/:API_COLLECTION_ID", apiCollectionController.UpdateMyApiCollection)
	}

	protected.GET("/regulated-entities", v510Controller.GetRegulatedEntities)
	protected.GET("/regulated-entities/:REGULATED_ENTITY_ID", v510Controller.GetRegulatedEntityById)
	protected.POST("/regulated-entities", v510Controller.CreateRegulatedEntity)
	protected.DELETE("/regulated-entities/:REGULATED_ENTITY_ID", v510Controller.DeleteRegulatedEntity)

	protected.POST("/banks/:BANK_ID/agents", agentController.CreateAgent)
	protected.PUT("/banks/:BANK_ID/agents/:AGENT_ID", agentController.UpdateAgentStatus)
	protected.GET("/banks/:BANK_ID/agents/:AGENT_ID", agentController.GetAgent)
	protected.GET("/banks/:BANK_ID/agents", agentController.GetAgents)

	protected.POST("/users/:USER_ID/non-personal/attributes", userAttributeController.CreateNonPersonalUserAttribute)
	protected.DELETE("/users/:USER_ID/non-personal/attributes/:USER_ATTRIBUTE_ID", userAttributeController.DeleteNonPersonalUserAttribute)
	protected.GET("/users/:USER_ID/non-personal/attributes", userAttributeController.GetNonPersonalUserAttributes)
	protected.GET("/users/:USER_ID/accounts/:BANK_ID", userAttributeController.GetAccountsHeldByUserAtBank)
	protected.GET("/users/:USER_ID/accounts", userAttributeController.GetAccountsHeldByUser)

	protected.GET("/users/provider/:PROVIDER/username/:USERNAME", userController.GetUserByProviderAndUsername)
	protected.GET("/users/provider/:PROVIDER/username/:USERNAME/lock-status", userController.GetUserLockStatus)
	protected.PUT("/users/provider/:PROVIDER/username/:USERNAME/lock-status", userController.UnlockUserByProviderAndUsername)
	protected.PUT("/users/provider/:PROVIDER/username/:USERNAME/lock", userController.LockUserByProviderAndUsername)
	management.PUT("/users/:USER_ID/validate", userController.ValidateUserByUserId)
	protected.POST("/users/provider/:PROVIDER/provider-id/:PROVIDER_ID/sync", userController.SyncExternalUser)

	management.GET("/system-integrity/custom-view-names-check", integrityCheckController.CustomViewNamesCheck)
	management.GET("/system-integrity/system-view-names-check", integrityCheckController.SystemViewNamesCheck)
	management.GET("/system-integrity/account-access-unique-index-1-check", integrityCheckController.AccountAccessUniqueIndexCheck)
	management.GET("/system-integrity/banks/:BANK_ID/account-currency-check", integrityCheckController.AccountCurrencyCheck)
	management.GET("/system-integrity/banks/:BANK_ID/orphaned-account-check", integrityCheckController.OrphanedAccountCheck)

	protected.GET("/banks/:BANK_ID/currencies", currencyController.GetCurrenciesAtBank)

	protected.GET("/users/:USER_ID/entitlements-and-permissions", entitlementController.GetEntitlementsAndPermissions)

	protected.GET("/mtls-client-certificate-info", certificateController.GetMtlsClientCertificateInfo)
	my.GET("/mtls-client-certificate-info", v510Controller.GetMtlsClientCertificateInfo)
	protected.GET("/mtls/client-certificate-info", consentController.GetMtlsClientCertificateInfo)

	protected.POST("/banks/:BANK_ID/atms/:ATM_ID/attributes", atmController.CreateAtmAttribute)
	protected.GET("/banks/:BANK_ID/atms/:ATM_ID/attributes", atmController.GetAtmAttributes)
	protected.GET("/banks/:BANK_ID/atms/:ATM_ID/attributes/:ATM_ATTRIBUTE_ID", atmController.GetAtmAttribute)
	protected.PUT("/banks/:BANK_ID/atms/:ATM_ID/attributes/:ATM_ATTRIBUTE_ID", atmController.UpdateAtmAttribute)
	protected.DELETE("/banks/:BANK_ID/atms/:ATM_ID/attributes/:ATM_ATTRIBUTE_ID", atmController.DeleteAtmAttribute)

	protected.PUT("/banks/:BANK_ID/consents/:CONSENT_ID/status", consentController.UpdateConsentStatusByConsent)
	protected.PUT("/banks/:BANK_ID/consents/:CONSENT_ID/account-access", consentController.UpdateConsentAccountAccessByConsentId)
	protected.PUT("/banks/:BANK_ID/consents/:CONSENT_ID/user-id", consentController.UpdateConsentUserIdByConsentId)
	my.GET("/banks/:BANK_ID/consents", consentController.GetMyConsentsByBank)
	my.GET("/consents", consentController.GetMyConsents)
	protected.GET("/banks/:BANK_ID/consents", consentController.GetConsentsAtBank)
	protected.GET("/consents", consentController.GetConsents)
	protected.GET("/banks/:BANK_ID/consents/:CONSENT_ID", consentController.GetConsentByConsentId)
	
	consumer := v510.Group("/consumer")
	consumer.Use(authMiddleware.OAuthAuth())
	{
		consumer.GET("/consents/:CONSENT_ID", consentController.GetConsentByConsentIdViaConsumer)
		consumer.DELETE("/consents/:CONSENT_ID", consentController.SelfRevokeConsent)
	}
	
	protected.DELETE("/banks/:BANK_ID/consents/:CONSENT_ID", consentController.RevokeConsentAtBank)
	my.DELETE("/consents/:CONSENT_ID", consentController.RevokeMyConsent)
	protected.POST("/banks/:BANK_ID/consents", consentController.CreateConsent)
	my.POST("/consents/:SCA_METHOD", consentController.CreateConsentImplicit)
	protected.PUT("/consents/:CONSENT_ID/status", consentController.UpdateConsentStatus)

	protected.GET("/banks/:BANK_ID/accounts/:ACCOUNT_ID/:VIEW_ID/counterparties", counterpartyController.GetCounterparties)
	protected.POST("/banks/:BANK_ID/accounts/:ACCOUNT_ID/:VIEW_ID/counterparties", counterpartyController.CreateCounterparty)
	protected.GET("/banks/:BANK_ID/accounts/:ACCOUNT_ID/:VIEW_ID/counterparties/:COUNTERPARTY_ID", counterpartyController.GetCounterpartyById)
	protected.PUT("/banks/:BANK_ID/accounts/:ACCOUNT_ID/:VIEW_ID/counterparties/:COUNTERPARTY_ID", counterpartyController.UpdateCounterparty)
	protected.DELETE("/banks/:BANK_ID/accounts/:ACCOUNT_ID/:VIEW_ID/counterparties/:COUNTERPARTY_ID", counterpartyController.DeleteCounterparty)

	protected.POST("/banks/:BANK_ID/accounts/:ACCOUNT_ID/:VIEW_ID/counterparties/:COUNTERPARTY_ID/limits", counterpartyLimitController.CreateCounterpartyLimit)
	protected.PUT("/banks/:BANK_ID/accounts/:ACCOUNT_ID/:VIEW_ID/counterparties/:COUNTERPARTY_ID/limits", counterpartyLimitController.UpdateCounterpartyLimit)
	protected.GET("/banks/:BANK_ID/accounts/:ACCOUNT_ID/:VIEW_ID/counterparties/:COUNTERPARTY_ID/limits", counterpartyLimitController.GetCounterpartyLimit)
	protected.GET("/banks/:BANK_ID/accounts/:ACCOUNT_ID/:VIEW_ID/counterparties/:COUNTERPARTY_ID/limit-status", counterpartyLimitController.GetCounterpartyLimitStatus)
	protected.DELETE("/banks/:BANK_ID/accounts/:ACCOUNT_ID/:VIEW_ID/counterparties/:COUNTERPARTY_ID/limits", counterpartyLimitController.DeleteCounterpartyLimit)


	management.GET("/aggregate-metrics", metricsController.GetAggregateMetrics)
	management.GET("/metrics", metricsController.GetMetrics)

	management.GET("/customers/customer-user-ids", customerController.GetCustomersForUserIdsOnly)
	management.GET("/customers/legal-name", customerController.GetCustomersByLegalName)

	protected.POST("/banks/:BANK_ID/atms", atmManagementController.CreateAtm)
	protected.PUT("/banks/:BANK_ID/atms/:ATM_ID", atmManagementController.UpdateAtm)
	protected.GET("/banks/:BANK_ID/atms", atmManagementController.GetAtms)
	protected.GET("/banks/:BANK_ID/atms/:ATM_ID", atmManagementController.GetAtm)
	protected.DELETE("/banks/:BANK_ID/atms/:ATM_ID", atmManagementController.DeleteAtm)

	management.POST("/consumers/dynamic-registration", consumerController.CreateConsumerDynamicRegistration)
	management.POST("/consumers", consumerController.CreateConsumer)
	my.POST("/consumers", consumerController.CreateMyConsumer)
	management.PUT("/consumers/:CONSUMER_ID/consumer/redirect_url", consumerController.UpdateConsumerRedirectURL)
	management.PUT("/consumers/:CONSUMER_ID/consumer/logo_url", consumerController.UpdateConsumerLogoURL)
	management.PUT("/consumers/:CONSUMER_ID/consumer/certificate", consumerController.UpdateConsumerCertificate)
	management.PUT("/consumers/:CONSUMER_ID/consumer/name", consumerController.UpdateConsumerName)
	management.GET("/consumers/:CONSUMER_ID", consumerController.GetConsumer)
	management.GET("/consumers", consumerController.GetConsumers)


	protected.GET("/user/current/consents/:CONSENT_ID", consentController.GetConsentByConsentId)
	management.GET("/consents/banks/:BANK_ID", consentController.GetConsentsAtBank)
	management.GET("/consents", consentController.GetConsents)

	protected.POST("/banks/:BANK_ID/accounts/:ACCOUNT_ID/views/:VIEW_ID/account-access/grant", accountAccessController.GrantUserAccessToViewById)
	protected.POST("/banks/:BANK_ID/accounts/:ACCOUNT_ID/views/:VIEW_ID/account-access/revoke", accountAccessController.RevokeUserAccessToViewById)
	protected.POST("/banks/:BANK_ID/accounts/:ACCOUNT_ID/views/:VIEW_ID/user-account-access", accountAccessController.CreateUserWithAccountAccessById)
	protected.GET("/users/:USER_ID/account-access", accountAccessController.GetAccountAccessByUserId)

	management.GET("/transaction-requests/:TRANSACTION_REQUEST_ID", transactionRequestController.GetTransactionRequestById)
	protected.GET("/banks/:BANK_ID/accounts/:ACCOUNT_ID/:VIEW_ID/transaction-requests", transactionRequestController.GetTransactionRequests)
	management.PUT("/transaction-requests/:TRANSACTION_REQUEST_ID", transactionRequestController.UpdateTransactionRequestStatus)

	protected.GET("/banks/:BANK_ID/accounts/:ACCOUNT_ID/views/:VIEW_ID/balances", balanceController.GetBankAccountBalances)
	protected.GET("/banks/:BANK_ID/balances", balanceController.GetBankAccountsBalances)
	protected.GET("/banks/:BANK_ID/views/:VIEW_ID/balances", balanceController.GetBankAccountsBalancesThroughView)
	protected.POST("/banks/:BANK_ID/accounts/:ACCOUNT_ID/balances", balanceController.CreateBankAccountBalance)
	protected.GET("/banks/:BANK_ID/accounts/:ACCOUNT_ID/balances/:BALANCE_ID", balanceController.GetBankAccountBalanceById)
	protected.GET("/banks/:BANK_ID/accounts/:ACCOUNT_ID/balances", balanceController.GetAllBankAccountBalances)
	protected.PUT("/banks/:BANK_ID/accounts/:ACCOUNT_ID/balances/:BALANCE_ID", balanceController.UpdateBankAccountBalance)
	protected.DELETE("/banks/:BANK_ID/accounts/:ACCOUNT_ID/balances/:BALANCE_ID", balanceController.DeleteBankAccountBalance)

	protected.POST("/banks/:BANK_ID/accounts/:ACCOUNT_ID/views/:VIEW_ID/target-views", customViewController.CreateCustomView)
	protected.PUT("/banks/:BANK_ID/accounts/:ACCOUNT_ID/views/:VIEW_ID/target-views/:TARGET_VIEW_ID", customViewController.UpdateCustomView)
	protected.GET("/banks/:BANK_ID/accounts/:ACCOUNT_ID/views/:VIEW_ID/target-views/:TARGET_VIEW_ID", customViewController.GetCustomView)
	protected.DELETE("/banks/:BANK_ID/accounts/:ACCOUNT_ID/views/:VIEW_ID/target-views/:TARGET_VIEW_ID", customViewController.DeleteCustomView)

	consumer.POST("/vrp-consent-requests", vrpController.CreateVRPConsentRequest)

	protected.POST("/regulated-entities/:REGULATED_ENTITY_ID/attributes", regulatedEntityAttributeController.CreateRegulatedEntityAttribute)
	protected.DELETE("/regulated-entities/:REGULATED_ENTITY_ID/attributes/:ATTRIBUTE_ID", regulatedEntityAttributeController.DeleteRegulatedEntityAttribute)
	protected.GET("/regulated-entities/:REGULATED_ENTITY_ID/attributes/:ATTRIBUTE_ID", regulatedEntityAttributeController.GetRegulatedEntityAttributeById)
	protected.GET("/regulated-entities/:REGULATED_ENTITY_ID/attributes", regulatedEntityAttributeController.GetAllRegulatedEntityAttributes)
	protected.PUT("/regulated-entities/:REGULATED_ENTITY_ID/attributes/:ATTRIBUTE_ID", regulatedEntityAttributeController.UpdateRegulatedEntityAttribute)

	protected.GET("/webui-props", webuiController.GetWebUIProps)

	protected.POST("/system-views/:VIEW_ID/permissions", systemViewController.AddSystemViewPermission)
	protected.DELETE("/system-views/:VIEW_ID/permissions/:PERMISSION_NAME", systemViewController.DeleteSystemViewPermission)
}
