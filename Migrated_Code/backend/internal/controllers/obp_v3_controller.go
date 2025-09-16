package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"obp-api-backend/internal/services"
	"obp-api-backend/internal/models"
)

type OBPv3Controller struct {
	bankService     services.BankService
	accountService  services.AccountService
	customerService services.CustomerService
}

func NewOBPv3Controller(
	bankService services.BankService,
	accountService services.AccountService,
	customerService services.CustomerService,
) *OBPv3Controller {
	return &OBPv3Controller{
		bankService:     bankService,
		accountService:  accountService,
		customerService: customerService,
	}
}

func (c *OBPv3Controller) GetAPIInfo(ctx *gin.Context) {
	apiInfo := map[string]interface{}{
		"version": "v3.1.0",
		"version_status": "STABLE",
		"git_commit": "unknown",
		"connector": "mapped",
	}
	ctx.JSON(http.StatusOK, apiInfo)
}

func (c *OBPv3Controller) GetConfig(ctx *gin.Context) {
	config := map[string]interface{}{
		"akka_ports": []string{"8080"},
		"elastic_search_enabled": false,
	}
	ctx.JSON(http.StatusOK, config)
}

func (c *OBPv3Controller) GetAdapterInfo(ctx *gin.Context) {
	adapterInfo := map[string]interface{}{
		"name": "OBP-API",
		"version": "v3.1.0",
	}
	ctx.JSON(http.StatusOK, adapterInfo)
}

func (c *OBPv3Controller) GetRateLimitingInfo(ctx *gin.Context) {
	rateLimiting := map[string]interface{}{
		"enabled": true,
		"technology": "REDIS",
	}
	ctx.JSON(http.StatusOK, rateLimiting)
}

func (c *OBPv3Controller) CreateAccountWebhook(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	_ = ctx.Param("viewId")
	
	var webhookData models.AccountWebhook
	if err := ctx.ShouldBindJSON(&webhookData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	webhookData.BankId = bankID
	webhookData.AccountId = accountID
	ctx.JSON(http.StatusCreated, webhookData)
}

func (c *OBPv3Controller) GetAccountWebhooks(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("accountId")
	_ = ctx.Param("viewId")
	
	webhooks := []models.AccountWebhook{}
	ctx.JSON(http.StatusOK, gin.H{"account_webhooks": webhooks})
}

func (c *OBPv3Controller) UpdateAccountWebhook(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	_ = ctx.Param("viewId")
	webhookID := ctx.Param("accountWebhookId")
	
	var webhookData models.AccountWebhook
	if err := ctx.ShouldBindJSON(&webhookData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	webhookData.AccountWebhookId = webhookID
	webhookData.BankId = bankID
	webhookData.AccountId = accountID
	ctx.JSON(http.StatusOK, webhookData)
}

func (c *OBPv3Controller) CreateProduct(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	var productData models.Product
	if err := ctx.ShouldBindJSON(&productData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	productData.BankId = bankID
	ctx.JSON(http.StatusCreated, productData)
}

func (c *OBPv3Controller) GetProducts(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	products := []models.Product{}
	ctx.JSON(http.StatusOK, gin.H{"products": products})
}

func (c *OBPv3Controller) GetProduct(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	productCode := ctx.Param("productCode")
	product := models.Product{
		ProductCode: productCode,
		BankId: bankID,
	}
	ctx.JSON(http.StatusOK, product)
}

func (c *OBPv3Controller) GetProductTree(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	productCode := ctx.Param("productCode")
	tree := map[string]interface{}{
		"product_code": productCode,
		"bank_id": bankID,
		"tree": []map[string]interface{}{},
	}
	ctx.JSON(http.StatusOK, tree)
}

func (c *OBPv3Controller) CreateProductAttribute(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	productCode := ctx.Param("productCode")
	var attrData models.ProductAttribute
	if err := ctx.ShouldBindJSON(&attrData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	attrData.BankId = bankID
	attrData.ProductCode = productCode
	ctx.JSON(http.StatusCreated, attrData)
}

func (c *OBPv3Controller) GetProductAttribute(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	productCode := ctx.Param("productCode")
	attrID := ctx.Param("productAttributeId")
	attr := models.ProductAttribute{
		ProductAttributeId: attrID,
		BankId: bankID,
		ProductCode: productCode,
	}
	ctx.JSON(http.StatusOK, attr)
}

func (c *OBPv3Controller) UpdateProductAttribute(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	productCode := ctx.Param("productCode")
	attrID := ctx.Param("productAttributeId")
	var attrData models.ProductAttribute
	if err := ctx.ShouldBindJSON(&attrData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	attrData.ProductAttributeId = attrID
	attrData.BankId = bankID
	attrData.ProductCode = productCode
	ctx.JSON(http.StatusOK, attrData)
}

func (c *OBPv3Controller) DeleteProductAttribute(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("productCode")
	_ = ctx.Param("productAttributeId")
	ctx.JSON(http.StatusOK, gin.H{"message": "Product attribute deleted"})
}

func (c *OBPv3Controller) CreateCustomerAttribute(ctx *gin.Context) {
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

func (c *OBPv3Controller) GetCustomerAttributes(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("customerId")
	attributes := []models.CustomerAttribute{}
	ctx.JSON(http.StatusOK, gin.H{"customer_attributes": attributes})
}

func (c *OBPv3Controller) UpdateCustomerAttribute(ctx *gin.Context) {
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

func (c *OBPv3Controller) DeleteCustomerAttribute(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("customerId")
	_ = ctx.Param("customerAttributeId")
	ctx.JSON(http.StatusOK, gin.H{"message": "Customer attribute deleted"})
}

func (c *OBPv3Controller) CreateMeeting(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	var meetingData models.Meeting
	if err := ctx.ShouldBindJSON(&meetingData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	meetingData.BankId = bankID
	ctx.JSON(http.StatusCreated, meetingData)
}

func (c *OBPv3Controller) GetMeetings(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	meetings := []models.Meeting{}
	ctx.JSON(http.StatusOK, gin.H{"meetings": meetings})
}

func (c *OBPv3Controller) GetMeeting(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	meetingID := ctx.Param("meetingId")
	meeting := models.Meeting{
		MeetingId: meetingID,
		BankId: bankID,
	}
	ctx.JSON(http.StatusOK, meeting)
}

func (c *OBPv3Controller) CreateCustomerAddress(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("customerId")
	var addressData map[string]interface{}
	if err := ctx.ShouldBindJSON(&addressData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, addressData)
}

func (c *OBPv3Controller) GetCustomerAddresses(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("customerId")
	addresses := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"customer_addresses": addresses})
}

func (c *OBPv3Controller) UpdateCustomerAddress(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("customerId")
	_ = ctx.Param("customerAddressId")
	var addressData map[string]interface{}
	if err := ctx.ShouldBindJSON(&addressData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, addressData)
}

func (c *OBPv3Controller) DeleteCustomerAddress(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("customerId")
	_ = ctx.Param("customerAddressId")
	ctx.JSON(http.StatusOK, gin.H{"message": "Customer address deleted"})
}

func (c *OBPv3Controller) CreateSystemView(ctx *gin.Context) {
	var viewData map[string]interface{}
	if err := ctx.ShouldBindJSON(&viewData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, viewData)
}

func (c *OBPv3Controller) GetSystemView(ctx *gin.Context) {
	viewID := ctx.Param("systemViewId")
	view := map[string]interface{}{
		"view_id": viewID,
		"name": "System View",
	}
	ctx.JSON(http.StatusOK, view)
}

func (c *OBPv3Controller) UpdateSystemView(ctx *gin.Context) {
	_ = ctx.Param("systemViewId")
	var viewData map[string]interface{}
	if err := ctx.ShouldBindJSON(&viewData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, viewData)
}

func (c *OBPv3Controller) DeleteSystemView(ctx *gin.Context) {
	_ = ctx.Param("systemViewId")
	ctx.JSON(http.StatusOK, gin.H{"message": "System view deleted"})
}

func (c *OBPv3Controller) GetCheckbookOrders(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("accountId")
	_ = ctx.Param("viewId")
	orders := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"checkbook_orders": orders})
}

func (c *OBPv3Controller) GetCreditCardOrders(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("accountId")
	_ = ctx.Param("viewId")
	orders := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"credit_card_orders": orders})
}

func (c *OBPv3Controller) GetTopAPIs(ctx *gin.Context) {
	apis := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"top_apis": apis})
}

func (c *OBPv3Controller) GetTopConsumers(ctx *gin.Context) {
	consumers := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"top_consumers": consumers})
}

func (c *OBPv3Controller) GetFirehoseCustomers(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	customers := []models.Customer{}
	ctx.JSON(http.StatusOK, gin.H{"customers": customers})
}

func (c *OBPv3Controller) GetBadLoginStatus(ctx *gin.Context) {
	username := ctx.Param("username")
	status := map[string]interface{}{
		"username": username,
		"bad_attempts": 0,
		"is_locked": false,
	}
	ctx.JSON(http.StatusOK, status)
}

func (c *OBPv3Controller) UnlockUser(ctx *gin.Context) {
	_ = ctx.Param("username")
	var unlockData map[string]interface{}
	if err := ctx.ShouldBindJSON(&unlockData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User unlocked"})
}

func (c *OBPv3Controller) SetCallsLimit(ctx *gin.Context) {
	_ = ctx.Param("consumerId")
	var limitData map[string]interface{}
	if err := ctx.ShouldBindJSON(&limitData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, limitData)
}

func (c *OBPv3Controller) GetCallsLimit(ctx *gin.Context) {
	consumerID := ctx.Param("consumerId")
	limit := map[string]interface{}{
		"consumer_id": consumerID,
		"per_second_call_limit": 1000,
		"per_minute_call_limit": 10000,
	}
	ctx.JSON(http.StatusOK, limit)
}

func (c *OBPv3Controller) CheckFundsAvailable(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("accountId")
	_ = ctx.Param("viewId")
	
	result := map[string]interface{}{
		"funds_available": "yes",
		"amount": "1000.00",
		"currency": "EUR",
	}
	ctx.JSON(http.StatusOK, result)
}

func (c *OBPv3Controller) GetConsumer(ctx *gin.Context) {
	consumerID := ctx.Param("consumerId")
	consumer := map[string]interface{}{
		"consumer_id": consumerID,
		"app_name": "Test App",
	}
	ctx.JSON(http.StatusOK, consumer)
}

func (c *OBPv3Controller) GetConsumersForCurrentUser(ctx *gin.Context) {
	consumers := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"consumers": consumers})
}

func (c *OBPv3Controller) GetConsumers(ctx *gin.Context) {
	consumers := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"consumers": consumers})
}

func (c *OBPv3Controller) GetConnectorLoopback(ctx *gin.Context) {
	loopback := map[string]interface{}{
		"connector": "mapped",
		"duration_ms": 1,
	}
	ctx.JSON(http.StatusOK, loopback)
}

func (c *OBPv3Controller) GetTransactionByIdForBankAccount(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("accountId")
	_ = ctx.Param("viewId")
	transactionID := ctx.Param("transactionId")
	
	transaction := models.Transaction{
		Id: transactionID,
		ThisAccount: ctx.Param("accountId"),
	}
	ctx.JSON(http.StatusOK, transaction)
}

func (c *OBPv3Controller) GetTransactionRequests(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("accountId")
	_ = ctx.Param("viewId")
	
	requests := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"transaction_requests": requests})
}

func (c *OBPv3Controller) GetCustomerByCustomerId(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	customerID := ctx.Param("customerId")
	
	customer := models.Customer{
		CustomerId: customerID,
		BankId: bankID,
	}
	ctx.JSON(http.StatusOK, customer)
}

func (c *OBPv3Controller) GetCustomerByCustomerNumber(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	customerNumber := ctx.Param("customerNumber")
	
	customer := models.Customer{
		Number: customerNumber,
		BankId: bankID,
	}
	ctx.JSON(http.StatusOK, customer)
}

func (c *OBPv3Controller) DeleteAccountWebhook(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("accountId")
	_ = ctx.Param("viewId")
	_ = ctx.Param("webhookId")
	ctx.JSON(http.StatusNoContent, nil)
}

func (c *OBPv3Controller) UpdateProduct(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	productCode := ctx.Param("productCode")
	var productData map[string]interface{}
	if err := ctx.ShouldBindJSON(&productData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	productData["bank_id"] = bankID
	productData["product_code"] = productCode
	ctx.JSON(http.StatusOK, productData)
}

func (c *OBPv3Controller) DeleteProduct(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("productCode")
	ctx.JSON(http.StatusNoContent, nil)
}

func (c *OBPv3Controller) UpdateMeeting(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	meetingID := ctx.Param("meetingId")
	var meetingData map[string]interface{}
	if err := ctx.ShouldBindJSON(&meetingData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	meetingData["bank_id"] = bankID
	meetingData["meeting_id"] = meetingID
	ctx.JSON(http.StatusOK, meetingData)
}

func (c *OBPv3Controller) DeleteMeeting(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("meetingId")
	ctx.JSON(http.StatusNoContent, nil)
}

func (c *OBPv3Controller) CreateTransactionComment(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	viewID := ctx.Param("viewId")
	var commentData map[string]interface{}
	if err := ctx.ShouldBindJSON(&commentData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	commentData["bank_id"] = bankID
	commentData["account_id"] = accountID
	commentData["view_id"] = viewID
	ctx.JSON(http.StatusCreated, commentData)
}

func (c *OBPv3Controller) GetTransactionComments(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("accountId")
	_ = ctx.Param("viewId")
	comments := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"transaction_comments": comments})
}

func (c *OBPv3Controller) DeleteTransactionComment(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("accountId")
	_ = ctx.Param("viewId")
	_ = ctx.Param("metadataId")
	ctx.JSON(http.StatusNoContent, nil)
}

func (c *OBPv3Controller) CreateOtherAccountMetadata(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	viewID := ctx.Param("viewId")
	var metadataData map[string]interface{}
	if err := ctx.ShouldBindJSON(&metadataData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	metadataData["bank_id"] = bankID
	metadataData["account_id"] = accountID
	metadataData["view_id"] = viewID
	ctx.JSON(http.StatusCreated, metadataData)
}

func (c *OBPv3Controller) GetOtherAccountMetadata(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("accountId")
	_ = ctx.Param("viewId")
	metadata := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"other_account_metadata": metadata})
}

func (c *OBPv3Controller) UpdateOtherAccountMetadata(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	viewID := ctx.Param("viewId")
	metadataID := ctx.Param("metadataId")
	var metadataData map[string]interface{}
	if err := ctx.ShouldBindJSON(&metadataData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	metadataData["bank_id"] = bankID
	metadataData["account_id"] = accountID
	metadataData["view_id"] = viewID
	metadataData["metadata_id"] = metadataID
	ctx.JSON(http.StatusOK, metadataData)
}

func (c *OBPv3Controller) DeleteOtherAccountMetadata(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("accountId")
	_ = ctx.Param("viewId")
	_ = ctx.Param("metadataId")
	ctx.JSON(http.StatusNoContent, nil)
}

func (c *OBPv3Controller) CreateTransactionTag(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	viewID := ctx.Param("viewId")
	transactionID := ctx.Param("transactionId")
	var tagData map[string]interface{}
	if err := ctx.ShouldBindJSON(&tagData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tagData["bank_id"] = bankID
	tagData["account_id"] = accountID
	tagData["view_id"] = viewID
	tagData["transaction_id"] = transactionID
	ctx.JSON(http.StatusCreated, tagData)
}

func (c *OBPv3Controller) GetTransactionTags(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("accountId")
	_ = ctx.Param("viewId")
	_ = ctx.Param("transactionId")
	tags := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"transaction_tags": tags})
}

func (c *OBPv3Controller) DeleteTransactionTag(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("accountId")
	_ = ctx.Param("viewId")
	_ = ctx.Param("transactionId")
	_ = ctx.Param("tagId")
	ctx.JSON(http.StatusNoContent, nil)
}

func (c *OBPv3Controller) CreateTransactionImage(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	viewID := ctx.Param("viewId")
	transactionID := ctx.Param("transactionId")
	var imageData map[string]interface{}
	if err := ctx.ShouldBindJSON(&imageData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	imageData["bank_id"] = bankID
	imageData["account_id"] = accountID
	imageData["view_id"] = viewID
	imageData["transaction_id"] = transactionID
	ctx.JSON(http.StatusCreated, imageData)
}

func (c *OBPv3Controller) GetTransactionImages(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("accountId")
	_ = ctx.Param("viewId")
	_ = ctx.Param("transactionId")
	images := []map[string]interface{}{}
	ctx.JSON(http.StatusOK, gin.H{"transaction_images": images})
}

func (c *OBPv3Controller) DeleteTransactionImage(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("accountId")
	_ = ctx.Param("viewId")
	_ = ctx.Param("transactionId")
	_ = ctx.Param("imageId")
	ctx.JSON(http.StatusNoContent, nil)
}

func (c *OBPv3Controller) CreateTransactionWhere(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	viewID := ctx.Param("viewId")
	transactionID := ctx.Param("transactionId")
	var whereData map[string]interface{}
	if err := ctx.ShouldBindJSON(&whereData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	whereData["bank_id"] = bankID
	whereData["account_id"] = accountID
	whereData["view_id"] = viewID
	whereData["transaction_id"] = transactionID
	ctx.JSON(http.StatusCreated, whereData)
}

func (c *OBPv3Controller) GetTransactionWhere(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("accountId")
	_ = ctx.Param("viewId")
	_ = ctx.Param("transactionId")
	where := map[string]interface{}{}
	ctx.JSON(http.StatusOK, where)
}

func (c *OBPv3Controller) UpdateTransactionWhere(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	viewID := ctx.Param("viewId")
	transactionID := ctx.Param("transactionId")
	var whereData map[string]interface{}
	if err := ctx.ShouldBindJSON(&whereData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	whereData["bank_id"] = bankID
	whereData["account_id"] = accountID
	whereData["view_id"] = viewID
	whereData["transaction_id"] = transactionID
	ctx.JSON(http.StatusOK, whereData)
}

func (c *OBPv3Controller) DeleteTransactionWhere(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("accountId")
	_ = ctx.Param("viewId")
	_ = ctx.Param("transactionId")
	ctx.JSON(http.StatusNoContent, nil)
}

func (c *OBPv3Controller) CreateOtherAccountMoreInfo(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	viewID := ctx.Param("viewId")
	otherAccountID := ctx.Param("otherAccountId")
	var moreInfoData map[string]interface{}
	if err := ctx.ShouldBindJSON(&moreInfoData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	moreInfoData["bank_id"] = bankID
	moreInfoData["account_id"] = accountID
	moreInfoData["view_id"] = viewID
	moreInfoData["other_account_id"] = otherAccountID
	ctx.JSON(http.StatusCreated, moreInfoData)
}

func (c *OBPv3Controller) GetOtherAccountMoreInfo(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("accountId")
	_ = ctx.Param("viewId")
	_ = ctx.Param("otherAccountId")
	moreInfo := map[string]interface{}{}
	ctx.JSON(http.StatusOK, moreInfo)
}

func (c *OBPv3Controller) UpdateOtherAccountMoreInfo(ctx *gin.Context) {
	bankID := ctx.Param("bankId")
	accountID := ctx.Param("accountId")
	viewID := ctx.Param("viewId")
	otherAccountID := ctx.Param("otherAccountId")
	var moreInfoData map[string]interface{}
	if err := ctx.ShouldBindJSON(&moreInfoData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	moreInfoData["bank_id"] = bankID
	moreInfoData["account_id"] = accountID
	moreInfoData["view_id"] = viewID
	moreInfoData["other_account_id"] = otherAccountID
	ctx.JSON(http.StatusOK, moreInfoData)
}

func (c *OBPv3Controller) DeleteOtherAccountMoreInfo(ctx *gin.Context) {
	_ = ctx.Param("bankId")
	_ = ctx.Param("accountId")
	_ = ctx.Param("viewId")
	_ = ctx.Param("otherAccountId")
	ctx.JSON(http.StatusNoContent, nil)
}
