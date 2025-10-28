package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/ashish-019-hash/obp-api-backend/internal/utils"
	"github.com/ashish-019-hash/obp-api-backend/internal/models"
	"github.com/ashish-019-hash/obp-api-backend/pkg/db"
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

	var existingBank models.Bank
	if err := db.GetDB().Where("bank_id = ?", req.ID).First(&existingBank).Error; err == nil {
		utils.SendErrorResponse(ctx, http.StatusConflict, "Bank ID already exists", "A bank with this ID already exists")
		return
	}

	bank := models.Bank{
		BankID:             req.ID,
		ShortName:          req.ShortName,
		FullName:           req.FullName,
		LogoURL:            req.Logo,
		WebsiteURL:         req.Website,
		BankRoutingScheme:  req.BankRoutingScheme,
		BankRoutingAddress: req.BankRoutingAddress,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	if err := db.GetDB().Create(&bank).Error; err != nil {
		utils.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to create bank", err.Error())
		return
	}

	response := BankResponse{
		ID:               bank.BankID,
		ShortName:        bank.ShortName,
		FullName:         bank.FullName,
		Logo:             bank.LogoURL,
		Website:          bank.WebsiteURL,
		BankRoutingScheme: bank.BankRoutingScheme,
		BankRoutingAddress: bank.BankRoutingAddress,
		CreatedAt:        bank.CreatedAt,
	}

	utils.SendJSONResponse(ctx, http.StatusCreated, response)
}

func (c *BankController) GetBanks(ctx *gin.Context) {
	var banks []models.Bank
	if err := db.GetDB().Find(&banks).Error; err != nil {
		utils.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve banks", err.Error())
		return
	}

	bankResponses := make([]BankResponse, len(banks))
	for i, bank := range banks {
		bankResponses[i] = BankResponse{
			ID:               bank.BankID,
			ShortName:        bank.ShortName,
			FullName:         bank.FullName,
			Logo:             bank.LogoURL,
			Website:          bank.WebsiteURL,
			BankRoutingScheme: bank.BankRoutingScheme,
			BankRoutingAddress: bank.BankRoutingAddress,
			CreatedAt:        bank.CreatedAt,
		}
	}

	response := BanksResponse{
		Banks: bankResponses,
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *BankController) GetBankById(ctx *gin.Context) {
	bankId := ctx.Param("bankId")

	var bank models.Bank
	if err := db.GetDB().Where("bank_id = ?", bankId).First(&bank).Error; err != nil {
		utils.SendErrorResponse(ctx, http.StatusNotFound, "Bank not found", "Bank with ID "+bankId+" does not exist")
		return
	}

	response := BankResponse{
		ID:               bank.BankID,
		ShortName:        bank.ShortName,
		FullName:         bank.FullName,
		Logo:             bank.LogoURL,
		Website:          bank.WebsiteURL,
		BankRoutingScheme: bank.BankRoutingScheme,
		BankRoutingAddress: bank.BankRoutingAddress,
		CreatedAt:        bank.CreatedAt,
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}
