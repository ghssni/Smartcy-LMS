package migrations

import (
	"github.com/ghssni/Smartcy-LMS/Enrollment-Service/internal/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// createEnrollmentsTableMigration is a function to create table enrollments
func createEnrollmentsTableMigration() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20240919012304_create_enrollments_table",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(
				&models.Enrollments{},
			)
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("enrollments")
		},
	}
}
