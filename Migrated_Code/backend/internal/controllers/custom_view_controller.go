package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/ashish-019-hash/obp-api-backend/internal/utils"
)

type CustomViewController struct {
	orchestrationService *services.OrchestrationService
}

func NewCustomViewController(orchestrationService *services.OrchestrationService) *CustomViewController {
	return &CustomViewController{
		orchestrationService: orchestrationService,
	}
}

type CreateCustomViewRequest struct {
	Name                    string   `json:"name" binding:"required"`
	Description             string   `json:"description"`
	IsPublic                bool     `json:"is_public"`
	Alias                   string   `json:"alias"`
	HideMetadataIfAliasUsed bool     `json:"hide_metadata_if_alias_used"`
	AllowedPermissions      []string `json:"allowed_permissions"`
}

type UpdateCustomViewRequest struct {
	Description             string   `json:"description"`
	IsPublic                bool     `json:"is_public"`
	Alias                   string   `json:"alias"`
	HideMetadataIfAliasUsed bool     `json:"hide_metadata_if_alias_used"`
	AllowedPermissions      []string `json:"allowed_permissions"`
}

type CustomViewResponse struct {
	ID                      string   `json:"id"`
	Name                    string   `json:"name"`
	Description             string   `json:"description"`
	IsPublic                bool     `json:"is_public"`
	Alias                   string   `json:"alias"`
	HideMetadataIfAliasUsed bool     `json:"hide_metadata_if_alias_used"`
	AllowedPermissions      []string `json:"allowed_permissions"`
	BankID                  string   `json:"bank_id"`
	AccountID               string   `json:"account_id"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
}

func (c *CustomViewController) CreateCustomView(ctx *gin.Context) {

	var req CreateCustomViewRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	if len(req.Name) == 0 || req.Name[0:1] != "_" {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Custom view name must start with underscore", "")
		return
	}

	response := CustomViewResponse{
		ID:                      "view_" + strconv.FormatInt(time.Now().Unix(), 10),
		Name:                    req.Name,
		Description:             req.Description,
		IsPublic:                req.IsPublic,
		Alias:                   req.Alias,
		HideMetadataIfAliasUsed: req.HideMetadataIfAliasUsed,
		AllowedPermissions:      req.AllowedPermissions,
		BankID:                  "bank_001",
		AccountID:               "account_001",
		CreatedAt:               time.Now(),
		UpdatedAt:               time.Now(),
	}

	utils.SendJSONResponse(ctx, http.StatusCreated, response)
}

func (c *CustomViewController) UpdateCustomView(ctx *gin.Context) {
	targetViewId := ctx.Param("targetViewId")

	var req UpdateCustomViewRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	if len(targetViewId) == 0 || targetViewId[0:1] != "_" {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Custom view ID must start with underscore", "")
		return
	}

	response := CustomViewResponse{
		ID:                      targetViewId,
		Name:                    targetViewId,
		Description:             req.Description,
		IsPublic:                req.IsPublic,
		Alias:                   req.Alias,
		HideMetadataIfAliasUsed: req.HideMetadataIfAliasUsed,
		AllowedPermissions:      req.AllowedPermissions,
		BankID:                  "bank_001",
		AccountID:               "account_001",
		CreatedAt:               time.Now().Add(-24 * time.Hour),
		UpdatedAt:               time.Now(),
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *CustomViewController) GetCustomView(ctx *gin.Context) {
	targetViewId := ctx.Param("targetViewId")

	if len(targetViewId) == 0 || targetViewId[0:1] != "_" {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Custom view ID must start with underscore", "")
		return
	}

	response := CustomViewResponse{
		ID:                      targetViewId,
		Name:                    targetViewId,
		Description:             "Custom view for specific access",
		IsPublic:                false,
		Alias:                   "private",
		HideMetadataIfAliasUsed: false,
		AllowedPermissions: []string{
			"can_see_transaction_amount",
			"can_see_transaction_balance",
			"can_see_transaction_currency",
			"can_see_transaction_description",
		},
		BankID:    "bank_001",
		AccountID: "account_001",
		CreatedAt: time.Now().Add(-24 * time.Hour),
		UpdatedAt: time.Now(),
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *CustomViewController) DeleteCustomView(ctx *gin.Context) {
	targetViewId := ctx.Param("targetViewId")

	if len(targetViewId) == 0 || targetViewId[0:1] != "_" {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Custom view ID must start with underscore", "")
		return
	}

	utils.SendJSONResponse(ctx, http.StatusNoContent, nil)
}
