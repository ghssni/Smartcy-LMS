package service

import (
	"context"
	"github.com/ghssni/Smartcy-LMS/User-Service/internal/models"
	"github.com/ghssni/Smartcy-LMS/User-Service/internal/repository"
	pb "github.com/ghssni/Smartcy-LMS/User-Service/pb/proto"
	"github.com/ghssni/Smartcy-LMS/User-Service/pkg"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
	"time"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	userRepo        repository.UserRepo
	activityLogRepo repository.UserActivityLogRepo
}

func NewUserService(userRepo repository.UserRepo, activityLogRepo repository.UserActivityLogRepo) *UserService {
	return &UserService{
		userRepo:        userRepo,
		activityLogRepo: activityLogRepo,
	}
}

func (s *UserService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	hashedPassword, err := pkg.HashPassword(req.RegisterInput.Password)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Error hashing password: %v", err)
	}

	user := &models.User{
		ID:        primitive.NewObjectID(),
		Name:      req.RegisterInput.Name,
		Email:     req.RegisterInput.Email,
		Password:  hashedPassword,
		Address:   req.RegisterInput.Address,
		Role:      req.RegisterInput.Role,
		Phone:     req.RegisterInput.Phone,
		Age:       req.RegisterInput.Age,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	_, err = s.userRepo.Register(ctx, user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error registering user: %v", err)
	}

	response := &pb.RegisterResponse{
		Meta: &pb.Meta{
			Code:    int32(codes.OK),
			Status:  http.StatusText(http.StatusOK),
			Message: "User registered successfully",
		},
		User: &pb.User{
			Id:        user.ID.Hex(),
			Name:      user.Name,
			Email:     user.Email,
			Address:   user.Address,
			Role:      user.Role,
			Phone:     user.Phone,
			Age:       user.Age,
			CreatedAt: timestamppb.New(user.CreatedAt.Time()),
			UpdatedAt: timestamppb.New(user.UpdatedAt.Time()),
		},
	}

	return response, nil
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := s.userRepo.Login(ctx, req.LoginInput.Email)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "User not found: %v", err)
	}

	if !pkg.CheckPasswordHash(req.LoginInput.Password, user.Password) {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid password")
	}

	log := &models.UserActivityLog{
		UserID:            user.ID,
		ActivityType:      "login",
		ActivityTimestamp: primitive.NewDateTimeFromTime(time.Now()),
	}

	_, err = s.activityLogRepo.CreateUserActivityLog(ctx, log)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error creating user activity log: %v", err)
	}

	response := &pb.LoginResponse{
		Meta: &pb.Meta{
			Code:    int32(codes.OK),
			Status:  http.StatusText(http.StatusOK),
			Message: "Login successful",
		},
		User: &pb.User{
			Id:        user.ID.Hex(),
			Name:      user.Name,
			Email:     user.Email,
			Address:   user.Address,
			Role:      user.Role,
			Phone:     user.Phone,
			Age:       user.Age,
			CreatedAt: timestamppb.New(user.CreatedAt.Time()),
			UpdatedAt: timestamppb.New(user.UpdatedAt.Time()),
		},
	}
	return response, nil
}

func (s *UserService) GetUserProfile(ctx context.Context, req *pb.GetUserProfileRequest) (*pb.GetUserProfileResponse, error) {
	user, err := s.userRepo.GetUserProfile(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "User not found: %v", err)
	}

	response := &pb.GetUserProfileResponse{
		Meta: &pb.Meta{
			Code:    int32(codes.OK),
			Status:  http.StatusText(http.StatusOK),
			Message: "User profile retrieved successfully",
		},
		UserProfile: &pb.UserProfile{
			Id:        user.ID.Hex(),
			Name:      user.Name,
			Email:     user.Email,
			Phone:     user.Phone,
			Age:       user.Age,
			Address:   user.Address,
			CreatedAt: timestamppb.New(user.CreatedAt.Time()),
			UpdatedAt: timestamppb.New(user.UpdatedAt.Time()),
		},
	}
	return response, nil
}

func (s *UserService) ForgotPassword(ctx context.Context, req *pb.ForgotPasswordRequest) (*pb.ForgotPasswordResponse, error) {
	_, err := s.userRepo.ForgotPassword(ctx, req.Email)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "User not found: %v", err)
	}

	response := &pb.ForgotPasswordResponse{
		Meta: &pb.Meta{
			Code:    int32(codes.OK),
			Status:  http.StatusText(http.StatusOK),
			Message: "Password reset link sent to email",
		},
	}
	return response, nil
}

func (s *UserService) UpdateUserProfile(ctx context.Context, req *pb.UpdateUserProfileRequest) (*pb.UpdateUserProfileResponse, error) {
	user := &models.User{
		Name:      req.Name,
		Email:     req.Email,
		Address:   req.Address,
		Phone:     req.Phone,
		Age:       req.Age,
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	updatedUser, err := s.userRepo.UpdateUserProfile(ctx, req.UserId, user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error updating user profile: %v", err)
	}

	response := &pb.UpdateUserProfileResponse{
		Meta: &pb.Meta{
			Code:    int32(codes.OK),
			Status:  http.StatusText(http.StatusOK),
			Message: "User profile updated successfully",
		},
		UserProfile: &pb.UserProfile{
			Id:        req.UserId,
			Name:      updatedUser.Name,
			Email:     updatedUser.Email,
			Phone:     updatedUser.Phone,
			Age:       updatedUser.Age,
			Address:   updatedUser.Address,
			CreatedAt: timestamppb.New(updatedUser.CreatedAt.Time()),
			UpdatedAt: timestamppb.New(time.Now()),
		},
	}
	return response, nil
}
