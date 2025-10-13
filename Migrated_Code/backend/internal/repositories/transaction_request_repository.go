package repositories

import (
	"github.com/ashish-019-hash/obp-api-backend/internal/models"
	"gorm.io/gorm"
)

type transactionRequestRepository struct {
	db *gorm.DB
}

func NewTransactionRequestRepository(db *gorm.DB) TransactionRequestRepository {
	return &transactionRequestRepository{db: db}
}

func (r *transactionRequestRepository) Create(request *models.TransactionRequest) error {
	return r.db.Create(request).Error
}

func (r *transactionRequestRepository) GetByID(id int64) (*models.TransactionRequest, error) {
	var request models.TransactionRequest
	err := r.db.First(&request, id).Error
	if err != nil {
		return nil, err
	}
	return &request, nil
}

func (r *transactionRequestRepository) GetByTransactionRequestID(requestID string) (*models.TransactionRequest, error) {
	var request models.TransactionRequest
	err := r.db.Where("transaction_request_id = ?", requestID).First(&request).Error
	if err != nil {
		return nil, err
	}
	return &request, nil
}

func (r *transactionRequestRepository) UpdateStatus(requestID string, status string) error {
	return r.db.Model(&models.TransactionRequest{}).Where("transaction_request_id = ?", requestID).Update("status", status).Error
}

func (r *transactionRequestRepository) GetByStatus(status string) ([]*models.TransactionRequest, error) {
	var requests []*models.TransactionRequest
	err := r.db.Where("status = ?", status).Order("created_at DESC").Find(&requests).Error
	return requests, err
}

func (r *transactionRequestRepository) Update(request *models.TransactionRequest) error {
	return r.db.Save(request).Error
}

func (r *transactionRequestRepository) Delete(id int64) error {
	return r.db.Delete(&models.TransactionRequest{}, id).Error
}

func (r *transactionRequestRepository) List(limit, offset int) ([]*models.TransactionRequest, error) {
	var requests []*models.TransactionRequest
	err := r.db.Limit(limit).Offset(offset).Order("created_at DESC").Find(&requests).Error
	return requests, err
}
