package repository

import (
	"context"
<<<<<<< HEAD
	"fmt"
	"github.com/ghssni/Smartcy-LMS/user-service/internal/models"
	"go.mongodb.org/mongo-driver/bson"
=======
	"github.com/ghssni/Smartcy-LMS/User-Service/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
>>>>>>> user-service
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type UserRepo interface {
	Register(ctx context.Context, user *models.User) (*mongo.InsertOneResult, error)
	FindUserByEmail(ctx context.Context, email string) (*models.User, error)
<<<<<<< HEAD
	FindUserByID(ctx context.Context, id string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	LogUserActivity(ctx context.Context, activity *models.UserActivityLog) (*mongo.InsertOneResult, error)
	FindUserActivityByUserID(ctx context.Context, userID string) ([]models.UserActivityLog, error)
=======
	Login(ctx context.Context, id string) (*models.User, error)
	GetUserProfile(ctx context.Context, id string) (*models.User, error)
	ForgotPassword(ctx context.Context, email string) (*models.User, error)
	UpdateUserProfile(ctx context.Context, id string, user *models.User) (*models.User, error)
>>>>>>> user-service
}

type userRepo struct {
	db *mongo.Database
}

func (r *userRepo) Register(ctx context.Context, user *models.User) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
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

	err := r.db.Collection("users").FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) Login(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := r.db.Collection("users").FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) GetUserProfile(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)

	err = r.db.Collection("users").FindOne(ctx, bson.M{"_id": objectId}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

<<<<<<< HEAD
func (r *userRepo) UpdateUser(ctx context.Context, user *models.User) error {
	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": bson.M{"token": user.Token}}

	res, err := r.db.Collection("users").UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed update user: %v", err)
	}

	if res.MatchedCount == 0 {
		return fmt.Errorf("user not found for ID: %v", user.ID)
	}

	if res.ModifiedCount == 0 {
		return fmt.Errorf("token not updated!: %v", user.ID)
	}

	return nil
}

func (r *userRepo) LogUserActivity(ctx context.Context, activity *models.UserActivityLog) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	result, err := r.db.Collection("user_activity_log").InsertOne(ctx, activity)
=======
func (r *userRepo) ForgotPassword(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := r.db.Collection("users").FindOne(ctx, bson.M{"email": email}).Decode(&user)
>>>>>>> user-service
	if err != nil {
		return nil, err
	}

<<<<<<< HEAD
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
=======
	return &user, nil
}

func (r *userRepo) UpdateUserProfile(ctx context.Context, id string, user *models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	update := bson.M{
		"$set": bson.M{
			"name":       user.Name,
			"email":      user.Email,
			"address":    user.Address,
			"role":       user.Role,
			"phone":      user.Phone,
			"age":        user.Age,
			"updated_at": time.Now(),
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	updatedUser := &models.User{}
	err = r.db.Collection("users").FindOneAndUpdate(ctx, bson.M{"_id": objectID}, update, opts).Decode(updatedUser)

	if err != nil {
		return nil, err
	}

	return updatedUser, nil
>>>>>>> user-service
}

func NewUserRepo(db *mongo.Database) UserRepo {
	return &userRepo{db: db}
}
