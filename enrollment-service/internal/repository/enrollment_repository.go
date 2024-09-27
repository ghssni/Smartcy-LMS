package repository

import (
	"errors"
	"github.com/ghssni/Smartcy-LMS/Enrollment-Service/internal/models"
	"gorm.io/gorm"
)

type EnrollmentRepository interface {
	CreateEnrollment(enrollment *models.Enrollments) error
	ExistingEnrollment(studentId string, courseId uint32) (*models.Enrollments, error)
	GetEnrollmentsById(enrollmentId uint32) (*models.Enrollments, error)
	GetEnrollmentsByStudentId(studentId string) ([]models.Enrollments, error)
	DeleteEnrollmentById(enrollment *models.Enrollments) error
	BeginTransaction() *gorm.DB
}

type enrollmentRepository struct {
	db *gorm.DB
}

func (r *enrollmentRepository) ExistingEnrollment(studentId string, courseId uint32) (*models.Enrollments, error) {
	var enrollment models.Enrollments
	err := r.db.Where("student_id = ? AND course_id = ? AND payment_status = ?", studentId, courseId, "PAID").First(&enrollment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &enrollment, nil
}

func (r *enrollmentRepository) CreateEnrollment(enrollment *models.Enrollments) error {
	return r.db.Create(enrollment).Error
}

// GetEnrollmentsById returns an enrollment by its ID
func (r *enrollmentRepository) GetEnrollmentsById(enrollmentId uint32) (*models.Enrollments, error) {
	var enrollment models.Enrollments
	err := r.db.Where("id = ?", enrollmentId).First(&enrollment).Error
	if err != nil {
		return nil, err
	}
	return &enrollment, nil
}

func (r *enrollmentRepository) GetEnrollmentsByStudentId(studentId string) ([]models.Enrollments, error) {
	var enrollments []models.Enrollments
	return enrollments, r.db.Where("student_id = ? AND payment_status = ? ", studentId, "PAID").Find(&enrollments).Error
}

func (r *enrollmentRepository) DeleteEnrollmentById(enrollment *models.Enrollments) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	var enroll models.Enrollments
	err := tx.Where("id = ?", enrollment.ID).Delete(&enroll).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	var payment models.Payments
	if err := tx.Where("enrollment_id = ?", enrollment.ID).First(&payment).Error; err != nil {
		tx.Rollback()
		return err
	}

	payment.Status = "DELETED"
	if err := r.db.Save(&payment).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (r *enrollmentRepository) BeginTransaction() *gorm.DB {
	return r.db.Begin()
}

func NewEnrollmentRepository(db *gorm.DB) EnrollmentRepository {
	return &enrollmentRepository{
		db: db,
	}
}
