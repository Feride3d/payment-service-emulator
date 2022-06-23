package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/Feride3d/payment-service-emulator/internal/storage/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Storage interface {
	CreatePayment(ctx context.Context, pt models.Payment) error
	UpdateStatus(ctx context.Context, id int, status string) (string, error)
	GetStatus(ctx context.Context, id int) (string, error)
	GetPaymentsByUserId(ctx context.Context, userID int) ([]*models.Payment, error)
	GetPaymentsByUserEmail(ctx context.Context, userEmail string) ([]*models.Payment, error)
	CancelPayment(ctx context.Context, id int) error
}

type Service struct {
	Payment Storage
}

type PaymentStore struct {
	conn *pgxpool.Pool
}

// New opens connection to database.
func New(dbURL string) (*PaymentStore, error) {
	conn, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	s := &PaymentStore{
		conn: conn,
	}
	return s, nil
}

// Close closes connection to database.
func (s *PaymentStore) Close() {
	if s.conn != nil {
		s.conn.Close()
	}
}

// CreatePayment creates payment in database
func (s *PaymentStore) CreatePayment(ctx context.Context, payment models.Payment) error {
	_, err := s.conn.Exec(ctx, `INSERT INTO payments (id, user_id, user_email, amount, currency, status)
VALUES ($1, $2, $3, $4, $5, $6)`,
		payment.ID, payment.UserID, payment.UserEmail, payment.Amount, payment.Currency, payment.Status)
	if err != nil {
		return fmt.Errorf("failed to create payment: %w", err)
	}
	return nil
}

// UpdateStatus updates payment status in database
func (s *PaymentStore) UpdateStatus(ctx context.Context, id int, status string) (string, error) {
	var changedStatus string
	now := time.Now()
	err := s.conn.QueryRow(ctx, `UPDATE payments SET status = $1, time_change = $2 WHERE id = $3 RETURNING status;`,
		status, now, id).Scan(&changedStatus)
	if err != nil {
		return "", fmt.Errorf("update status: %w", err)
	}
	return changedStatus, nil
}

// GetStatus receives status of payment by payment ID from database
func (s *PaymentStore) GetStatus(ctx context.Context, id int) (string, error) {
	var actualStatus string
	err := s.conn.QueryRow(ctx, `SELECT status FROM payments WHERE id = $1;`, id).Scan(&actualStatus)
	if err != nil {
		return "", fmt.Errorf("get status: %w", err)
	}
	return actualStatus, nil
}

// GetPaymentsByUserId receives payments by user ID from database
func (s *PaymentStore) GetPaymentsByUserId(ctx context.Context, userID int) ([]*models.Payment, error) {
	payment := &models.Payment{}
	var payments []*models.Payment
	rows, err := s.conn.Query(ctx, `SELECT * FROM payments WHERE user_id = $1;`, userID)
	if err != nil {
		return nil, fmt.Errorf("get status bu user id: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&payment.ID, &payment.UserID, &payment.UserEmail, &payment.Amount,
			&payment.Currency, &payment.TimeCreation, &payment.TimeChange, &payment.Status)
		if err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}

		payments = append(payments, payment)
	}
	return payments, nil
}

// GetPaymentsByUserEmail receives payments by user email from database
func (s *PaymentStore) GetPaymentsByUserEmail(ctx context.Context, userEmail string) ([]*models.Payment, error) {
	payment := &models.Payment{}
	var payments []*models.Payment
	rows, err := s.conn.Query(ctx, `SELECT * FROM payments WHERE user_email = $1;`, userEmail)
	if err != nil {
		return nil, fmt.Errorf("get status bu user email: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&payment.ID, &payment.UserID, &payment.UserEmail, &payment.Amount,
			&payment.Currency, &payment.TimeCreation, &payment.TimeChange, &payment.Status)
		if err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}

		payments = append(payments, payment)
	}
	return payments, nil
}

// CancelPayment deletes payment from database
func (s *PaymentStore) CancelPayment(ctx context.Context, id int) error {
	_, err := s.conn.Exec(ctx, "DELETE FROM payments WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to cancel payment: %w", err)
	}
	return nil
}
