package service

import (
	"context"
	"errors"
	"time"
	"user-service/internal/models"
	"user-service/internal/repository"
	"user-service/pkg"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	RegisterUser(ctx context.Context, req models.RegisterInput) (*models.User, error)
	LoginUser(ctx context.Context, req *models.LoginInput) (*models.User, error)
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

	token, err := pkg.GenerateToken(user.ID.Hex(), u.jwtSecret)
	if err != nil {
		return nil, err
	}

	user.Token = token

	return user, nil
}

func NewUserService(userRepo repository.UserRepo, jwtSecret string) UserService {
	return &userService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}
