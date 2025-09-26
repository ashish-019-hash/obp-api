package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/ashish-019-hash/obp-api-backend/internal/utils"
)

type AuthContextController struct {
	authService *services.AuthenticationService
}

func NewAuthContextController(authService *services.AuthenticationService) *AuthContextController {
	return &AuthContextController{
		authService: authService,
	}
}

type CreateAuthContextRequest struct {
	Key   string `json:"key" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type AuthContextResponse struct {
	ContextID  string    `json:"context_id"`
	UserID     string    `json:"user_id"`
	ConsumerID string    `json:"consumer_id"`
	Key        string    `json:"key"`
	Value      string    `json:"value"`
	CreatedAt  time.Time `json:"created_at"`
}

func (c *AuthContextController) CreateUserAuthContext(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		utils.SendErrorResponse(ctx, http.StatusUnauthorized, "User context not found", "")
		return
	}

	consumerID, exists := ctx.Get("consumer_id")
	if !exists {
		utils.SendErrorResponse(ctx, http.StatusUnauthorized, "Consumer context not found", "")
		return
	}

	var req CreateAuthContextRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	if err := c.authService.CreateUserAuthContext(userID.(string), consumerID.(string), req.Key, req.Value); err != nil {
		utils.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to create auth context", err.Error())
		return
	}

	response := AuthContextResponse{
		ContextID:  "ctx_" + userID.(string) + "_" + req.Key,
		UserID:     userID.(string),
		ConsumerID: consumerID.(string),
		Key:        req.Key,
		Value:      req.Value,
		CreatedAt:  time.Now(),
	}

	utils.SendJSONResponse(ctx, http.StatusCreated, response)
}

func (c *AuthContextController) GetUserAuthContexts(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		utils.SendErrorResponse(ctx, http.StatusUnauthorized, "User context not found", "")
		return
	}

	contexts, err := c.authService.GetUserAuthContexts(userID.(string))
	if err != nil {
		utils.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve auth contexts", err.Error())
		return
	}

	contextResponses := make([]AuthContextResponse, len(contexts))
	for i, context := range contexts {
		contextResponses[i] = AuthContextResponse{
			ContextID:  context.ContextID,
			UserID:     context.UserID,
			ConsumerID: context.ConsumerID,
			Key:        context.Key,
			Value:      context.Value,
			CreatedAt:  context.CreatedAt,
		}
	}

	utils.SendJSONResponse(ctx, http.StatusOK, gin.H{
		"contexts": contextResponses,
	})
}
