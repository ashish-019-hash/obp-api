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

type UserAttributeResponse struct {
	UserAttributeID string `json:"user_attribute_id"`
	Name            string `json:"name"`
	Type            string `json:"type"`
	Value           string `json:"value"`
	IsActive        bool   `json:"is_active"`
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

	response := UserAttributeResponse{
		UserAttributeID: "attr_" + strconv.FormatInt(time.Now().Unix(), 10),
		Name:            req.Name,
		Type:            req.Type,
		Value:           req.Value,
		IsActive:        req.IsActive,
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
