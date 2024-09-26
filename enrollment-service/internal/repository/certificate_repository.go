package repository

import (
	"context"
	"github.com/ghssni/Smartcy-LMS/Enrollment-Service/internal/models"
	"gorm.io/gorm"
)

type CertificateRepository interface {
	CreateCertificate(ctx context.Context, certificate *models.Certificate) (*models.Certificate, error)
	GetCertificateByEnrollmentId(ctx context.Context, enrollmentId uint32) (*models.Certificate, error)
}

type certificateRepository struct {
	db *gorm.DB
}

func (r *certificateRepository) CreateCertificate(ctx context.Context, certificate *models.Certificate) (*models.Certificate, error) {
	//TODO implement me
	panic("implement me")
}

func (r *certificateRepository) GetCertificateByEnrollmentId(ctx context.Context, enrollmentId uint32) (*models.Certificate, error) {
	//TODO implement me
	panic("implement me")
}

func NewCertificateRepository(db *gorm.DB) CertificateRepository {
	return &certificateRepository{
		db: db,
	}
}
