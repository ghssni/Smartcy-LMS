package models

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type Enrollments struct {
	ID            uint32    `gorm:"primaryKey;autoIncrement" json:"id"`
	StudentID     string    `json:"student_id"`
	CourseID      uint32    `json:"course_id"`
	PaymentStatus string    `json:"payment_status"`
	EnrolledAt    time.Time `json:"enrolled_at"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type EnrollmentInput struct {
	StudentID     string    `json:"student_id" validate:"required"`
	CourseID      uint32    `json:"course_id" validate:"required"`
	PaymentStatus string    `json:"payment_status" validate:"required" default:"Pending"`
	EnrolledAt    time.Time `json:"enrolled_at" validate:"required"`
	CreatedAt     time.Time `json:"created_at" validate:"required"`
	UpdatedAt     time.Time `json:"updated_at" validate:"required"`
}

// Validate is a function to validate EnrollmentInput
func (ei *EnrollmentInput) Validate() error {
	validate := validator.New()
	return validate.Struct(ei)
}
