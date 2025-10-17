package repositories

import (
	"context"
	"database/sql"
	"obp-api-backend/internal/models"
	"obp-api-backend/pkg/db"
)

type resourceUserRepository struct{}

func NewResourceUserRepository() ResourceUserRepository {
	return &resourceUserRepository{}
}

func (r *resourceUserRepository) Create(ctx context.Context, user *models.ResourceUser) error {
	database := db.GetDB()
	query := `
		INSERT INTO resource_users (
			user_id, auth_user_id, provider, provider_id, name, email, company, is_deleted,
			last_marketing_agreement_signed_date, terms_accepted_date, privacy_accepted_date,
			created_by_consent_id, created_by_user_invitation_id, created_at, updated_at
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := database.ExecContext(ctx, query,
		user.UserID,
		user.AuthUserID,
		user.Provider,
		user.ProviderID,
		user.Name,
		user.Email,
		user.Company,
		user.IsDeleted,
		user.LastMarketingAgreementSignedDate,
		user.TermsAcceptedDate,
		user.PrivacyAcceptedDate,
		user.CreatedByConsentID,
		user.CreatedByUserInvitationID,
		user.CreatedAt,
		user.UpdatedAt,
	)
	return err
}

func (r *resourceUserRepository) GetByID(ctx context.Context, userID string) (*models.ResourceUser, error) {
	database := db.GetDB()
	query := `
		SELECT user_id, auth_user_id, provider, provider_id, name, email, company, is_deleted,
		       last_marketing_agreement_signed_date, terms_accepted_date, privacy_accepted_date,
		       created_by_consent_id, created_by_user_invitation_id, created_at, updated_at
		FROM resource_users
		WHERE user_id = ?
	`

	user := &models.ResourceUser{}
	err := database.QueryRowContext(ctx, query, userID).Scan(
		&user.UserID,
		&user.AuthUserID,
		&user.Provider,
		&user.ProviderID,
		&user.Name,
		&user.Email,
		&user.Company,
		&user.IsDeleted,
		&user.LastMarketingAgreementSignedDate,
		&user.TermsAcceptedDate,
		&user.PrivacyAcceptedDate,
		&user.CreatedByConsentID,
		&user.CreatedByUserInvitationID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return user, err
}

func (r *resourceUserRepository) GetByProviderID(ctx context.Context, provider, providerID string) (*models.ResourceUser, error) {
	database := db.GetDB()
	query := `
		SELECT user_id, auth_user_id, provider, provider_id, name, email, company, is_deleted,
		       last_marketing_agreement_signed_date, terms_accepted_date, privacy_accepted_date,
		       created_by_consent_id, created_by_user_invitation_id, created_at, updated_at
		FROM resource_users
		WHERE provider = ? AND provider_id = ?
	`

	user := &models.ResourceUser{}
	err := database.QueryRowContext(ctx, query, provider, providerID).Scan(
		&user.UserID,
		&user.AuthUserID,
		&user.Provider,
		&user.ProviderID,
		&user.Name,
		&user.Email,
		&user.Company,
		&user.IsDeleted,
		&user.LastMarketingAgreementSignedDate,
		&user.TermsAcceptedDate,
		&user.PrivacyAcceptedDate,
		&user.CreatedByConsentID,
		&user.CreatedByUserInvitationID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return user, err
}

func (r *resourceUserRepository) Update(ctx context.Context, user *models.ResourceUser) error {
	database := db.GetDB()
	query := `
		UPDATE resource_users
		SET auth_user_id = ?, provider = ?, provider_id = ?, name = ?, email = ?, company = ?, 
		    is_deleted = ?, last_marketing_agreement_signed_date = ?, terms_accepted_date = ?,
		    privacy_accepted_date = ?, created_by_consent_id = ?, created_by_user_invitation_id = ?,
		    updated_at = ?
		WHERE user_id = ?
	`
	_, err := database.ExecContext(ctx, query,
		user.AuthUserID,
		user.Provider,
		user.ProviderID,
		user.Name,
		user.Email,
		user.Company,
		user.IsDeleted,
		user.LastMarketingAgreementSignedDate,
		user.TermsAcceptedDate,
		user.PrivacyAcceptedDate,
		user.CreatedByConsentID,
		user.CreatedByUserInvitationID,
		user.UpdatedAt,
		user.UserID,
	)
	return err
}

func (r *resourceUserRepository) Delete(ctx context.Context, userID string) error {
	database := db.GetDB()
	query := `UPDATE resource_users SET is_deleted = TRUE WHERE user_id = ?`
	_, err := database.ExecContext(ctx, query, userID)
	return err
}
