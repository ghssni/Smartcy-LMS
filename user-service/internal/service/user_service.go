package service

import (
	"context"
	"errors"
	"github.com/ghssni/Smartcy-LMS/pkg"
	"github.com/ghssni/Smartcy-LMS/user-service/internal/models"
	"github.com/ghssni/Smartcy-LMS/user-service/internal/repository"
	"github.com/labstack/gommon/log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	RegisterUser(ctx context.Context, req models.RegisterInput) (*models.User, error)
	LoginUser(ctx context.Context, req *models.LoginInput) (*models.User, error)
	FindUserByEmail(ctx context.Context, email string) (*models.User, error)
	RegisterActivity(ctx context.Context, userId string, activityType string) error
}

type userService struct {
	userRepo  repository.UserRepo
	jwtSecret string
}

func (u *userService) RegisterUser(ctx context.Context, req models.RegisterInput) (*models.User, error) {
	hashPassword, err := pkg.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	objectId := primitive.NewObjectID()
	user := models.User{
		ID:        objectId,
		Email:     req.Email,
		Password:  hashPassword,
		Address:   req.Address,
		Name:      req.Name,
		Role:      req.Role,
		Age:       req.Age,
		Phone:     req.Phone,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	_, err = u.userRepo.SaveUser(ctx, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userService) LoginUser(ctx context.Context, req *models.LoginInput) (*models.User, error) {
	user, err := u.userRepo.FindUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if !pkg.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	//email, err := u.FindUserByEmail(ctx, req.Email)
	//if err != nil {
	//	return nil, err
	//}
	token, err := pkg.GenerateToken(user.ID.Hex(), user.Email, u.jwtSecret)
	if err != nil {
		return nil, err
	}
	user.Token = token

	// Update user token
	//err = u.userRepo.UpdateUser(ctx, user)
	//if err != nil {
	//	return nil, err
	//}
	// Register user activity
	err = u.RegisterActivity(ctx, user.ID.Hex(), "login")
	if err != nil {
		log.Printf("Error inserting user activity: %v", err)
	}

	return user, nil
}

func (u *userService) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := u.userRepo.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userService) RegisterActivity(ctx context.Context, userId string, activityType string) error {
	objectId := primitive.NewObjectID()
	activity := models.UserActivityLog{
		ID:                objectId,
		UserID:            userId,
		ActivityType:      activityType,
		ActivityTimestamp: time.Now().Format("2006-01-02 15:04:05"),
	}

	_, err := u.userRepo.LogUserActivity(ctx, &activity)
	if err != nil {
		return err
	}

	return nil
}

func NewUserService(userRepo repository.UserRepo, jwtSecret string) UserService {
	return &userService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}
