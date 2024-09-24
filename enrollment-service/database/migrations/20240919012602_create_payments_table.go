package migrations

import (
	"github.com/ghssni/Smartcy-LMS/enrollment-service/internal/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// Fungsi migrasi untuk create_payments_table
func createPaymentsTableMigration() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20240919012602_create_payments_table",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(
				&models.Payments{},
			)
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("payments")
		},
	}
}
