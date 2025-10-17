package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"obp-api-backend/internal/models"
	"obp-api-backend/internal/repositories"
	"obp-api-backend/internal/services"
)

func AuthMiddleware(authService services.AuthService) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "missing authorization header"})
			c.Abort()
			return
		}

		var user *models.ResourceUser
		var consumer *models.Consumer
		var err error

		if strings.HasPrefix(authHeader, "DirectLogin ") {
			token := strings.TrimPrefix(authHeader, "DirectLogin ")
			token = strings.TrimPrefix(token, "token=")
			token = strings.Trim(token, "\"")
			user, consumer, err = authService.ValidateDirectLoginToken(c.Request.Context(), token)
		} else {
			c.JSON(401, gin.H{"error": "unsupported authentication type"})
			c.Abort()
			return
		}

		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Set("consumer", consumer)
		c.Next()
	})
}

func RequireEntitlement(entitlementRepo repositories.EntitlementRepository, roleName string, bankID string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.JSON(403, gin.H{"error": "forbidden: not authenticated"})
			c.Abort()
			return
		}

		resourceUser := user.(*models.ResourceUser)

		hasRole, err := entitlementRepo.HasRole(c.Request.Context(), resourceUser.UserID, roleName, bankID)
		if err != nil || !hasRole {
			c.JSON(403, gin.H{"error": "forbidden: missing required entitlement " + roleName})
			c.Abort()
			return
		}

		c.Next()
	}
}
