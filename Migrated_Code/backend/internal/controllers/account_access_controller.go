package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/ashish-019-hash/obp-api-backend/internal/utils"
)

type AccountAccessController struct {
	orchestrationService *services.OrchestrationService
}

func NewAccountAccessController(orchestrationService *services.OrchestrationService) *AccountAccessController {
	return &AccountAccessController{
		orchestrationService: orchestrationService,
	}
}

type GrantAccessRequest struct {
	UserID string `json:"user_id" binding:"required"`
	ViewID string `json:"view_id" binding:"required"`
}

type RevokeAccessRequest struct {
	UserID string `json:"user_id" binding:"required"`
	ViewID string `json:"view_id" binding:"required"`
}

type CreateUserAccountAccessRequest struct {
	Provider string `json:"provider" binding:"required"`
	Username string `json:"username" binding:"required"`
	ViewID   string `json:"view_id" binding:"required"`
}

type ViewResponse struct {
	ID                    string   `json:"id"`
	ShortName             string   `json:"short_name"`
	Description           string   `json:"description"`
	IsPublic              bool     `json:"is_public"`
	Alias                 string   `json:"alias"`
	HideMetadataIfAliasUsed bool   `json:"hide_metadata_if_alias_used"`
	CanAddComment         bool     `json:"can_add_comment"`
	CanAddCorporateLocation bool   `json:"can_add_corporate_location"`
	CanAddImageURL        bool     `json:"can_add_image_url"`
	CanAddMoreInfo        bool     `json:"can_add_more_info"`
	CanAddOpenCorporatesURL bool   `json:"can_add_open_corporates_url"`
	CanAddPhysicalLocation bool    `json:"can_add_physical_location"`
	CanAddPrivateAlias    bool     `json:"can_add_private_alias"`
	CanAddPublicAlias     bool     `json:"can_add_public_alias"`
	CanAddTag             bool     `json:"can_add_tag"`
	CanAddURL             bool     `json:"can_add_url"`
	CanAddWhereTag        bool     `json:"can_add_where_tag"`
	CanDeleteComment      bool     `json:"can_delete_comment"`
	CanDeleteCorporateLocation bool `json:"can_delete_corporate_location"`
	CanDeleteImageURL     bool     `json:"can_delete_image_url"`
	CanDeleteMoreInfo     bool     `json:"can_delete_more_info"`
	CanDeleteOpenCorporatesURL bool `json:"can_delete_open_corporates_url"`
	CanDeletePhysicalLocation bool  `json:"can_delete_physical_location"`
	CanDeleteTag          bool     `json:"can_delete_tag"`
	CanDeleteURL          bool     `json:"can_delete_url"`
	CanDeleteWhereTag     bool     `json:"can_delete_where_tag"`
	CanEditOwnerComment   bool     `json:"can_edit_owner_comment"`
	CanSeeImageURL        bool     `json:"can_see_image_url"`
	CanSeeMoreInfo        bool     `json:"can_see_more_info"`
	CanSeeOpenCorporatesURL bool   `json:"can_see_open_corporates_url"`
	CanSeeOtherAccountBankName bool `json:"can_see_other_account_bank_name"`
	CanSeeOtherAccountIBAN bool     `json:"can_see_other_account_iban"`
	CanSeeOtherAccountKind bool     `json:"can_see_other_account_kind"`
	CanSeeOtherAccountMetadata bool `json:"can_see_other_account_metadata"`
	CanSeeOtherAccountNationalIdentifier bool `json:"can_see_other_account_national_identifier"`
	CanSeeOtherAccountNumber bool `json:"can_see_other_account_number"`
	CanSeeOtherAccountRoutingAddress bool `json:"can_see_other_account_routing_address"`
	CanSeeOtherAccountRoutingScheme bool `json:"can_see_other_account_routing_scheme"`
	CanSeeOwnerComment    bool     `json:"can_see_owner_comment"`
	CanSeePhysicalLocation bool    `json:"can_see_physical_location"`
	CanSeePrivateAlias    bool     `json:"can_see_private_alias"`
	CanSeePublicAlias     bool     `json:"can_see_public_alias"`
	CanSeeTag             bool     `json:"can_see_tag"`
	CanSeeTransactionAmount bool   `json:"can_see_transaction_amount"`
	CanSeeTransactionBalance bool  `json:"can_see_transaction_balance"`
	CanSeeTransactionCurrency bool `json:"can_see_transaction_currency"`
	CanSeeTransactionDescription bool `json:"can_see_transaction_description"`
	CanSeeTransactionFinishDate bool `json:"can_see_transaction_finish_date"`
	CanSeeTransactionMetadata bool `json:"can_see_transaction_metadata"`
	CanSeeTransactionOtherBankAccount bool `json:"can_see_transaction_other_bank_account"`
	CanSeeTransactionStartDate bool `json:"can_see_transaction_start_date"`
	CanSeeTransactionThisBankAccount bool `json:"can_see_transaction_this_bank_account"`
	CanSeeTransactionType bool     `json:"can_see_transaction_type"`
	CanSeeURL             bool     `json:"can_see_url"`
	CanSeeWhereTag        bool     `json:"can_see_where_tag"`
}

type RevokedResponse struct {
	Revoked bool `json:"revoked"`
}

type AccountAccessResponse struct {
	Accounts []AccountMinimal `json:"accounts"`
}

type AccountMinimal struct {
	ID       string `json:"id"`
	Label    string `json:"label"`
	BankID   string `json:"bank_id"`
	ViewsAvailable []string `json:"views_available"`
}

func (c *AccountAccessController) GrantUserAccessToViewById(ctx *gin.Context) {
	var req GrantAccessRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	response := ViewResponse{
		ID:          req.ViewID,
		ShortName:   "Owner View",
		Description: "Owner view of account",
		IsPublic:    false,
		Alias:       "private",
		HideMetadataIfAliasUsed: false,
		CanAddComment: true,
		CanSeeTransactionAmount: true,
		CanSeeTransactionBalance: true,
		CanSeeTransactionCurrency: true,
		CanSeeTransactionDescription: true,
		CanSeeTransactionFinishDate: true,
		CanSeeTransactionStartDate: true,
		CanSeeTransactionType: true,
	}

	utils.SendJSONResponse(ctx, http.StatusCreated, response)
}

func (c *AccountAccessController) RevokeUserAccessToViewById(ctx *gin.Context) {
	var req RevokeAccessRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	response := RevokedResponse{
		Revoked: true,
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *AccountAccessController) CreateUserWithAccountAccessById(ctx *gin.Context) {
	var req CreateUserAccountAccessRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	if req.Provider != "dauth" {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Provider must be 'dauth'", "")
		return
	}

	response := ViewResponse{
		ID:          req.ViewID,
		ShortName:   "DAuth View",
		Description: "DAuth user view of account",
		IsPublic:    false,
		Alias:       "private",
		HideMetadataIfAliasUsed: false,
		CanSeeTransactionAmount: true,
		CanSeeTransactionBalance: true,
		CanSeeTransactionCurrency: true,
		CanSeeTransactionDescription: true,
	}

	utils.SendJSONResponse(ctx, http.StatusCreated, response)
}

func (c *AccountAccessController) GetAccountAccessByUserId(ctx *gin.Context) {

	response := AccountAccessResponse{
		Accounts: []AccountMinimal{
			{
				ID:     "account_001",
				Label:  "Main Account",
				BankID: "bank_001",
				ViewsAvailable: []string{"owner", "public"},
			},
			{
				ID:     "account_002",
				Label:  "Savings Account",
				BankID: "bank_001",
				ViewsAvailable: []string{"owner"},
			},
		},
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}
