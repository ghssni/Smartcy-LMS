package migrations

import (
	"github.com/ghssni/Smartcy-LMS/Email-Service/internal/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func CreateEmailsTableMigration() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20240919012304_create_Emails_table",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(models.EmailLogs{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("emails")
		},
	}
}
