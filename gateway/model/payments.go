package model

import (
	"time"
)

type Payments struct {
	ID                     uint32    `gorm:"primaryKey;autoIncrement" json:"id"`
	EnrollmentID           uint32    `json:"enrollment_id"`
	ExternalID             string    `json:"external_id"`
	UserID                 string    `json:"user_id"`
	IsHigh                 bool      `json:"is_high"`
	PaymentMethod          string    `json:"payment_method"`
	Status                 string    `json:"status"`
	MerchantName           string    `json:"merchant_name"`
	Amount                 float64   `json:"amount"`
	PaidAmount             float64   `json:"paid_amount"`
	BankCode               string    `json:"bank_code"`
	PaidAt                 string    `json:"paid_at"`
	PayerEmail             string    `json:"payer_email"`
	Description            string    `json:"description"`
	AdjustedReceivedAmount float64   `json:"adjusted_received_amount"`
	FeesPaidAmount         float64   `json:"fees_paid_amount"`
	Updated                time.Time `json:"updated"`
	Created                time.Time `json:"created"`
	Currency               string    `json:"currency"`
	PaymentChannel         string    `json:"payment_channel"`
	PaymentDestination     string    `json:"payment_destination"`
	InvoiceUrl             string    `json:"invoice_url"`
}
