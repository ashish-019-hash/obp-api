package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ashish-019-hash/obp-api-backend/internal/controllers"
	"github.com/ashish-019-hash/obp-api-backend/internal/middleware"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
)

func SetupV400Routes(router *gin.Engine, orchestrationService *services.OrchestrationService, authMiddleware *middleware.AuthMiddleware) {
	bankController := controllers.NewBankController(orchestrationService)
	userController := controllers.NewUserController(orchestrationService)
	accountController := controllers.NewAccountController(orchestrationService)
	atmController := controllers.NewAtmController(orchestrationService)
	atmManagementController := controllers.NewAtmManagementController(orchestrationService)
	
	v400 := router.Group("/obp/v4.0.0")
	
	protected := v400.Group("")
	protected.Use(authMiddleware.MultiAuth())
	{
		protected.GET("/banks", bankController.GetBanks)
		protected.GET("/banks/:BANK_ID", bankController.GetBankById)
		protected.POST("/banks", bankController.CreateBank)
		
		protected.GET("/banks/:BANK_ID/accounts/:ACCOUNT_ID/views/:VIEW_ID", accountController.GetCoreAccountByIdThroughView)
		
		protected.POST("/banks/:BANK_ID/atms", atmManagementController.CreateAtm)
		protected.PUT("/banks/:BANK_ID/atms/:ATM_ID", atmManagementController.UpdateAtm)
		protected.GET("/banks/:BANK_ID/atms", atmManagementController.GetAtms)
		protected.GET("/banks/:BANK_ID/atms/:ATM_ID", atmManagementController.GetAtm)
		protected.DELETE("/banks/:BANK_ID/atms/:ATM_ID", atmManagementController.DeleteAtm)
		
		protected.POST("/banks/:BANK_ID/atms/:ATM_ID/attributes", atmController.CreateAtmAttribute)
		protected.GET("/banks/:BANK_ID/atms/:ATM_ID/attributes", atmController.GetAtmAttributes)
		protected.PUT("/banks/:BANK_ID/atms/:ATM_ID/attributes/:ATM_ATTRIBUTE_ID", atmController.UpdateAtmAttribute)
		protected.DELETE("/banks/:BANK_ID/atms/:ATM_ID/attributes/:ATM_ATTRIBUTE_ID", atmController.DeleteAtmAttribute)
		
		protected.POST("/users", userController.CreateUser)
		protected.GET("/users", userController.GetUsers)
	}
}
