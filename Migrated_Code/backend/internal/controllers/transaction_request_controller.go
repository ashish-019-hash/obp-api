package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/ashish-019-hash/obp-api-backend/internal/utils"
)

type TransactionRequestController struct {
	orchestrationService *services.OrchestrationService
}

func NewTransactionRequestController(orchestrationService *services.OrchestrationService) *TransactionRequestController {
	return &TransactionRequestController{
		orchestrationService: orchestrationService,
	}
}

type TransactionRequestResponse struct {
	ID          string                    `json:"id"`
	Type        string                    `json:"type"`
	From        TransactionRequestAccount `json:"from"`
	Details     TransactionRequestDetails `json:"details"`
	Body        TransactionRequestBody    `json:"body"`
	Status      string                    `json:"status"`
	StartDate   time.Time                 `json:"start_date"`
	EndDate     time.Time                 `json:"end_date"`
	Challenge   TransactionRequestChallenge `json:"challenge"`
	Charge      TransactionRequestCharge  `json:"charge"`
}

type TransactionRequestAccount struct {
	BankID    string `json:"bank_id"`
	AccountID string `json:"account_id"`
}

type TransactionRequestDetails struct {
	ToSandboxTan    TransactionRequestAccount `json:"to_sandbox_tan"`
	ToSepa          TransactionRequestAccount `json:"to_sepa"`
	ToCounterparty  TransactionRequestCounterparty `json:"to_counterparty"`
	ToTransferToPhone TransactionRequestPhone `json:"to_transfer_to_phone"`
	ToTransferToAtm TransactionRequestAtm   `json:"to_transfer_to_atm"`
	Value           TransactionRequestValue `json:"value"`
	Description     string                  `json:"description"`
}

type TransactionRequestCounterparty struct {
	CounterpartyID string `json:"counterparty_id"`
}

type TransactionRequestPhone struct {
	PhoneNumber string `json:"phone_number"`
}

type TransactionRequestAtm struct {
	AtmID string `json:"atm_id"`
}

type TransactionRequestValue struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type TransactionRequestBody struct {
	To          TransactionRequestAccount `json:"to"`
	Value       TransactionRequestValue   `json:"value"`
	Description string                    `json:"description"`
}

type TransactionRequestChallenge struct {
	ID               string `json:"id"`
	AllowedAttempts  int    `json:"allowed_attempts"`
	ChallengeType    string `json:"challenge_type"`
}

type TransactionRequestCharge struct {
	Summary string                  `json:"summary"`
	Value   TransactionRequestValue `json:"value"`
}

type TransactionRequestsResponse struct {
	TransactionRequests []TransactionRequestResponse `json:"transaction_requests"`
}

type UpdateTransactionRequestStatusRequest struct {
	Status string `json:"status" binding:"required"`
}

type TransactionRequestStatusResponse struct {
	TransactionRequestID string `json:"transaction_request_id"`
	Status               string `json:"status"`
}

func (c *TransactionRequestController) GetTransactionRequestById(ctx *gin.Context) {
	transactionRequestId := ctx.Param("transactionRequestId")

	response := TransactionRequestResponse{
		ID:   transactionRequestId,
		Type: "SANDBOX_TAN",
		From: TransactionRequestAccount{
			BankID:    "bank_001",
			AccountID: "account_001",
		},
		Details: TransactionRequestDetails{
			ToSandboxTan: TransactionRequestAccount{
				BankID:    "bank_002",
				AccountID: "account_002",
			},
			Value: TransactionRequestValue{
				Currency: "EUR",
				Amount:   "100.00",
			},
			Description: "Test transaction",
		},
		Body: TransactionRequestBody{
			To: TransactionRequestAccount{
				BankID:    "bank_002",
				AccountID: "account_002",
			},
			Value: TransactionRequestValue{
				Currency: "EUR",
				Amount:   "100.00",
			},
			Description: "Test transaction",
		},
		Status:    "INITIATED",
		StartDate: time.Now(),
		EndDate:   time.Now().Add(24 * time.Hour),
		Challenge: TransactionRequestChallenge{
			ID:              "challenge_001",
			AllowedAttempts: 3,
			ChallengeType:   "SANDBOX_TAN",
		},
		Charge: TransactionRequestCharge{
			Summary: "Transaction charge",
			Value: TransactionRequestValue{
				Currency: "EUR",
				Amount:   "0.50",
			},
		},
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *TransactionRequestController) GetTransactionRequests(ctx *gin.Context) {
	bankId := ctx.Param("bankId")
	accountId := ctx.Param("accountId")

	response := TransactionRequestsResponse{
		TransactionRequests: []TransactionRequestResponse{
			{
				ID:   "tr_001",
				Type: "SANDBOX_TAN",
				From: TransactionRequestAccount{
					BankID:    bankId,
					AccountID: accountId,
				},
				Details: TransactionRequestDetails{
					ToSandboxTan: TransactionRequestAccount{
						BankID:    "bank_002",
						AccountID: "account_002",
					},
					Value: TransactionRequestValue{
						Currency: "EUR",
						Amount:   "100.00",
					},
					Description: "Test transaction 1",
				},
				Status:    "INITIATED",
				StartDate: time.Now(),
				EndDate:   time.Now().Add(24 * time.Hour),
			},
			{
				ID:   "tr_002",
				Type: "SEPA",
				From: TransactionRequestAccount{
					BankID:    bankId,
					AccountID: accountId,
				},
				Details: TransactionRequestDetails{
					ToSepa: TransactionRequestAccount{
						BankID:    "bank_003",
						AccountID: "account_003",
					},
					Value: TransactionRequestValue{
						Currency: "EUR",
						Amount:   "250.00",
					},
					Description: "Test transaction 2",
				},
				Status:    "COMPLETED",
				StartDate: time.Now().Add(-2 * time.Hour),
				EndDate:   time.Now().Add(-1 * time.Hour),
			},
		},
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *TransactionRequestController) UpdateTransactionRequestStatus(ctx *gin.Context) {
	transactionRequestId := ctx.Param("transactionRequestId")

	var req UpdateTransactionRequestStatusRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	validStatuses := []string{"INITIATED", "COMPLETED", "CANCELLED", "REJECTED"}
	isValid := false
	for _, status := range validStatuses {
		if req.Status == status {
			isValid = true
			break
		}
	}

	if !isValid {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid transaction request status", "")
		return
	}

	response := TransactionRequestStatusResponse{
		TransactionRequestID: transactionRequestId,
		Status:               req.Status,
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}
