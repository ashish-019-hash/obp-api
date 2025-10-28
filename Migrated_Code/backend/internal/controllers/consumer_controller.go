package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/ashish-019-hash/obp-api-backend/internal/utils"
)

type ConsumerController struct {
	orchestrationService *services.OrchestrationService
}

func NewConsumerController(orchestrationService *services.OrchestrationService) *ConsumerController {
	return &ConsumerController{
		orchestrationService: orchestrationService,
	}
}

type ConsumerJwtRequest struct {
	JWT string `json:"jwt" binding:"required"`
}

type CreateConsumerRequest struct {
	AppName           string `json:"app_name" binding:"required"`
	AppType           string `json:"app_type" binding:"required"`
	Description       string `json:"description" binding:"required"`
	DeveloperEmail    string `json:"developer_email" binding:"required"`
	Company           string `json:"company"`
	RedirectURL       string `json:"redirect_url"`
	ClientCertificate string `json:"client_certificate"`
	LogoURL           string `json:"logo_url"`
	Enabled           bool   `json:"enabled"`
}

type ConsumerRedirectURLRequest struct {
	RedirectURL string `json:"redirect_url" binding:"required"`
}

type ConsumerLogoURLRequest struct {
	LogoURL string `json:"logo_url" binding:"required"`
}

type ConsumerCertificateRequest struct {
	Certificate string `json:"certificate" binding:"required"`
}

type ConsumerNameRequest struct {
	AppName string `json:"app_name" binding:"required"`
}

type ConsumerResponse struct {
	ConsumerID        string `json:"consumer_id"`
	Key               string `json:"key"`
	Secret            string `json:"secret"`
	AppName           string `json:"app_name"`
	AppType           string `json:"app_type"`
	Description       string `json:"description"`
	DeveloperEmail    string `json:"developer_email"`
	Company           string `json:"company"`
	RedirectURL       string `json:"redirect_url"`
	ClientCertificate string `json:"client_certificate"`
	LogoURL           string `json:"logo_url"`
	Enabled           bool   `json:"enabled"`
	CreatedByUserId   string `json:"created_by_user_id"`
}

type ConsumersResponse struct {
	Consumers []ConsumerResponse `json:"consumers"`
}

func (c *ConsumerController) CreateConsumerDynamicRegistration(ctx *gin.Context) {
	var req ConsumerJwtRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	response := ConsumerResponse{
		ConsumerID:      "consumer_" + strconv.FormatInt(time.Now().Unix(), 10),
		Key:             "key_" + strconv.FormatInt(time.Now().Unix(), 10),
		Secret:          "secret_" + strconv.FormatInt(time.Now().Unix(), 10),
		AppName:         "Dynamic App",
		AppType:         "Confidential",
		Description:     "Dynamically registered application",
		DeveloperEmail:  "dev@example.com",
		Company:         "Example Corp",
		RedirectURL:     "https://example.com/callback",
		Enabled:         true,
		CreatedByUserId: "system",
	}

	utils.SendJSONResponse(ctx, http.StatusCreated, response)
}

func (c *ConsumerController) CreateConsumer(ctx *gin.Context) {
	var req CreateConsumerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	response := ConsumerResponse{
		ConsumerID:        "consumer_" + strconv.FormatInt(time.Now().Unix(), 10),
		Key:               "key_" + strconv.FormatInt(time.Now().Unix(), 10),
		Secret:            "secret_" + strconv.FormatInt(time.Now().Unix(), 10),
		AppName:           req.AppName,
		AppType:           req.AppType,
		Description:       req.Description,
		DeveloperEmail:    req.DeveloperEmail,
		Company:           req.Company,
		RedirectURL:       req.RedirectURL,
		ClientCertificate: req.ClientCertificate,
		LogoURL:           req.LogoURL,
		Enabled:           req.Enabled,
		CreatedByUserId:   "user_123",
	}

	utils.SendJSONResponse(ctx, http.StatusCreated, response)
}

func (c *ConsumerController) CreateMyConsumer(ctx *gin.Context) {
	var req CreateConsumerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	response := ConsumerResponse{
		ConsumerID:        "consumer_" + strconv.FormatInt(time.Now().Unix(), 10),
		Key:               "key_" + strconv.FormatInt(time.Now().Unix(), 10),
		Secret:            "secret_" + strconv.FormatInt(time.Now().Unix(), 10),
		AppName:           req.AppName,
		AppType:           req.AppType,
		Description:       req.Description,
		DeveloperEmail:    req.DeveloperEmail,
		Company:           req.Company,
		RedirectURL:       req.RedirectURL,
		ClientCertificate: req.ClientCertificate,
		LogoURL:           req.LogoURL,
		Enabled:           req.Enabled,
		CreatedByUserId:   "current_user",
	}

	utils.SendJSONResponse(ctx, http.StatusCreated, response)
}

func (c *ConsumerController) UpdateConsumerRedirectURL(ctx *gin.Context) {
	consumerId := ctx.Param("consumerId")

	var req ConsumerRedirectURLRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	response := ConsumerResponse{
		ConsumerID:  consumerId,
		RedirectURL: req.RedirectURL,
		AppName:     "Updated App",
		Enabled:     true,
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *ConsumerController) UpdateConsumerLogoURL(ctx *gin.Context) {
	consumerId := ctx.Param("consumerId")

	var req ConsumerLogoURLRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	response := ConsumerResponse{
		ConsumerID: consumerId,
		LogoURL:    req.LogoURL,
		AppName:    "Updated App",
		Enabled:    true,
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *ConsumerController) UpdateConsumerCertificate(ctx *gin.Context) {
	consumerId := ctx.Param("consumerId")

	var req ConsumerCertificateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	response := ConsumerResponse{
		ConsumerID:        consumerId,
		ClientCertificate: req.Certificate,
		AppName:           "Updated App",
		Enabled:           true,
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *ConsumerController) UpdateConsumerName(ctx *gin.Context) {
	consumerId := ctx.Param("consumerId")

	var req ConsumerNameRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	response := ConsumerResponse{
		ConsumerID: consumerId,
		AppName:    req.AppName,
		Enabled:    true,
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *ConsumerController) GetConsumer(ctx *gin.Context) {
	consumerId := ctx.Param("consumerId")

	response := ConsumerResponse{
		ConsumerID:      consumerId,
		Key:             "key_123",
		AppName:         "Example App",
		AppType:         "Confidential",
		Description:     "Example application",
		DeveloperEmail:  "dev@example.com",
		Company:         "Example Corp",
		RedirectURL:     "https://example.com/callback",
		Enabled:         true,
		CreatedByUserId: "user_123",
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *ConsumerController) GetConsumers(ctx *gin.Context) {
	response := ConsumersResponse{
		Consumers: []ConsumerResponse{
			{
				ConsumerID:      "consumer_001",
				Key:             "key_001",
				AppName:         "Example App 1",
				AppType:         "Confidential",
				Description:     "First example application",
				DeveloperEmail:  "dev1@example.com",
				Company:         "Example Corp",
				RedirectURL:     "https://example.com/callback1",
				Enabled:         true,
				CreatedByUserId: "user_123",
			},
			{
				ConsumerID:      "consumer_002",
				Key:             "key_002",
				AppName:         "Example App 2",
				AppType:         "Public",
				Description:     "Second example application",
				DeveloperEmail:  "dev2@example.com",
				Company:         "Example Corp",
				RedirectURL:     "https://example.com/callback2",
				Enabled:         false,
				CreatedByUserId: "user_456",
			},
		},
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}
