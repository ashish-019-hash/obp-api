package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/ashish-019-hash/obp-api-backend/internal/utils"
)

type CounterpartyController struct {
	orchestrationService *services.OrchestrationService
}

func NewCounterpartyController(orchestrationService *services.OrchestrationService) *CounterpartyController {
	return &CounterpartyController{
		orchestrationService: orchestrationService,
	}
}

type CounterpartyLimitRequest struct {
	Currency                         string  `json:"currency" binding:"required"`
	MaxSingleAmount                  string  `json:"max_single_amount" binding:"required"`
	MaxMonthlyAmount                 string  `json:"max_monthly_amount" binding:"required"`
	MaxNumberOfMonthlyTransactions   int     `json:"max_number_of_monthly_transactions" binding:"required"`
	MaxYearlyAmount                  string  `json:"max_yearly_amount" binding:"required"`
	MaxNumberOfYearlyTransactions    int     `json:"max_number_of_yearly_transactions" binding:"required"`
	MaxTotalAmount                   string  `json:"max_total_amount" binding:"required"`
	MaxNumberOfTransactions          int     `json:"max_number_of_transactions" binding:"required"`
}

type CounterpartyLimitResponse struct {
	CounterpartyLimitID              string  `json:"counterparty_limit_id"`
	BankID                           string  `json:"bank_id"`
	AccountID                        string  `json:"account_id"`
	ViewID                           string  `json:"view_id"`
	CounterpartyID                   string  `json:"counterparty_id"`
	Currency                         string  `json:"currency"`
	MaxSingleAmount                  string  `json:"max_single_amount"`
	MaxMonthlyAmount                 string  `json:"max_monthly_amount"`
	MaxNumberOfMonthlyTransactions   int     `json:"max_number_of_monthly_transactions"`
	MaxYearlyAmount                  string  `json:"max_yearly_amount"`
	MaxNumberOfYearlyTransactions    int     `json:"max_number_of_yearly_transactions"`
	MaxTotalAmount                   string  `json:"max_total_amount"`
	MaxNumberOfTransactions          int     `json:"max_number_of_transactions"`
}

type CounterpartyLimitStatusResponse struct {
	CounterpartyLimitID              string                    `json:"counterparty_limit_id"`
	BankID                           string                    `json:"bank_id"`
	AccountID                        string                    `json:"account_id"`
	ViewID                           string                    `json:"view_id"`
	CounterpartyID                   string                    `json:"counterparty_id"`
	Currency                         string                    `json:"currency"`
	MaxSingleAmount                  string                    `json:"max_single_amount"`
	MaxMonthlyAmount                 string                    `json:"max_monthly_amount"`
	MaxNumberOfMonthlyTransactions   int                       `json:"max_number_of_monthly_transactions"`
	MaxYearlyAmount                  string                    `json:"max_yearly_amount"`
	MaxNumberOfYearlyTransactions    int                       `json:"max_number_of_yearly_transactions"`
	MaxTotalAmount                   string                    `json:"max_total_amount"`
	MaxNumberOfTransactions          int                       `json:"max_number_of_transactions"`
	Status                           CounterpartyLimitStatus   `json:"status"`
}

type CounterpartyLimitStatus struct {
	CurrencyStatus                        string  `json:"currency_status"`
	MaxMonthlyAmountStatus                string  `json:"max_monthly_amount_status"`
	MaxNumberOfMonthlyTransactionsStatus  int     `json:"max_number_of_monthly_transactions_status"`
	MaxYearlyAmountStatus                 string  `json:"max_yearly_amount_status"`
	MaxNumberOfYearlyTransactionsStatus   int     `json:"max_number_of_yearly_transactions_status"`
	MaxTotalAmountStatus                  string  `json:"max_total_amount_status"`
	MaxNumberOfTransactionsStatus         int     `json:"max_number_of_transactions_status"`
}

func (c *CounterpartyController) CreateCounterpartyLimit(ctx *gin.Context) {
	bankId := ctx.Param("BANK_ID")
	accountId := ctx.Param("ACCOUNT_ID")
	viewId := ctx.Param("VIEW_ID")
	counterpartyId := ctx.Param("COUNTERPARTY_ID")

	var req CounterpartyLimitRequest
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
		Currency:                         req.Currency,
		MaxSingleAmount:                  req.MaxSingleAmount,
		MaxMonthlyAmount:                 req.MaxMonthlyAmount,
		MaxNumberOfMonthlyTransactions:   req.MaxNumberOfMonthlyTransactions,
		MaxYearlyAmount:                  req.MaxYearlyAmount,
		MaxNumberOfYearlyTransactions:    req.MaxNumberOfYearlyTransactions,
		MaxTotalAmount:                   req.MaxTotalAmount,
		MaxNumberOfTransactions:          req.MaxNumberOfTransactions,
	}

	utils.SendJSONResponse(ctx, http.StatusCreated, response)
}

func (c *CounterpartyController) UpdateCounterpartyLimit(ctx *gin.Context) {
	bankId := ctx.Param("BANK_ID")
	accountId := ctx.Param("ACCOUNT_ID")
	viewId := ctx.Param("VIEW_ID")
	counterpartyId := ctx.Param("COUNTERPARTY_ID")

	var req CounterpartyLimitRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	response := CounterpartyLimitResponse{
		CounterpartyLimitID:              "limit_existing",
		BankID:                           bankId,
		AccountID:                        accountId,
		ViewID:                           viewId,
		CounterpartyID:                   counterpartyId,
		Currency:                         req.Currency,
		MaxSingleAmount:                  req.MaxSingleAmount,
		MaxMonthlyAmount:                 req.MaxMonthlyAmount,
		MaxNumberOfMonthlyTransactions:   req.MaxNumberOfMonthlyTransactions,
		MaxYearlyAmount:                  req.MaxYearlyAmount,
		MaxNumberOfYearlyTransactions:    req.MaxNumberOfYearlyTransactions,
		MaxTotalAmount:                   req.MaxTotalAmount,
		MaxNumberOfTransactions:          req.MaxNumberOfTransactions,
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *CounterpartyController) GetCounterpartyLimit(ctx *gin.Context) {
	bankId := ctx.Param("BANK_ID")
	accountId := ctx.Param("ACCOUNT_ID")
	viewId := ctx.Param("VIEW_ID")
	counterpartyId := ctx.Param("COUNTERPARTY_ID")

	response := CounterpartyLimitResponse{
		CounterpartyLimitID:              "limit_existing",
		BankID:                           bankId,
		AccountID:                        accountId,
		ViewID:                           viewId,
		CounterpartyID:                   counterpartyId,
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

func (c *CounterpartyController) GetCounterpartyLimitStatus(ctx *gin.Context) {
	bankId := ctx.Param("BANK_ID")
	accountId := ctx.Param("ACCOUNT_ID")
	viewId := ctx.Param("VIEW_ID")
	counterpartyId := ctx.Param("COUNTERPARTY_ID")

	response := CounterpartyLimitStatusResponse{
		CounterpartyLimitID:              "limit_existing",
		BankID:                           bankId,
		AccountID:                        accountId,
		ViewID:                           viewId,
		CounterpartyID:                   counterpartyId,
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

func (c *CounterpartyController) DeleteCounterpartyLimit(ctx *gin.Context) {
	utils.SendJSONResponse(ctx, http.StatusNoContent, nil)
}

type CounterpartyResponse struct {
	CounterpartyID string    `json:"counterparty_id"`
	Name           string    `json:"name"`
	BankID         string    `json:"bank_id"`
	AccountID      string    `json:"account_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type CreateCounterpartyRequest struct {
	Name      string `json:"name" binding:"required"`
	BankID    string `json:"bank_id" binding:"required"`
	AccountID string `json:"account_id" binding:"required"`
}

type UpdateCounterpartyRequest struct {
	Name string `json:"name" binding:"required"`
}

func (c *CounterpartyController) GetCounterparties(ctx *gin.Context) {
	response := gin.H{
		"counterparties": []CounterpartyResponse{
			{
				CounterpartyID: "counterparty_001",
				Name:           "Counterparty One",
				BankID:         "bank_001",
				AccountID:      "account_001",
				CreatedAt:      time.Now().Add(-48 * time.Hour),
				UpdatedAt:      time.Now().Add(-24 * time.Hour),
			},
		},
	}
	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *CounterpartyController) CreateCounterparty(ctx *gin.Context) {
	var req CreateCounterpartyRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	response := CounterpartyResponse{
		CounterpartyID: "counterparty_" + strconv.FormatInt(time.Now().Unix(), 10),
		Name:           req.Name,
		BankID:         req.BankID,
		AccountID:      req.AccountID,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	utils.SendJSONResponse(ctx, http.StatusCreated, response)
}

func (c *CounterpartyController) GetCounterpartyById(ctx *gin.Context) {
	counterpartyId := ctx.Param("COUNTERPARTY_ID")

	response := CounterpartyResponse{
		CounterpartyID: counterpartyId,
		Name:           "Example Counterparty",
		BankID:         "bank_001",
		AccountID:      "account_001",
		CreatedAt:      time.Now().Add(-24 * time.Hour),
		UpdatedAt:      time.Now(),
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *CounterpartyController) UpdateCounterparty(ctx *gin.Context) {
	counterpartyId := ctx.Param("COUNTERPARTY_ID")
	
	var req UpdateCounterpartyRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	response := CounterpartyResponse{
		CounterpartyID: counterpartyId,
		Name:           req.Name,
		BankID:         "bank_001",
		AccountID:      "account_001",
		CreatedAt:      time.Now().Add(-24 * time.Hour),
		UpdatedAt:      time.Now(),
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *CounterpartyController) DeleteCounterparty(ctx *gin.Context) {
	utils.SendJSONResponse(ctx, http.StatusNoContent, nil)
}
