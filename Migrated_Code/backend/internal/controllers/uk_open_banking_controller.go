package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"obp-api-backend/internal/services"
)

type UKOpenBankingController struct {
	accountService services.AccountService
	balanceService services.BalanceService
	paymentService services.PaymentService
}

func NewUKOpenBankingController(
	accountService services.AccountService,
	balanceService services.BalanceService,
	paymentService services.PaymentService,
) *UKOpenBankingController {
	return &UKOpenBankingController{
		accountService: accountService,
		balanceService: balanceService,
		paymentService: paymentService,
	}
}

func (c *UKOpenBankingController) GetAccounts(ctx *gin.Context) {
	accounts, err := c.accountService.GetAccountsForUser(ctx.Request.Context(), "current-user")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ukAccounts := make([]map[string]interface{}, len(accounts))
	for i, account := range accounts {
		ukAccounts[i] = map[string]interface{}{
			"AccountId":      account.AccountId,
			"Currency":       account.Currency,
			"AccountType":    "Personal",
			"AccountSubType": "CurrentAccount",
			"Nickname":       account.Label,
		}
	}

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"Account": ukAccounts,
		},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) GetAccountBalances(ctx *gin.Context) {
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
				"Currency": "GBP",
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

func (c *UKOpenBankingController) GetAccountTransactions(ctx *gin.Context) {
	accountID := ctx.Param("AccountId")

	transactions, err := c.accountService.GetTransactionsByAccountID(ctx.Request.Context(), accountID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"Transaction": transactions,
		},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) CreateDomesticPaymentConsents(ctx *gin.Context) {
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

func (c *UKOpenBankingController) CreateDomesticPayments(ctx *gin.Context) {
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

func (c *UKOpenBankingController) GetDomesticPaymentConsent(ctx *gin.Context) {
	consentID := ctx.Param("ConsentId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"ConsentId": consentID,
			"Status":    "Authorised",
		},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) CreateFundsConfirmationConsents(ctx *gin.Context) {
	var consentRequest map[string]interface{}
	if err := ctx.ShouldBindJSON(&consentRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	consentID := "funds-consent-123"

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"ConsentId": consentID,
			"Status":    "AwaitingAuthorisation",
		},
	}
	ctx.JSON(http.StatusCreated, response)
}

func (c *UKOpenBankingController) CreateFundsConfirmations(ctx *gin.Context) {
	var fundsRequest map[string]interface{}
	if err := ctx.ShouldBindJSON(&fundsRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"FundsAvailable": true,
		},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) CreateInternationalPaymentConsents(ctx *gin.Context) {
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

func (c *UKOpenBankingController) CreateInternationalPayments(ctx *gin.Context) {
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

func (c *UKOpenBankingController) GetInternationalPaymentConsent(ctx *gin.Context) {
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

func (c *UKOpenBankingController) GetInternationalPaymentConsentFundsConfirmation(ctx *gin.Context) {
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

func (c *UKOpenBankingController) GetInternationalPayment(ctx *gin.Context) {
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

func (c *UKOpenBankingController) GetAccountStatements(ctx *gin.Context) {
	_ = ctx.Param("AccountId")
	statements := []map[string]interface{}{}

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"Statement": statements,
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) GetAccountStatement(ctx *gin.Context) {
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

func (c *UKOpenBankingController) GetAccountStatementFile(ctx *gin.Context) {
	_ = ctx.Param("AccountId")
	_ = ctx.Param("StatementId")

	ctx.Header("Content-Type", "application/pdf")
	ctx.Data(http.StatusOK, "application/pdf", []byte("PDF content"))
}

func (c *UKOpenBankingController) GetAccountStatementTransactions(ctx *gin.Context) {
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

func (c *UKOpenBankingController) GetStatements(ctx *gin.Context) {
	statements := []map[string]interface{}{}

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"Statement": statements,
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) GetAccountStandingOrders(ctx *gin.Context) {
	_ = ctx.Param("AccountId")
	standingOrders := []map[string]interface{}{}

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"StandingOrder": standingOrders,
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) GetStandingOrders(ctx *gin.Context) {
	standingOrders := []map[string]interface{}{}

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"StandingOrder": standingOrders,
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) CreateInternationalScheduledPaymentConsents(ctx *gin.Context) {
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

func (c *UKOpenBankingController) CreateInternationalScheduledPayments(ctx *gin.Context) {
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

func (c *UKOpenBankingController) GetInternationalScheduledPaymentConsent(ctx *gin.Context) {
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

func (c *UKOpenBankingController) GetInternationalScheduledPaymentConsentFundsConfirmation(ctx *gin.Context) {
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

func (c *UKOpenBankingController) GetInternationalScheduledPayment(ctx *gin.Context) {
	paymentID := ctx.Param("internationalScheduledPaymentId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"InternationalScheduledPaymentId": paymentID,
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) CreateInternationalStandingOrderConsents(ctx *gin.Context) {
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

func (c *UKOpenBankingController) CreateInternationalStandingOrders(ctx *gin.Context) {
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

func (c *UKOpenBankingController) GetInternationalStandingOrderConsent(ctx *gin.Context) {
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

func (c *UKOpenBankingController) GetInternationalStandingOrder(ctx *gin.Context) {
	paymentID := ctx.Param("internationalStandingOrderPaymentId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"InternationalStandingOrderPaymentId": paymentID,
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) GetAccountOffers(ctx *gin.Context) {
	_ = ctx.Param("AccountId")
	offers := []map[string]interface{}{}

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"Offer": offers,
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) GetOffers(ctx *gin.Context) {
	offers := []map[string]interface{}{}

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"Offer": offers,
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) CreateDomesticStandingOrderConsents(ctx *gin.Context) {
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

func (c *UKOpenBankingController) CreateDomesticStandingOrders(ctx *gin.Context) {
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

func (c *UKOpenBankingController) GetDomesticStandingOrderConsent(ctx *gin.Context) {
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

func (c *UKOpenBankingController) GetDomesticStandingOrder(ctx *gin.Context) {
	paymentID := ctx.Param("domesticStandingOrderId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"DomesticStandingOrderId": paymentID,
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) GetAccountDirectDebits(ctx *gin.Context) {
	_ = ctx.Param("AccountId")
	directDebits := []map[string]interface{}{}

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"DirectDebit": directDebits,
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) GetDirectDebits(ctx *gin.Context) {
	directDebits := []map[string]interface{}{}

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"DirectDebit": directDebits,
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) CreateAccountAccessConsents(ctx *gin.Context) {
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

func (c *UKOpenBankingController) DeleteAccountAccessConsent(ctx *gin.Context) {
	_ = ctx.Param("consentId")
	ctx.JSON(http.StatusNoContent, nil)
}

func (c *UKOpenBankingController) GetAccountAccessConsent(ctx *gin.Context) {
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

func (c *UKOpenBankingController) CreateDomesticScheduledPaymentConsents(ctx *gin.Context) {
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

func (c *UKOpenBankingController) CreateDomesticScheduledPayments(ctx *gin.Context) {
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

func (c *UKOpenBankingController) GetDomesticScheduledPaymentConsent(ctx *gin.Context) {
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

func (c *UKOpenBankingController) GetDomesticScheduledPayment(ctx *gin.Context) {
	paymentID := ctx.Param("domesticScheduledPaymentId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"DomesticScheduledPaymentId": paymentID,
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) GetBalances(ctx *gin.Context) {
	balances := []map[string]interface{}{}

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"Balance": balances,
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) GetAccountBeneficiaries(ctx *gin.Context) {
	_ = ctx.Param("AccountId")
	beneficiaries := []map[string]interface{}{}

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"Beneficiary": beneficiaries,
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) GetBeneficiaries(ctx *gin.Context) {
	beneficiaries := []map[string]interface{}{}

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"Beneficiary": beneficiaries,
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) GetTransactions(ctx *gin.Context) {
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

func (c *UKOpenBankingController) GetFundsConfirmationConsent(ctx *gin.Context) {
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

func (c *UKOpenBankingController) DeleteFundsConfirmationConsent(ctx *gin.Context) {
	_ = ctx.Param("consentId")
	ctx.JSON(http.StatusNoContent, nil)
}

func (c *UKOpenBankingController) CreateFilePaymentConsents(ctx *gin.Context) {
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

func (c *UKOpenBankingController) CreateFilePaymentConsentFile(ctx *gin.Context) {
	_ = ctx.Param("consentId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"FileType": "UK.OBIE.pain.001.001.08",
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) CreateFilePayments(ctx *gin.Context) {
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

func (c *UKOpenBankingController) GetFilePaymentConsent(ctx *gin.Context) {
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

func (c *UKOpenBankingController) GetFilePaymentConsentFile(ctx *gin.Context) {
	_ = ctx.Param("consentId")

	ctx.Header("Content-Type", "application/xml")
	ctx.Data(http.StatusOK, "application/xml", []byte("<xml>File content</xml>"))
}

func (c *UKOpenBankingController) GetFilePayment(ctx *gin.Context) {
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

func (c *UKOpenBankingController) GetFilePaymentReportFile(ctx *gin.Context) {
	_ = ctx.Param("filePaymentId")

	ctx.Header("Content-Type", "application/xml")
	ctx.Data(http.StatusOK, "application/xml", []byte("<xml>Report content</xml>"))
}

func (c *UKOpenBankingController) GetAccount(ctx *gin.Context) {
	accountID := ctx.Param("AccountId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"Account": []map[string]interface{}{
				{
					"AccountId": accountID,
				},
			},
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) GetProducts(ctx *gin.Context) {
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

func (c *UKOpenBankingController) GetAccountProduct(ctx *gin.Context) {
	accountID := ctx.Param("AccountId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"Product": []map[string]interface{}{
				{
					"AccountId": accountID,
				},
			},
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) GetParty(ctx *gin.Context) {
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

func (c *UKOpenBankingController) GetAccountParty(ctx *gin.Context) {
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

func (c *UKOpenBankingController) GetAccountScheduledPayments(ctx *gin.Context) {
	_ = ctx.Param("AccountId")
	scheduledPayments := []map[string]interface{}{}

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"ScheduledPayment": scheduledPayments,
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) GetScheduledPayments(ctx *gin.Context) {
	scheduledPayments := []map[string]interface{}{}

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"ScheduledPayment": scheduledPayments,
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) CreateDomesticVRPConsents(ctx *gin.Context) {
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

func (c *UKOpenBankingController) CreateDomesticVRPs(ctx *gin.Context) {
	var vrpData map[string]interface{}
	if err := ctx.ShouldBindJSON(&vrpData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"Data":  vrpData,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusCreated, response)
}

func (c *UKOpenBankingController) GetDomesticVRPConsent(ctx *gin.Context) {
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

func (c *UKOpenBankingController) GetDomesticVRPConsentFundsConfirmation(ctx *gin.Context) {
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

func (c *UKOpenBankingController) DeleteDomesticVRPConsent(ctx *gin.Context) {
	_ = ctx.Param("consentId")
	ctx.JSON(http.StatusNoContent, nil)
}

func (c *UKOpenBankingController) GetDomesticVRP(ctx *gin.Context) {
	vrpID := ctx.Param("domesticVRPId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"DomesticVRPId": vrpID,
			"Status":        "AcceptedSettlementCompleted",
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) GetDomesticVRPDetails(ctx *gin.Context) {
	vrpID := ctx.Param("domesticVRPId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"DomesticVRPId":  vrpID,
			"PaymentDetails": map[string]interface{}{},
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) CreateCallbackUrl(ctx *gin.Context) {
	var callbackData map[string]interface{}
	if err := ctx.ShouldBindJSON(&callbackData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"Data":  callbackData,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusCreated, response)
}

func (c *UKOpenBankingController) UpdateCallbackUrl(ctx *gin.Context) {
	callbackUrlID := ctx.Param("callbackUrlId")

	var callbackData map[string]interface{}
	if err := ctx.ShouldBindJSON(&callbackData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	callbackData["callback_url_id"] = callbackUrlID
	response := map[string]interface{}{
		"Data":  callbackData,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) GetCallbackUrls(ctx *gin.Context) {
	callbackUrls := []map[string]interface{}{}

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"CallbackUrl": callbackUrls,
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) DeleteCallbackUrl(ctx *gin.Context) {
	_ = ctx.Param("callbackUrlId")
	ctx.JSON(http.StatusNoContent, nil)
}

func (c *UKOpenBankingController) CreateEventSubscription(ctx *gin.Context) {
	var subscriptionData map[string]interface{}
	if err := ctx.ShouldBindJSON(&subscriptionData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"Data":  subscriptionData,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusCreated, response)
}

func (c *UKOpenBankingController) UpdateEventSubscription(ctx *gin.Context) {
	eventSubscriptionID := ctx.Param("eventSubscriptionId")

	var subscriptionData map[string]interface{}
	if err := ctx.ShouldBindJSON(&subscriptionData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subscriptionData["event_subscription_id"] = eventSubscriptionID
	response := map[string]interface{}{
		"Data":  subscriptionData,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) GetEventSubscriptions(ctx *gin.Context) {
	subscriptions := []map[string]interface{}{}

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"EventSubscription": subscriptions,
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) DeleteEventSubscription(ctx *gin.Context) {
	_ = ctx.Param("eventSubscriptionId")
	ctx.JSON(http.StatusNoContent, nil)
}

func (c *UKOpenBankingController) CreateAggregatedPolling(ctx *gin.Context) {
	var pollingData map[string]interface{}
	if err := ctx.ShouldBindJSON(&pollingData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"Data":  pollingData,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusCreated, response)
}

func (c *UKOpenBankingController) GetAggregatedPolling(ctx *gin.Context) {
	pollingID := ctx.Param("pollingId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"PollingId": pollingID,
			"Status":    "Active",
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) DeleteAggregatedPolling(ctx *gin.Context) {
	_ = ctx.Param("pollingId")
	ctx.JSON(http.StatusNoContent, nil)
}

func (c *UKOpenBankingController) CreatePaymentOrderConsents(ctx *gin.Context) {
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

func (c *UKOpenBankingController) GetPaymentOrderConsent(ctx *gin.Context) {
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

func (c *UKOpenBankingController) DeletePaymentOrderConsent(ctx *gin.Context) {
	_ = ctx.Param("consentId")
	ctx.JSON(http.StatusNoContent, nil)
}

func (c *UKOpenBankingController) GetAccountTransactionsByStatementId(ctx *gin.Context) {
	_ = ctx.Param("AccountId")
	_ = ctx.Param("StatementId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"Transaction": []map[string]interface{}{},
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) GetTransactionsByStatementId(ctx *gin.Context) {
	_ = ctx.Param("StatementId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"Transaction": []map[string]interface{}{},
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) GetDomesticPaymentConsentFundsConfirmation(ctx *gin.Context) {
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

func (c *UKOpenBankingController) GetDomesticScheduledPaymentConsentFundsConfirmation(ctx *gin.Context) {
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

func (c *UKOpenBankingController) GetDomesticStandingOrderConsentFundsConfirmation(ctx *gin.Context) {
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

func (c *UKOpenBankingController) GetInternationalStandingOrderConsentFundsConfirmation(ctx *gin.Context) {
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

func (c *UKOpenBankingController) GetFilePaymentConsentFundsConfirmation(ctx *gin.Context) {
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
func (c *UKOpenBankingController) GetDomesticPayment(ctx *gin.Context) {
	domesticPaymentID := ctx.Param("DomesticPaymentId")

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"DomesticPaymentId": domesticPaymentID,
			"Status":            "AcceptedSettlementCompleted",
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) CreateEventNotification(ctx *gin.Context) {
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

func (c *UKOpenBankingController) GetEventNotifications(ctx *gin.Context) {
	notifications := []map[string]interface{}{}

	response := map[string]interface{}{
		"Data": map[string]interface{}{
			"EventNotification": notifications,
		},
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *UKOpenBankingController) CreateDomesticVRP(ctx *gin.Context) {
	var vrpData map[string]interface{}
	if err := ctx.ShouldBindJSON(&vrpData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"Data":  vrpData,
		"Meta":  map[string]interface{}{},
		"Links": map[string]interface{}{},
	}
	ctx.JSON(http.StatusCreated, response)
}
