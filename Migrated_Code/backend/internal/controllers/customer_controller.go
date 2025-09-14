package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/ashish-019-hash/obp-api-backend/internal/utils"
)

type CustomerController struct {
	orchestrationService *services.OrchestrationService
}

func NewCustomerController(orchestrationService *services.OrchestrationService) *CustomerController {
	return &CustomerController{
		orchestrationService: orchestrationService,
	}
}

type CustomerIdsResponse struct {
	CustomerIds []string `json:"customer_ids"`
}

type CustomerLegalNameRequest struct {
	LegalName string `json:"legal_name" binding:"required"`
}

type CustomerResponse struct {
	CustomerID   string `json:"customer_id"`
	LegalName    string `json:"legal_name"`
	MobileNumber string `json:"mobile_number"`
	Email        string `json:"email"`
	FaceImage    string `json:"face_image"`
	DateOfBirth  string `json:"date_of_birth"`
	RelationshipStatus string `json:"relationship_status"`
	Dependents   int    `json:"dependents"`
	DobbOfDependents []string `json:"dobb_of_dependents"`
	CreditRating string `json:"credit_rating"`
	CreditLimit  string `json:"credit_limit"`
	HighestEducationAttained string `json:"highest_education_attained"`
	EmploymentStatus string `json:"employment_status"`
	KycStatus    bool   `json:"kyc_status"`
	LastOkDate   string `json:"last_ok_date"`
}

func (c *CustomerController) GetCustomersForUserIdsOnly(ctx *gin.Context) {
	response := CustomerIdsResponse{
		CustomerIds: []string{
			"customer_123",
			"customer_456",
			"customer_789",
		},
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *CustomerController) GetCustomersByLegalName(ctx *gin.Context) {

	var req CustomerLegalNameRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	response := []CustomerResponse{
		{
			CustomerID:   "customer_123",
			LegalName:    req.LegalName,
			MobileNumber: "+1234567890",
			Email:        "customer@example.com",
			FaceImage:    "https://example.com/face.jpg",
			DateOfBirth:  "1990-01-01",
			RelationshipStatus: "Single",
			Dependents:   0,
			DobbOfDependents: []string{},
			CreditRating: "AAA",
			CreditLimit:  "50000.00",
			HighestEducationAttained: "Bachelor",
			EmploymentStatus: "Employed",
			KycStatus:    true,
			LastOkDate:   "2023-01-01",
		},
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}
