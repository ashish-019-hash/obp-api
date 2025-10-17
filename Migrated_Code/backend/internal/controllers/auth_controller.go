package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"obp-api-backend/internal/models"
	"obp-api-backend/internal/services"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (ctrl *AuthController) DirectLogin(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if !strings.HasPrefix(authHeader, "DirectLogin ") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid authorization header format"})
		return
	}

	params := parseDirectLoginHeader(strings.TrimPrefix(authHeader, "DirectLogin "))

	username, ok := params["username"]
	if !ok || username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username is required"})
		return
	}

	password, ok := params["password"]
	if !ok || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password is required"})
		return
	}

	consumerKey, ok := params["consumer_key"]
	if !ok || consumerKey == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "consumer_key is required"})
		return
	}

	token, err := ctrl.authService.DirectLogin(
		c.Request.Context(),
		username,
		password,
		consumerKey,
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

type CreateUserRequest struct {
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (ctrl *AuthController) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authUser, resourceUser, err := ctrl.authService.CreateUser(
		c.Request.Context(),
		req.Username,
		req.Email,
		req.Password,
		req.FirstName,
		req.LastName,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"auth_user":     authUser,
		"resource_user": resourceUser,
	})
}

type CreateConsumerRequest struct {
	Name           string `json:"name" binding:"required"`
	AppType        string `json:"app_type"`
	Description    string `json:"description"`
	DeveloperEmail string `json:"developer_email" binding:"required,email"`
}

func (ctrl *AuthController) CreateConsumer(c *gin.Context) {
	var req CreateConsumerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, exists := c.Get("user")
	var createdByUserID *string
	if exists {
		resourceUser := user.(*models.ResourceUser)
		createdByUserID = &resourceUser.UserID
	}

	consumer, err := ctrl.authService.CreateConsumer(
		c.Request.Context(),
		req.Name,
		req.AppType,
		req.Description,
		req.DeveloperEmail,
		createdByUserID,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"consumer_id":     consumer.ConsumerID,
		"consumer_key":    consumer.ConsumerKey,
		"consumer_secret": consumer.ConsumerSecret,
		"message":         "Please save the consumer_secret as it will not be shown again",
	})
}

func (ctrl *AuthController) GetCurrentUser(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not authenticated"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func parseDirectLoginHeader(header string) map[string]string {
	params := make(map[string]string)
	parts := strings.Split(header, ",")
	for _, part := range parts {
		kv := strings.SplitN(strings.TrimSpace(part), "=", 2)
		if len(kv) == 2 {
			params[kv[0]] = strings.Trim(kv[1], "\"")
		}
	}
	return params
}
