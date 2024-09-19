package seeder

import (
	"enrollment-service/internal/models"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/sirupsen/logrus"
)

// CertificateSeeder is a function to seed certificate data
func CertificateSeeder() {
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

	logrus.Println("Seed all success")
}
