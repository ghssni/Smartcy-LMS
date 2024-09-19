package migrations

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// Migration function for create_user_activity_log_collection
func createUserActivityLogCollectionMigration(database *mongo.Database, indexField string) *Migration {
	return &Migration{
		ID: "20240919001742_create_user_activity_log_collection",
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

			logrus.Printf("Migration: %s completed. Index created on field: %s", "create_user_activity_log_collection", indexField)
			return nil
		},
		Rollback: func() error {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			err := database.Collection("user_activity_log").Drop(ctx)
			if err != nil {
				return err
			}

			logrus.Printf("Rollback: %s completed", "create_user_activity_log_collection")
			return nil
		},
	}
}
