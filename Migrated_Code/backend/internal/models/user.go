package models

import "time"

type User struct {
	UserPrimaryKey                    string     `json:"user_primary_key"`
	UserId                           string     `json:"user_id"`
	Provider                         string     `json:"provider"`
	EmailAddress                     string     `json:"email_address"`
	Name                             string     `json:"name"`
	CreatedByConsentId               *string    `json:"created_by_consent_id,omitempty"`
	CreatedByUserInvitationId        *string    `json:"created_by_user_invitation_id,omitempty"`
	IsDeleted                        *bool      `json:"is_deleted,omitempty"`
	LastMarketingAgreementSignedDate *time.Time `json:"last_marketing_agreement_signed_date,omitempty"`
}

func NewUser() *User {
	return &User{}
}
