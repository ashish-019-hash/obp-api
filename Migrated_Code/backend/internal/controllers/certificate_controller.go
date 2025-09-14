package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/ashish-019-hash/obp-api-backend/internal/utils"
)

type CertificateController struct {
	orchestrationService *services.OrchestrationService
}

func NewCertificateController(orchestrationService *services.OrchestrationService) *CertificateController {
	return &CertificateController{
		orchestrationService: orchestrationService,
	}
}

type CertificateInfoResponse struct {
	CommonName   string    `json:"common_name"`
	Organization string    `json:"organization"`
	Email        string    `json:"email"`
	ValidFrom    time.Time `json:"valid_from"`
	ValidTo      time.Time `json:"valid_to"`
	SerialNumber string    `json:"serial_number"`
	Issuer       string    `json:"issuer"`
}

func (c *CertificateController) GetMtlsClientCertificateInfo(ctx *gin.Context) {
	response := CertificateInfoResponse{
		CommonName:   "client.example.com",
		Organization: "Example Corp",
		Email:        "admin@example.com",
		ValidFrom:    time.Now().Add(-30 * 24 * time.Hour),
		ValidTo:      time.Now().Add(365 * 24 * time.Hour),
		SerialNumber: "123456789",
		Issuer:       "Example CA",
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}
