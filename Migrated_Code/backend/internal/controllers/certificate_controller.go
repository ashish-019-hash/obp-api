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
	if ctx.Request.TLS == nil || len(ctx.Request.TLS.PeerCertificates) == 0 {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "No client certificate provided", "")
		return
	}

	cert := ctx.Request.TLS.PeerCertificates[0]
	
	var email string
	if len(cert.EmailAddresses) > 0 {
		email = cert.EmailAddresses[0]
	}

	var organization string
	if len(cert.Subject.Organization) > 0 {
		organization = cert.Subject.Organization[0]
	}

	response := CertificateInfoResponse{
		CommonName:   cert.Subject.CommonName,
		Organization: organization,
		Email:        email,
		ValidFrom:    cert.NotBefore,
		ValidTo:      cert.NotAfter,
		SerialNumber: cert.SerialNumber.String(),
		Issuer:       cert.Issuer.CommonName,
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}
