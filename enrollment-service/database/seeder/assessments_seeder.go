package seeder

import (
	"enrollment-service/internal/models"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/sirupsen/logrus"
)

// AssessmentsSeeder is a function to seed assessments data
func AssessmentsSeeder() {
	var assessments []models.Assessments
	gofakeit.Seed(0)
	for i := 1; i < 15; i++ {
		assessments = append(assessments, models.Assessments{
			ID:             uint(i),
			EnrollmentID:   uint(i),
			Score:          gofakeit.Number(0, 100),
			AssessmentType: gofakeit.RandomString([]string{"Quiz", "Mid Exam", "Final Exam"}),
			TakenAt:        gofakeit.Date(),
			CreatedAt:      gofakeit.Date(),
			UpdatedAt:      gofakeit.Date(),
		})
	}
	logrus.Println("Assessments seeder success")
}
