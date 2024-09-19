package seeder

import (
	"enrollment-service/internal/models"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/sirupsen/logrus"
)

func EnrollmentsSeeder() {
	var enrollments []models.Enrollments
	go gofakeit.Seed(0)
	for i := 1; i < 15; i++ {
		enrollments = append(enrollments, models.Enrollments{
			ID:            uint(i),
			StudentID:     gofakeit.UUID(),
			CourseID:      uint(i),
			PaymentStatus: gofakeit.RandomString([]string{"Pending", "Success", "Failed"}),
			EnrolledAt:    gofakeit.Date(),
			CreatedAt:     gofakeit.Date(),
			UpdatedAt:     gofakeit.Date(),
		})
	}
	logrus.Println("Enrollments seeder success")
}
