package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/ashish-019-hash/obp-api-backend/internal/models"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/ashish-019-hash/obp-api-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *services.AuthenticationService
}

func NewAuthController(authService *services.AuthenticationService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

type DirectLoginRequest struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	ConsumerKey string `json:"consumer_key" binding:"required"`
}

type DirectLoginResponse struct {
	Token     string `json:"token"`
	TokenType string `json:"token_type"`
	ExpiresIn int64  `json:"expires_in"`
}

type ConsumerRegistrationRequest struct {
	Name           string `json:"name" binding:"required"`
	Description    string `json:"description"`
	DeveloperEmail string `json:"developer_email" binding:"required,email"`
	RedirectURL    string `json:"redirect_url"`
	AppType        string `json:"app_type"`
}

type ConsumerRegistrationResponse struct {
	ConsumerID     string `json:"consumer_id"`
	ConsumerKey    string `json:"consumer_key"`
	ConsumerSecret string `json:"consumer_secret"`
	Name           string `json:"name"`
	AppType        string `json:"app_type"`
	DeveloperEmail string `json:"developer_email"`
	RedirectURL    string `json:"redirect_url"`
	IsActive       bool   `json:"is_active"`
	CreatedAt      string `json:"created_at"`
}

type UserInfoResponse struct {
	UserID       string `json:"user_id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Provider     string `json:"provider"`
	ProviderID   string `json:"provider_id"`
	IsActive     bool   `json:"is_active"`
	ConsentGiven bool   `json:"consent_given"`
}

func (ac *AuthController) DirectLogin(c *gin.Context) {
	var req DirectLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
		return
	}

	token, err := ac.authService.CreateDirectLoginToken(req.Username, req.Password, req.ConsumerKey)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusUnauthorized, "Authentication failed", err.Error())
		return
	}

	tokenConfig, _ := ac.authService.GetTokenConfiguration("DirectLogin")
	expiresIn := int64(tokenConfig.ExpirationSeconds)

	response := DirectLoginResponse{
		Token:     token,
		TokenType: "DirectLogin",
		ExpiresIn: expiresIn,
	}

	utils.SendJSONResponse(c, http.StatusOK, response)
}

func (ac *AuthController) RegisterConsumer(c *gin.Context) {
	var req ConsumerRegistrationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
		return
	}

	if req.AppType == "" {
		req.AppType = "Web"
	}

	consumer, err := ac.authService.CreateConsumer(req.Name, req.Description, req.DeveloperEmail, req.RedirectURL, req.AppType)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to create consumer", err.Error())
		return
	}

	response := ConsumerRegistrationResponse{
		ConsumerID:     consumer.ConsumerID,
		ConsumerKey:    consumer.ConsumerKey,
		ConsumerSecret: consumer.ConsumerSecret,
		Name:           consumer.Name,
		AppType:        consumer.AppType,
		DeveloperEmail: consumer.DeveloperEmail,
		RedirectURL:    consumer.RedirectURL,
		IsActive:       consumer.IsActive,
		CreatedAt:      consumer.CreatedAt.Format(time.RFC3339),
	}

	utils.SendJSONResponse(c, http.StatusCreated, response)
}

func (ac *AuthController) GetCurrentUser(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		utils.SendErrorResponse(c, http.StatusUnauthorized, "User not authenticated", "")
		return
	}

	userObj, ok := user.(*models.User)
	if !ok {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Invalid user context", "")
		return
	}

	response := UserInfoResponse{
		UserID:       userObj.UserID,
		Username:     userObj.UserID, // Using UserID as username for now
		Email:        userObj.Email,
		FirstName:    userObj.FirstName,
		LastName:     userObj.LastName,
		Provider:     userObj.Provider,
		ProviderID:   userObj.ProviderID,
		IsActive:     userObj.IsActive,
		ConsentGiven: userObj.ConsentGiven,
	}

	utils.SendJSONResponse(c, http.StatusOK, response)
}

func (ac *AuthController) OAuthInitiate(c *gin.Context) {
	consumerKey := c.Query("oauth_consumer_key")
	callbackURL := c.Query("oauth_callback")

	if consumerKey == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Missing oauth_consumer_key", "")
		return
	}

	requestToken, err := ac.authService.CreateOAuthRequestToken(consumerKey, callbackURL)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusUnauthorized, "Failed to create request token", err.Error())
		return
	}

	response := map[string]string{
		"oauth_token":              requestToken.TokenValue,
		"oauth_token_secret":       requestToken.TokenSecret,
		"oauth_callback_confirmed": "true",
	}

	utils.SendJSONResponse(c, http.StatusOK, response)
}

func (ac *AuthController) OAuthToken(c *gin.Context) {
	oauthToken := c.Query("oauth_token")
	oauthVerifier := c.Query("oauth_verifier")

	if oauthToken == "" || oauthVerifier == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Missing oauth_token or oauth_verifier", "")
		return
	}

	accessToken, err := ac.authService.CreateOAuthAccessToken(oauthToken, oauthVerifier)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusUnauthorized, "Failed to create access token", err.Error())
		return
	}

	response := map[string]string{
		"oauth_token":        accessToken.TokenValue,
		"oauth_token_secret": accessToken.TokenSecret,
	}

	utils.SendJSONResponse(c, http.StatusOK, response)
}

func (ac *AuthController) OAuthAuthorize(c *gin.Context) {
	oauthToken := c.Query("oauth_token")
	if oauthToken == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Missing oauth_token", "")
		return
	}

	verifier, err := ac.authService.AuthorizeOAuthToken(oauthToken)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid oauth_token", err.Error())
		return
	}

	response := map[string]string{
		"oauth_token":    oauthToken,
		"oauth_verifier": verifier,
		"status":         "authorized",
	}

	utils.SendJSONResponse(c, http.StatusOK, response)
}

func (ac *AuthController) CreateUser(c *gin.Context) {
	var req struct {
		Username    string `json:"username" binding:"required"`
		Password    string `json:"password" binding:"required"`
		Email       string `json:"email" binding:"required,email"`
		FirstName   string `json:"first_name" binding:"required"`
		LastName    string `json:"last_name" binding:"required"`
		Provider    string `json:"provider"`
		ProviderID  string `json:"provider_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
		return
	}

	if req.Provider == "" {
		req.Provider = "local"
	}
	if req.ProviderID == "" {
		req.ProviderID = req.Username
	}

	user, err := ac.authService.CreateUser(req.Username, req.Password, req.Email, req.FirstName, req.LastName, req.Provider, req.ProviderID)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to create user", err.Error())
		return
	}

	response := UserInfoResponse{
		UserID:       user.UserID,
		Username:     req.Username,
		Email:        user.Email,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Provider:     user.Provider,
		ProviderID:   user.ProviderID,
		IsActive:     user.IsActive,
		ConsentGiven: user.ConsentGiven,
	}

	utils.SendJSONResponse(c, http.StatusCreated, response)
}

func (ac *AuthController) GetLoginAttempts(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "50")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 50
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	attempts, err := ac.authService.GetLoginAttempts(limit, offset)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve login attempts", err.Error())
		return
	}

	utils.SendJSONResponse(c, http.StatusOK, map[string]interface{}{
		"login_attempts": attempts,
		"limit":          limit,
		"offset":         offset,
	})
}
