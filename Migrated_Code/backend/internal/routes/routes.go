package routes

import (
	"github.com/ashish-019-hash/obp-api-backend/internal/controllers"
	"github.com/ashish-019-hash/obp-api-backend/internal/middleware"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, orchestrationService *services.OrchestrationService, authMiddleware *middleware.AuthMiddleware) {
	healthController := controllers.NewHealthController()

	router.GET("/health", healthController.HealthCheck)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := router.Group("/api/v1")
	{
		v1.GET("/health", healthController.HealthCheck)
		
	}

	SetupV400Routes(router, orchestrationService, authMiddleware)
	SetupV510Routes(router, orchestrationService, authMiddleware)
}
