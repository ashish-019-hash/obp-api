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

	router.GET("/root", v510Controller.GetRoot)
	router.GET("/ui/suggested-session-timeout", v510Controller.GetSuggestedSessionTimeout)
	router.GET("/well-known", v510Controller.GetWellKnown)
	router.GET("/waiting-for-godot", v510Controller.WaitingForGodot)

	router.GET("/tags", v510Controller.GetApiTags)

	router.GET("/banks/:bankId/accounts/:accountId/views/:viewId", accountController.GetCoreAccountByIdThroughView)

	router.GET("/management/api-collections", apiCollectionController.GetApiCollections)
	router.POST("/management/api-collections", apiCollectionController.CreateApiCollection)
	router.GET("/my/api-collections", apiCollectionController.GetMyApiCollections)
	router.PUT("/my/api-collections/:collectionId", apiCollectionController.UpdateMyApiCollection)

	router.GET("/regulated-entities", v510Controller.GetRegulatedEntities)
	router.GET("/regulated-entities/:regulatedEntityId", v510Controller.GetRegulatedEntityById)
	router.POST("/regulated-entities", v510Controller.CreateRegulatedEntity)
	router.DELETE("/regulated-entities/:regulatedEntityId", v510Controller.DeleteRegulatedEntity)

	router.POST("/banks/:bankId/agents", agentController.CreateAgent)
	router.PUT("/banks/:bankId/agents/:agentId", agentController.UpdateAgentStatus)
	router.GET("/banks/:bankId/agents/:agentId", agentController.GetAgent)
	router.GET("/banks/:bankId/agents", agentController.GetAgents)

	router.POST("/users/:userId/non-personal/attributes", userAttributeController.CreateNonPersonalUserAttribute)
	router.DELETE("/users/:userId/non-personal/attributes/:userAttributeId", userAttributeController.DeleteNonPersonalUserAttribute)
	router.GET("/users/:userId/non-personal/attributes", userAttributeController.GetNonPersonalUserAttributes)
	router.GET("/users/:userId/banks/:bankId/accounts-held", userAttributeController.GetAccountsHeldByUserAtBank)
	router.GET("/users/:userId/accounts-held", userAttributeController.GetAccountsHeldByUser)

	router.GET("/users/provider/:provider/username/:username", userController.GetUserByProviderAndUsername)
	router.GET("/users/provider/:provider/username/:username/lock-status", userController.GetUserLockStatus)
	router.PUT("/users/provider/:provider/username/:username/lock-status", userController.UnlockUserByProviderAndUsername)
	router.POST("/users/provider/:provider/username/:username/locks", userController.LockUserByProviderAndUsername)
	router.PUT("/management/users/:userId", userController.ValidateUserByUserId)
	router.POST("/users/provider/:provider/provider-id/:providerId/sync", userController.SyncExternalUser)

	router.GET("/management/system/integrity/custom-view-names-check", integrityCheckController.CustomViewNamesCheck)
	router.GET("/management/system/integrity/system-view-names-check", integrityCheckController.SystemViewNamesCheck)
	router.GET("/management/system/integrity/account-access-unique-index-1-check", integrityCheckController.AccountAccessUniqueIndexCheck)
	router.GET("/management/system/integrity/banks/:bankId/account-currency-check", integrityCheckController.AccountCurrencyCheck)
	router.GET("/management/system/integrity/banks/:bankId/orphaned-account-check", integrityCheckController.OrphanedAccountCheck)

	router.GET("/banks/:bankId/currencies", currencyController.GetCurrenciesAtBank)

	router.GET("/users/:userId/entitlements-and-permissions", entitlementController.GetEntitlementsAndPermissions)

	router.GET("/mtls-client-certificate-info", certificateController.GetMtlsClientCertificateInfo)
	router.GET("/my/mtls/certificate/current", v510Controller.GetMtlsClientCertificateInfo)
	router.GET("/mtls/client-certificate-info", consentController.GetMtlsClientCertificateInfo)

	router.POST("/banks/:bankId/atms/:atmId/attributes", atmController.CreateAtmAttribute)
	router.GET("/banks/:bankId/atms/:atmId/attributes", atmController.GetAtmAttributes)
	router.GET("/banks/:bankId/atms/:atmId/attributes/:atmAttributeId", atmController.GetAtmAttribute)
	router.PUT("/banks/:bankId/atms/:atmId/attributes/:atmAttributeId", atmController.UpdateAtmAttribute)
	router.DELETE("/banks/:bankId/atms/:atmId/attributes/:atmAttributeId", atmController.DeleteAtmAttribute)

	router.PUT("/banks/:bankId/consents/:consentId/status", consentController.UpdateConsentStatusByConsent)
	router.PUT("/banks/:bankId/consents/:consentId/account-access", consentController.UpdateConsentAccountAccessByConsentId)
	router.PUT("/banks/:bankId/consents/:consentId/user-id", consentController.UpdateConsentUserIdByConsentId)
	router.GET("/banks/:bankId/my/consents", consentController.GetMyConsentsByBank)
	router.GET("/my/consents", consentController.GetMyConsents)
	router.GET("/banks/:bankId/consents", consentController.GetConsentsAtBank)
	router.GET("/consents", consentController.GetConsents)
	router.GET("/banks/:bankId/consents/:consentId", consentController.GetConsentByConsentId)
	router.GET("/consumer/consents/:consentId", consentController.GetConsentByConsentIdViaConsumer)
	router.DELETE("/banks/:bankId/consents/:consentId", consentController.RevokeConsentAtBank)
	router.DELETE("/consumer/consents/:consentId", consentController.SelfRevokeConsent)
	router.DELETE("/my/consents/:consentId", consentController.RevokeMyConsent)
	router.POST("/banks/:bankId/consents", consentController.CreateConsent)
	router.POST("/my/consents/:scaMethod", consentController.CreateConsentImplicit)

	router.GET("/banks/:bankId/accounts/:accountId/views/:viewId/counterparties", counterpartyController.GetCounterparties)
	router.POST("/banks/:bankId/accounts/:accountId/views/:viewId/counterparties", counterpartyController.CreateCounterparty)
	router.GET("/banks/:bankId/accounts/:accountId/views/:viewId/counterparties/:counterpartyId", counterpartyController.GetCounterpartyById)
	router.PUT("/banks/:bankId/accounts/:accountId/views/:viewId/counterparties/:counterpartyId", counterpartyController.UpdateCounterparty)
	router.DELETE("/banks/:bankId/accounts/:accountId/views/:viewId/counterparties/:counterpartyId", counterpartyController.DeleteCounterparty)

	router.POST("/banks/:bankId/accounts/:accountId/views/:viewId/counterparties/:counterpartyId/limits", counterpartyLimitController.CreateCounterpartyLimit)
	router.PUT("/banks/:bankId/accounts/:accountId/views/:viewId/counterparties/:counterpartyId/limits", counterpartyLimitController.UpdateCounterpartyLimit)
	router.GET("/banks/:bankId/accounts/:accountId/views/:viewId/counterparties/:counterpartyId/limits", counterpartyLimitController.GetCounterpartyLimit)
	router.GET("/banks/:bankId/accounts/:accountId/views/:viewId/counterparties/:counterpartyId/limit-status", counterpartyLimitController.GetCounterpartyLimitStatus)
	router.DELETE("/banks/:bankId/accounts/:accountId/views/:viewId/counterparties/:counterpartyId/limits", counterpartyLimitController.DeleteCounterpartyLimit)


	router.GET("/management/aggregate-metrics", metricsController.GetAggregateMetrics)
	router.GET("/management/metrics", metricsController.GetMetrics)

	router.GET("/users/current/customers/customer_ids", customerController.GetCustomersForUserIdsOnly)
	router.POST("/banks/:bankId/customers/legal-name", customerController.GetCustomersByLegalName)

	router.POST("/banks/:bankId/atms", atmManagementController.CreateAtm)
	router.PUT("/banks/:bankId/atms/:atmId", atmManagementController.UpdateAtm)
	router.GET("/banks/:bankId/atms", atmManagementController.GetAtms)
	router.GET("/banks/:bankId/atms/:atmId", atmManagementController.GetAtm)
	router.DELETE("/banks/:bankId/atms/:atmId", atmManagementController.DeleteAtm)

	router.POST("/dynamic-registration/consumers", consumerController.CreateConsumerDynamicRegistration)
	router.POST("/management/consumers", consumerController.CreateConsumer)
	router.POST("/my/consumers", consumerController.CreateMyConsumer)
	router.PUT("/management/consumers/:consumerId/consumer/redirect_url", consumerController.UpdateConsumerRedirectURL)
	router.PUT("/management/consumers/:consumerId/consumer/logo_url", consumerController.UpdateConsumerLogoURL)
	router.PUT("/management/consumers/:consumerId/consumer/certificate", consumerController.UpdateConsumerCertificate)
	router.PUT("/management/consumers/:consumerId/consumer/name", consumerController.UpdateConsumerName)
	router.GET("/management/consumers/:consumerId", consumerController.GetConsumer)
	router.GET("/management/consumers", consumerController.GetConsumers)


	router.GET("/user/current/consents/:consentId", consentController.GetConsentByConsentId)
	router.GET("/management/consents/banks/:bankId", consentController.GetConsentsAtBank)
	router.GET("/management/consents", consentController.GetConsents)

	router.POST("/banks/:bankId/accounts/:accountId/views/:viewId/account-access/grant", accountAccessController.GrantUserAccessToViewById)
	router.POST("/banks/:bankId/accounts/:accountId/views/:viewId/account-access/revoke", accountAccessController.RevokeUserAccessToViewById)
	router.POST("/banks/:bankId/accounts/:accountId/views/:viewId/user-account-access", accountAccessController.CreateUserWithAccountAccessById)
	router.GET("/users/:userId/account-access", accountAccessController.GetAccountAccessByUserId)

	router.GET("/management/transaction-requests/:transactionRequestId", transactionRequestController.GetTransactionRequestById)
	router.GET("/banks/:bankId/accounts/:accountId/:viewId/transaction-requests", transactionRequestController.GetTransactionRequests)
	router.PUT("/management/transaction-requests/:transactionRequestId", transactionRequestController.UpdateTransactionRequestStatus)

	router.GET("/banks/:bankId/accounts/:accountId/views/:viewId/balances", balanceController.GetBankAccountBalances)
	router.GET("/banks/:bankId/balances", balanceController.GetBankAccountsBalances)
	router.GET("/banks/:bankId/views/:viewId/balances", balanceController.GetBankAccountsBalancesThroughView)
	router.POST("/banks/:bankId/accounts/:accountId/balances", balanceController.CreateBankAccountBalance)
	router.GET("/banks/:bankId/accounts/:accountId/balances/:balanceId", balanceController.GetBankAccountBalanceById)
	router.GET("/banks/:bankId/accounts/:accountId/balances", balanceController.GetAllBankAccountBalances)
	router.PUT("/banks/:bankId/accounts/:accountId/balances/:balanceId", balanceController.UpdateBankAccountBalance)
	router.DELETE("/banks/:bankId/accounts/:accountId/balances/:balanceId", balanceController.DeleteBankAccountBalance)

	router.POST("/banks/:bankId/accounts/:accountId/views/:viewId/target-views", customViewController.CreateCustomView)
	router.PUT("/banks/:bankId/accounts/:accountId/views/:viewId/target-views/:targetViewId", customViewController.UpdateCustomView)
	router.GET("/banks/:bankId/accounts/:accountId/views/:viewId/target-views/:targetViewId", customViewController.GetCustomView)
	router.DELETE("/banks/:bankId/accounts/:accountId/views/:viewId/target-views/:targetViewId", customViewController.DeleteCustomView)

	router.POST("/consumer/vrp-consent-requests", vrpController.CreateVRPConsentRequest)

	router.POST("/regulated-entities/:regulatedEntityId/attributes", regulatedEntityAttributeController.CreateRegulatedEntityAttribute)
	router.DELETE("/regulated-entities/:regulatedEntityId/attributes/:attributeId", regulatedEntityAttributeController.DeleteRegulatedEntityAttribute)
	router.GET("/regulated-entities/:regulatedEntityId/attributes/:attributeId", regulatedEntityAttributeController.GetRegulatedEntityAttributeById)
	router.GET("/regulated-entities/:regulatedEntityId/attributes", regulatedEntityAttributeController.GetAllRegulatedEntityAttributes)
	router.PUT("/regulated-entities/:regulatedEntityId/attributes/:attributeId", regulatedEntityAttributeController.UpdateRegulatedEntityAttribute)

	router.GET("/webui-props", webuiController.GetWebUIProps)

	router.POST("/system-views/:viewId/permissions", systemViewController.AddSystemViewPermission)
	router.DELETE("/system-views/:viewId/permissions/:permissionName", systemViewController.DeleteSystemViewPermission)
}
