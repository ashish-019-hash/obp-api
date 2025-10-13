package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/ashish-019-hash/obp-api-backend/internal/utils"
)

type TagController struct {
	orchestrationService *services.OrchestrationService
}

func NewTagController(orchestrationService *services.OrchestrationService) *TagController {
	return &TagController{
		orchestrationService: orchestrationService,
	}
}

type TagResponse struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

type TagsResponse struct {
	Tags []TagResponse `json:"tags"`
}

func (c *TagController) GetApiTags(ctx *gin.Context) {
	response := TagsResponse{
		Tags: []TagResponse{
			{
				ID:    "tag_001",
				Value: "Account",
			},
			{
				ID:    "tag_002",
				Value: "Transaction",
			},
			{
				ID:    "tag_003",
				Value: "Customer",
			},
			{
				ID:    "tag_004",
				Value: "Bank",
			},
			{
				ID:    "tag_005",
				Value: "Payment",
			},
		},
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}
