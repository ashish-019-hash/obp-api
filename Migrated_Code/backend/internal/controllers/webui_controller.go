package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/ashish-019-hash/obp-api-backend/internal/utils"
)

type WebUIController struct {
	orchestrationService *services.OrchestrationService
}

func NewWebUIController(orchestrationService *services.OrchestrationService) *WebUIController {
	return &WebUIController{
		orchestrationService: orchestrationService,
	}
}

type WebUIPropsResponse struct {
	WebUIProps []WebUIProp `json:"webui_props"`
}

type WebUIProp struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	ID    string `json:"id"`
}

func (c *WebUIController) GetWebUIProps(ctx *gin.Context) {
	active := ctx.DefaultQuery("active", "false")

	response := WebUIPropsResponse{
		WebUIProps: []WebUIProp{
			{
				Name:  "webui_api_explorer_url",
				Value: "https://apiexplorer.openbankproject.com",
				ID:    "webui_prop_001",
			},
			{
				Name:  "webui_api_documentation_url",
				Value: "https://github.com/OpenBankProject/OBP-API/wiki",
				ID:    "webui_prop_002",
			},
			{
				Name:  "webui_api_manager_url",
				Value: "https://apimanager.openbankproject.com",
				ID:    "webui_prop_003",
			},
		},
	}

	if active == "true" {
		response.WebUIProps = append(response.WebUIProps, WebUIProp{
			Name:  "webui_default_bank_id",
			Value: "gh.29.uk",
			ID:    "default",
		})
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}
