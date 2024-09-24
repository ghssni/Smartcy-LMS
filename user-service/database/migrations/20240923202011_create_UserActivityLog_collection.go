package migrations

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// Migrations is a struct to define migration
type Migrations struct {
	ID       string
	Migrate  func() error
	Rollback func() error
}

// Migration function for create_UserActivityLog_collection
func createUserActivityLogCollectionMigration(database *mongo.Database, indexField string) *Migration {
	return &Migration{
		ID: "20240923202011_create_UserActivityLog_collection",
		Migrate: func() error {
			collection := database.Collection("user_activity_log")
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			indexOptions := options.Index().SetUnique(true)
			indexModel := mongo.IndexModel{
				Keys:    bson.M{indexField: 1},
				Options: indexOptions,
			}

			_, err := collection.Indexes().CreateOne(ctx, indexModel)
			if err != nil {
				return err
			}

			logrus.Printf("Migration: %s completed. Index created on field: %s", "create_UserActivityLog_collection", indexField)
			return nil
		},
		Rollback: func() error {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			err := database.Collection("user_activity_log").Drop(ctx)
			if err != nil {
				return err
			}

			logrus.Printf("Rollback: %s completed", "create_UserActivityLog_collection")
			return nil
		},
	}
}
