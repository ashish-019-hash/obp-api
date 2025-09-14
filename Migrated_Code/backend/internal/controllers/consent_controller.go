package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/ashish-019-hash/obp-api-backend/internal/utils"
)

type ConsentController struct {
	orchestrationService *services.OrchestrationService
}

func NewConsentController(orchestrationService *services.OrchestrationService) *ConsentController {
	return &ConsentController{
		orchestrationService: orchestrationService,
	}
}

type UpdateConsentStatusRequest struct {
	Status string `json:"status" binding:"required"`
}

type ConsentResponse struct {
	ConsentID string `json:"consent_id"`
	JWT       string `json:"jwt"`
	Status    string `json:"status"`
}

type CreateConsentRequest struct {
	Everything  bool                    `json:"everything"`
	Views       []ConsentViewRequest    `json:"views"`
	Entitlements []ConsentEntitlementRequest `json:"entitlements"`
	ConsumerID  *string                 `json:"consumer_id,omitempty"`
	ValidFrom   *time.Time              `json:"valid_from,omitempty"`
	TimeToLive  *int                    `json:"time_to_live,omitempty"`
}

type ConsentViewRequest struct {
	BankID    string `json:"bank_id"`
	AccountID string `json:"account_id"`
	ViewID    string `json:"view_id"`
}

type ConsentEntitlementRequest struct {
	BankID   string `json:"bank_id"`
	RoleName string `json:"role_name"`
}

func (c *ConsentController) UpdateConsentStatus(ctx *gin.Context) {
	consentId := ctx.Param("consentId")

	var req UpdateConsentStatusRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	validStatuses := []string{"INITIATED", "ACCEPTED", "REJECTED", "REVOKED", "EXPIRED", "AUTHORISED"}
	isValid := false
	for _, status := range validStatuses {
		if req.Status == status {
			isValid = true
			break
		}
	}

	if !isValid {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid consent status", "")
		return
	}

	response := ConsentResponse{
		ConsentID: consentId,
		JWT:       "eyJhbGciOiJIUzI1NiJ9...",
		Status:    req.Status,
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *ConsentController) GetMyConsentsByBank(ctx *gin.Context) {

	response := gin.H{
		"consents": []ConsentResponse{},
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *ConsentController) GetMyConsents(ctx *gin.Context) {
	response := gin.H{
		"consents": []ConsentResponse{},
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *ConsentController) GetConsentsAtBank(ctx *gin.Context) {

	response := gin.H{
		"consents": []ConsentResponse{},
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *ConsentController) GetConsents(ctx *gin.Context) {
	response := gin.H{
		"consents": []ConsentResponse{},
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *ConsentController) GetConsentByConsentId(ctx *gin.Context) {
	consentId := ctx.Param("consentId")

	response := ConsentResponse{
		ConsentID: consentId,
		JWT:       "eyJhbGciOiJIUzI1NiJ9...",
		Status:    "AUTHORISED",
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *ConsentController) GetConsentByConsentIdViaConsumer(ctx *gin.Context) {
	consentId := ctx.Param("consentId")

	response := ConsentResponse{
		ConsentID: consentId,
		JWT:       "eyJhbGciOiJIUzI1NiJ9...",
		Status:    "AUTHORISED",
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *ConsentController) RevokeConsentAtBank(ctx *gin.Context) {

	utils.SendJSONResponse(ctx, http.StatusNoContent, nil)
}

func (c *ConsentController) SelfRevokeConsent(ctx *gin.Context) {
	utils.SendJSONResponse(ctx, http.StatusNoContent, nil)
}

func (c *ConsentController) RevokeMyConsent(ctx *gin.Context) {
	utils.SendJSONResponse(ctx, http.StatusNoContent, nil)
}

func (c *ConsentController) CreateConsent(ctx *gin.Context) {
	scaMethod := ctx.Param("scaMethod")

	validMethods := []string{"SMS", "EMAIL", "IMPLICIT"}
	isValid := false
	for _, method := range validMethods {
		if scaMethod == method {
			isValid = true
			break
		}
	}

	if !isValid {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid SCA method", "")
		return
	}

	var req CreateConsentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	response := ConsentResponse{
		ConsentID: "consent_" + strconv.FormatInt(time.Now().Unix(), 10),
		JWT:       "eyJhbGciOiJIUzI1NiJ9...",
		Status:    "INITIATED",
	}

	utils.SendJSONResponse(ctx, http.StatusCreated, response)
}

func (c *ConsentController) GetMtlsClientCertificateInfo(ctx *gin.Context) {
	response := gin.H{
		"certificate_info": gin.H{
			"subject":    "CN=example.com",
			"issuer":     "CN=Example CA",
			"valid_from": "2023-01-01T00:00:00Z",
			"valid_to":   "2024-01-01T00:00:00Z",
		},
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}
