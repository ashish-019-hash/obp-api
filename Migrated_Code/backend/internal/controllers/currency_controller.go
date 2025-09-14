package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/ashish-019-hash/obp-api-backend/internal/utils"
)

type CurrencyController struct {
	orchestrationService *services.OrchestrationService
}

func NewCurrencyController(orchestrationService *services.OrchestrationService) *CurrencyController {
	return &CurrencyController{
		orchestrationService: orchestrationService,
	}
}

type CurrencyResponse struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type CurrenciesResponse struct {
	Currencies []CurrencyResponse `json:"currencies"`
}

func (c *CurrencyController) GetCurrenciesAtBank(ctx *gin.Context) {
	response := CurrenciesResponse{
		Currencies: []CurrencyResponse{
			{
				Code: "EUR",
				Name: "Euro",
			},
			{
				Code: "USD",
				Name: "US Dollar",
			},
			{
				Code: "GBP",
				Name: "British Pound",
			},
			{
				Code: "JPY",
				Name: "Japanese Yen",
			},
			{
				Code: "CHF",
				Name: "Swiss Franc",
			},
		},
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}
