package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/ashish-019-hash/obp-api-backend/internal/utils"
)

type AtmController struct {
	orchestrationService *services.OrchestrationService
}

func NewAtmController(orchestrationService *services.OrchestrationService) *AtmController {
	return &AtmController{
		orchestrationService: orchestrationService,
	}
}

type AtmAttributeRequest struct {
	Name     string `json:"name" binding:"required"`
	Type     string `json:"type" binding:"required"`
	Value    string `json:"value" binding:"required"`
	IsActive bool   `json:"is_active"`
}

type AtmAttributeResponse struct {
	AtmAttributeID string `json:"atm_attribute_id"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	Value          string `json:"value"`
	IsActive       bool   `json:"is_active"`
}

func (c *AtmController) CreateAtmAttribute(ctx *gin.Context) {
	var req AtmAttributeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	validTypes := []string{"STRING", "INTEGER", "DOUBLE", "DATE_WITH_DAY"}
	isValid := false
	for _, validType := range validTypes {
		if req.Type == validType {
			isValid = true
			break
		}
	}

	if !isValid {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid attribute type", "")
		return
	}

	response := AtmAttributeResponse{
		AtmAttributeID: "attr_" + strconv.FormatInt(time.Now().Unix(), 10),
		Name:           req.Name,
		Type:           req.Type,
		Value:          req.Value,
		IsActive:       req.IsActive,
	}

	utils.SendJSONResponse(ctx, http.StatusCreated, response)
}

func (c *AtmController) GetAtmAttributes(ctx *gin.Context) {
	response := gin.H{
		"atm_attributes": []AtmAttributeResponse{},
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *AtmController) GetAtmAttribute(ctx *gin.Context) {
	atmAttributeId := ctx.Param("ATM_ATTRIBUTE_ID")

	response := AtmAttributeResponse{
		AtmAttributeID: atmAttributeId,
		Name:           "Example Attribute",
		Type:           "STRING",
		Value:          "Example Value",
		IsActive:       true,
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *AtmController) UpdateAtmAttribute(ctx *gin.Context) {
	atmAttributeId := ctx.Param("ATM_ATTRIBUTE_ID")

	var req AtmAttributeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	validTypes := []string{"STRING", "INTEGER", "DOUBLE", "DATE_WITH_DAY"}
	isValid := false
	for _, validType := range validTypes {
		if req.Type == validType {
			isValid = true
			break
		}
	}

	if !isValid {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid attribute type", "")
		return
	}

	response := AtmAttributeResponse{
		AtmAttributeID: atmAttributeId,
		Name:           req.Name,
		Type:           req.Type,
		Value:          req.Value,
		IsActive:       req.IsActive,
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *AtmController) DeleteAtmAttribute(ctx *gin.Context) {
	utils.SendJSONResponse(ctx, http.StatusNoContent, nil)
}
