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
	CheckUserEntitlement(userID, roleName string) (bool, error)
	CheckConsumerScope(consumerID, roleName string) (bool, error)
	IsUserLocked(userID string) (bool, error)
	CheckViewPermission(viewID, permissionName string) (bool, error)
	ValidateAuthenticationTypeForOperation(operationID, authType string) (bool, error)
}

type DAuthService interface {
	ValidateDAuthToken(tokenString string) (*models.User, *models.Consumer, error)
	IsDAuthEnabled() bool
}

type GatewayLoginService interface {
	ValidateGatewayToken(tokenString string) (*models.User, *models.Consumer, error)
	IsGatewayLoginEnabled() bool
}

type AuthMiddleware struct {
	authService        AuthService
	dauthService       DAuthService
	gatewayService     GatewayLoginService
	rateLimiter        *services.RateLimiter
	jwtSecret          string
}

func NewAuthMiddleware(authService AuthService, rateLimiter *services.RateLimiter, dauthService DAuthService, gatewayService GatewayLoginService, jwtSecret string) *AuthMiddleware {
	return &AuthMiddleware{
		authService:    authService,
		dauthService:   dauthService,
		gatewayService: gatewayService,
		rateLimiter:    rateLimiter,
		jwtSecret:      jwtSecret,
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

func (am *AuthMiddleware) DAuthAuth() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		if !am.dauthService.IsDAuthEnabled() {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "DAuth authentication is disabled",
				"code":  "OBP-20008",
			})
			c.Abort()
			return
		}

		tokenString := extractDAuthToken(c)
		if tokenString == "" {
			am.recordFailedAttempt(c, "", "DAuth", "No token provided")
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "DAuth token required",
				"code":  "OBP-20009",
			})
			c.Abort()
			return
		}

		user, consumer, err := am.dauthService.ValidateDAuthToken(tokenString)
		if err != nil {
			am.recordFailedAttempt(c, "", "DAuth", err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid DAuth token",
				"code":  "OBP-20010",
			})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Set("user_id", user.UserID)
		c.Set("consumer", consumer)
		c.Set("consumer_id", consumer.ConsumerID)
		c.Set("auth_method", "DAuth")
		
		am.recordSuccessfulAttempt(c, user.UserID, "DAuth")
		c.Next()
	})
}

func (am *AuthMiddleware) GatewayLoginAuth() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		if !am.gatewayService.IsGatewayLoginEnabled() {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Gateway Login authentication is disabled",
				"code":  "OBP-20015",
			})
			c.Abort()
			return
		}

		tokenString := extractGatewayToken(c)
		if tokenString == "" {
			am.recordFailedAttempt(c, "", "GatewayLogin", "No token provided")
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Gateway Login token required",
				"code":  "OBP-20016",
			})
			c.Abort()
			return
		}

		user, consumer, err := am.gatewayService.ValidateGatewayToken(tokenString)
		if err != nil {
			am.recordFailedAttempt(c, "", "GatewayLogin", err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid Gateway Login token",
				"code":  "OBP-20017",
			})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Set("user_id", user.UserID)
		c.Set("consumer", consumer)
		c.Set("consumer_id", consumer.ConsumerID)
		c.Set("auth_method", "GatewayLogin")
		
		am.recordSuccessfulAttempt(c, user.UserID, "GatewayLogin")
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

		if dauthToken := extractDAuthToken(c); dauthToken != "" && am.dauthService.IsDAuthEnabled() {
			user, consumer, err := am.dauthService.ValidateDAuthToken(dauthToken)
			if err == nil {
				c.Set("user", user)
				c.Set("user_id", user.UserID)
				c.Set("consumer", consumer)
				c.Set("consumer_id", consumer.ConsumerID)
				c.Set("auth_method", "DAuth")
				am.recordSuccessfulAttempt(c, user.UserID, "DAuth")
				c.Next()
				return
			}
		}

		if gatewayToken := extractGatewayToken(c); gatewayToken != "" && am.gatewayService.IsGatewayLoginEnabled() {
			user, consumer, err := am.gatewayService.ValidateGatewayToken(gatewayToken)
			if err == nil {
				c.Set("user", user)
				c.Set("user_id", user.UserID)
				c.Set("consumer", consumer)
				c.Set("consumer_id", consumer.ConsumerID)
				c.Set("auth_method", "GatewayLogin")
				am.recordSuccessfulAttempt(c, user.UserID, "GatewayLogin")
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
			"error": "Authentication required. Supported methods: OAuth, DirectLogin, JWT, DAuth, GatewayLogin",
			"code":  "OBP-20007",
		})
		c.Abort()
	})
}

func (am *AuthMiddleware) RequireEntitlement(requiredRole string) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "User context not found",
				"code":  "OBP-20001",
			})
			c.Abort()
			return
		}

		if !am.checkUserEntitlement(userID.(string), requiredRole) {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Insufficient entitlements. Required role: " + requiredRole,
				"code":  "OBP-20006",
			})
			c.Abort()
			return
		}

		c.Next()
	})
}

func (am *AuthMiddleware) RequireScope(requiredScope string) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		consumerID, exists := c.Get("consumer_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Consumer context not found",
				"code":  "OBP-20010",
			})
			c.Abort()
			return
		}

		hasScope, err := am.authService.CheckConsumerScope(consumerID.(string), requiredScope)
		if err != nil || !hasScope {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Insufficient scope. Required scope: " + requiredScope,
				"code":  "OBP-20011",
			})
			c.Abort()
			return
		}

		c.Next()
	})
}

func (am *AuthMiddleware) RequireViewPermission(viewID, permissionName string) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		hasPermission, err := am.authService.CheckViewPermission(viewID, permissionName)
		if err != nil || !hasPermission {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Insufficient view permissions. Required permission: " + permissionName,
				"code":  "OBP-20012",
			})
			c.Abort()
			return
		}

		c.Next()
	})
}

func (am *AuthMiddleware) ValidateAuthType(operationID, authType string) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		isValid, err := am.authService.ValidateAuthenticationTypeForOperation(operationID, authType)
		if err != nil || !isValid {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Authentication type not allowed for this operation",
				"code":  "OBP-20013",
			})
			c.Abort()
			return
		}

		c.Next()
	})
}

func (am *AuthMiddleware) CheckUserLock() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.Next()
			return
		}

		isLocked, err := am.authService.IsUserLocked(userID.(string))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to check user lock status",
				"code":  "OBP-50000",
			})
			c.Abort()
			return
		}

		if isLocked {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "User account is locked",
				"code":  "OBP-20014",
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

func extractDAuthToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	
	if strings.HasPrefix(authHeader, "DAuth ") {
		return strings.TrimSpace(strings.TrimPrefix(authHeader, "DAuth "))
	}
	
	dauthHeader := c.GetHeader("DAuth")
	if dauthHeader != "" {
		return strings.TrimSpace(dauthHeader)
	}
	
	return ""
}

func extractGatewayToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	
	if strings.HasPrefix(authHeader, "GatewayLogin") {
		tokenPart := strings.TrimPrefix(authHeader, "GatewayLogin")
		tokenPart = strings.TrimSpace(tokenPart)
		
		if strings.HasPrefix(tokenPart, "token=") {
			token := strings.TrimPrefix(tokenPart, "token=")
			return strings.Trim(token, "\"")
		}
		
		return tokenPart
	}
	
	gatewayHeader := c.GetHeader("GatewayLogin")
	if gatewayHeader != "" {
		if strings.HasPrefix(gatewayHeader, "token=") {
			token := strings.TrimPrefix(gatewayHeader, "token=")
			return strings.Trim(token, "\"")
		}
		return gatewayHeader
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
	hasEntitlement, err := am.authService.CheckUserEntitlement(userID, requiredRole)
	if err != nil {
		return false
	}
	return hasEntitlement
}

func (am *AuthMiddleware) isRateLimited(identifier string) bool {
	return am.rateLimiter.IsLimited(identifier)
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
