package repository

import (
	"context"
	"github.com/ghssni/Smartcy-LMS/user-service/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UserRepo interface {
	SaveUser(ctx context.Context, user *models.User) (*mongo.InsertOneResult, error)
	FindUserByEmail(ctx context.Context, email string) (*models.User, error)
	FindUserByID(ctx context.Context, id string) (*models.User, error)
	LogUserActivity(ctx context.Context, activity *models.UserActivityLog) (*mongo.InsertOneResult, error)
	FindUserActivityByUserID(ctx context.Context, userID string) ([]models.UserActivityLog, error)
}

type userRepo struct {
	db *mongo.Database
}

func (r *userRepo) SaveUser(ctx context.Context, user *models.User) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	result, err := r.db.Collection("users").InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *userRepo) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := r.db.Collection("users").FindOne(ctx, map[string]string{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) FindUserByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := r.db.Collection("users").FindOne(ctx, map[string]string{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) LogUserActivity(ctx context.Context, activity *models.UserActivityLog) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	result, err := r.db.Collection("user_activity_log").InsertOne(ctx, activity)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *userRepo) FindUserActivityByUserID(ctx context.Context, userID string) ([]models.UserActivityLog, error) {
	var activities []models.UserActivityLog
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cursor, err := r.db.Collection("user_activity_log").Find(ctx, map[string]string{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var activity models.UserActivityLog
		err := cursor.Decode(&activity)
		if err != nil {
			return nil, err
		}

		activities = append(activities, activity)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return activities, nil
}

func NewUserRepo(db *mongo.Database) UserRepo {
	return &userRepo{db: db}
}
