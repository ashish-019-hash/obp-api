package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/ashish-019-hash/obp-api-backend/internal/utils"
)

type VRPController struct {
	orchestrationService *services.OrchestrationService
}

func NewVRPController(orchestrationService *services.OrchestrationService) *VRPController {
	return &VRPController{
		orchestrationService: orchestrationService,
	}
}

type VRPConsentRequest struct {
	FromAccount    VRPAccount    `json:"from_account" binding:"required"`
	ToAccount      VRPAccount    `json:"to_account" binding:"required"`
	MaxSingleAmount string       `json:"max_single_amount" binding:"required"`
	MaxMonthlyAmount string      `json:"max_monthly_amount" binding:"required"`
	MaxYearlyAmount string       `json:"max_yearly_amount" binding:"required"`
	MaxNumberOfMonthlyTransactions int `json:"max_number_of_monthly_transactions" binding:"required"`
	MaxNumberOfYearlyTransactions int  `json:"max_number_of_yearly_transactions" binding:"required"`
	Currency       string        `json:"currency" binding:"required"`
	ValidFrom      *time.Time    `json:"valid_from,omitempty"`
	TimeToLive     *int          `json:"time_to_live,omitempty"`
}

type VRPAccount struct {
	BankRouting    VRPBankRouting    `json:"bank_routing" binding:"required"`
	AccountRouting VRPAccountRouting `json:"account_routing" binding:"required"`
}

type VRPBankRouting struct {
	Scheme  string `json:"scheme" binding:"required"`
	Address string `json:"address" binding:"required"`
}

type VRPAccountRouting struct {
	Scheme  string `json:"scheme" binding:"required"`
	Address string `json:"address" binding:"required"`
}

type VRPConsentResponse struct {
	ConsentRequestID string    `json:"consent_request_id"`
	Status           string    `json:"status"`
	FromAccount      VRPAccount `json:"from_account"`
	ToAccount        VRPAccount `json:"to_account"`
	MaxSingleAmount  string    `json:"max_single_amount"`
	MaxMonthlyAmount string    `json:"max_monthly_amount"`
	MaxYearlyAmount  string    `json:"max_yearly_amount"`
	MaxNumberOfMonthlyTransactions int `json:"max_number_of_monthly_transactions"`
	MaxNumberOfYearlyTransactions int  `json:"max_number_of_yearly_transactions"`
	Currency         string    `json:"currency"`
	ValidFrom        time.Time `json:"valid_from"`
	ValidUntil       time.Time `json:"valid_until"`
	CreatedAt        time.Time `json:"created_at"`
}

func (c *VRPController) CreateVRPConsentRequest(ctx *gin.Context) {
	var req VRPConsentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	validFrom := time.Now()
	if req.ValidFrom != nil {
		validFrom = *req.ValidFrom
	}

	defaultTTL := 3600 // Default from configuration service would be better
	timeToLive := defaultTTL
	if req.TimeToLive != nil {
		timeToLive = *req.TimeToLive
	}

	validUntil := validFrom.Add(time.Duration(timeToLive) * time.Second)

	response := VRPConsentResponse{
		ConsentRequestID: "vrp_consent_" + strconv.FormatInt(time.Now().Unix(), 10),
		Status:           "INITIATED",
		FromAccount:      req.FromAccount,
		ToAccount:        req.ToAccount,
		MaxSingleAmount:  req.MaxSingleAmount,
		MaxMonthlyAmount: req.MaxMonthlyAmount,
		MaxYearlyAmount:  req.MaxYearlyAmount,
		MaxNumberOfMonthlyTransactions: req.MaxNumberOfMonthlyTransactions,
		MaxNumberOfYearlyTransactions:  req.MaxNumberOfYearlyTransactions,
		Currency:         req.Currency,
		ValidFrom:        validFrom,
		ValidUntil:       validUntil,
		CreatedAt:        time.Now(),
	}

	utils.SendJSONResponse(ctx, http.StatusCreated, response)
}
