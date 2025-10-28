package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ashish-019-hash/obp-api-backend/internal/models"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/ashish-019-hash/obp-api-backend/internal/utils"
)

type AgentController struct {
	orchestrationService *services.OrchestrationService
}

func NewAgentController(orchestrationService *services.OrchestrationService) *AgentController {
	return &AgentController{
		orchestrationService: orchestrationService,
	}
}

type CreateAgentRequest struct {
	LegalName         string `json:"legal_name" binding:"required"`
	MobilePhoneNumber string `json:"mobile_phone_number" binding:"required"`
	AgentNumber       string `json:"agent_number" binding:"required"`
	Currency          string `json:"currency" binding:"required"`
}

type UpdateAgentStatusRequest struct {
	IsPendingAgent   bool `json:"is_pending_agent"`
	IsConfirmedAgent bool `json:"is_confirmed_agent"`
}

type AgentResponse struct {
	AgentID           string              `json:"agent_id"`
	LegalName         string              `json:"legal_name"`
	MobilePhoneNumber string              `json:"mobile_phone_number"`
	AgentNumber       string              `json:"agent_number"`
	IsPendingAgent    bool                `json:"is_pending_agent"`
	IsConfirmedAgent  bool                `json:"is_confirmed_agent"`
	BankAccount       *models.BankAccount `json:"bank_account,omitempty"`
}

func (c *AgentController) CreateAgent(ctx *gin.Context) {
	var req CreateAgentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	response := AgentResponse{
		AgentID:           "agent_" + strconv.FormatInt(time.Now().Unix(), 10),
		LegalName:         req.LegalName,
		MobilePhoneNumber: req.MobilePhoneNumber,
		AgentNumber:       req.AgentNumber,
		IsPendingAgent:    true,
		IsConfirmedAgent:  false,
	}

	utils.SendJSONResponse(ctx, http.StatusCreated, response)
}

func (c *AgentController) UpdateAgentStatus(ctx *gin.Context) {
	agentId := ctx.Param("agentId")

	var req UpdateAgentStatusRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	response := AgentResponse{
		AgentID:          agentId,
		IsPendingAgent:   req.IsPendingAgent,
		IsConfirmedAgent: req.IsConfirmedAgent,
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *AgentController) GetAgent(ctx *gin.Context) {
	agentId := ctx.Param("agentId")

	response := AgentResponse{
		AgentID:           agentId,
		LegalName:         "Example Agent",
		MobilePhoneNumber: "+1234567890",
		AgentNumber:       "AGT001",
		IsPendingAgent:    false,
		IsConfirmedAgent:  true,
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *AgentController) GetAgents(ctx *gin.Context) {
	limit := ctx.DefaultQuery("limit", "10")
	offset := ctx.DefaultQuery("offset", "0")

	response := gin.H{
		"agents": []AgentResponse{},
		"pagination": gin.H{
			"limit":  limit,
			"offset": offset,
			"total":  0,
		},
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}
