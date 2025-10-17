package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"obp-api-backend/internal/services"
)

type BahrainOBFController struct {
	accountService services.AccountService
	balanceService services.BalanceService
	paymentService services.PaymentService
}

func NewBahrainOBFController(
	accountService services.AccountService,
	balanceService services.BalanceService,
	paymentService services.PaymentService,
) *BahrainOBFController {
	return &BahrainOBFController{
		accountService: accountService,
		balanceService: balanceService,
		paymentService: paymentService,
	}
}

func (c *BahrainOBFController) GetAccounts(ctx *gin.Context) {
	accounts, err := c.accountService.GetAccountsForUser(ctx.Request.Context(), "current-user")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"Account": accounts,
		},
		"Links": map[string]interface{}{
			"Self": "/bahrain-obf/v1.0.0/accounts",
		},
		"Meta": map[string]interface{}{
			"TotalPages": 1,
		},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetAccount(ctx *gin.Context) {
	accountID := ctx.Param("AccountId")

	account, err := c.accountService.GetAccountByID(ctx.Request.Context(), accountID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"Account": account,
		},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetAccountBalances(ctx *gin.Context) {
	accountID := ctx.Param("AccountId")

	currentBalance, err := c.balanceService.CalculateCurrentBalance(ctx.Request.Context(), accountID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	balances := []map[string]interface{}{
		{
			"AccountId": accountID,
			"Amount": map[string]interface{}{
				"Amount":   currentBalance.String(),
				"Currency": "BHD",
			},
			"CreditDebitIndicator": "Credit",
			"Type":                 "ClosingAvailable",
		},
	}

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"Balance": balances,
		},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) CreateDomesticPaymentConsent(ctx *gin.Context) {
	var consentRequest map[string]interface{}
	if err := ctx.ShouldBindJSON(&consentRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	consentID := "consent-123"

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"ConsentId": consentID,
			"Status":    "AwaitingAuthorisation",
		},
	}
	ctx.JSON(http.StatusCreated, response)
}

func (c *BahrainOBFController) CreateDomesticPayment(ctx *gin.Context) {
	var paymentRequest map[string]interface{}
	if err := ctx.ShouldBindJSON(&paymentRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	paymentID, err := c.paymentService.InitiatePayment(ctx.Request.Context(), paymentRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"DomesticPaymentId": paymentID,
			"Status":            "AcceptedSettlementInProcess",
		},
	}
	ctx.JSON(http.StatusCreated, response)
}

func (c *BahrainOBFController) GetBalances(ctx *gin.Context) {
	balances := []map[string]interface{}{}

	response := map[string]interface{}{
		"Data":  balances,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) CreateDomesticPaymentConsents(ctx *gin.Context) {
	var consentData map[string]interface{}
	if err := ctx.ShouldBindJSON(&consentData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"Data":  consentData,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusCreated, response)
}

func (c *BahrainOBFController) CreateDomesticPayments(ctx *gin.Context) {
	var paymentData map[string]interface{}
	if err := ctx.ShouldBindJSON(&paymentData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"Data":  paymentData,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusCreated, response)
}

func (c *BahrainOBFController) CreateFilePaymentConsents(ctx *gin.Context) {
	var consentData map[string]interface{}
	if err := ctx.ShouldBindJSON(&consentData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"Data":  consentData,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusCreated, response)
}

func (c *BahrainOBFController) CreateFilePayments(ctx *gin.Context) {
	var paymentData map[string]interface{}
	if err := ctx.ShouldBindJSON(&paymentData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"Data":  paymentData,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusCreated, response)
}

func (c *BahrainOBFController) CreateInternationalPaymentConsents(ctx *gin.Context) {
	var consentData map[string]interface{}
	if err := ctx.ShouldBindJSON(&consentData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"Data":  consentData,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusCreated, response)
}

func (c *BahrainOBFController) CreateInternationalPayments(ctx *gin.Context) {
	var paymentData map[string]interface{}
	if err := ctx.ShouldBindJSON(&paymentData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"Data":  paymentData,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusCreated, response)
}

func (c *BahrainOBFController) CreateDomesticFutureDatedPaymentConsents(ctx *gin.Context) {
	var consentData map[string]interface{}
	if err := ctx.ShouldBindJSON(&consentData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"Data":  consentData,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusCreated, response)
}

func (c *BahrainOBFController) CreateDomesticFutureDatedPayments(ctx *gin.Context) {
	var paymentData map[string]interface{}
	if err := ctx.ShouldBindJSON(&paymentData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"Data":  paymentData,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusCreated, response)
}

func (c *BahrainOBFController) GetAccountStatements(ctx *gin.Context) {
	_ = ctx.Param("AccountId")
	statements := []map[string]interface{}{}

	response := map[string]interface{}{
		"Data":  statements,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetAccountStatement(ctx *gin.Context) {
	accountID := ctx.Param("AccountId")
	statementID := ctx.Param("StatementId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"Statement": []map[string]interface{}{
				{
					"AccountId":   accountID,
					"StatementId": statementID,
				},
			},
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetAccountStandingOrders(ctx *gin.Context) {
	_ = ctx.Param("AccountId")
	standingOrders := []map[string]interface{}{}

	response := map[string]interface{}{
		"Data":  standingOrders,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetAccountOffers(ctx *gin.Context) {
	_ = ctx.Param("AccountId")
	offers := []map[string]interface{}{}

	response := map[string]interface{}{
		"Data":  offers,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetAccountDirectDebits(ctx *gin.Context) {
	_ = ctx.Param("AccountId")
	directDebits := []map[string]interface{}{}

	response := map[string]interface{}{
		"Data":  directDebits,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetAccountBeneficiaries(ctx *gin.Context) {
	_ = ctx.Param("AccountId")
	beneficiaries := []map[string]interface{}{}

	response := map[string]interface{}{
		"Data":  beneficiaries,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetAccountParty(ctx *gin.Context) {
	_ = ctx.Param("AccountId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"Party": map[string]interface{}{
				"PartyId":   "PARTY1",
				"PartyType": "Individual",
			},
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetAccountSupplementaryAccountInfo(ctx *gin.Context) {
	_ = ctx.Param("AccountId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"SupplementaryAccountInfo": map[string]interface{}{
				"AccountId": ctx.Param("AccountId"),
			},
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetFilePaymentConsentFile(ctx *gin.Context) {
	_ = ctx.Param("consentId")

	ctx.Header("Content-Type", "application/xml")
	ctx.Data(http.StatusOK, "application/xml", []byte("<xml>File content</xml>"))
}

func (c *BahrainOBFController) CreateFilePaymentConsentFile(ctx *gin.Context) {
	_ = ctx.Param("consentId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"FileType": "BH.OBF.pain.001.001.08",
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetFilePaymentConsent(ctx *gin.Context) {
	consentID := ctx.Param("consentId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"ConsentId": consentID,
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetInternationalPayment(ctx *gin.Context) {
	paymentID := ctx.Param("internationalPaymentId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"InternationalPaymentId": paymentID,
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetInternationalPaymentDetails(ctx *gin.Context) {
	paymentID := ctx.Param("internationalPaymentId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"InternationalPaymentId": paymentID,
			"PaymentDetails":         map[string]interface{}{},
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetAccountStatementFile(ctx *gin.Context) {
	_ = ctx.Param("AccountId")
	_ = ctx.Param("StatementId")

	ctx.Header("Content-Type", "application/pdf")
	ctx.Data(http.StatusOK, "application/pdf", []byte("PDF content"))
}

func (c *BahrainOBFController) GetAccountStatementTransactions(ctx *gin.Context) {
	_ = ctx.Param("AccountId")
	_ = ctx.Param("StatementId")
	transactions := []map[string]interface{}{}

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"Transaction": transactions,
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetStatements(ctx *gin.Context) {
	statements := []map[string]interface{}{}

	response := map[string]interface{}{
		"Data":  statements,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetDomesticFutureDatedPayment(ctx *gin.Context) {
	paymentID := ctx.Param("domesticFutureDatedPaymentId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"DomesticFutureDatedPaymentId": paymentID,
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) UpdateDomesticFutureDatedPayment(ctx *gin.Context) {
	paymentID := ctx.Param("domesticFutureDatedPaymentId")
	var updateData map[string]interface{}
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"DomesticFutureDatedPaymentId": paymentID,
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetDomesticFutureDatedPaymentDetails(ctx *gin.Context) {
	paymentID := ctx.Param("domesticFutureDatedPaymentId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"DomesticFutureDatedPaymentId": paymentID,
			"PaymentDetails":               map[string]interface{}{},
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetStandingOrders(ctx *gin.Context) {
	standingOrders := []map[string]interface{}{}

	response := map[string]interface{}{
		"Data":  standingOrders,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetOffers(ctx *gin.Context) {
	offers := []map[string]interface{}{}

	response := map[string]interface{}{
		"Data":  offers,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetDomesticPayment(ctx *gin.Context) {
	paymentID := ctx.Param("domesticPaymentId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"DomesticPaymentId": paymentID,
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetDomesticPaymentDetails(ctx *gin.Context) {
	paymentID := ctx.Param("domesticPaymentId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"DomesticPaymentId": paymentID,
			"PaymentDetails":    map[string]interface{}{},
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetDirectDebits(ctx *gin.Context) {
	directDebits := []map[string]interface{}{}

	response := map[string]interface{}{
		"Data":  directDebits,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetDomesticPaymentConsentFundsConfirmation(ctx *gin.Context) {
	consentID := ctx.Param("consentId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"ConsentId":      consentID,
			"FundsAvailable": true,
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetDomesticPaymentConsent(ctx *gin.Context) {
	consentID := ctx.Param("consentId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"ConsentId": consentID,
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetAccountFutureDatedPayments(ctx *gin.Context) {
	_ = ctx.Param("AccountId")
	payments := []map[string]interface{}{}

	response := map[string]interface{}{
		"Data":  payments,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetFutureDatedPayments(ctx *gin.Context) {
	payments := []map[string]interface{}{}

	response := map[string]interface{}{
		"Data":  payments,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetBeneficiaries(ctx *gin.Context) {
	beneficiaries := []map[string]interface{}{}

	response := map[string]interface{}{
		"Data":  beneficiaries,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetFilePayment(ctx *gin.Context) {
	paymentID := ctx.Param("filePaymentId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"FilePaymentId": paymentID,
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetFilePaymentDetails(ctx *gin.Context) {
	paymentID := ctx.Param("filePaymentId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"FilePaymentId":  paymentID,
			"PaymentDetails": map[string]interface{}{},
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetFilePaymentReportFile(ctx *gin.Context) {
	_ = ctx.Param("filePaymentId")

	ctx.Header("Content-Type", "application/xml")
	ctx.Data(http.StatusOK, "application/xml", []byte("<xml>Report content</xml>"))
}

func (c *BahrainOBFController) GetAccountTransactions(ctx *gin.Context) {
	_ = ctx.Param("AccountId")
	transactions := []map[string]interface{}{}

	response := map[string]interface{}{
		"Data":  transactions,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetTransactions(ctx *gin.Context) {
	transactions := []map[string]interface{}{}

	response := map[string]interface{}{
		"Data":  transactions,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) UpdateAccountAccessConsent(ctx *gin.Context) {
	consentID := ctx.Param("consentId")
	var updateData map[string]interface{}
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"ConsentId": consentID,
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) CreateAccountAccessConsent(ctx *gin.Context) {
	var consentData map[string]interface{}
	if err := ctx.ShouldBindJSON(&consentData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"Data":  consentData,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusCreated, response)
}

func (c *BahrainOBFController) CreateEventNotification(ctx *gin.Context) {
	var eventData map[string]interface{}
	if err := ctx.ShouldBindJSON(&eventData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"Data":  eventData,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusCreated, response)
}

func (c *BahrainOBFController) GetAccountParties(ctx *gin.Context) {
	_ = ctx.Param("AccountId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"Party": []map[string]interface{}{
				{
					"PartyId":   "PARTY1",
					"PartyType": "Individual",
				},
			},
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetParty(ctx *gin.Context) {
	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"Party": map[string]interface{}{
				"PartyId":   "PARTY1",
				"PartyType": "Individual",
			},
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) CreateInternationalScheduledPaymentConsents(ctx *gin.Context) {
	var consentData map[string]interface{}
	if err := ctx.ShouldBindJSON(&consentData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"Data":  consentData,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusCreated, response)
}

func (c *BahrainOBFController) CreateInternationalScheduledPayments(ctx *gin.Context) {
	var paymentData map[string]interface{}
	if err := ctx.ShouldBindJSON(&paymentData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"Data":  paymentData,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusCreated, response)
}

func (c *BahrainOBFController) GetInternationalScheduledPaymentConsent(ctx *gin.Context) {
	consentID := ctx.Param("consentId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"ConsentId": consentID,
			"Status":    "Authorised",
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetInternationalScheduledPayment(ctx *gin.Context) {
	paymentID := ctx.Param("internationalScheduledPaymentId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"InternationalScheduledPaymentId": paymentID,
			"Status":                          "Pending",
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) CreateInternationalStandingOrderConsents(ctx *gin.Context) {
	var consentData map[string]interface{}
	if err := ctx.ShouldBindJSON(&consentData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"Data":  consentData,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusCreated, response)
}

func (c *BahrainOBFController) CreateInternationalStandingOrders(ctx *gin.Context) {
	var orderData map[string]interface{}
	if err := ctx.ShouldBindJSON(&orderData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"Data":  orderData,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusCreated, response)
}

func (c *BahrainOBFController) GetInternationalStandingOrderConsent(ctx *gin.Context) {
	consentID := ctx.Param("consentId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"ConsentId": consentID,
			"Status":    "Authorised",
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetInternationalStandingOrder(ctx *gin.Context) {
	orderID := ctx.Param("internationalStandingOrderPaymentId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"InternationalStandingOrderId": orderID,
			"Status":                       "Active",
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) CreateDomesticStandingOrderConsents(ctx *gin.Context) {
	var consentData map[string]interface{}
	if err := ctx.ShouldBindJSON(&consentData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"Data":  consentData,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusCreated, response)
}

func (c *BahrainOBFController) CreateDomesticStandingOrders(ctx *gin.Context) {
	var orderData map[string]interface{}
	if err := ctx.ShouldBindJSON(&orderData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"Data":  orderData,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusCreated, response)
}

func (c *BahrainOBFController) GetDomesticStandingOrderConsent(ctx *gin.Context) {
	consentID := ctx.Param("consentId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"ConsentId": consentID,
			"Status":    "Authorised",
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetDomesticStandingOrder(ctx *gin.Context) {
	orderID := ctx.Param("domesticStandingOrderId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"DomesticStandingOrderId": orderID,
			"Status":                  "Active",
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) CreateDomesticScheduledPaymentConsents(ctx *gin.Context) {
	var consentData map[string]interface{}
	if err := ctx.ShouldBindJSON(&consentData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"Data":  consentData,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusCreated, response)
}

func (c *BahrainOBFController) CreateDomesticScheduledPayments(ctx *gin.Context) {
	var paymentData map[string]interface{}
	if err := ctx.ShouldBindJSON(&paymentData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"Data":  paymentData,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusCreated, response)
}

func (c *BahrainOBFController) GetDomesticScheduledPaymentConsent(ctx *gin.Context) {
	consentID := ctx.Param("consentId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"ConsentId": consentID,
			"Status":    "Authorised",
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetDomesticScheduledPayment(ctx *gin.Context) {
	paymentID := ctx.Param("domesticScheduledPaymentId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"DomesticScheduledPaymentId": paymentID,
			"Status":                     "Pending",
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) CreateAccountAccessConsents(ctx *gin.Context) {
	var consentData map[string]interface{}
	if err := ctx.ShouldBindJSON(&consentData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"Data":  consentData,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusCreated, response)
}

func (c *BahrainOBFController) GetAccountAccessConsent(ctx *gin.Context) {
	consentID := ctx.Param("consentId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"ConsentId": consentID,
			"Status":    "Authorised",
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) DeleteAccountAccessConsent(ctx *gin.Context) {
	_ = ctx.Param("consentId")
	ctx.JSON(http.StatusNoContent, nil)
}

func (c *BahrainOBFController) CreateFundsConfirmationConsents(ctx *gin.Context) {
	var consentData map[string]interface{}
	if err := ctx.ShouldBindJSON(&consentData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"Data":  consentData,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusCreated, response)
}

func (c *BahrainOBFController) CreateFundsConfirmations(ctx *gin.Context) {
	var confirmationData map[string]interface{}
	if err := ctx.ShouldBindJSON(&confirmationData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"Data":  confirmationData,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusCreated, response)
}

func (c *BahrainOBFController) GetFundsConfirmationConsent(ctx *gin.Context) {
	consentID := ctx.Param("consentId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"ConsentId": consentID,
			"Status":    "Authorised",
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) DeleteFundsConfirmationConsent(ctx *gin.Context) {
	_ = ctx.Param("consentId")
	ctx.JSON(http.StatusNoContent, nil)
}

func (c *BahrainOBFController) GetProducts(ctx *gin.Context) {
	products := []map[string]interface{}{}

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"Product": products,
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetAccountProduct(ctx *gin.Context) {
	accountID := ctx.Param("AccountId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"AccountId": accountID,
			"Product":   map[string]interface{}{},
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetInternationalPaymentConsentFundsConfirmation(ctx *gin.Context) {
	consentID := ctx.Param("consentId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"ConsentId":      consentID,
			"FundsAvailable": true,
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetInternationalPaymentConsent(ctx *gin.Context) {
	consentID := ctx.Param("consentId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"ConsentId": consentID,
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetDomesticFutureDatedPaymentCancellationConsent(ctx *gin.Context) {
	consentID := ctx.Param("consentId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"ConsentId": consentID,
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) CreateDomesticFutureDatedPaymentCancellationConsent(ctx *gin.Context) {
	var consentData map[string]interface{}
	if err := ctx.ShouldBindJSON(&consentData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"Data":  consentData,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusCreated, response)
}

func (c *BahrainOBFController) GetDomesticFutureDatedPaymentConsent(ctx *gin.Context) {
	consentID := ctx.Param("consentId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"ConsentId": consentID,
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) PatchDomesticFutureDatedPayment(ctx *gin.Context) {
	paymentID := ctx.Param("DomesticFutureDatedPaymentId")

	var patchData map[string]interface{}
	if err := ctx.ShouldBindJSON(&patchData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"DomesticFutureDatedPaymentId": paymentID,
			"Status":                       "Updated",
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *BahrainOBFController) GetAccountSupplementaryAccountInfoNew(ctx *gin.Context) {
	accountID := ctx.Param("AccountId")
	info := map[string]interface{}{
		"Data": map[string]interface{}{
			"Account": map[string]interface{}{
				"AccountId": accountID,
				"SupplementaryData": map[string]interface{}{
					"AccountType":    "Personal",
					"AccountSubType": "CurrentAccount",
				},
			},
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, info)
}

func (c *BahrainOBFController) GetDomesticFutureDatedPaymentDetailsNew(ctx *gin.Context) {
	paymentID := ctx.Param("DomesticFutureDatedPaymentId")
	details := map[string]interface{}{
		"Data": map[string]interface{}{
			"DomesticFutureDatedPaymentId": paymentID,
			"Status":                       "Pending",
			"CreationDateTime":             "2023-01-01T00:00:00Z",
			"StatusUpdateDateTime":         "2023-01-01T00:00:00Z",
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, details)
}

func (c *BahrainOBFController) PatchDomesticFutureDatedPaymentNew(ctx *gin.Context) {
	paymentID := ctx.Param("DomesticFutureDatedPaymentId")
	var patchData map[string]interface{}
	if err := ctx.ShouldBindJSON(&patchData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	patchData["payment_id"] = paymentID
	ctx.JSON(http.StatusOK, patchData)
}
