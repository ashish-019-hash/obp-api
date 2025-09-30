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
	ConsentID string    `json:"consent_id"`
	JWT       string    `json:"jwt"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
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
		CreatedAt: time.Now().Add(-24 * time.Hour),
		ExpiresAt: time.Now().Add(24 * time.Hour),
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
		CreatedAt: time.Now().Add(-24 * time.Hour),
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *ConsentController) GetConsentByConsentIdViaConsumer(ctx *gin.Context) {
	consentId := ctx.Param("consentId")

	response := ConsentResponse{
		ConsentID: consentId,
		JWT:       "eyJhbGciOiJIUzI1NiJ9...",
		Status:    "AUTHORISED",
		CreatedAt: time.Now().Add(-24 * time.Hour),
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *ConsentController) RevokeConsentAtBank(ctx *gin.Context) {

	utils.SendJSONResponse(ctx, http.StatusNoContent, nil)
}

func (c *ConsentController) SelfRevokeConsent(ctx *gin.Context) {
	utils.SendJSONResponse(ctx, http.StatusNoContent, nil)
}

func (c *ConsentController) CreateConsentImplicit(ctx *gin.Context) {
	_ = ctx.Param("scaMethod") // SCA method for consent creation

	var req CreateConsentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	defaultTTL := 3600 // Default from configuration service would be better
	ttl := defaultTTL
	if req.TimeToLive != nil {
		ttl = *req.TimeToLive
	}

	response := ConsentResponse{
		ConsentID: "consent_" + strconv.FormatInt(time.Now().Unix(), 10),
		Status:    "INITIATED",
		JWT:       "eyJhbGciOiJIUzI1NiJ9...",
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(time.Duration(ttl) * time.Second),
	}

	utils.SendJSONResponse(ctx, http.StatusCreated, response)
}

func (c *ConsentController) RevokeMyConsent(ctx *gin.Context) {
	utils.SendJSONResponse(ctx, http.StatusNoContent, nil)
}

func (c *ConsentController) CreateConsent(ctx *gin.Context) {
	var req CreateConsentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	response := ConsentResponse{
		ConsentID: "consent_" + strconv.FormatInt(time.Now().Unix(), 10),
		JWT:       "eyJhbGciOiJIUzI1NiJ9...",
		Status:    "INITIATED",
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(24 * time.Hour),
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

func (c *ConsentController) UpdateConsentStatusByConsent(ctx *gin.Context) {
	consentId := ctx.Param("consentId")
	
	var req gin.H
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	response := gin.H{
		"consent_id": consentId,
		"status":     req["status"],
		"updated_at": "2024-01-01T00:00:00Z",
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *ConsentController) UpdateConsentAccountAccessByConsentId(ctx *gin.Context) {
	consentId := ctx.Param("consentId")
	
	var req gin.H
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	response := gin.H{
		"consent_id":      consentId,
		"account_access":  req["account_access"],
		"updated_at":      "2024-01-01T00:00:00Z",
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *ConsentController) UpdateConsentUserIdByConsentId(ctx *gin.Context) {
	consentId := ctx.Param("consentId")
	
	var req gin.H
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	response := gin.H{
		"consent_id": consentId,
		"user_id":    req["user_id"],
		"updated_at": "2024-01-01T00:00:00Z",
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}
