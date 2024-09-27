package repository

import (
	"github.com/ghssni/Smartcy-LMS/Enrollment-Service/internal/models"
	"gorm.io/gorm"
	"time"
)

type AssessmentsRepository interface {
	CreateAssessment(enrollmentId uint32, score int32, assessmentType string, takenAt, createdAt, updatedAt time.Time) (*models.Assessments, error)
	GetAssessmentByStudentId(id uint32, enrollmentId uint32) (*models.Assessments, error)
}

type assessmentsRepository struct {
	db *gorm.DB
}

func (r *assessmentsRepository) CreateAssessment(enrollmentId uint32, score int32, assessmentType string, takenAt, createdAt, updatedAt time.Time) (*models.Assessments, error) {
	var assessment = &models.Assessments{
		EnrollmentID:   enrollmentId,
		Score:          uint32(score),
		AssessmentType: assessmentType,
		TakenAt:        takenAt,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
	}

	if err := r.db.Create(assessment).Error; err != nil {
		return nil, err
	}

	return assessment, nil
}

func (r *assessmentsRepository) GetAssessmentByStudentId(id uint32, enrollmentId uint32) (*models.Assessments, error) {
	var assessment models.Assessments
	err := r.db.Where("id = ? AND enrollment_id  = ?", id, enrollmentId).First(&assessment).Error
	if err != nil {
		return nil, err
	}

	return &assessment, nil
}

func NewAssessmentsRepository(db *gorm.DB) AssessmentsRepository {
	return &assessmentsRepository{
		db: db,
	}
}
