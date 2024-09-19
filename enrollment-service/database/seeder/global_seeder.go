package seeder

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SeedAll(db *gorm.DB) {
	// Seed all data
	EnrollmentsSeeder()
	CertificateSeeder()
	AssessmentsSeeder()
	PaymentsSeeder()
	logrus.Println("Seed all success")
}
