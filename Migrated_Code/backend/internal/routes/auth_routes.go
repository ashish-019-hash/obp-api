package routes

import (
	"github.com/ashish-019-hash/obp-api-backend/internal/controllers"
	"github.com/ashish-019-hash/obp-api-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.Engine, authController *controllers.AuthController, authMiddleware *middleware.AuthMiddleware) {
	auth := router.Group("/auth")
	{
		auth.POST("/direct-login", authController.DirectLogin)
		auth.POST("/consumers", authController.RegisterConsumer)
	}

	authProtected := router.Group("/auth")
	authProtected.Use(authMiddleware.MultiAuth())
	{
		authProtected.POST("/users", authController.CreateUser)
	}

	oauth := router.Group("/oauth")
	{
		oauth.POST("/initiate", authController.OAuthInitiate)
		oauth.POST("/token", authController.OAuthToken)
		oauth.GET("/authorize", authController.OAuthAuthorize)
	}

	obpAuth := router.Group("/my")
	{
		obpAuth.POST("/logins/direct", authController.DirectLogin)
	}

	v510 := router.Group("/obp/v5.1.0")
	v510Protected := v510.Group("")
	v510Protected.Use(authMiddleware.MultiAuth())
	
	v510Protected.GET("/users/current", authController.GetCurrentUser)

	admin := router.Group("/management")
	admin.Use(authMiddleware.MultiAuth())
	admin.Use(authMiddleware.RequireEntitlement("CanGetLoginAttempts"))
	{
		admin.GET("/login-attempts", authController.GetLoginAttempts)
	}
}
