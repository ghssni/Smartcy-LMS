package seeder

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/ghssni/Smartcy-LMS/enrollment-service/internal/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func LearningProgressSeeder(db *gorm.DB) {
	var learningProgress []models.LearningProgress

	gofakeit.Seed(0)

	for i := 1; i < 15; i++ {
		learningProgress = append(learningProgress, models.LearningProgress{
			ID:           uint(i),
			EnrollmentID: uint(i),
			LessonID:     uint(i),
			Status:       gofakeit.RandomString([]string{"Not Started", "In Progress", "Completed"}),
			CompletedAt:  gofakeit.Date(),
			CreatedAt:    gofakeit.Date(),
		})
	}
	if err := db.Create(&learningProgress).Error; err != nil {
		logrus.Fatalf("Failed to seed learning progress: %v", err)
	} else {
		logrus.Println("Learning Progress seeder success")
	}
}
