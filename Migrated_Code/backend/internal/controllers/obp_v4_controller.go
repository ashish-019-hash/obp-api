package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"obp-api-backend/internal/services"
	"obp-api-backend/internal/models"
)

type OBPv4Controller struct {
	bankService        services.BankService
	accountService     services.AccountService
	transactionService services.TransactionService
	customerService    services.CustomerService
}

func NewOBPv4Controller(
	bankService services.BankService,
	accountService services.AccountService,
	transactionService services.TransactionService,
	customerService services.CustomerService,
) *OBPv4Controller {
	return &OBPv4Controller{
		bankService:        bankService,
		accountService:     accountService,
		transactionService: transactionService,
		customerService:    customerService,
	}
}

func (c *OBPv4Controller) GetAPIInfo(ctx *gin.Context) {
	apiInfo := map[string]interface{}{
		"version": "v4.0.0",
		"version_status": "STABLE",
		"git_commit": "unknown",
		"connector": "mapped",
	}
	ctx.JSON(http.StatusOK, apiInfo)
}

func (c *OBPv4Controller) GetDatabaseInfo(ctx *gin.Context) {
	dbInfo := map[string]interface{}{
		"database": "SQLite",
		"version": "3.x",
	}
	ctx.JSON(http.StatusOK, dbInfo)
}

func (c *OBPv4Controller) GetLogoutLink(ctx *gin.Context) {
	logoutLink := map[string]interface{}{
		"logout_link": "/logout",
	}
	ctx.JSON(http.StatusOK, logoutLink)
}

func (c *OBPv4Controller) CreateUserWithRoles(ctx *gin.Context) {
	var userData map[string]interface{}
	if err := ctx.ShouldBindJSON(&userData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, userData)
}

func (c *OBPv4Controller) GetEntitlements(ctx *gin.Context) {
	_ = ctx.Param("userId")
	entitlements := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"entitlements": entitlements})
}

func (c *OBPv4Controller) GetEntitlementsForBank(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	entitlements := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"entitlements": entitlements})
}

func (c *OBPv4Controller) LockUser(ctx *gin.Context) {
	_ = ctx.Param("userId")
	var lockData map[string]interface{}
	if err := ctx.ShouldBindJSON(&lockData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User locked"})
}

func (c *OBPv4Controller) GetSystemDynamicEntities(ctx *gin.Context) {
	entities := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"dynamic_entities": entities})
}

func (c *OBPv4Controller) CreateSystemDynamicEntity(ctx *gin.Context) {
	var entityData map[string]interface{}
	if err := ctx.ShouldBindJSON(&entityData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, entityData)
}

func (c *OBPv4Controller) UpdateSystemDynamicEntity(ctx *gin.Context) {
	_ = ctx.Param("dynamicEntityId")
	var entityData map[string]interface{}
	if err := ctx.ShouldBindJSON(&entityData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, entityData)
}

func (c *OBPv4Controller) DeleteSystemDynamicEntity(ctx *gin.Context) {
	_ = ctx.Param("dynamicEntityId")
	ctx.JSON(http.StatusOK, gin.H{"message": "System dynamic entity deleted"})
}

func (c *OBPv4Controller) GetBankDynamicEntities(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	entities := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"dynamic_entities": entities})
}

func (c *OBPv4Controller) CreateBankDynamicEntity(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	var entityData map[string]interface{}
	if err := ctx.ShouldBindJSON(&entityData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, entityData)
}

func (c *OBPv4Controller) UpdateBankDynamicEntity(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("dynamicEntityId")
	var entityData map[string]interface{}
	if err := ctx.ShouldBindJSON(&entityData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, entityData)
}

func (c *OBPv4Controller) DeleteBankDynamicEntity(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("dynamicEntityId")
	ctx.JSON(http.StatusOK, gin.H{"message": "Bank dynamic entity deleted"})
}

func (c *OBPv4Controller) GetMyDynamicEntities(ctx *gin.Context) {
	entities := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"dynamic_entities": entities})
}

func (c *OBPv4Controller) UpdateMyDynamicEntity(ctx *gin.Context) {
	_ = ctx.Param("dynamicEntityId")
	var entityData map[string]interface{}
	if err := ctx.ShouldBindJSON(&entityData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, entityData)
}

func (c *OBPv4Controller) DeleteMyDynamicEntity(ctx *gin.Context) {
	_ = ctx.Param("dynamicEntityId")
	ctx.JSON(http.StatusOK, gin.H{"message": "My dynamic entity deleted"})
}

func (c *OBPv4Controller) CreateAccountTransactionRequest(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("accountId")
	_ = ctx.Param("viewId")
	
	var requestData map[string]interface{}
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusCreated, requestData)
}

func (c *OBPv4Controller) CreateAccountOTPTransactionRequest(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("accountId")
	_ = ctx.Param("viewId")
	
	var requestData map[string]interface{}
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusCreated, requestData)
}

func (c *OBPv4Controller) CreateCounterpartyTransactionRequest(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("accountId")
	_ = ctx.Param("viewId")
	
	var requestData map[string]interface{}
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusCreated, requestData)
}

func (c *OBPv4Controller) CreateSimpleTransactionRequest(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("accountId")
	_ = ctx.Param("viewId")
	
	var requestData map[string]interface{}
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusCreated, requestData)
}

func (c *OBPv4Controller) CreateSEPATransactionRequest(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("accountId")
	_ = ctx.Param("viewId")
	
	var requestData map[string]interface{}
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusCreated, requestData)
}

func (c *OBPv4Controller) CreateRefundTransactionRequest(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("accountId")
	_ = ctx.Param("viewId")
	
	var requestData map[string]interface{}
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusCreated, requestData)
}

func (c *OBPv4Controller) CreateFreeFormTransactionRequest(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("accountId")
	_ = ctx.Param("viewId")
	
	var requestData map[string]interface{}
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusCreated, requestData)
}

func (c *OBPv4Controller) CreateAgentCashWithdrawalTransactionRequest(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("accountId")
	_ = ctx.Param("viewId")
	
	var requestData map[string]interface{}
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusCreated, requestData)
}

func (c *OBPv4Controller) CreateCardTransactionRequest(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("accountId")
	_ = ctx.Param("viewId")
	
	var requestData map[string]interface{}
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusCreated, requestData)
}

func (c *OBPv4Controller) AnswerTransactionRequestChallenge(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("accountId")
	_ = ctx.Param("viewId")
	_ = ctx.Param("transactionRequestType")
	_ = ctx.Param("transactionRequestId")
	
	var challengeData map[string]interface{}
	if err := ctx.ShouldBindJSON(&challengeData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusOK, challengeData)
}

func (c *OBPv4Controller) CreateSettlementAccount(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	var accountData map[string]interface{}
	if err := ctx.ShouldBindJSON(&accountData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, accountData)
}

func (c *OBPv4Controller) GetSettlementAccounts(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	accounts := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"settlement_accounts": accounts})
}

func (c *OBPv4Controller) GetDoubleEntryTransaction(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("accountId")
	_ = ctx.Param("viewId")
	transactionID := ctx.Param("transactionId")
	
	transaction := map[string]interface{}{
		"transaction_id": transactionID,
		"double_entry": true,
	}
	ctx.JSON(http.StatusOK, transaction)
}

func (c *OBPv4Controller) GetBalancingTransaction(ctx *gin.Context) {
	transactionID := ctx.Param("transactionId")
	
	transaction := map[string]interface{}{
		"transaction_id": transactionID,
		"balancing": true,
	}
	ctx.JSON(http.StatusOK, transaction)
}

func (c *OBPv4Controller) IBANChecker(ctx *gin.Context) {
	var ibanData map[string]interface{}
	if err := ctx.ShouldBindJSON(&ibanData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	result := map[string]interface{}{
		"is_valid": true,
		"iban": ibanData["iban"],
	}
	ctx.JSON(http.StatusOK, result)
}

func (c *OBPv4Controller) GetCallContext(ctx *gin.Context) {
	callContext := map[string]interface{}{
		"correlation_id": "12345",
		"session_id": "session_12345",
	}
	ctx.JSON(http.StatusOK, callContext)
}

func (c *OBPv4Controller) VerifyRequestSignResponse(ctx *gin.Context) {
	var signData map[string]interface{}
	if err := ctx.ShouldBindJSON(&signData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	result := map[string]interface{}{
		"is_valid": true,
	}
	ctx.JSON(http.StatusOK, result)
}

func (c *OBPv4Controller) AddAccount(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	var accountData models.BankAccount
	if err := ctx.ShouldBindJSON(&accountData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	accountData.BankId = bankID
	ctx.JSON(http.StatusCreated, accountData)
}

func (c *OBPv4Controller) UpdateAccountLabel(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("accountId")
	var labelData map[string]interface{}
	if err := ctx.ShouldBindJSON(&labelData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, labelData)
}

func (c *OBPv4Controller) CreateCustomerAttribute(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	customerID := ctx.Param("customerId")
	var attrData models.CustomerAttribute
	if err := ctx.ShouldBindJSON(&attrData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	attrData.BankId = bankID
	attrData.CustomerId = customerID
	ctx.JSON(http.StatusCreated, attrData)
}

func (c *OBPv4Controller) UpdateCustomerAttribute(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	customerID := ctx.Param("customerId")
	attrID := ctx.Param("customerAttributeId")
	var attrData models.CustomerAttribute
	if err := ctx.ShouldBindJSON(&attrData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	attrData.CustomerAttributeId = attrID
	attrData.BankId = bankID
	attrData.CustomerId = customerID
	ctx.JSON(http.StatusOK, attrData)
}

func (c *OBPv4Controller) GetCustomerAttributes(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("customerId")
	attributes := []models.CustomerAttribute{}
	ctx.JSON(http.StatusOK, gin.H{"customer_attributes": attributes})
}

func (c *OBPv4Controller) GetCustomerAttributeById(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	customerID := ctx.Param("customerId")
	attrID := ctx.Param("customerAttributeId")
	attr := models.CustomerAttribute{
		CustomerAttributeId: attrID,
		BankId: bankID,
		CustomerId: customerID,
	}
	ctx.JSON(http.StatusOK, attr)
}

func (c *OBPv4Controller) GetCustomersByAttributes(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	customers := []models.Customer{}
	ctx.JSON(http.StatusOK, gin.H{"customers": customers})
}

func (c *OBPv4Controller) CreateOrUpdateTransactionRequestAttributeDefinition(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	attributeDefinitionID := ctx.Param("attributeDefinitionId")
	
	var attrDefData map[string]interface{}
	if err := ctx.ShouldBindJSON(&attrDefData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	attrDefData["bank_id"] = bankID
	attrDefData["attribute_definition_id"] = attributeDefinitionID
	ctx.JSON(http.StatusOK, attrDefData)
}

func (c *OBPv4Controller) GetTransactionRequestAttributeDefinition(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	attributeDefinitionID := ctx.Param("attributeDefinitionId")
	
	attrDef := map[string]interface{}{
		"bank_id": bankID,
		"attribute_definition_id": attributeDefinitionID,
		"name": "attribute_definition_name",
	}
	ctx.JSON(http.StatusOK, attrDef)
}

func (c *OBPv4Controller) DeleteTransactionRequestAttributeDefinition(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	attributeDefinitionID := ctx.Param("attributeDefinitionId")
	
	ctx.JSON(http.StatusOK, gin.H{"message": "Transaction request attribute definition deleted", "bank_id": bankID, "attribute_definition_id": attributeDefinitionID})
}

func (c *OBPv4Controller) CreateResetPasswordUrl(ctx *gin.Context) {
	var resetData map[string]interface{}
	if err := ctx.ShouldBindJSON(&resetData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, resetData)
}

func (c *OBPv4Controller) GetCurrentUserId(ctx *gin.Context) {
	user := map[string]interface{}{
		"user_id": "current_user_123",
		"username": "current_user",
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *OBPv4Controller) GetUserByUserId(ctx *gin.Context) {
	userID := ctx.Param("userId")
	user := map[string]interface{}{
		"user_id": userID,
		"username": "user_" + userID,
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *OBPv4Controller) GetUserByUsername(ctx *gin.Context) {
	username := ctx.Param("username")
	user := map[string]interface{}{
		"username": username,
		"user_id": "user_123",
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *OBPv4Controller) GetUsersByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	users := []map[string]interface{}{
		{
			"email": email,
			"user_id": "user_123",
		},
	}
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func (c *OBPv4Controller) GetUsers(ctx *gin.Context) {
	users := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func (c *OBPv4Controller) CreateUserInvitation(ctx *gin.Context) {
	var invitationData map[string]interface{}
	if err := ctx.ShouldBindJSON(&invitationData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, invitationData)
}

func (c *OBPv4Controller) GetUserInvitationAnonymous(ctx *gin.Context) {
	secretLink := ctx.Param("secretLink")
	invitation := map[string]interface{}{
		"secret_link": secretLink,
		"status": "pending",
	}
	ctx.JSON(http.StatusOK, invitation)
}

func (c *OBPv4Controller) GetUserInvitation(ctx *gin.Context) {
	invitationID := ctx.Param("userInvitationId")
	invitation := map[string]interface{}{
		"invitation_id": invitationID,
		"status": "pending",
	}
	ctx.JSON(http.StatusOK, invitation)
}

func (c *OBPv4Controller) GetUserInvitations(ctx *gin.Context) {
	invitations := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"user_invitations": invitations})
}

func (c *OBPv4Controller) DeleteUser(ctx *gin.Context) {
	userID := ctx.Param("userId")
	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted", "user_id": userID})
}

func (c *OBPv4Controller) GetBanks(ctx *gin.Context) {
	banks := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"banks": banks})
}

func (c *OBPv4Controller) GetBank(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	bank := map[string]interface{}{
		"bank_id": bankID,
		"full_name": "Bank " + bankID,
	}
	ctx.JSON(http.StatusOK, bank)
}

func (c *OBPv4Controller) CreateBank(ctx *gin.Context) {
	var bankData map[string]interface{}
	if err := ctx.ShouldBindJSON(&bankData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, bankData)
}

func (c *OBPv4Controller) GetAccountByIdCore(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	
	account := map[string]interface{}{
		"bank_id": bankID,
		"account_id": accountID,
		"label": "Account " + accountID,
	}
	ctx.JSON(http.StatusOK, account)
}

func (c *OBPv4Controller) GetAccountByIdFull(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	viewID := ctx.Param("viewId")
	
	account := map[string]interface{}{
		"bank_id": bankID,
		"account_id": accountID,
		"view_id": viewID,
		"label": "Account " + accountID,
	}
	ctx.JSON(http.StatusOK, account)
}

func (c *OBPv4Controller) GetAccountByAccountRouting(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	scheme := ctx.Param("scheme")
	address := ctx.Param("address")
	
	account := map[string]interface{}{
		"bank_id": bankID,
		"scheme": scheme,
		"address": address,
	}
	ctx.JSON(http.StatusOK, account)
}

func (c *OBPv4Controller) GetAccountsByAccountRoutingRegex(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountRoutingRegex := ctx.Param("accountRoutingRegex")
	
	accounts := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"accounts": accounts, "bank_id": bankID, "regex": accountRoutingRegex})
}

func (c *OBPv4Controller) GetBankAccountsBalancesForCurrentUser(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	balances := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"balances": balances, "bank_id": bankID})
}

func (c *OBPv4Controller) GetBankAccountBalancesForCurrentUser(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	
	balances := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"balances": balances, "bank_id": bankID, "account_id": accountID})
}

func (c *OBPv4Controller) GetFirehoseAccountsAtOneBank(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	viewID := ctx.Param("viewId")
	
	accounts := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"accounts": accounts, "bank_id": bankID, "view_id": viewID})
}

func (c *OBPv4Controller) GetFastFirehoseAccountsAtOneBank(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	viewID := ctx.Param("viewId")
	
	accounts := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"accounts": accounts, "bank_id": bankID, "view_id": viewID})
}

func (c *OBPv4Controller) GetCustomersByCustomerPhoneNumber(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	mobilePhoneNumber := ctx.Param("mobilePhoneNumber")
	
	customers := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"customers": customers, "bank_id": bankID, "mobile_phone_number": mobilePhoneNumber})
}

func (c *OBPv4Controller) CreateDirectDebitManagement(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	
	var debitData map[string]interface{}
	if err := ctx.ShouldBindJSON(&debitData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	debitData["bank_id"] = bankID
	debitData["account_id"] = accountID
	ctx.JSON(http.StatusCreated, debitData)
}

func (c *OBPv4Controller) CreateStandingOrderManagement(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	
	var orderData map[string]interface{}
	if err := ctx.ShouldBindJSON(&orderData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	orderData["bank_id"] = bankID
	orderData["account_id"] = accountID
	ctx.JSON(http.StatusCreated, orderData)
}

func (c *OBPv4Controller) RevokeGrantUserAccessToViews(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	viewID := ctx.Param("viewId")
	
	var revokeData map[string]interface{}
	if err := ctx.ShouldBindJSON(&revokeData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	revokeData["bank_id"] = bankID
	revokeData["account_id"] = accountID
	revokeData["view_id"] = viewID
	ctx.JSON(http.StatusOK, revokeData)
}

func (c *OBPv4Controller) AddTagForViewOnAccount(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	viewID := ctx.Param("viewId")
	
	var tagData map[string]interface{}
	if err := ctx.ShouldBindJSON(&tagData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	tagData["bank_id"] = bankID
	tagData["account_id"] = accountID
	tagData["view_id"] = viewID
	ctx.JSON(http.StatusCreated, tagData)
}

func (c *OBPv4Controller) DeleteTagForViewOnAccount(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	viewID := ctx.Param("viewId")
	tagID := ctx.Param("tagId")
	
	ctx.JSON(http.StatusOK, gin.H{"message": "Tag deleted", "bank_id": bankID, "account_id": accountID, "view_id": viewID, "tag_id": tagID})
}

func (c *OBPv4Controller) GetTagsForViewOnAccount(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	viewID := ctx.Param("viewId")
	
	tags := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"tags": tags, "bank_id": bankID, "account_id": accountID, "view_id": viewID})
}

func (c *OBPv4Controller) GetAccountTransactionRequest(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	transactionRequestID := ctx.Param("transactionRequestId")
	
	transactionRequest := map[string]interface{}{
		"id": transactionRequestID,
		"type": "ACCOUNT",
		"from": map[string]interface{}{
			"bank_id": bankID,
			"account_id": accountID,
		},
		"status": "INITIATED",
	}
	ctx.JSON(http.StatusOK, transactionRequest)
}

func (c *OBPv4Controller) GetAccountOTPTransactionRequest(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	transactionRequestID := ctx.Param("transactionRequestId")
	
	transactionRequest := map[string]interface{}{
		"id": transactionRequestID,
		"type": "ACCOUNT_OTP",
		"from": map[string]interface{}{
			"bank_id": bankID,
			"account_id": accountID,
		},
		"status": "INITIATED",
	}
	ctx.JSON(http.StatusOK, transactionRequest)
}

func (c *OBPv4Controller) GetCounterpartyTransactionRequest(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	transactionRequestID := ctx.Param("transactionRequestId")
	
	transactionRequest := map[string]interface{}{
		"id": transactionRequestID,
		"type": "COUNTERPARTY",
		"from": map[string]interface{}{
			"bank_id": bankID,
			"account_id": accountID,
		},
		"status": "INITIATED",
	}
	ctx.JSON(http.StatusOK, transactionRequest)
}

func (c *OBPv4Controller) GetSEPATransactionRequest(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	transactionRequestID := ctx.Param("transactionRequestId")
	
	transactionRequest := map[string]interface{}{
		"id": transactionRequestID,
		"type": "SEPA",
		"from": map[string]interface{}{
			"bank_id": bankID,
			"account_id": accountID,
		},
		"status": "INITIATED",
	}
	ctx.JSON(http.StatusOK, transactionRequest)
}

func (c *OBPv4Controller) GetSimpleTransactionRequest(ctx *gin.Context) {
	transactionRequestID := ctx.Param("transactionRequestId")
	
	transactionRequest := map[string]interface{}{
		"id": transactionRequestID,
		"type": "SIMPLE",
		"status": "INITIATED",
	}
	ctx.JSON(http.StatusOK, transactionRequest)
}

func (c *OBPv4Controller) GetFreeFormTransactionRequest(ctx *gin.Context) {
	transactionRequestID := ctx.Param("transactionRequestId")
	
	transactionRequest := map[string]interface{}{
		"id": transactionRequestID,
		"type": "FREE_FORM",
		"status": "INITIATED",
	}
	ctx.JSON(http.StatusOK, transactionRequest)
}

func (c *OBPv4Controller) AnswerAccountTransactionRequestChallenge(ctx *gin.Context) {
	transactionRequestID := ctx.Param("transactionRequestId")
	
	var challengeAnswer map[string]interface{}
	if err := ctx.ShouldBindJSON(&challengeAnswer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	response := map[string]interface{}{
		"transaction_request_id": transactionRequestID,
		"challenge_answer": challengeAnswer,
		"status": "COMPLETED",
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *OBPv4Controller) AnswerAccountOTPTransactionRequestChallenge(ctx *gin.Context) {
	transactionRequestID := ctx.Param("transactionRequestId")
	
	var challengeAnswer map[string]interface{}
	if err := ctx.ShouldBindJSON(&challengeAnswer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	response := map[string]interface{}{
		"transaction_request_id": transactionRequestID,
		"challenge_answer": challengeAnswer,
		"status": "COMPLETED",
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *OBPv4Controller) AnswerCounterpartyTransactionRequestChallenge(ctx *gin.Context) {
	transactionRequestID := ctx.Param("transactionRequestId")
	
	var challengeAnswer map[string]interface{}
	if err := ctx.ShouldBindJSON(&challengeAnswer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	response := map[string]interface{}{
		"transaction_request_id": transactionRequestID,
		"challenge_answer": challengeAnswer,
		"status": "COMPLETED",
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *OBPv4Controller) AnswerSEPATransactionRequestChallenge(ctx *gin.Context) {
	transactionRequestID := ctx.Param("transactionRequestId")
	
	var challengeAnswer map[string]interface{}
	if err := ctx.ShouldBindJSON(&challengeAnswer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	response := map[string]interface{}{
		"transaction_request_id": transactionRequestID,
		"challenge_answer": challengeAnswer,
		"status": "COMPLETED",
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *OBPv4Controller) AnswerSimpleTransactionRequestChallenge(ctx *gin.Context) {
	transactionRequestID := ctx.Param("transactionRequestId")
	
	var challengeAnswer map[string]interface{}
	if err := ctx.ShouldBindJSON(&challengeAnswer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	response := map[string]interface{}{
		"transaction_request_id": transactionRequestID,
		"challenge_answer": challengeAnswer,
		"status": "COMPLETED",
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *OBPv4Controller) AnswerFreeFormTransactionRequestChallenge(ctx *gin.Context) {
	transactionRequestID := ctx.Param("transactionRequestId")
	
	var challengeAnswer map[string]interface{}
	if err := ctx.ShouldBindJSON(&challengeAnswer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	response := map[string]interface{}{
		"transaction_request_id": transactionRequestID,
		"challenge_answer": challengeAnswer,
		"status": "COMPLETED",
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *OBPv4Controller) GetTransactionRequestTypes(ctx *gin.Context) {
	transactionRequestTypes := []map[string]interface{}{
		{"value": "ACCOUNT"},
		{"value": "ACCOUNT_OTP"},
		{"value": "COUNTERPARTY"},
		{"value": "SEPA"},
		{"value": "SIMPLE"},
		{"value": "FREE_FORM"},
	}
	ctx.JSON(http.StatusOK, gin.H{"transaction_request_types": transactionRequestTypes})
}

func (c *OBPv4Controller) GetTransactionRequests(ctx *gin.Context) {
	transactionRequests := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"transaction_requests": transactionRequests})
}

func (c *OBPv4Controller) CreateTransactionRequestAttributeDefinition(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	var attributeData map[string]interface{}
	if err := ctx.ShouldBindJSON(&attributeData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	attributeData["bank_id"] = bankID
	ctx.JSON(http.StatusCreated, attributeData)
}

func (c *OBPv4Controller) GetTransactionRequestAttributeDefinitions(ctx *gin.Context) {
	definitions := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"transaction_request_attribute_definitions": definitions})
}

func (c *OBPv4Controller) UpdateTransactionRequestAttributeDefinition(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	attributeDefinitionID := ctx.Param("attributeDefinitionId")
	
	var attributeData map[string]interface{}
	if err := ctx.ShouldBindJSON(&attributeData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	attributeData["bank_id"] = bankID
	attributeData["attribute_definition_id"] = attributeDefinitionID
	ctx.JSON(http.StatusOK, attributeData)
}
