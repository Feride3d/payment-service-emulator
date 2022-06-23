package service

import (
	"context"
	"fmt"

	"github.com/Feride3d/payment-service-emulator/internal/storage"
	"github.com/Feride3d/payment-service-emulator/internal/storage/models"
)

type PaymentService struct {
	db storage.Storage
}

func NewPaymentService(db storage.Storage) *PaymentService {
	return &PaymentService{db: db}
}

// CreatePayment creates new payment
func (s *PaymentService) CreatePayment(ctx context.Context, pt models.Payment) error {
	err := s.db.CreatePayment(ctx, pt)
	if err != nil {
		return fmt.Errorf("create payment: %w", err)
	}
	return nil
}

// UpdateStatus updates status of payment
func (s *PaymentService) UpdateStatus(ctx context.Context, id int, status string) (string, error) {
	newStatus, err := s.db.UpdateStatus(ctx, id, status)
	if err != nil {
		return "", fmt.Errorf("update status: %w", err)
	}
	return newStatus, nil
}

// GetStatus receives status of payment by payment ID
func (s *PaymentService) GetStatus(ctx context.Context, id int) (string, error) {
	status, err := s.db.GetStatus(ctx, id)
	if err != nil {
		return "", fmt.Errorf("receive status: %w", err)
	}
	return status, nil
}

// GetPaymentsByUserId receives payments by user ID
func (s *PaymentService) GetPaymentsByUserId(ctx context.Context, userID int) ([]*models.Payment, error) {
	p, err := s.db.GetPaymentsByUserId(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("receive payments by user id: %w", err)
	}
	return p, nil
}

// GetPaymentsByUserEmail receives payments by user email
func (s *PaymentService) GetPaymentsByUserEmail(ctx context.Context, userEmail string) ([]*models.Payment, error) {
	p, err := s.db.GetPaymentsByUserEmail(ctx, userEmail)
	if err != nil {
		return nil, fmt.Errorf("receive payments by user email: %w", err)
	}
	return p, nil
}

// CancelPayment deletes payment by id
func (s *PaymentService) CancelPayment(ctx context.Context, id int) error {
	err := s.db.CancelPayment(ctx, id)
	if err != nil {
		return fmt.Errorf("cancel payment: %w", err)
	}
	return nil
}
