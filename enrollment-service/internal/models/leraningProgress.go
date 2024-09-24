package models

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type LearningProgress struct {
	ID           uint32    `gorm:"primaryKey;autoIncrement" json:"id"`
	EnrollmentID uint32    `json:"enrollment_id"`
	LessonID     uint32    `json:"lesson_id"`
	Status       string    `json:"status"`
	CompletedAt  time.Time `json:"completed_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type LearningProgressInput struct {
	EnrollmentID uint32    `json:"enrollment_id" validate:"required"`
	LessonID     uint32    `json:"lesson_id" validate:"required"`
	Status       string    `json:"status" validate:"required"`
	CompletedAt  time.Time `json:"completed_at" validate:"required"`
	CreatedAt    time.Time `json:"created_at" validate:"required"`
	UpdatedAt    time.Time `json:"updated_at" validate:"required"`
}

// Validate is a function to validate LearningProgressInput
func (li *LearningProgressInput) Validate() error {
	validate := validator.New()
	return validate.Struct(li)
}
