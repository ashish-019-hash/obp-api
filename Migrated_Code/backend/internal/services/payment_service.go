package services

import (
	"context"
	"fmt"
	"time"
)

type PaymentService interface {
	InitiatePayment(ctx context.Context, paymentRequest map[string]interface{}) (string, error)
}

type paymentService struct {
	transactionService TransactionService
}

func NewPaymentService(transactionService TransactionService) PaymentService {
	return &paymentService{
		transactionService: transactionService,
	}
}

func (s *paymentService) InitiatePayment(ctx context.Context, paymentRequest map[string]interface{}) (string, error) {
	paymentID := fmt.Sprintf("payment-%d", time.Now().Unix())
	return paymentID, nil
}
