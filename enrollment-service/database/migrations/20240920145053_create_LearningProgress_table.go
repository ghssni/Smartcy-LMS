package migrations

import (
	"github.com/ghssni/Smartcy-LMS/enrollment-service/internal/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// Fungsi migrasi untuk create_LearningProgress_table
func createLearningprogressTableMigration() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20240920145053_create_LearningProgress_table",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(models.LearningProgress{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("learning_progress")
		},
	}
}
