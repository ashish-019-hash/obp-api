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

	v310 := router.Group("/obp/v3.1.0")
	v310Protected := v310.Group("")
	v310Protected.Use(authMiddleware.MultiAuth())
	{
		v310Protected.POST("/users", authController.CreateUser)
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

	v310Protected.GET("/users/current", authController.GetCurrentUser)

	admin := router.Group("/management")
	admin.Use(authMiddleware.MultiAuth())
	admin.Use(authMiddleware.RequireEntitlement("CanGetLoginAttempts"))
	{
		admin.GET("/login-attempts", authController.GetLoginAttempts)
	}
}
