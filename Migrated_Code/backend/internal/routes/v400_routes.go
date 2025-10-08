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
		protected.GET("/banks/:bankId", bankController.GetBankById)
		protected.POST("/banks", bankController.CreateBank)
		
		protected.GET("/banks/:bankId/accounts/:accountId/views/:viewId", accountController.GetCoreAccountByIdThroughView)
		
		protected.POST("/banks/:bankId/atms", atmManagementController.CreateAtm)
		protected.PUT("/banks/:bankId/atms/:atmId", atmManagementController.UpdateAtm)
		protected.GET("/banks/:bankId/atms", atmManagementController.GetAtms)
		protected.GET("/banks/:bankId/atms/:atmId", atmManagementController.GetAtm)
		protected.DELETE("/banks/:bankId/atms/:atmId", atmManagementController.DeleteAtm)
		
		protected.POST("/banks/:bankId/atms/:atmId/attributes", atmController.CreateAtmAttribute)
		protected.GET("/banks/:bankId/atms/:atmId/attributes", atmController.GetAtmAttributes)
		protected.PUT("/banks/:bankId/atms/:atmId/attributes/:atmAttributeId", atmController.UpdateAtmAttribute)
		protected.DELETE("/banks/:bankId/atms/:atmId/attributes/:atmAttributeId", atmController.DeleteAtmAttribute)
		
		protected.POST("/users", userController.CreateUser)
		protected.GET("/users", userController.GetUsers)
	}
}
