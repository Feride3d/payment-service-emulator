package models

import (
	"time"
)

// Payment represents a payment
type Payment struct {
	ID           int       `db:"id" json:"id"`
	UserID       int       `db:"user_id" json:"userId"`
	UserEmail    string    `db:"user_email" json:"userEmail"`
	Amount       int       `db:"amount" json:"amount"`
	Currency     string    `db:"currency" json:"currency"`
	TimeCreation time.Time `db:"time_creation" json:"timeCreation"`
	TimeChange   time.Time `db:"time_change" json:"timeChange"`
	Status       string    `db:"status" json:"status"`
}

// PaymentStatus represents a status of payment
type PaymentStatus string

const (
	// StatusNew means a payment is new
	StatusNew PaymentStatus = "new"
	// StatusSuccess informs you that a payment is paid
	StatusSuccess PaymentStatus = "success"
	// StatusFailed means a payment has failed
	StatusFailed PaymentStatus = "failed"
	// StatusError means that something went wrong at the time of creating the payment
	StatusError PaymentStatus = "error"
	// StatusCancelled means a payment is cancelled
	StatusCancelled PaymentStatus = "cancelled"
)

// PaymentList represents a list of payments.
type PaymentList struct {
	Payments []Payment `json:"Payments"`
}
