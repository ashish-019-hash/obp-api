package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/ashish-019-hash/obp-api-backend/internal/utils"
)

type SystemViewController struct {
	orchestrationService *services.OrchestrationService
}

func NewSystemViewController(orchestrationService *services.OrchestrationService) *SystemViewController {
	return &SystemViewController{
		orchestrationService: orchestrationService,
	}
}

type CreateViewPermissionRequest struct {
	PermissionName string `json:"permission_name" binding:"required"`
	ExtraData      string `json:"extra_data"`
}

type ViewPermissionResponse struct {
	PermissionID   string    `json:"permission_id"`
	ViewID         string    `json:"view_id"`
	PermissionName string    `json:"permission_name"`
	ExtraData      string    `json:"extra_data"`
	CreatedAt      time.Time `json:"created_at"`
}

func (c *SystemViewController) AddSystemViewPermission(ctx *gin.Context) {
	viewId := ctx.Param("VIEW_ID")

	var req CreateViewPermissionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	validPermissions := []string{
		"can_see_transaction_amount",
		"can_see_transaction_balance",
		"can_see_transaction_currency",
		"can_see_transaction_description",
		"can_see_transaction_start_date",
		"can_see_transaction_finish_date",
		"can_see_transaction_type",
		"can_see_transaction_metadata",
		"can_see_other_account_bank_name",
		"can_see_other_account_number",
		"can_see_other_account_kind",
		"can_see_other_account_iban",
		"can_see_other_account_routing_scheme",
		"can_see_other_account_routing_address",
		"can_see_other_account_national_identifier",
		"can_see_other_account_metadata",
		"can_add_comment",
		"can_add_tag",
		"can_add_image_url",
		"can_add_url",
		"can_add_more_info",
		"can_add_private_alias",
		"can_add_public_alias",
		"can_add_corporate_location",
		"can_add_physical_location",
		"can_add_where_tag",
		"can_delete_comment",
		"can_delete_tag",
		"can_delete_image_url",
		"can_delete_url",
		"can_delete_more_info",
		"can_delete_corporate_location",
		"can_delete_physical_location",
		"can_delete_where_tag",
		"can_edit_owner_comment",
		"can_see_image_url",
		"can_see_url",
		"can_see_more_info",
		"can_see_physical_location",
		"can_see_private_alias",
		"can_see_public_alias",
		"can_see_tag",
		"can_see_where_tag",
		"can_see_owner_comment",
		"can_see_open_corporates_url",
		"can_add_open_corporates_url",
		"can_delete_open_corporates_url",
	}

	isValid := false
	for _, permission := range validPermissions {
		if req.PermissionName == permission {
			isValid = true
			break
		}
	}

	if !isValid {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid permission name", "")
		return
	}

	response := ViewPermissionResponse{
		PermissionID:   "perm_" + strconv.FormatInt(time.Now().Unix(), 10),
		ViewID:         viewId,
		PermissionName: req.PermissionName,
		ExtraData:      req.ExtraData,
		CreatedAt:      time.Now(),
	}

	utils.SendJSONResponse(ctx, http.StatusCreated, response)
}

func (c *SystemViewController) DeleteSystemViewPermission(ctx *gin.Context) {

	utils.SendJSONResponse(ctx, http.StatusNoContent, nil)
}
