package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/ashish-019-hash/obp-api-backend/internal/models"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	ValidateJWTToken(tokenString string) (*models.User, error)
	ValidateOAuthToken(tokenString string) (*models.User, *models.Consumer, error)
	ValidateDirectLoginToken(tokenString string) (*models.User, *models.Consumer, error)
	RecordLoginAttempt(userID, username, ipAddress, userAgent, authMethod string, success bool, failureReason string) error
}

type AuthMiddleware struct {
	authService AuthService
	jwtSecret   string
}

func NewAuthMiddleware(authService AuthService, jwtSecret string) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
		jwtSecret:   jwtSecret,
	}
}

func (am *AuthMiddleware) JWTAuth() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		tokenString := extractToken(c)
		if tokenString == "" {
			am.recordFailedAttempt(c, "", "JWT", "No token provided")
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization token required",
				"code":  "OBP-20001",
			})
			c.Abort()
			return
		}

		user, err := am.authService.ValidateJWTToken(tokenString)
		if err != nil {
			am.recordFailedAttempt(c, "", "JWT", err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired token",
				"code":  "OBP-20002",
			})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Set("user_id", user.UserID)
		c.Set("auth_method", "JWT")
		
		am.recordSuccessfulAttempt(c, user.UserID, "JWT")
		c.Next()
	})
}

func (am *AuthMiddleware) OAuthAuth() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		tokenString := extractToken(c)
		if tokenString == "" {
			am.recordFailedAttempt(c, "", "OAuth", "No token provided")
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "OAuth token required",
				"code":  "OBP-20003",
			})
			c.Abort()
			return
		}

		user, consumer, err := am.authService.ValidateOAuthToken(tokenString)
		if err != nil {
			am.recordFailedAttempt(c, "", "OAuth", err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid OAuth token",
				"code":  "OBP-20004",
			})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Set("user_id", user.UserID)
		c.Set("consumer", consumer)
		c.Set("consumer_id", consumer.ConsumerID)
		c.Set("auth_method", "OAuth")
		
		am.recordSuccessfulAttempt(c, user.UserID, "OAuth")
		c.Next()
	})
}

func (am *AuthMiddleware) DirectLoginAuth() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		tokenString := extractDirectLoginToken(c)
		if tokenString == "" {
			am.recordFailedAttempt(c, "", "DirectLogin", "No token provided")
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "DirectLogin token required",
				"code":  "OBP-20005",
			})
			c.Abort()
			return
		}

		user, consumer, err := am.authService.ValidateDirectLoginToken(tokenString)
		if err != nil {
			am.recordFailedAttempt(c, "", "DirectLogin", err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid DirectLogin token",
				"code":  "OBP-20006",
			})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Set("user_id", user.UserID)
		c.Set("consumer", consumer)
		c.Set("consumer_id", consumer.ConsumerID)
		c.Set("auth_method", "DirectLogin")
		
		am.recordSuccessfulAttempt(c, user.UserID, "DirectLogin")
		c.Next()
	})
}

func (am *AuthMiddleware) MultiAuth() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		if directLoginToken := extractDirectLoginToken(c); directLoginToken != "" {
			user, consumer, err := am.authService.ValidateDirectLoginToken(directLoginToken)
			if err == nil {
				c.Set("user", user)
				c.Set("user_id", user.UserID)
				c.Set("consumer", consumer)
				c.Set("consumer_id", consumer.ConsumerID)
				c.Set("auth_method", "DirectLogin")
				am.recordSuccessfulAttempt(c, user.UserID, "DirectLogin")
				c.Next()
				return
			}
		}

		if oauthToken := extractToken(c); oauthToken != "" {
			user, consumer, err := am.authService.ValidateOAuthToken(oauthToken)
			if err == nil {
				c.Set("user", user)
				c.Set("user_id", user.UserID)
				c.Set("consumer", consumer)
				c.Set("consumer_id", consumer.ConsumerID)
				c.Set("auth_method", "OAuth")
				am.recordSuccessfulAttempt(c, user.UserID, "OAuth")
				c.Next()
				return
			}

			user, err = am.authService.ValidateJWTToken(oauthToken)
			if err == nil {
				c.Set("user", user)
				c.Set("user_id", user.UserID)
				c.Set("auth_method", "JWT")
				am.recordSuccessfulAttempt(c, user.UserID, "JWT")
				c.Next()
				return
			}
		}

		am.recordFailedAttempt(c, "", "MultiAuth", "No valid authentication method")
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Authentication required. Supported methods: OAuth, DirectLogin, JWT",
			"code":  "OBP-20007",
		})
		c.Abort()
	})
}

func (am *AuthMiddleware) RequireEntitlement(requiredRole string) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "User context not found",
				"code":  "OBP-20008",
			})
			c.Abort()
			return
		}

		userObj, ok := user.(*models.User)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Invalid user context",
				"code":  "OBP-50001",
			})
			c.Abort()
			return
		}

		hasEntitlement := am.checkUserEntitlement(userObj.UserID, requiredRole)
		if !hasEntitlement {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Insufficient permissions. Required role: " + requiredRole,
				"code":  "OBP-20009",
			})
			c.Abort()
			return
		}

		c.Next()
	})
}

func (am *AuthMiddleware) RateLimiting() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		consumerID, exists := c.Get("consumer_id")
		if !exists {
			consumerID = c.ClientIP()
		}

		if am.isRateLimited(consumerID.(string)) {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded",
				"code":  "OBP-10001",
			})
			c.Abort()
			return
		}

		c.Next()
	})
}


func extractToken(c *gin.Context) string {
	bearerToken := c.GetHeader("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		token := strings.Split(bearerToken, " ")[1]
		return token
	}
	return ""
}

func extractDirectLoginToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	
	if strings.HasPrefix(authHeader, "DirectLogin") {
		parts := strings.Split(authHeader, "token=")
		if len(parts) == 2 {
			return strings.TrimSpace(parts[1])
		}
	}
	
	directLoginHeader := c.GetHeader("DirectLogin")
	if directLoginHeader != "" {
		parts := strings.Split(directLoginHeader, "token=")
		if len(parts) == 2 {
			return strings.TrimSpace(parts[1])
		}
		return directLoginHeader
	}
	
	return ""
}

func (am *AuthMiddleware) recordSuccessfulAttempt(c *gin.Context, userID, authMethod string) {
	go func() {
		am.authService.RecordLoginAttempt(
			userID,
			"", // username not available in middleware
			c.ClientIP(),
			c.GetHeader("User-Agent"),
			authMethod,
			true,
			"",
		)
	}()
}

func (am *AuthMiddleware) recordFailedAttempt(c *gin.Context, userID, authMethod, reason string) {
	go func() {
		am.authService.RecordLoginAttempt(
			userID,
			"", // username not available in middleware
			c.ClientIP(),
			c.GetHeader("User-Agent"),
			authMethod,
			false,
			reason,
		)
	}()
}

func (am *AuthMiddleware) checkUserEntitlement(userID, requiredRole string) bool {
	return true
}

func (am *AuthMiddleware) isRateLimited(identifier string) bool {
	return false
}

type JWTClaims struct {
	UserID     string `json:"user_id"`
	ConsumerID string `json:"consumer_id,omitempty"`
	AuthMethod string `json:"auth_method"`
	jwt.RegisteredClaims
}

func (am *AuthMiddleware) GenerateJWT(userID, consumerID, authMethod string, duration time.Duration) (string, error) {
	claims := JWTClaims{
		UserID:     userID,
		ConsumerID: consumerID,
		AuthMethod: authMethod,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "OBP-API-Backend",
			Subject:   userID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(am.jwtSecret))
}

func (am *AuthMiddleware) ValidateJWT(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(am.jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrInvalidKey
}
