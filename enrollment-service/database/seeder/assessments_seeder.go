package seeder

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/ghssni/Smartcy-LMS/Enrollment-Service/internal/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// AssessmentsSeeder is a function to seed assessments data
func AssessmentsSeeder(db *gorm.DB) {
	var assessments []models.Assessments
	gofakeit.Seed(0)
	for i := 1; i < 15; i++ {
		assessments = append(assessments, models.Assessments{
			ID:             uint32(uint(i)),
			EnrollmentID:   uint32(uint(i)),
			Score:          uint32(gofakeit.Number(0, 100)),
			AssessmentType: gofakeit.RandomString([]string{"Quiz", "Mid Exam", "Final Exam"}),
			TakenAt:        gofakeit.Date(),
			CreatedAt:      gofakeit.Date(),
			UpdatedAt:      gofakeit.Date(),
		})
	}
	if err := db.Create(&assessments).Error; err != nil {
		logrus.Fatalf("Failed to seed assessments: %v", err)
	} else {
		logrus.Println("Assessments seeder success")
	}
}
