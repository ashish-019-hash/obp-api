package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/ashish-019-hash/obp-api-backend/internal/utils"
)

type V510Controller struct {
	orchestrationService *services.OrchestrationService
}

func NewV510Controller(orchestrationService *services.OrchestrationService) *V510Controller {
	return &V510Controller{
		orchestrationService: orchestrationService,
	}
}

type UserAttributeRequest struct {
	Name     string `json:"name" binding:"required"`
	Type     string `json:"type" binding:"required"`
	Value    string `json:"value" binding:"required"`
	IsActive bool   `json:"is_active"`
}


type RegulatedEntityRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	EntityType  string `json:"entity_type" binding:"required"`
}

type RegulatedEntityResponse struct {
	RegulatedEntityID string `json:"regulated_entity_id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	EntityType        string `json:"entity_type"`
}

func (c *V510Controller) GetRoot(ctx *gin.Context) {
	response := gin.H{
		"version":        "v5.1.0",
		"version_status": "BLEEDING_EDGE",
		"git_commit":     "unknown",
		"hosted_by": gin.H{
			"organisation": "Open Bank Project",
			"email":        "contact@openbankproject.com",
			"phone":        "+49 (0)30 8145 3994",
		},
		"hosted_at": gin.H{
			"organisation":         "Example",
			"organisation_website": "https://www.example.com",
		},
		"energy_source": gin.H{
			"organisation":         "Example",
			"organisation_website": "https://www.example.com",
		},
	}
	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *V510Controller) GetSuggestedSessionTimeout(ctx *gin.Context) {
	response := gin.H{
		"suggested_session_timeout": "300",
	}
	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *V510Controller) GetWellKnown(ctx *gin.Context) {
	response := gin.H{
		"well_known_uris": []gin.H{
			{
				"name": "keycloak",
				"uri":  "https://keycloak.example.com/.well-known/openid_configuration",
			},
		},
	}
	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *V510Controller) WaitingForGodot(ctx *gin.Context) {
	sleepParam := ctx.DefaultQuery("sleep", "0")
	sleepMs := 0
	if sleep, err := strconv.Atoi(sleepParam); err == nil {
		sleepMs = sleep
	}

	if sleepMs > 0 {
		time.Sleep(time.Duration(sleepMs) * time.Millisecond)
	}

	response := gin.H{
		"sleep_in_milliseconds": sleepMs,
	}
	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *V510Controller) GetRegulatedEntities(ctx *gin.Context) {
	response := gin.H{
		"regulated_entities": []RegulatedEntityResponse{},
	}
	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *V510Controller) GetRegulatedEntityById(ctx *gin.Context) {
	regulatedEntityId := ctx.Param("regulatedEntityId")

	response := RegulatedEntityResponse{
		RegulatedEntityID: regulatedEntityId,
		Name:              "Example Entity",
		Description:       "Example regulated entity",
		EntityType:        "BANK",
	}
	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *V510Controller) CreateRegulatedEntity(ctx *gin.Context) {
	var req RegulatedEntityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}
	
	if req.EntityType == "" {
		req.EntityType = "BANK"
	}

	response := RegulatedEntityResponse{
		RegulatedEntityID: "entity_" + strconv.FormatInt(time.Now().Unix(), 10),
		Name:              req.Name,
		Description:       req.Description,
		EntityType:        req.EntityType,
	}

	utils.SendJSONResponse(ctx, http.StatusCreated, response)
}

func (c *V510Controller) DeleteRegulatedEntity(ctx *gin.Context) {
	utils.SendJSONResponse(ctx, http.StatusNoContent, nil)
}

func (c *V510Controller) CreateNonPersonalUserAttribute(ctx *gin.Context) {
	var req UserAttributeRequest
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

	response := gin.H{
		"user_attribute_id": "attr_" + strconv.FormatInt(time.Now().Unix(), 10),
		"name":              req.Name,
		"type":              req.Type,
		"value":             req.Value,
		"is_active":         req.IsActive,
	}

	utils.SendJSONResponse(ctx, http.StatusCreated, response)
}

func (c *V510Controller) DeleteNonPersonalUserAttribute(ctx *gin.Context) {
	utils.SendJSONResponse(ctx, http.StatusNoContent, nil)
}

func (c *V510Controller) GetNonPersonalUserAttributes(ctx *gin.Context) {
	response := gin.H{
		"user_attributes": []UserAttributeResponse{},
	}
	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *V510Controller) GetAccountsHeldByUserAtBank(ctx *gin.Context) {
	response := gin.H{
		"accounts": []gin.H{},
	}
	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *V510Controller) GetAccountsHeldByUser(ctx *gin.Context) {
	response := gin.H{
		"accounts": []gin.H{},
	}
	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *V510Controller) GetEntitlementsAndPermissions(ctx *gin.Context) {
	response := gin.H{
		"entitlements": []gin.H{},
		"permissions":  []gin.H{},
	}
	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *V510Controller) CustomViewNamesCheck(ctx *gin.Context) {
	response := gin.H{
		"is_everything_ok": true,
		"incorrect_views":  []gin.H{},
	}
	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *V510Controller) SystemViewNamesCheck(ctx *gin.Context) {
	response := gin.H{
		"is_everything_ok": true,
		"incorrect_views":  []gin.H{},
	}
	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *V510Controller) AccountAccessUniqueIndexCheck(ctx *gin.Context) {
	response := gin.H{
		"is_everything_ok":   true,
		"duplicated_records": []gin.H{},
	}
	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *V510Controller) AccountCurrencyCheck(ctx *gin.Context) {
	response := gin.H{
		"is_everything_ok": true,
		"currencies":       []string{},
	}
	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *V510Controller) OrphanedAccountCheck(ctx *gin.Context) {
	response := gin.H{
		"is_everything_ok":    true,
		"orphaned_accounts":   []string{},
		"orphaned_account_ids": []string{},
	}
	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *V510Controller) GetCurrenciesAtBank(ctx *gin.Context) {
	response := gin.H{
		"currencies": []gin.H{
			{"code": "EUR", "name": "Euro"},
			{"code": "USD", "name": "US Dollar"},
			{"code": "GBP", "name": "British Pound"},
		},
	}
	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *V510Controller) GetAllApiCollections(ctx *gin.Context) {
	response := gin.H{
		"api_collections": []gin.H{},
	}
	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *V510Controller) GetApiTags(ctx *gin.Context) {
	response := gin.H{
		"tags": []string{
			"Account",
			"Transaction",
			"Customer",
			"Bank",
			"ATM",
			"Branch",
			"Card",
			"Consent",
			"Consumer",
			"Counterparty",
			"Entitlement",
			"Metric",
			"Product",
			"Role",
			"User",
			"View",
			"Webhook",
		},
	}
	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *V510Controller) GetCoreAccountByIdThroughView(ctx *gin.Context) {
	bankId := ctx.Param("bankId")
	accountId := ctx.Param("accountId")
	viewId := ctx.Param("viewId")

	response := gin.H{
		"id":       accountId,
		"bank_id":  bankId,
		"label":    "Main Account",
		"number":   "123456789",
		"owners": []gin.H{
			{
				"id":           "user_001",
				"provider":     "http://127.0.0.1:8080",
				"display_name": "John Doe",
			},
		},
		"type":     "CURRENT",
		"balance": gin.H{
			"currency": "EUR",
			"amount":   "1000.00",
		},
		"account_routings": []gin.H{
			{
				"scheme":  "AccountNumber",
				"address": "123456789",
			},
		},
		"account_rules": []gin.H{},
		"tags":          []gin.H{},
		"views_available": []gin.H{
			{
				"id":          viewId,
				"short_name":  "Owner",
				"description": "Owner view of account",
				"is_public":   false,
			},
		},
	}
	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *V510Controller) UpdateMyApiCollection(ctx *gin.Context) {
	apiCollectionId := ctx.Param("apiCollectionId")

	var req gin.H
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	response := gin.H{
		"api_collection_id":   apiCollectionId,
		"api_collection_name": req["api_collection_name"],
		"is_sharable":         req["is_sharable"],
		"description":         req["description"],
		"api_collection_endpoints": []gin.H{},
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *V510Controller) GetMtlsClientCertificateInfo(ctx *gin.Context) {
	response := gin.H{
		"certificate_info": gin.H{
			"subject": gin.H{
				"common_name":    "example.com",
				"organization":   "Example Corp",
				"country":        "US",
				"email_address":  "admin@example.com",
			},
			"issuer": gin.H{
				"common_name":  "Example CA",
				"organization": "Example Certificate Authority",
				"country":      "US",
			},
			"serial_number": "123456789",
			"not_before":    time.Now().Add(-365 * 24 * time.Hour).Format(time.RFC3339),
			"not_after":     time.Now().Add(365 * 24 * time.Hour).Format(time.RFC3339),
			"fingerprint":   "AA:BB:CC:DD:EE:FF:00:11:22:33:44:55:66:77:88:99:AA:BB:CC:DD",
		},
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}
