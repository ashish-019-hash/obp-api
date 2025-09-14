package routes

import (
	"github.com/ashish-019-hash/obp-api-backend/internal/controllers"
	"github.com/ashish-019-hash/obp-api-backend/internal/middleware"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(orchestrationService *services.OrchestrationService) *gin.Engine {
	r := gin.New()

	r.Use(middleware.Logger())
	r.Use(middleware.CORS())
	r.Use(gin.Recovery())

	healthController := controllers.NewHealthController()

	r.GET("/health", healthController.HealthCheck)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/api/v1")
	{
		v1.GET("/health", healthController.HealthCheck)
		
	}

	SetupV510Routes(r, orchestrationService)

	return r
}
