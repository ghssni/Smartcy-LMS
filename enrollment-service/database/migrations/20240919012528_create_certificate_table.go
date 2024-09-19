package migrations

import (
	"enrollment-service/internal/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// createCertificateTableMigration is a function to create table certificates
func createCertificateTableMigration() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20240919012528_create_certificate_table",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(
				&models.Certificate{},
			)
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("certificates")
		},
	}
}
