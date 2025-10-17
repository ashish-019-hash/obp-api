package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"obp-api-backend/internal/services"
)

type AdditionalRegulatoryController struct {
	accountService services.AccountService
	balanceService services.BalanceService
	paymentService services.PaymentService
}

func NewAdditionalRegulatoryController(
	accountService services.AccountService,
	balanceService services.BalanceService,
	paymentService services.PaymentService,
) *AdditionalRegulatoryController {
	return &AdditionalRegulatoryController{
		accountService: accountService,
		balanceService: balanceService,
		paymentService: paymentService,
	}
}

func (c *AdditionalRegulatoryController) GetHealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "All regulatory API endpoints are operational",
	})
}

func (c *AdditionalRegulatoryController) GetAPIVersions(ctx *gin.Context) {
	versions := map[string]interface{}{
		"supported_versions": []string{
			"v3.1.0", "v4.0.0", "v5.1.0",
		},
		"regulatory_standards": []string{
			"OBP Core", "Berlin Group PSD2", "UK Open Banking",
			"Australian CDR", "Bahrain OBF", "Polish API",
			"STET API", "MxOF API",
		},
	}
	ctx.JSON(http.StatusOK, versions)
}
