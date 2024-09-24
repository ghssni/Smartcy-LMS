package repository

import (
	"context"
	"github.com/ghssni/Smartcy-LMS/User-Service/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type UserRepo interface {
	Register(ctx context.Context, user *models.User) (*mongo.InsertOneResult, error)
	FindUserByEmail(ctx context.Context, email string) (*models.User, error)
	Login(ctx context.Context, id string) (*models.User, error)
	GetUserProfile(ctx context.Context, id string) (*models.User, error)
	ForgotPassword(ctx context.Context, email string) (*models.User, error)
	UpdateUserProfile(ctx context.Context, id string, user *models.User) (*models.User, error)
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

func (r *userRepo) ForgotPassword(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := r.db.Collection("users").FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}

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
}

func NewUserRepo(db *mongo.Database) UserRepo {
	return &userRepo{db: db}
}
