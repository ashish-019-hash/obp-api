package models

import "time"

type ResourceUser struct {
	UserID                           string     `json:"user_id"`
	AuthUserID                       *string    `json:"auth_user_id,omitempty"`
	Provider                         string     `json:"provider"`
	ProviderID                       string     `json:"provider_id"`
	Name                             string     `json:"name"`
	Email                            string     `json:"email"`
	Company                          string     `json:"company,omitempty"`
	IsDeleted                        bool       `json:"is_deleted"`
	LastMarketingAgreementSignedDate *time.Time `json:"last_marketing_agreement_signed_date,omitempty"`
	TermsAcceptedDate                *time.Time `json:"terms_accepted_date,omitempty"`
	PrivacyAcceptedDate              *time.Time `json:"privacy_accepted_date,omitempty"`
	CreatedByConsentID               *string    `json:"created_by_consent_id,omitempty"`
	CreatedByUserInvitationID        *string    `json:"created_by_user_invitation_id,omitempty"`
	CreatedAt                        time.Time  `json:"created_at"`
	UpdatedAt                        time.Time  `json:"updated_at"`
}

type User = ResourceUser

func NewUser() *ResourceUser {
	return &ResourceUser{}
}
