package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ashish-019-hash/obp-api-backend/internal/controllers"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
)

func SetupV510Routes(router *gin.Engine, orchestrationService *services.OrchestrationService) {
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

	v510.GET("/tags", v510Controller.GetApiTags)

	v510.POST("/banks", bankController.CreateBank)
	v510.GET("/banks", bankController.GetBanks)
	v510.GET("/banks/:bankId", bankController.GetBankById)

	v510.POST("/users", userController.CreateUser)
	v510.GET("/users", userController.GetUsers)

	v510.GET("/banks/:bankId/accounts/:accountId/views/:viewId", accountController.GetCoreAccountByIdThroughView)

	v510.GET("/management/api-collections", apiCollectionController.GetApiCollections)
	v510.POST("/management/api-collections", apiCollectionController.CreateApiCollection)
	v510.GET("/my/api-collections", apiCollectionController.GetMyApiCollections)
	v510.PUT("/my/api-collections/:collectionId", apiCollectionController.UpdateMyApiCollection)

	v510.GET("/regulated-entities", v510Controller.GetRegulatedEntities)
	v510.GET("/regulated-entities/:regulatedEntityId", v510Controller.GetRegulatedEntityById)
	v510.POST("/regulated-entities", v510Controller.CreateRegulatedEntity)
	v510.DELETE("/regulated-entities/:regulatedEntityId", v510Controller.DeleteRegulatedEntity)

	v510.POST("/banks/:bankId/agents", agentController.CreateAgent)
	v510.PUT("/banks/:bankId/agents/:agentId", agentController.UpdateAgentStatus)
	v510.GET("/banks/:bankId/agents/:agentId", agentController.GetAgent)
	v510.GET("/banks/:bankId/agents", agentController.GetAgents)

	v510.POST("/users/:userId/non-personal/attributes", userAttributeController.CreateNonPersonalUserAttribute)
	v510.DELETE("/users/:userId/non-personal/attributes/:userAttributeId", userAttributeController.DeleteNonPersonalUserAttribute)
	v510.GET("/users/:userId/non-personal/attributes", userAttributeController.GetNonPersonalUserAttributes)
	v510.GET("/users/:userId/banks/:bankId/accounts-held", userAttributeController.GetAccountsHeldByUserAtBank)
	v510.GET("/users/:userId/accounts-held", userAttributeController.GetAccountsHeldByUser)

	v510.GET("/users/provider/:provider/username/:username", userController.GetUserByProviderAndUsername)
	v510.GET("/users/provider/:provider/username/:username/lock-status", userController.GetUserLockStatus)
	v510.PUT("/users/provider/:provider/username/:username/lock-status", userController.UnlockUserByProviderAndUsername)
	v510.POST("/users/provider/:provider/username/:username/locks", userController.LockUserByProviderAndUsername)
	v510.PUT("/management/users/:userId", userController.ValidateUserByUserId)
	v510.POST("/users/provider/:provider/provider-id/:providerId/sync", userController.SyncExternalUser)

	v510.GET("/management/system/integrity/custom-view-names-check", integrityCheckController.CustomViewNamesCheck)
	v510.GET("/management/system/integrity/system-view-names-check", integrityCheckController.SystemViewNamesCheck)
	v510.GET("/management/system/integrity/account-access-unique-index-1-check", integrityCheckController.AccountAccessUniqueIndexCheck)
	v510.GET("/management/system/integrity/banks/:bankId/account-currency-check", integrityCheckController.AccountCurrencyCheck)
	v510.GET("/management/system/integrity/banks/:bankId/orphaned-account-check", integrityCheckController.OrphanedAccountCheck)

	v510.GET("/banks/:bankId/currencies", currencyController.GetCurrenciesAtBank)

	v510.GET("/users/:userId/entitlements-and-permissions", entitlementController.GetEntitlementsAndPermissions)

	v510.GET("/mtls-client-certificate-info", certificateController.GetMtlsClientCertificateInfo)
	v510.GET("/my/mtls/certificate/current", v510Controller.GetMtlsClientCertificateInfo)
	v510.GET("/mtls/client-certificate-info", consentController.GetMtlsClientCertificateInfo)

	v510.POST("/banks/:bankId/atms/:atmId/attributes", atmController.CreateAtmAttribute)
	v510.GET("/banks/:bankId/atms/:atmId/attributes", atmController.GetAtmAttributes)
	v510.GET("/banks/:bankId/atms/:atmId/attributes/:atmAttributeId", atmController.GetAtmAttribute)
	v510.PUT("/banks/:bankId/atms/:atmId/attributes/:atmAttributeId", atmController.UpdateAtmAttribute)
	v510.DELETE("/banks/:bankId/atms/:atmId/attributes/:atmAttributeId", atmController.DeleteAtmAttribute)

	v510.PUT("/banks/:bankId/consents/:consentId/status", consentController.UpdateConsentStatusByConsent)
	v510.PUT("/banks/:bankId/consents/:consentId/account-access", consentController.UpdateConsentAccountAccessByConsentId)
	v510.PUT("/banks/:bankId/consents/:consentId/user-id", consentController.UpdateConsentUserIdByConsentId)
	v510.GET("/banks/:bankId/my/consents", consentController.GetMyConsentsByBank)
	v510.GET("/my/consents", consentController.GetMyConsents)
	v510.GET("/banks/:bankId/consents", consentController.GetConsentsAtBank)
	v510.GET("/consents", consentController.GetConsents)
	v510.GET("/banks/:bankId/consents/:consentId", consentController.GetConsentByConsentId)
	v510.GET("/consumer/consents/:consentId", consentController.GetConsentByConsentIdViaConsumer)
	v510.DELETE("/banks/:bankId/consents/:consentId", consentController.RevokeConsentAtBank)
	v510.DELETE("/consumer/consents/:consentId", consentController.SelfRevokeConsent)
	v510.DELETE("/my/consents/:consentId", consentController.RevokeMyConsent)
	v510.POST("/banks/:bankId/consents", consentController.CreateConsent)
	v510.POST("/my/consents/:scaMethod", consentController.CreateConsentImplicit)
	v510.PUT("/consents/:consentId/status", consentController.UpdateConsentStatus)

	v510.GET("/banks/:bankId/accounts/:accountId/views/:viewId/counterparties", counterpartyController.GetCounterparties)
	v510.POST("/banks/:bankId/accounts/:accountId/views/:viewId/counterparties", counterpartyController.CreateCounterparty)
	v510.GET("/banks/:bankId/accounts/:accountId/views/:viewId/counterparties/:counterpartyId", counterpartyController.GetCounterpartyById)
	v510.PUT("/banks/:bankId/accounts/:accountId/views/:viewId/counterparties/:counterpartyId", counterpartyController.UpdateCounterparty)
	v510.DELETE("/banks/:bankId/accounts/:accountId/views/:viewId/counterparties/:counterpartyId", counterpartyController.DeleteCounterparty)

	v510.POST("/banks/:bankId/accounts/:accountId/views/:viewId/counterparties/:counterpartyId/limits", counterpartyLimitController.CreateCounterpartyLimit)
	v510.PUT("/banks/:bankId/accounts/:accountId/views/:viewId/counterparties/:counterpartyId/limits", counterpartyLimitController.UpdateCounterpartyLimit)
	v510.GET("/banks/:bankId/accounts/:accountId/views/:viewId/counterparties/:counterpartyId/limits", counterpartyLimitController.GetCounterpartyLimit)
	v510.GET("/banks/:bankId/accounts/:accountId/views/:viewId/counterparties/:counterpartyId/limit-status", counterpartyLimitController.GetCounterpartyLimitStatus)
	v510.DELETE("/banks/:bankId/accounts/:accountId/views/:viewId/counterparties/:counterpartyId/limits", counterpartyLimitController.DeleteCounterpartyLimit)


	v510.GET("/management/aggregate-metrics", metricsController.GetAggregateMetrics)
	v510.GET("/management/metrics", metricsController.GetMetrics)

	v510.GET("/users/current/customers/customer_ids", customerController.GetCustomersForUserIdsOnly)
	v510.POST("/banks/:bankId/customers/legal-name", customerController.GetCustomersByLegalName)

	v510.POST("/banks/:bankId/atms", atmManagementController.CreateAtm)
	v510.PUT("/banks/:bankId/atms/:atmId", atmManagementController.UpdateAtm)
	v510.GET("/banks/:bankId/atms", atmManagementController.GetAtms)
	v510.GET("/banks/:bankId/atms/:atmId", atmManagementController.GetAtm)
	v510.DELETE("/banks/:bankId/atms/:atmId", atmManagementController.DeleteAtm)

	v510.POST("/dynamic-registration/consumers", consumerController.CreateConsumerDynamicRegistration)
	v510.POST("/management/consumers", consumerController.CreateConsumer)
	v510.POST("/my/consumers", consumerController.CreateMyConsumer)
	v510.PUT("/management/consumers/:consumerId/consumer/redirect_url", consumerController.UpdateConsumerRedirectURL)
	v510.PUT("/management/consumers/:consumerId/consumer/logo_url", consumerController.UpdateConsumerLogoURL)
	v510.PUT("/management/consumers/:consumerId/consumer/certificate", consumerController.UpdateConsumerCertificate)
	v510.PUT("/management/consumers/:consumerId/consumer/name", consumerController.UpdateConsumerName)
	v510.GET("/management/consumers/:consumerId", consumerController.GetConsumer)
	v510.GET("/management/consumers", consumerController.GetConsumers)


	v510.GET("/user/current/consents/:consentId", consentController.GetConsentByConsentId)
	v510.GET("/management/consents/banks/:bankId", consentController.GetConsentsAtBank)
	v510.GET("/management/consents", consentController.GetConsents)

	v510.POST("/banks/:bankId/accounts/:accountId/views/:viewId/account-access/grant", accountAccessController.GrantUserAccessToViewById)
	v510.POST("/banks/:bankId/accounts/:accountId/views/:viewId/account-access/revoke", accountAccessController.RevokeUserAccessToViewById)
	v510.POST("/banks/:bankId/accounts/:accountId/views/:viewId/user-account-access", accountAccessController.CreateUserWithAccountAccessById)
	v510.GET("/users/:userId/account-access", accountAccessController.GetAccountAccessByUserId)

	v510.GET("/management/transaction-requests/:transactionRequestId", transactionRequestController.GetTransactionRequestById)
	v510.GET("/banks/:bankId/accounts/:accountId/:viewId/transaction-requests", transactionRequestController.GetTransactionRequests)
	v510.PUT("/management/transaction-requests/:transactionRequestId", transactionRequestController.UpdateTransactionRequestStatus)

	v510.GET("/banks/:bankId/accounts/:accountId/views/:viewId/balances", balanceController.GetBankAccountBalances)
	v510.GET("/banks/:bankId/balances", balanceController.GetBankAccountsBalances)
	v510.GET("/banks/:bankId/views/:viewId/balances", balanceController.GetBankAccountsBalancesThroughView)
	v510.POST("/banks/:bankId/accounts/:accountId/balances", balanceController.CreateBankAccountBalance)
	v510.GET("/banks/:bankId/accounts/:accountId/balances/:balanceId", balanceController.GetBankAccountBalanceById)
	v510.GET("/banks/:bankId/accounts/:accountId/balances", balanceController.GetAllBankAccountBalances)
	v510.PUT("/banks/:bankId/accounts/:accountId/balances/:balanceId", balanceController.UpdateBankAccountBalance)
	v510.DELETE("/banks/:bankId/accounts/:accountId/balances/:balanceId", balanceController.DeleteBankAccountBalance)

	v510.POST("/banks/:bankId/accounts/:accountId/views/:viewId/target-views", customViewController.CreateCustomView)
	v510.PUT("/banks/:bankId/accounts/:accountId/views/:viewId/target-views/:targetViewId", customViewController.UpdateCustomView)
	v510.GET("/banks/:bankId/accounts/:accountId/views/:viewId/target-views/:targetViewId", customViewController.GetCustomView)
	v510.DELETE("/banks/:bankId/accounts/:accountId/views/:viewId/target-views/:targetViewId", customViewController.DeleteCustomView)

	v510.POST("/consumer/vrp-consent-requests", vrpController.CreateVRPConsentRequest)

	v510.POST("/regulated-entities/:regulatedEntityId/attributes", regulatedEntityAttributeController.CreateRegulatedEntityAttribute)
	v510.DELETE("/regulated-entities/:regulatedEntityId/attributes/:attributeId", regulatedEntityAttributeController.DeleteRegulatedEntityAttribute)
	v510.GET("/regulated-entities/:regulatedEntityId/attributes/:attributeId", regulatedEntityAttributeController.GetRegulatedEntityAttributeById)
	v510.GET("/regulated-entities/:regulatedEntityId/attributes", regulatedEntityAttributeController.GetAllRegulatedEntityAttributes)
	v510.PUT("/regulated-entities/:regulatedEntityId/attributes/:attributeId", regulatedEntityAttributeController.UpdateRegulatedEntityAttribute)

	v510.GET("/webui-props", webuiController.GetWebUIProps)

	v510.POST("/system-views/:viewId/permissions", systemViewController.AddSystemViewPermission)
	v510.DELETE("/system-views/:viewId/permissions/:permissionName", systemViewController.DeleteSystemViewPermission)
}
