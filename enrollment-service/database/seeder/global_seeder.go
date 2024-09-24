package seeder

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SeedAll(db *gorm.DB) {
	// Seed all data
	EnrollmentsSeeder(db)
	CertificateSeeder(db)
	AssessmentsSeeder(db)
	PaymentsSeeder(db)
	LearningProgressSeeder(db)
	logrus.Println("Seed all success")
}
