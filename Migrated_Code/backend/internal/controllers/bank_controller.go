package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/ashish-019-hash/obp-api-backend/internal/utils"
)

type BankController struct {
	orchestrationService *services.OrchestrationService
}

func NewBankController(orchestrationService *services.OrchestrationService) *BankController {
	return &BankController{
		orchestrationService: orchestrationService,
	}
}

type CreateBankRequest struct {
	ID               string `json:"id" binding:"required"`
	ShortName        string `json:"short_name" binding:"required"`
	FullName         string `json:"full_name" binding:"required"`
	Logo             string `json:"logo"`
	Website          string `json:"website"`
	BankRoutingScheme string `json:"bank_routing_scheme"`
	BankRoutingAddress string `json:"bank_routing_address"`
}

type BankResponse struct {
	ID               string    `json:"id"`
	ShortName        string    `json:"short_name"`
	FullName         string    `json:"full_name"`
	Logo             string    `json:"logo"`
	Website          string    `json:"website"`
	BankRoutingScheme string   `json:"bank_routing_scheme"`
	BankRoutingAddress string  `json:"bank_routing_address"`
	CreatedAt        time.Time `json:"created_at"`
}

type BanksResponse struct {
	Banks []BankResponse `json:"banks"`
}

func (c *BankController) CreateBank(ctx *gin.Context) {
	var req CreateBankRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	if len(req.ID) < 3 || len(req.ID) > 50 {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid bank ID", "Bank ID must be between 3 and 50 characters")
		return
	}

	response := BankResponse{
		ID:               req.ID,
		ShortName:        req.ShortName,
		FullName:         req.FullName,
		Logo:             req.Logo,
		Website:          req.Website,
		BankRoutingScheme: req.BankRoutingScheme,
		BankRoutingAddress: req.BankRoutingAddress,
		CreatedAt:        time.Now(),
	}

	utils.SendJSONResponse(ctx, http.StatusCreated, response)
}

func (c *BankController) GetBanks(ctx *gin.Context) {
	response := BanksResponse{
		Banks: []BankResponse{
			{
				ID:               "bank_001",
				ShortName:        "OBP Bank",
				FullName:         "Open Bank Project Demo Bank",
				Logo:             "https://static.openbankproject.com/images/sandbox/bank_x.png",
				Website:          "https://www.example.com",
				BankRoutingScheme: "OBP",
				BankRoutingAddress: "obp.bank.001",
				CreatedAt:        time.Now(),
			},
		},
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *BankController) GetBankById(ctx *gin.Context) {
	bankId := ctx.Param("bankId")

	response := BankResponse{
		ID:               bankId,
		ShortName:        "OBP Bank",
		FullName:         "Open Bank Project Demo Bank",
		Logo:             "https://static.openbankproject.com/images/sandbox/bank_x.png",
		Website:          "https://www.example.com",
		BankRoutingScheme: "OBP",
		BankRoutingAddress: "obp.bank." + bankId,
		CreatedAt:        time.Now(),
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}
