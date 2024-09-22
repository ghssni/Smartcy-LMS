package seeder

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/ghssni/Smartcy-LMS/enrollment-service/internal/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

// PaymentsSeeder is a function to seed payments data
func PaymentsSeeder(db *gorm.DB) {
	var payments []models.Payments
	gofakeit.Seed(11)

	for i := 1; i < 15; i++ {
		payments = append(payments, models.Payments{
			ID:                     uint32(uint(i)),
			ExternalID:             gofakeit.UUID(),
			EnrollmentID:           uint32(uint(i)),
			UserID:                 gofakeit.UUID(),
			IsHigh:                 false,
			PaymentMethod:          "",
			Status:                 "PENDING",
			MerchantName:           "",
			Amount:                 gofakeit.Price(1000, 5000),
			PaidAmount:             gofakeit.Price(1000, 5000),
			BankCode:               gofakeit.CreditCard().Number,
			PayerEmail:             gofakeit.Email(),
			Description:            gofakeit.Sentence(10),
			AdjustedReceivedAmount: gofakeit.Price(1000, 5000),
			FeesPaidAmount:         gofakeit.Price(1000, 5000),
			Updated:                time.Time{},
			Created:                time.Time{},
			Currency:               "IDR",
			PaymentChannel:         gofakeit.CreditCard().Type,
			PaymentDestination:     "",
		})
	}

	if err := db.Create(&payments).Error; err != nil {
		logrus.Fatalf("Failed to seed payments: %v", err)
	} else {
		logrus.Println("Payments seeder success")
	}
}
