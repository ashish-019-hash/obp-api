package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/ashish-019-hash/obp-api-backend/internal/utils"
)

type BalanceController struct {
	orchestrationService *services.OrchestrationService
}

func NewBalanceController(orchestrationService *services.OrchestrationService) *BalanceController {
	return &BalanceController{
		orchestrationService: orchestrationService,
	}
}

type BalanceResponse struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
	Date     time.Time `json:"date"`
}

type BalancesResponse struct {
	Balances []BalanceResponse `json:"balances"`
}

type AccountBalanceResponse struct {
	AccountID string            `json:"account_id"`
	Balances  []BalanceResponse `json:"balances"`
}

type AccountBalancesResponse struct {
	Accounts []AccountBalanceResponse `json:"accounts"`
}

type CreateBalanceRequest struct {
	BalanceType   string `json:"balance_type" binding:"required"`
	BalanceAmount string `json:"balance_amount" binding:"required"`
}

func (c *BalanceController) GetBankAccountBalances(ctx *gin.Context) {

	response := BalancesResponse{
		Balances: []BalanceResponse{
			{
				ID:       "balance_001",
				Type:     "CURRENT",
				Currency: "EUR",
				Amount:   "1500.50",
				Date:     time.Now(),
			},
			{
				ID:       "balance_002",
				Type:     "AVAILABLE",
				Currency: "EUR",
				Amount:   "1450.50",
				Date:     time.Now(),
			},
		},
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *BalanceController) GetBankAccountsBalances(ctx *gin.Context) {

	response := AccountBalancesResponse{
		Accounts: []AccountBalanceResponse{
			{
				AccountID: "account_001",
				Balances: []BalanceResponse{
					{
						ID:       "balance_001",
						Type:     "CURRENT",
						Currency: "EUR",
						Amount:   "1500.50",
						Date:     time.Now(),
					},
				},
			},
			{
				AccountID: "account_002",
				Balances: []BalanceResponse{
					{
						ID:       "balance_003",
						Type:     "CURRENT",
						Currency: "USD",
						Amount:   "2000.00",
						Date:     time.Now(),
					},
				},
			},
		},
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *BalanceController) GetBankAccountsBalancesThroughView(ctx *gin.Context) {

	response := AccountBalancesResponse{
		Accounts: []AccountBalanceResponse{
			{
				AccountID: "account_001",
				Balances: []BalanceResponse{
					{
						ID:       "balance_001",
						Type:     "CURRENT",
						Currency: "EUR",
						Amount:   "1500.50",
						Date:     time.Now(),
					},
				},
			},
		},
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *BalanceController) CreateBankAccountBalance(ctx *gin.Context) {

	var req CreateBalanceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	response := BalanceResponse{
		ID:       "balance_" + strconv.FormatInt(time.Now().Unix(), 10),
		Type:     req.BalanceType,
		Currency: "EUR",
		Amount:   req.BalanceAmount,
		Date:     time.Now(),
	}

	utils.SendJSONResponse(ctx, http.StatusCreated, response)
}

func (c *BalanceController) GetBankAccountBalanceById(ctx *gin.Context) {
	balanceId := ctx.Param("BALANCE_ID")

	response := BalanceResponse{
		ID:       balanceId,
		Type:     "CURRENT",
		Currency: "EUR",
		Amount:   "1500.50",
		Date:     time.Now(),
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *BalanceController) GetAllBankAccountBalances(ctx *gin.Context) {

	response := BalancesResponse{
		Balances: []BalanceResponse{
			{
				ID:       "balance_001",
				Type:     "CURRENT",
				Currency: "EUR",
				Amount:   "1500.50",
				Date:     time.Now(),
			},
			{
				ID:       "balance_002",
				Type:     "AVAILABLE",
				Currency: "EUR",
				Amount:   "1450.50",
				Date:     time.Now(),
			},
		},
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *BalanceController) UpdateBankAccountBalance(ctx *gin.Context) {
	balanceId := ctx.Param("BALANCE_ID")

	var req CreateBalanceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	response := BalanceResponse{
		ID:       balanceId,
		Type:     req.BalanceType,
		Currency: "EUR",
		Amount:   req.BalanceAmount,
		Date:     time.Now(),
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *BalanceController) DeleteBankAccountBalance(ctx *gin.Context) {
	utils.SendJSONResponse(ctx, http.StatusNoContent, nil)
}
