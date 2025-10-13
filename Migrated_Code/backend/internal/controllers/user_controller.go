package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/ashish-019-hash/obp-api-backend/internal/utils"
	"github.com/ashish-019-hash/obp-api-backend/internal/models"
	"github.com/ashish-019-hash/obp-api-backend/pkg/db"
)

type UserController struct {
	orchestrationService *services.OrchestrationService
}

func NewUserController(orchestrationService *services.OrchestrationService) *UserController {
	return &UserController{
		orchestrationService: orchestrationService,
	}
}

type UserResponse struct {
	UserID       string `json:"user_id"`
	Username     string `json:"username"`
	Provider     string `json:"provider"`
	Email        string `json:"email"`
	DisplayName  string `json:"display_name"`
	IsLocked     bool   `json:"is_locked"`
	IsValidated  bool   `json:"is_validated"`
}

type UserLockStatusResponse struct {
	IsLocked      bool      `json:"is_locked"`
	LockReason    string    `json:"lock_reason"`
	LockedAt      time.Time `json:"locked_at"`
	BadAttempts   int       `json:"bad_attempts"`
}

type UserValidatedResponse struct {
	IsValidated bool `json:"is_validated"`
}

func (c *UserController) GetUserByProviderAndUsername(ctx *gin.Context) {
	provider := ctx.Param("PROVIDER")
	username := ctx.Param("USERNAME")

	response := UserResponse{
		UserID:      "user_" + strconv.FormatInt(time.Now().Unix(), 10),
		Username:    username,
		Provider:    provider,
		Email:       username + "@example.com",
		DisplayName: username,
		IsLocked:    false,
		IsValidated: true,
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *UserController) GetUserLockStatus(ctx *gin.Context) {

	response := UserLockStatusResponse{
		IsLocked:    false,
		LockReason:  "",
		LockedAt:    time.Now(),
		BadAttempts: 0,
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *UserController) UnlockUserByProviderAndUsername(ctx *gin.Context) {

	response := UserLockStatusResponse{
		IsLocked:    false,
		LockReason:  "",
		LockedAt:    time.Now(),
		BadAttempts: 0,
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *UserController) LockUserByProviderAndUsername(ctx *gin.Context) {

	response := UserLockStatusResponse{
		IsLocked:    true,
		LockReason:  "Manual lock",
		LockedAt:    time.Now(),
		BadAttempts: 0,
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *UserController) ValidateUserByUserId(ctx *gin.Context) {

	response := UserValidatedResponse{
		IsValidated: true,
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *UserController) SyncExternalUser(ctx *gin.Context) {
	provider := ctx.Param("PROVIDER")
	providerId := ctx.Param("PROVIDER_ID")

	response := UserResponse{
		UserID:      "user_" + strconv.FormatInt(time.Now().Unix(), 10),
		Username:    providerId,
		Provider:    provider,
		Email:       providerId + "@" + provider,
		DisplayName: providerId,
		IsLocked:    false,
		IsValidated: true,
	}

	utils.SendJSONResponse(ctx, http.StatusCreated, response)
}

type UsersResponse struct {
	Users []UserResponse `json:"users"`
}

func (c *UserController) GetUsers(ctx *gin.Context) {
	var users []models.User
	if err := db.GetDB().Find(&users).Error; err != nil {
		utils.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve users", err.Error())
		return
	}

	userResponses := make([]UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = UserResponse{
			UserID:      user.UserID,
			Username:    user.Name,
			Provider:    user.Provider,
			Email:       user.EmailAddress,
			DisplayName: user.Name,
			IsLocked:    false,
			IsValidated: true,
		}
	}

	response := UsersResponse{
		Users: userResponses,
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

type CreateUserRequest struct {
	Username   string `json:"username" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
	FirstName  string `json:"first_name" binding:"required"`
	LastName   string `json:"last_name" binding:"required"`
	Provider   string `json:"provider"`
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var req CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	if req.Provider == "" {
		req.Provider = "local"
	}

	var existingUser models.User
	if err := db.GetDB().Where("email_address = ? OR name = ?", req.Email, req.Username).First(&existingUser).Error; err == nil {
		utils.SendErrorResponse(ctx, http.StatusConflict, "User already exists", "A user with this email or username already exists")
		return
	}

	userID := "user_" + strconv.FormatInt(time.Now().Unix(), 10)
	user := models.User{
		UserID:         userID,
		Provider:       req.Provider,
		EmailAddress:   req.Email,
		Name:           req.Username,
		IsOriginalUser: true,
		IsConsentUser:  false,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := db.GetDB().Create(&user).Error; err != nil {
		utils.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to create user", err.Error())
		return
	}

	response := UserResponse{
		UserID:      user.UserID,
		Username:    user.Name,
		Provider:    user.Provider,
		Email:       user.EmailAddress,
		DisplayName: req.FirstName + " " + req.LastName,
		IsLocked:    false,
		IsValidated: true,
	}

	utils.SendJSONResponse(ctx, http.StatusCreated, response)
}
