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

	router.GET("/root", v510Controller.GetRoot)
	router.GET("/ui/suggested-session-timeout", v510Controller.GetSuggestedSessionTimeout)
	router.GET("/well-known", v510Controller.GetWellKnown)
	router.GET("/waiting-for-godot", v510Controller.WaitingForGodot)

	router.GET("/regulated-entities", v510Controller.GetRegulatedEntities)
	router.GET("/regulated-entities/:regulatedEntityId", v510Controller.GetRegulatedEntityById)
	router.POST("/regulated-entities", v510Controller.CreateRegulatedEntity)
	router.DELETE("/regulated-entities/:regulatedEntityId", v510Controller.DeleteRegulatedEntity)

	router.POST("/banks/:bankId/agents", agentController.CreateAgent)
	router.PUT("/banks/:bankId/agents/:agentId", agentController.UpdateAgentStatus)
	router.GET("/banks/:bankId/agents/:agentId", agentController.GetAgent)
	router.GET("/banks/:bankId/agents", agentController.GetAgents)

	router.POST("/users/:userId/non-personal/attributes", v510Controller.CreateNonPersonalUserAttribute)
	router.DELETE("/users/:userId/non-personal/attributes/:userAttributeId", v510Controller.DeleteNonPersonalUserAttribute)
	router.GET("/users/:userId/non-personal/attributes", v510Controller.GetNonPersonalUserAttributes)

	router.GET("/users/:userId/banks/:bankId/accounts-held", v510Controller.GetAccountsHeldByUserAtBank)
	router.GET("/users/:userId/accounts-held", v510Controller.GetAccountsHeldByUser)
	router.GET("/users/:userId/entitlements-and-permissions", v510Controller.GetEntitlementsAndPermissions)

	router.GET("/management/system/integrity/custom-view-names-check", v510Controller.CustomViewNamesCheck)
	router.GET("/management/system/integrity/system-view-names-check", v510Controller.SystemViewNamesCheck)
	router.GET("/management/system/integrity/account-access-unique-index-1-check", v510Controller.AccountAccessUniqueIndexCheck)
	router.GET("/management/system/integrity/banks/:bankId/account-currency-check", v510Controller.AccountCurrencyCheck)
	router.GET("/management/system/integrity/banks/:bankId/orphaned-account-check", v510Controller.OrphanedAccountCheck)

	router.GET("/banks/:bankId/currencies", v510Controller.GetCurrenciesAtBank)

	router.GET("/management/api-collections", v510Controller.GetAllApiCollections)

	router.POST("/banks/:bankId/atms/:atmId/attributes", atmController.CreateAtmAttribute)
	router.GET("/banks/:bankId/atms/:atmId/attributes", atmController.GetAtmAttributes)
	router.GET("/banks/:bankId/atms/:atmId/attributes/:atmAttributeId", atmController.GetAtmAttribute)
	router.PUT("/banks/:bankId/atms/:atmId/attributes/:atmAttributeId", atmController.UpdateAtmAttribute)
	router.DELETE("/banks/:bankId/atms/:atmId/attributes/:atmAttributeId", atmController.DeleteAtmAttribute)

	router.PUT("/management/banks/:bankId/consents/:consentId", consentController.UpdateConsentStatus)
	router.GET("/banks/:bankId/my/consents", consentController.GetMyConsentsByBank)
	router.GET("/my/consents", consentController.GetMyConsents)
	router.GET("/management/consents/banks/:bankId", consentController.GetConsentsAtBank)
	router.GET("/management/consents", consentController.GetConsents)
	router.GET("/user/current/consents/:consentId", consentController.GetConsentByConsentId)
	router.GET("/consumer/current/consents/:consentId", consentController.GetConsentByConsentIdViaConsumer)
	router.DELETE("/banks/:bankId/consents/:consentId", consentController.RevokeConsentAtBank)
	router.DELETE("/my/consent/current", consentController.SelfRevokeConsent)
	router.DELETE("/my/consents/:consentId", consentController.RevokeMyConsent)
	router.POST("/my/consents/:scaMethod", consentController.CreateConsent)
	router.GET("/my/mtls/certificate/current", consentController.GetMtlsClientCertificateInfo)

	router.POST("/banks/:bankId/accounts/:accountId/views/:viewId/counterparties/:counterpartyId/limits", counterpartyController.CreateCounterpartyLimit)
	router.PUT("/banks/:bankId/accounts/:accountId/views/:viewId/counterparties/:counterpartyId/limits", counterpartyController.UpdateCounterpartyLimit)
	router.GET("/banks/:bankId/accounts/:accountId/views/:viewId/counterparties/:counterpartyId/limits", counterpartyController.GetCounterpartyLimit)
	router.GET("/banks/:bankId/accounts/:accountId/views/:viewId/counterparties/:counterpartyId/limit-status", counterpartyController.GetCounterpartyLimitStatus)
	router.DELETE("/banks/:bankId/accounts/:accountId/views/:viewId/counterparties/:counterpartyId/limits", counterpartyController.DeleteCounterpartyLimit)
}
