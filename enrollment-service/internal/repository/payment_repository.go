package repository

import (
	"context"
	"errors"
	"github.com/ghssni/Smartcy-LMS/enrollment-service/internal/models"
	"gorm.io/gorm"
	"time"
)

type PaymentRepository interface {
	CreatePayment(ctx context.Context, payment *models.Payments) error
	GetPaymentByEnrollmentId(ctx context.Context, enrollmentId string) (*models.Payments, error)
	UpdatePaymentStatusByWebhook(ctx context.Context, invoiceId string, updatedPayment models.Payments) (*models.Payments, error)
	UpdateExpiredPaymentStatus() error
}

type paymentRepository struct {
	db *gorm.DB
}

func (r *paymentRepository) GetPaymentByEnrollmentId(ctx context.Context, enrollmentId string) (*models.Payments, error) {
	var payment models.Payments
	if err := r.db.Where("enrollment_id = ?", enrollmentId).First(&payment).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *paymentRepository) CreatePayment(ctx context.Context, payment *models.Payments) error {

	if err := r.db.Create(payment).Error; err != nil {
		return err
	}
	return nil
}

func (r *paymentRepository) UpdatePaymentStatusByWebhook(ctx context.Context, invoiceId string, updatedPayment models.Payments) (*models.Payments, error) {
	var payment models.Payments
	var enrollment models.Enrollments
	if err := r.db.Where("external_id = ?", invoiceId).First(&payment).Error; err != nil {
		return nil, err
	}

	if err := r.db.Model(&payment).Updates(updatedPayment).Error; err != nil {
		return nil, err
	}

	if err := r.db.Where("id = ?", payment.EnrollmentID).First(&enrollment).Error; err != nil {
		return nil, err
	}

	enrollment.PaymentStatus = updatedPayment.Status
	if err := r.db.Save(&enrollment).Error; err != nil {
		return nil, err
	}

	return &payment, nil
}

// UpdateExpiredPaymentStatus updates the status of expired payments
func (r *paymentRepository) UpdateExpiredPaymentStatus() error {
	now := time.Now()
	var payment models.Payments
	if err := r.db.Where("created < ? AND status = ?", now.Add(-24*time.Hour), "PENDING").First(&payment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	// update status to EXPIRED
	if err := r.db.Model(&payment).Update("status", "EXPIRED").Error; err != nil {
		return err
	}
	if err := r.db.Model(&models.Enrollments{}).
		Where("student_id = ? AND created < ? AND payment_status = ?", payment.UserID, now.Add(-24*time.Hour), "PENDING").
		Update("payment_status", "EXPIRED").Error; err != nil {
		return err
	}

	return nil
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{
		db: db,
	}
}
