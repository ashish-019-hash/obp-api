package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/ashish-019-hash/obp-api-backend/internal/utils"
)

type ScopeController struct {
	authService *services.AuthenticationService
}

func NewScopeController(authService *services.AuthenticationService) *ScopeController {
	return &ScopeController{
		authService: authService,
	}
}

type CreateScopeRequest struct {
	ConsumerID string  `json:"consumer_id" binding:"required"`
	RoleName   string  `json:"role_name" binding:"required"`
	BankID     *string `json:"bank_id"`
}

type ScopeResponse struct {
	ScopeID    string    `json:"scope_id"`
	ConsumerID string    `json:"consumer_id"`
	RoleName   string    `json:"role_name"`
	BankID     string    `json:"bank_id"`
	IsActive   bool      `json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
}

func (c *ScopeController) CreateScope(ctx *gin.Context) {
	var req CreateScopeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	if err := c.authService.CreateScope(req.ConsumerID, req.RoleName, req.BankID); err != nil {
		utils.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to create scope", err.Error())
		return
	}

	response := ScopeResponse{
		ScopeID:    "scope_" + req.ConsumerID + "_" + req.RoleName,
		ConsumerID: req.ConsumerID,
		RoleName:   req.RoleName,
		IsActive:   true,
		CreatedAt:  time.Now(),
	}

	if req.BankID != nil {
		response.BankID = *req.BankID
	}

	utils.SendJSONResponse(ctx, http.StatusCreated, response)
}

func (c *ScopeController) GetConsumerScopes(ctx *gin.Context) {
	consumerID := ctx.Param("consumerId")

	scopes, err := c.authService.GetScopesByConsumerID(consumerID)
	if err != nil {
		utils.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve scopes", err.Error())
		return
	}

	scopeResponses := make([]ScopeResponse, len(scopes))
	for i, scope := range scopes {
		scopeResponses[i] = ScopeResponse{
			ScopeID:    scope.ScopeID,
			ConsumerID: scope.ConsumerID,
			RoleName:   scope.RoleName,
			BankID:     scope.BankID,
			IsActive:   scope.IsActive,
			CreatedAt:  scope.CreatedAt,
		}
	}

	utils.SendJSONResponse(ctx, http.StatusOK, gin.H{
		"scopes": scopeResponses,
	})
}
