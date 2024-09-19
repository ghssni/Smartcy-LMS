package models

import (
	"time"
)

type Assessments struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	EnrollmentID   uint      `json:"enrollment_id"`
	Score          int       `json:"score"`
	AssessmentType string    `json:"assessment_type"`
	TakenAt        time.Time `json:"taken_at"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
