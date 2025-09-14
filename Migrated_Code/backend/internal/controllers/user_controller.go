package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/ashish-019-hash/obp-api-backend/internal/utils"
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
	provider := ctx.Param("provider")
	username := ctx.Param("username")

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
	provider := ctx.Param("provider")
	providerId := ctx.Param("providerId")

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
