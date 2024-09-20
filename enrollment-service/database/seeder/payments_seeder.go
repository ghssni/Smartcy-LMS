package seeder

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/ghssni/Smartcy-LMS/enrollment-service/internal/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// PaymentsSeeder is a function to seed payments data
func PaymentsSeeder(db *gorm.DB) {
	var payments []models.Payments
	gofakeit.Seed(0)

	for i := 1; i < 15; i++ {
		payments = append(payments, models.Payments{
			ID:                uint(i),
			EnrollmentID:      uint(i),
			Amount:            gofakeit.Price(1000000, 10000000),
			TransactionStatus: gofakeit.RandomString([]string{"Pending", "Success", "Failed"}),
			TransactionDate:   gofakeit.Date(),
			InvoiceID:         gofakeit.UUID(),
			PaymentMethod:     gofakeit.RandomString([]string{"Credit Card", "Bank Transfer", "Virtual Account"}),
			PaymentProvider:   gofakeit.RandomString([]string{"Visa", "MasterCard", "E-Wallet"}),
			Description:       gofakeit.Sentence(10),
			CreatedAt:         gofakeit.Date(),
			UpdatedAt:         gofakeit.Date(),
			DeletedAt:         nil,
		})
	}

	if err := db.Create(&payments).Error; err != nil {
		logrus.Fatalf("Failed to seed payments: %v", err)
	} else {
		logrus.Println("Payments seeder success")
	}
}
