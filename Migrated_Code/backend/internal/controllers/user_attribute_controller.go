package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/ashish-019-hash/obp-api-backend/internal/utils"
)

type UserAttributeController struct {
	orchestrationService *services.OrchestrationService
}

func NewUserAttributeController(orchestrationService *services.OrchestrationService) *UserAttributeController {
	return &UserAttributeController{
		orchestrationService: orchestrationService,
	}
}

type CreateUserAttributeRequest struct {
	Name  string `json:"name" binding:"required"`
	Type  string `json:"type" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type UserAttributeResponse struct {
	UserAttributeId string    `json:"user_attribute_id"`
	UserId          string    `json:"user_id"`
	Name            string    `json:"name"`
	Type            string    `json:"type"`
	Value           string    `json:"value"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type UserAttributesResponse struct {
	UserAttributes []UserAttributeResponse `json:"user_attributes"`
}

func (c *UserAttributeController) CreateNonPersonalUserAttribute(ctx *gin.Context) {
	userId := ctx.Param("USER_ID")

	var req CreateUserAttributeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	validTypes := []string{"STRING", "INTEGER", "DOUBLE", "DATE_WITH_DAY"}
	isValidType := false
	for _, validType := range validTypes {
		if req.Type == validType {
			isValidType = true
			break
		}
	}

	if !isValidType {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid attribute type", "Type must be one of: STRING, INTEGER, DOUBLE, DATE_WITH_DAY")
		return
	}

	response := UserAttributeResponse{
		UserAttributeId: "attr_" + strconv.FormatInt(time.Now().Unix(), 10),
		UserId:          userId,
		Name:            req.Name,
		Type:            req.Type,
		Value:           req.Value,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	utils.SendJSONResponse(ctx, http.StatusCreated, response)
}

func (c *UserAttributeController) DeleteNonPersonalUserAttribute(ctx *gin.Context) {
	utils.SendJSONResponse(ctx, http.StatusNoContent, nil)
}

func (c *UserAttributeController) GetNonPersonalUserAttributes(ctx *gin.Context) {
	userId := ctx.Param("USER_ID")

	response := UserAttributesResponse{
		UserAttributes: []UserAttributeResponse{
			{
				UserAttributeId: "attr_001",
				UserId:          userId,
				Name:            "TAX_NUMBER",
				Type:            "STRING",
				Value:           "123456789",
				CreatedAt:       time.Now().Add(-24 * time.Hour),
				UpdatedAt:       time.Now(),
			},
			{
				UserAttributeId: "attr_002",
				UserId:          userId,
				Name:            "CREDIT_SCORE",
				Type:            "INTEGER",
				Value:           "750",
				CreatedAt:       time.Now().Add(-12 * time.Hour),
				UpdatedAt:       time.Now(),
			},
		},
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *UserAttributeController) GetAccountsHeldByUserAtBank(ctx *gin.Context) {
	userId := ctx.Param("USER_ID")
	bankId := ctx.Param("BANK_ID")

	response := gin.H{
		"user_id": userId,
		"bank_id": bankId,
		"accounts": []gin.H{
			{
				"id":           "account_001",
				"bank_id":      bankId,
				"label":        "Savings Account",
				"number":       "123456789",
				"type":         "SAVINGS",
				"balance": gin.H{
					"currency": "EUR",
					"amount":   "1500.50",
				},
				"account_routings": []gin.H{
					{
						"scheme":  "IBAN",
						"address": "DE89370400440532013000",
					},
				},
			},
		},
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *UserAttributeController) GetAccountsHeldByUser(ctx *gin.Context) {
	userId := ctx.Param("USER_ID")

	response := gin.H{
		"user_id": userId,
		"accounts": []gin.H{
			{
				"id":           "account_001",
				"bank_id":      "bank_001",
				"label":        "Savings Account",
				"number":       "123456789",
				"type":         "SAVINGS",
				"balance": gin.H{
					"currency": "EUR",
					"amount":   "1500.50",
				},
			},
			{
				"id":           "account_002",
				"bank_id":      "bank_002",
				"label":        "Checking Account",
				"number":       "987654321",
				"type":         "CURRENT",
				"balance": gin.H{
					"currency": "GBP",
					"amount":   "2300.75",
				},
			},
		},
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}
