package repository

import (
	"github.com/ghssni/Smartcy-LMS/Enrollment-Service/proto/assessments"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"time"
)

type AssessmentsRepository interface {
	CreateAssessment(enrollmentId uint32, score int32, assessmentType string, takenAt, createdAt, updatedAt time.Time) (*assessments.Assessments, error)
	GetAssessmentByStudentId(id uint32, enrollmentId uint32) (*assessments.Assessments, error)
}

type assessmentsRepository struct {
	db *gorm.DB
}

func (r *assessmentsRepository) CreateAssessment(enrollmentId uint32, score int32, assessmentType string, takenAt, createdAt, updatedAt time.Time) (*assessments.Assessments, error) {
	var assessment = &assessments.Assessments{
		EnrollmentId:   enrollmentId,
		Score:          score,
		AssessmentType: assessmentType,
		TakenAt:        timestamppb.New(takenAt),
		CreatedAt:      timestamppb.New(createdAt),
		UpdatedAt:      timestamppb.New(updatedAt),
	}

	if err := r.db.Create(assessment).Error; err != nil {
		return nil, err
	}

	return assessment, nil
}

func (r *assessmentsRepository) GetAssessmentByStudentId(id uint32, enrollmentId uint32) (*assessments.Assessments, error) {
	var assessment assessments.Assessments
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
