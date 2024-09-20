package seeder

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/ghssni/Smartcy-LMS/enrollment-service/internal/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// CertificateSeeder is a function to seed certificate data
func CertificateSeeder(db *gorm.DB) {
	var certificates []models.Certificate
	gofakeit.Seed(0)

	for i := 1; i < 15; i++ {
		certificates = append(certificates, models.Certificate{
			ID:             uint(i),
			EnrollmentID:   uint(i),
			CertificateURL: gofakeit.URL(),
			IssuedAt:       gofakeit.Date(),
			CreatedAt:      gofakeit.Date(),
			UpdatedAt:      gofakeit.Date(),
		})
	}

	if err := db.Create(&certificates).Error; err != nil {
		logrus.Fatalf("Failed to seed certificates: %v", err)
	} else {
		logrus.Println("Certificate seeder success")
	}

}
