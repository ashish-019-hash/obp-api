package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/ashish-019-hash/obp-api-backend/internal/utils"
)

type EntitlementController struct {
	orchestrationService *services.OrchestrationService
}

func NewEntitlementController(orchestrationService *services.OrchestrationService) *EntitlementController {
	return &EntitlementController{
		orchestrationService: orchestrationService,
	}
}

type EntitlementResponse struct {
	EntitlementId string `json:"entitlement_id"`
	RoleName      string `json:"role_name"`
	BankId        string `json:"bank_id"`
}

type PermissionResponse struct {
	Permission string `json:"permission"`
	ViewId     string `json:"view_id"`
}

type EntitlementsAndPermissionsResponse struct {
	Entitlements []EntitlementResponse `json:"entitlements"`
	Permissions  []PermissionResponse  `json:"permissions"`
}

func (c *EntitlementController) GetEntitlementsAndPermissions(ctx *gin.Context) {
	response := EntitlementsAndPermissionsResponse{
		Entitlements: []EntitlementResponse{
			{
				EntitlementId: "ent_001",
				RoleName:      "CanGetAnyUser",
				BankId:        "",
			},
			{
				EntitlementId: "ent_002",
				RoleName:      "CanCreateAccount",
				BankId:        "bank_001",
			},
		},
		Permissions: []PermissionResponse{
			{
				Permission: "can_see_transaction_amount",
				ViewId:     "owner",
			},
			{
				Permission: "can_see_transaction_balance",
				ViewId:     "owner",
			},
		},
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}
