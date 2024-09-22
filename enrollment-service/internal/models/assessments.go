package models

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type Assessments struct {
	ID             uint32    `gorm:"primaryKey;autoIncrement" json:"id"`
	EnrollmentID   uint32    `json:"enrollment_id"`
	Score          uint32    `json:"score"`
	AssessmentType string    `json:"assessment_type"`
	TakenAt        time.Time `json:"taken_at"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type AssessmentsInput struct {
	EnrollmentID   uint      `json:"enrollment_id" validate:"required"`
	Score          uint32    `json:"score" validate:"required"`
	AssessmentType string    `json:"assessment_type" validate:"required"`
	TakenAt        time.Time `json:"taken_at" validate:"required"`
	CreatedAt      time.Time `json:"created_at" validate:"required"`
	UpdatedAt      time.Time `json:"updated_at" validate:"required"`
}

// Validate is a function to validate AssessmentsInput
func (ai *AssessmentsInput) Validate() error {
	validate := validator.New()
	return validate.Struct(ai)
}
