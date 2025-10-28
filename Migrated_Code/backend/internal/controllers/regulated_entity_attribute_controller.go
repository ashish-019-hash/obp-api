package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/ashish-019-hash/obp-api-backend/internal/utils"
)

type RegulatedEntityAttributeController struct {
	orchestrationService *services.OrchestrationService
}

func NewRegulatedEntityAttributeController(orchestrationService *services.OrchestrationService) *RegulatedEntityAttributeController {
	return &RegulatedEntityAttributeController{
		orchestrationService: orchestrationService,
	}
}

type RegulatedEntityAttributeRequest struct {
	Name          string `json:"name" binding:"required"`
	AttributeType string `json:"attribute_type" binding:"required"`
	Value         string `json:"value" binding:"required"`
	IsActive      bool   `json:"is_active"`
}

type RegulatedEntityAttributeResponse struct {
	AttributeID   string    `json:"attribute_id"`
	EntityID      string    `json:"entity_id"`
	Name          string    `json:"name"`
	AttributeType string    `json:"attribute_type"`
	Value         string    `json:"value"`
	IsActive      bool      `json:"is_active"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type RegulatedEntityAttributesResponse struct {
	Attributes []RegulatedEntityAttributeResponse `json:"attributes"`
}

func (c *RegulatedEntityAttributeController) CreateRegulatedEntityAttribute(ctx *gin.Context) {
	entityId := ctx.Param("entityId")

	var req RegulatedEntityAttributeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	validTypes := []string{"STRING", "INTEGER", "DOUBLE", "DATE_WITH_DAY"}
	isValid := false
	for _, validType := range validTypes {
		if req.AttributeType == validType {
			isValid = true
			break
		}
	}

	if !isValid {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid attribute type", "")
		return
	}

	response := RegulatedEntityAttributeResponse{
		AttributeID:   "attr_" + strconv.FormatInt(time.Now().Unix(), 10),
		EntityID:      entityId,
		Name:          req.Name,
		AttributeType: req.AttributeType,
		Value:         req.Value,
		IsActive:      req.IsActive,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	utils.SendJSONResponse(ctx, http.StatusCreated, response)
}

func (c *RegulatedEntityAttributeController) DeleteRegulatedEntityAttribute(ctx *gin.Context) {
	utils.SendJSONResponse(ctx, http.StatusNoContent, nil)
}

func (c *RegulatedEntityAttributeController) GetRegulatedEntityAttributeById(ctx *gin.Context) {
	entityId := ctx.Param("entityId")
	attributeId := ctx.Param("attributeId")

	response := RegulatedEntityAttributeResponse{
		AttributeID:   attributeId,
		EntityID:      entityId,
		Name:          "Example Attribute",
		AttributeType: "STRING",
		Value:         "Example Value",
		IsActive:      true,
		CreatedAt:     time.Now().Add(-24 * time.Hour),
		UpdatedAt:     time.Now(),
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *RegulatedEntityAttributeController) GetAllRegulatedEntityAttributes(ctx *gin.Context) {
	entityId := ctx.Param("entityId")

	response := RegulatedEntityAttributesResponse{
		Attributes: []RegulatedEntityAttributeResponse{
			{
				AttributeID:   "attr_001",
				EntityID:      entityId,
				Name:          "License Number",
				AttributeType: "STRING",
				Value:         "LIC123456",
				IsActive:      true,
				CreatedAt:     time.Now().Add(-48 * time.Hour),
				UpdatedAt:     time.Now().Add(-24 * time.Hour),
			},
			{
				AttributeID:   "attr_002",
				EntityID:      entityId,
				Name:          "Establishment Date",
				AttributeType: "DATE_WITH_DAY",
				Value:         "2020-01-15",
				IsActive:      true,
				CreatedAt:     time.Now().Add(-48 * time.Hour),
				UpdatedAt:     time.Now().Add(-24 * time.Hour),
			},
		},
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *RegulatedEntityAttributeController) UpdateRegulatedEntityAttribute(ctx *gin.Context) {
	entityId := ctx.Param("entityId")
	attributeId := ctx.Param("attributeId")

	var req RegulatedEntityAttributeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	validTypes := []string{"STRING", "INTEGER", "DOUBLE", "DATE_WITH_DAY"}
	isValid := false
	for _, validType := range validTypes {
		if req.AttributeType == validType {
			isValid = true
			break
		}
	}

	if !isValid {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid attribute type", "")
		return
	}

	response := RegulatedEntityAttributeResponse{
		AttributeID:   attributeId,
		EntityID:      entityId,
		Name:          req.Name,
		AttributeType: req.AttributeType,
		Value:         req.Value,
		IsActive:      req.IsActive,
		CreatedAt:     time.Now().Add(-24 * time.Hour),
		UpdatedAt:     time.Now(),
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}
