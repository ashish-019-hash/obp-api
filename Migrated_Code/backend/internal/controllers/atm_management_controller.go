package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/ashish-019-hash/obp-api-backend/internal/utils"
)

type AtmManagementController struct {
	orchestrationService *services.OrchestrationService
}

func NewAtmManagementController(orchestrationService *services.OrchestrationService) *AtmManagementController {
	return &AtmManagementController{
		orchestrationService: orchestrationService,
	}
}

type AtmRequest struct {
	ID          string      `json:"id"`
	BankID      string      `json:"bank_id" binding:"required"`
	Name        string      `json:"name" binding:"required"`
	Address     AtmAddress  `json:"address" binding:"required"`
	Location    AtmLocation `json:"location" binding:"required"`
	Meta        AtmMeta     `json:"meta"`
}

type AtmAddress struct {
	Line1      string `json:"line_1"`
	Line2      string `json:"line_2"`
	Line3      string `json:"line_3"`
	City       string `json:"city"`
	County     string `json:"county"`
	State      string `json:"state"`
	PostCode   string `json:"post_code"`
	CountryCode string `json:"country_code"`
}

type AtmLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type AtmMeta struct {
	License AtmLicense `json:"license"`
}

type AtmLicense struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AtmResponse struct {
	ID          string      `json:"id"`
	BankID      string      `json:"bank_id"`
	Name        string      `json:"name"`
	Address     AtmAddress  `json:"address"`
	Location    AtmLocation `json:"location"`
	Meta        AtmMeta     `json:"meta"`
	Attributes  []AtmAttributeResponse `json:"attributes"`
}

type AtmsResponse struct {
	Atms []AtmResponse `json:"atms"`
}

func (c *AtmManagementController) CreateAtm(ctx *gin.Context) {
	bankId := ctx.Param("bankId")

	var req AtmRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	if req.BankID != bankId {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Bank ID mismatch", "")
		return
	}

	response := AtmResponse{
		ID:         "atm_" + strconv.FormatInt(time.Now().Unix(), 10),
		BankID:     req.BankID,
		Name:       req.Name,
		Address:    req.Address,
		Location:   req.Location,
		Meta:       req.Meta,
		Attributes: []AtmAttributeResponse{},
	}

	utils.SendJSONResponse(ctx, http.StatusCreated, response)
}

func (c *AtmManagementController) UpdateAtm(ctx *gin.Context) {
	bankId := ctx.Param("bankId")
	atmId := ctx.Param("atmId")

	var req AtmRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	if req.BankID != bankId {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Bank ID mismatch", "")
		return
	}

	response := AtmResponse{
		ID:         atmId,
		BankID:     req.BankID,
		Name:       req.Name,
		Address:    req.Address,
		Location:   req.Location,
		Meta:       req.Meta,
		Attributes: []AtmAttributeResponse{},
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *AtmManagementController) GetAtms(ctx *gin.Context) {
	bankId := ctx.Param("bankId")

	response := AtmsResponse{
		Atms: []AtmResponse{
			{
				ID:     "atm_001",
				BankID: bankId,
				Name:   "Main Branch ATM",
				Address: AtmAddress{
					Line1:       "123 Main St",
					City:        "New York",
					State:       "NY",
					PostCode:    "10001",
					CountryCode: "US",
				},
				Location: AtmLocation{
					Latitude:  40.7128,
					Longitude: -74.0060,
				},
				Meta: AtmMeta{
					License: AtmLicense{
						ID:   "license_001",
						Name: "Standard ATM License",
					},
				},
				Attributes: []AtmAttributeResponse{},
			},
		},
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *AtmManagementController) GetAtm(ctx *gin.Context) {
	bankId := ctx.Param("bankId")
	atmId := ctx.Param("atmId")

	response := AtmResponse{
		ID:     atmId,
		BankID: bankId,
		Name:   "Main Branch ATM",
		Address: AtmAddress{
			Line1:       "123 Main St",
			City:        "New York",
			State:       "NY",
			PostCode:    "10001",
			CountryCode: "US",
		},
		Location: AtmLocation{
			Latitude:  40.7128,
			Longitude: -74.0060,
		},
		Meta: AtmMeta{
			License: AtmLicense{
				ID:   "license_001",
				Name: "Standard ATM License",
			},
		},
		Attributes: []AtmAttributeResponse{},
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *AtmManagementController) DeleteAtm(ctx *gin.Context) {
	utils.SendJSONResponse(ctx, http.StatusNoContent, nil)
}
