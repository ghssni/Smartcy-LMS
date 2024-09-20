package seeder

import (
	"context"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/ghssni/Smartcy-LMS/pkg"
	"github.com/ghssni/Smartcy-LMS/user-service/internal/models"
	"github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func SeedUsers(db *mongo.Database) {
	userCollection := db.Collection("users")

	var users []interface{}

	for i := 1; i <= 15; i++ {
		gofakeit.Seed(0)
		password, _ := pkg.HashPassword("test12345")

		user := models.User{
			Name:      gofakeit.Name(),
			Email:     gofakeit.Email(),
			Address:   gofakeit.Address().Address,
			Age:       gofakeit.Number(18, 60),
			Phone:     gofakeit.Phone(),
			Role:      gofakeit.RandomString([]string{"student", "instructor"}),
			Password:  password,
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
			UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		users = append(users, user)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := userCollection.InsertMany(ctx, users, options.InsertMany().SetOrdered(false))
	if err != nil {
		logrus.Fatalf("Seed users failed: %v", err)
		return
	}

	logrus.Println("Seed users success")
}
