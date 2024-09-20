package models

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type Certificate struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	EnrollmentID   uint      `json:"enrollment_id"`
	IssuedAt       time.Time `json:"issued_at"`
	CertificateURL string    `json:"certificate_url"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type CertificateInput struct {
	EnrollmentID   uint      `json:"enrollment_id" validate:"required"`
	IssuedAt       time.Time `json:"issued_at" validate:"required"`
	CertificateURL string    `json:"certificate_url" validate:"required"`
	CreatedAt      time.Time `json:"created_at" validate:"required"`
	UpdatedAt      time.Time `json:"updated_at" validate:"required"`
}

// Validate is a function to validate CertificateInput
func (ci *CertificateInput) Validate() error {
	validate := validator.New()
	return validate.Struct(ci)
}
