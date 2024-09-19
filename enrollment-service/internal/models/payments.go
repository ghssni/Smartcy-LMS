package models

import (
	"time"
)

type Payments struct {
	ID                uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	EnrollmentID      uint       `json:"enrollment_id"`
	Amount            float64    `json:"amount"`
	TransactionStatus string     `json:"transaction_status"`
	TransactionDate   time.Time  `json:"transaction_date"`
	InvoiceID         string     `json:"invoice_id"`
	PaymentMethod     string     `json:"payment_method"`
	PaymentProvider   string     `json:"payment_provider"`
	Description       string     `json:"description"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	DeletedAt         *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}
