package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/ashish-019-hash/obp-api-backend/internal/utils"
)

type CounterpartyLimitController struct {
	orchestrationService *services.OrchestrationService
}

func NewCounterpartyLimitController(orchestrationService *services.OrchestrationService) *CounterpartyLimitController {
	return &CounterpartyLimitController{
		orchestrationService: orchestrationService,
	}
}

type CounterpartyLimitCreateRequest struct {
	CurrencyCode                string `json:"currency_code" binding:"required"`
	MaxSingleAmount             string `json:"max_single_amount" binding:"required"`
	MaxMonthlyAmount            string `json:"max_monthly_amount" binding:"required"`
	MaxYearlyAmount             string `json:"max_yearly_amount" binding:"required"`
	MaxNumberOfMonthlyTransactions int `json:"max_number_of_monthly_transactions" binding:"required"`
	MaxNumberOfYearlyTransactions  int `json:"max_number_of_yearly_transactions" binding:"required"`
}

func (c *CounterpartyLimitController) CreateCounterpartyLimit(ctx *gin.Context) {
	bankId := ctx.Param("BANK_ID")
	accountId := ctx.Param("ACCOUNT_ID")
	viewId := ctx.Param("VIEW_ID")
	counterpartyId := ctx.Param("COUNTERPARTY_ID")

	var req CounterpartyLimitCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	response := CounterpartyLimitResponse{
		CounterpartyLimitID:              "limit_" + strconv.FormatInt(time.Now().Unix(), 10),
		BankID:                           bankId,
		AccountID:                        accountId,
		ViewID:                           viewId,
		CounterpartyID:                   counterpartyId,
		Currency:                         req.CurrencyCode,
		MaxSingleAmount:                  req.MaxSingleAmount,
		MaxMonthlyAmount:                 req.MaxMonthlyAmount,
		MaxNumberOfMonthlyTransactions:   req.MaxNumberOfMonthlyTransactions,
		MaxYearlyAmount:                  req.MaxYearlyAmount,
		MaxNumberOfYearlyTransactions:    req.MaxNumberOfYearlyTransactions,
		MaxTotalAmount:                   "100000.00",
		MaxNumberOfTransactions:          500,
	}

	utils.SendJSONResponse(ctx, http.StatusCreated, response)
}

func (c *CounterpartyLimitController) UpdateCounterpartyLimit(ctx *gin.Context) {
	counterpartyLimitId := ctx.Param("COUNTERPARTY_LIMIT_ID")

	var req CounterpartyLimitCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	response := CounterpartyLimitResponse{
		CounterpartyLimitID:              counterpartyLimitId,
		BankID:                           "bank_001",
		AccountID:                        "account_001",
		ViewID:                           "owner",
		CounterpartyID:                   "counterparty_001",
		Currency:                         req.CurrencyCode,
		MaxSingleAmount:                  req.MaxSingleAmount,
		MaxMonthlyAmount:                 req.MaxMonthlyAmount,
		MaxNumberOfMonthlyTransactions:   req.MaxNumberOfMonthlyTransactions,
		MaxYearlyAmount:                  req.MaxYearlyAmount,
		MaxNumberOfYearlyTransactions:    req.MaxNumberOfYearlyTransactions,
		MaxTotalAmount:                   "100000.00",
		MaxNumberOfTransactions:          500,
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *CounterpartyLimitController) GetCounterpartyLimit(ctx *gin.Context) {
	counterpartyLimitId := ctx.Param("COUNTERPARTY_LIMIT_ID")

	response := CounterpartyLimitResponse{
		CounterpartyLimitID:              counterpartyLimitId,
		BankID:                           "bank_001",
		AccountID:                        "account_001",
		ViewID:                           "owner",
		CounterpartyID:                   "counterparty_001",
		Currency:                         "EUR",
		MaxSingleAmount:                  "1000.00",
		MaxMonthlyAmount:                 "5000.00",
		MaxNumberOfMonthlyTransactions:   10,
		MaxYearlyAmount:                  "50000.00",
		MaxNumberOfYearlyTransactions:    100,
		MaxTotalAmount:                   "100000.00",
		MaxNumberOfTransactions:          500,
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *CounterpartyLimitController) GetCounterpartyLimitStatus(ctx *gin.Context) {
	counterpartyLimitId := ctx.Param("COUNTERPARTY_LIMIT_ID")

	response := CounterpartyLimitStatusResponse{
		CounterpartyLimitID:              counterpartyLimitId,
		BankID:                           "bank_001",
		AccountID:                        "account_001",
		ViewID:                           "owner",
		CounterpartyID:                   "counterparty_001",
		Currency:                         "EUR",
		MaxSingleAmount:                  "1000.00",
		MaxMonthlyAmount:                 "5000.00",
		MaxNumberOfMonthlyTransactions:   10,
		MaxYearlyAmount:                  "50000.00",
		MaxNumberOfYearlyTransactions:    100,
		MaxTotalAmount:                   "100000.00",
		MaxNumberOfTransactions:          500,
		Status: CounterpartyLimitStatus{
			CurrencyStatus:                        "EUR",
			MaxMonthlyAmountStatus:                "1500.00",
			MaxNumberOfMonthlyTransactionsStatus:  3,
			MaxYearlyAmountStatus:                 "15000.00",
			MaxNumberOfYearlyTransactionsStatus:   30,
			MaxTotalAmountStatus:                  "15000.00",
			MaxNumberOfTransactionsStatus:         30,
		},
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *CounterpartyLimitController) DeleteCounterpartyLimit(ctx *gin.Context) {
	utils.SendJSONResponse(ctx, http.StatusNoContent, nil)
}
