package memstorage

import (
	"context"
	"sync"

	"github.com/Feride3d/payment-service-emulator/internal/storage/models"
)

type MemStorage struct {
	mu       sync.Mutex // many threads can come and work with one variable so we use a mutex for minimize a risk of data race
	payments models.PaymentList
	lastID   int
}

// New return memstore
func New() *MemStorage {
	return &MemStorage{}
}

func (s *MemStorage) CreatePayment(ctx context.Context, payment models.Payment) (models.Payment, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.lastID++
	payment.ID = s.lastID
	s.payments.Payments = append(s.payments.Payments, payment)
	return payment, nil
}

func (s *MemStorage) CancelPayment(ctx context.Context, id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, t := range s.payments.Payments {
		if t.ID == id {
			// copy payments that are stored in the store, if done differently, an array, that is not protected by a mutex, will be returned
			s.payments.Payments = append(s.payments.Payments[:i], s.payments.Payments[i+1:]...)
			break
		}
	}
	return nil
}
