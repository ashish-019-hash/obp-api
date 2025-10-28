package repositories

import (
	"context"
	"database/sql"
	"time"

	"obp-api-backend/internal/models"
	"obp-api-backend/pkg/db"
)

type customerRepository struct {
	db *sql.DB
}

func NewCustomerRepository() CustomerRepository {
	return &customerRepository{
		db: db.GetDB(),
	}
}

func (r *customerRepository) Create(ctx context.Context, customer *models.Customer) error {
	query := `INSERT INTO customers (customer_id, bank_id, customer_number, legal_name, mobile_phone_number, email, 
			  face_image_url, date_of_birth, relationship_status, dependents, dob_of_dependents, highest_education_attained, 
			  employment_status, kyc_status, last_ok_date, credit_rating_rating, credit_rating_source, credit_limit_currency, credit_limit_amount) 
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	
	dobDependentsStr := ""
	if len(customer.DobOfDependents) > 0 {
		dobDependentsStr = customer.DobOfDependents[0].Format("2006-01-02")
	}
	
	_, err := r.db.ExecContext(ctx, query,
		customer.CustomerId,
		customer.BankId,
		customer.Number,
		customer.LegalName,
		customer.MobileNumber,
		customer.Email,
		customer.FaceImage.Url,
		customer.DateOfBirth,
		customer.RelationshipStatus,
		customer.Dependents,
		dobDependentsStr,
		customer.HighestEducationAttained,
		customer.EmploymentStatus,
		customer.KycStatus,
		customer.LastOkDate,
		"", // credit_rating_rating - not available in model
		"", // credit_rating_source - not available in model
		"", // credit_limit_currency - not available in model
		"", // credit_limit_amount - not available in model
	)
	return err
}

func (r *customerRepository) GetByID(ctx context.Context, customerID string) (*models.Customer, error) {
	query := `SELECT customer_id, bank_id, customer_number, legal_name, mobile_phone_number, email, 
			  face_image_url, date_of_birth, relationship_status, dependents, dob_of_dependents, highest_education_attained, 
			  employment_status, kyc_status, last_ok_date, credit_rating_rating, credit_rating_source, credit_limit_currency, credit_limit_amount 
			  FROM customers WHERE customer_id = ?`
	
	customer := &models.Customer{}
	var dobDependentsStr, creditRatingRating, creditRatingSource, creditLimitCurrency, creditLimitAmount string
	
	err := r.db.QueryRowContext(ctx, query, customerID).Scan(
		&customer.CustomerId,
		&customer.BankId,
		&customer.Number,
		&customer.LegalName,
		&customer.MobileNumber,
		&customer.Email,
		&customer.FaceImage.Url,
		&customer.DateOfBirth,
		&customer.RelationshipStatus,
		&customer.Dependents,
		&dobDependentsStr,
		&customer.HighestEducationAttained,
		&customer.EmploymentStatus,
		&customer.KycStatus,
		&customer.LastOkDate,
		&creditRatingRating,
		&creditRatingSource,
		&creditLimitCurrency,
		&creditLimitAmount,
	)
	
	if err != nil {
		return nil, err
	}
	
	customer.DobOfDependents = make([]time.Time, 0)
	
	return customer, nil
}

func (r *customerRepository) GetByBankID(ctx context.Context, bankID string) ([]*models.Customer, error) {
	query := `SELECT customer_id, bank_id, customer_number, legal_name, mobile_phone_number, email, 
			  face_image_url, date_of_birth, relationship_status, dependents, dob_of_dependents, highest_education_attained, 
			  employment_status, kyc_status, last_ok_date, credit_rating_rating, credit_rating_source, credit_limit_currency, credit_limit_amount 
			  FROM customers WHERE bank_id = ?`
	
	rows, err := r.db.QueryContext(ctx, query, bankID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var customers []*models.Customer
	for rows.Next() {
		customer := &models.Customer{}
		var dobDependentsStr, creditRatingRating, creditRatingSource, creditLimitCurrency, creditLimitAmount string
		
		err := rows.Scan(
			&customer.CustomerId,
			&customer.BankId,
			&customer.Number,
			&customer.LegalName,
			&customer.MobileNumber,
			&customer.Email,
			&customer.FaceImage.Url,
			&customer.DateOfBirth,
			&customer.RelationshipStatus,
			&customer.Dependents,
			&dobDependentsStr,
			&customer.HighestEducationAttained,
			&customer.EmploymentStatus,
			&customer.KycStatus,
			&customer.LastOkDate,
			&creditRatingRating,
			&creditRatingSource,
			&creditLimitCurrency,
			&creditLimitAmount,
		)
		if err != nil {
			return nil, err
		}
		
		customer.DobOfDependents = make([]time.Time, 0)
		
		customers = append(customers, customer)
	}
	
	return customers, rows.Err()
}

func (r *customerRepository) Update(ctx context.Context, customer *models.Customer) error {
	query := `UPDATE customers SET customer_number = ?, legal_name = ?, mobile_phone_number = ?, email = ?, 
			  face_image_url = ?, date_of_birth = ?, relationship_status = ?, dependents = ?, dob_of_dependents = ?, 
			  highest_education_attained = ?, employment_status = ?, kyc_status = ?, last_ok_date = ? 
			  WHERE customer_id = ?`
	
	dobDependentsStr := ""
	if len(customer.DobOfDependents) > 0 {
		dobDependentsStr = customer.DobOfDependents[0].Format("2006-01-02")
	}
	
	_, err := r.db.ExecContext(ctx, query,
		customer.Number,
		customer.LegalName,
		customer.MobileNumber,
		customer.Email,
		customer.FaceImage.Url,
		customer.DateOfBirth,
		customer.RelationshipStatus,
		customer.Dependents,
		dobDependentsStr,
		customer.HighestEducationAttained,
		customer.EmploymentStatus,
		customer.KycStatus,
		customer.LastOkDate,
		customer.CustomerId,
	)
	return err
}

func (r *customerRepository) Delete(ctx context.Context, customerID string) error {
	query := `DELETE FROM customers WHERE customer_id = ?`
	_, err := r.db.ExecContext(ctx, query, customerID)
	return err
}
