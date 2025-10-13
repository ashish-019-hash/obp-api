package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/ashish-019-hash/obp-api-backend/internal/utils"
)

type AccountController struct {
	orchestrationService *services.OrchestrationService
}

func NewAccountController(orchestrationService *services.OrchestrationService) *AccountController {
	return &AccountController{
		orchestrationService: orchestrationService,
	}
}

type CoreAccountResponse struct {
	ID              string                 `json:"id"`
	BankID          string                 `json:"bank_id"`
	Label           string                 `json:"label"`
	Number          string                 `json:"number"`
	Type            string                 `json:"type"`
	Balance         AccountBalanceInfo     `json:"balance"`
	IBAN            string                 `json:"IBAN"`
	SwiftBIC        string                 `json:"swift_bic"`
	BankRouting     AccountBankRouting     `json:"bank_routing"`
	AccountRouting  AccountRoutingInfo     `json:"account_routing"`
	AccountRules    []AccountRule          `json:"account_rules"`
	AccountHolders  []AccountHolder        `json:"account_holders"`
	Views           []AccountView          `json:"views"`
	Tags            []string               `json:"tags"`
}

type AccountBalanceInfo struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type AccountBankRouting struct {
	Scheme  string `json:"scheme"`
	Address string `json:"address"`
}

type AccountRoutingInfo struct {
	Scheme  string `json:"scheme"`
	Address string `json:"address"`
}

type AccountRule struct {
	Scheme string `json:"scheme"`
	Value  string `json:"value"`
}

type AccountHolder struct {
	Name           string `json:"name"`
	IsAlias        bool   `json:"is_alias"`
	IsPrivateAlias bool   `json:"is_private_alias"`
	IsPublicAlias  bool   `json:"is_public_alias"`
}

type AccountView struct {
	ID          string `json:"id"`
	ShortName   string `json:"short_name"`
	Description string `json:"description"`
	IsPublic    bool   `json:"is_public"`
}

func (c *AccountController) GetCoreAccountByIdThroughView(ctx *gin.Context) {
	bankId := ctx.Param("BANK_ID")
	accountId := ctx.Param("ACCOUNT_ID")
	viewId := ctx.Param("VIEW_ID")

	response := CoreAccountResponse{
		ID:     accountId,
		BankID: bankId,
		Label:  "Main Account",
		Number: "123456789",
		Type:   "CURRENT",
		Balance: AccountBalanceInfo{
			Currency: "EUR",
			Amount:   "1500.50",
		},
		IBAN:     "DE89370400440532013000",
		SwiftBIC: "DEUTDEFF",
		BankRouting: AccountBankRouting{
			Scheme:  "IBAN",
			Address: "DE89370400440532013000",
		},
		AccountRouting: AccountRoutingInfo{
			Scheme:  "AccountNumber",
			Address: "123456789",
		},
		AccountRules: []AccountRule{
			{
				Scheme: "OVERDRAFT_LIMIT",
				Value:  "1000.00",
			},
		},
		AccountHolders: []AccountHolder{
			{
				Name:           "John Doe",
				IsAlias:        false,
				IsPrivateAlias: false,
				IsPublicAlias:  false,
			},
		},
		Views: []AccountView{
			{
				ID:          viewId,
				ShortName:   "Owner",
				Description: "Full access view",
				IsPublic:    false,
			},
		},
		Tags: []string{"personal", "main"},
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}
