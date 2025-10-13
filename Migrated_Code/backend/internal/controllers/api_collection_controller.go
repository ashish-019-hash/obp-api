package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/ashish-019-hash/obp-api-backend/internal/utils"
)

type ApiCollectionController struct {
	orchestrationService *services.OrchestrationService
}

func NewApiCollectionController(orchestrationService *services.OrchestrationService) *ApiCollectionController {
	return &ApiCollectionController{
		orchestrationService: orchestrationService,
	}
}

type ApiCollectionResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsShared    bool      `json:"is_shared"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ApiCollectionsResponse struct {
	ApiCollections []ApiCollectionResponse `json:"api_collections"`
}

type CreateApiCollectionRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	IsShared    bool   `json:"is_shared"`
}

func (c *ApiCollectionController) GetApiCollections(ctx *gin.Context) {
	response := ApiCollectionsResponse{
		ApiCollections: []ApiCollectionResponse{
			{
				ID:          "collection_001",
				Name:        "Banking APIs",
				Description: "Core banking functionality",
				IsShared:    true,
				CreatedAt:   time.Now().Add(-24 * time.Hour),
				UpdatedAt:   time.Now(),
			},
			{
				ID:          "collection_002",
				Name:        "Payment APIs",
				Description: "Payment processing endpoints",
				IsShared:    false,
				CreatedAt:   time.Now().Add(-12 * time.Hour),
				UpdatedAt:   time.Now(),
			},
		},
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *ApiCollectionController) CreateApiCollection(ctx *gin.Context) {
	var req CreateApiCollectionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	response := ApiCollectionResponse{
		ID:          "collection_" + strconv.FormatInt(time.Now().Unix(), 10),
		Name:        req.Name,
		Description: req.Description,
		IsShared:    req.IsShared,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	utils.SendJSONResponse(ctx, http.StatusCreated, response)
}

func (c *ApiCollectionController) GetMyApiCollections(ctx *gin.Context) {
	response := ApiCollectionsResponse{
		ApiCollections: []ApiCollectionResponse{
			{
				ID:          "my_collection_001",
				Name:        "My Banking APIs",
				Description: "Personal banking collection",
				IsShared:    false,
				CreatedAt:   time.Now().Add(-24 * time.Hour),
				UpdatedAt:   time.Now(),
			},
		},
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *ApiCollectionController) UpdateMyApiCollection(ctx *gin.Context) {
	collectionId := ctx.Param("API_COLLECTION_ID")

	var req CreateApiCollectionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	response := ApiCollectionResponse{
		ID:          collectionId,
		Name:        req.Name,
		Description: req.Description,
		IsShared:    req.IsShared,
		CreatedAt:   time.Now().Add(-24 * time.Hour),
		UpdatedAt:   time.Now(),
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}
