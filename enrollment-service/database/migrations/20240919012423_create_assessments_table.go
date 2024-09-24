package migrations

import (
	"github.com/ghssni/Smartcy-LMS/enrollment-service/internal/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// createAssessmentsTableMigration is a function to create table assessments
func createAssessmentsTableMigration() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20240919012423_create_assessments_table",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(
				&models.Assessments{},
			)
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("assessments")
		},
	}
}
