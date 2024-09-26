package seeder

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/ghssni/Smartcy-LMS/Enrollment-Service/internal/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func EnrollmentsSeeder(db *gorm.DB) {
	var enrollments []models.Enrollments
	go gofakeit.Seed(0)
	for i := 1; i < 15; i++ {
		enrollments = append(enrollments, models.Enrollments{
			ID:            uint32(uint(i)),
			StudentID:     gofakeit.UUID(),
			CourseID:      uint32(uint(i)),
			PaymentStatus: gofakeit.RandomString([]string{"Pending", "Success", "Failed"}),
			EnrolledAt:    gofakeit.Date(),
			CreatedAt:     gofakeit.Date(),
			UpdatedAt:     gofakeit.Date(),
		})
	}
	if err := db.Create(&enrollments).Error; err != nil {
		logrus.Fatalf("Failed to seed enrollments: %v", err)
	} else {
		logrus.Println("Enrollments seeder success")
	}
}
