package handler

import (
	"context"
	"fmt"
	"gateway-service/config"
	"gateway-service/constans"
	"gateway-service/model"
	"gateway-service/pb"
	"gateway-service/server/middlewares"
	"gateway-service/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net/http"
	"time"
)

type UserHandler struct {
	// Add the user service client here
	userService  pb.UserServiceClient
	emailService pb.EmailServiceClient
}

func NewUserHandler(userService pb.UserServiceClient, emailService pb.EmailServiceClient) *UserHandler {
	return &UserHandler{
		userService:  userService,
		emailService: emailService,
	}
}

func (h *UserHandler) Register(c echo.Context) error {
	user := new(model.UserRequest)

	// Bind the user struct
	err := c.Bind(&user)
	if err != nil {
		return utils.HandleError(c, constans.ErrBadRequest, "Invalid request body")
	}

	// Validate the user struct
	err = config.Validator.Struct(user)
	if err != nil {
		// Format the validation errors
		errors := utils.FormatValidationErrors(err)
		return utils.HandleValidationError(c, errors)
	}

	// Get user by email
	req1 := pb.GetUserByEmailRequest{
		Email: user.Email,
	}

	_, err = h.userService.GetUserByEmail(c.Request().Context(), &req1)
	if err == nil {
		return utils.HandleError(c, constans.ErrConflict, "Email already registered")
	}

	createdAt := time.Now()
	updatedAt := time.Now()

	req2 := pb.RegisterRequest{
		RegisterInput: &pb.RegisterInput{
			Name:      user.Name,
			Email:     user.Email,
			Password:  user.Password,
			Address:   user.Address,
			Role:      user.Role,
			Phone:     user.Phone,
			Age:       user.Age,
			CreatedAt: timestamppb.New(createdAt),
			UpdatedAt: timestamppb.New(updatedAt),
		},
	}

	// Do the gRPC call
	res, err := h.userService.Register(c.Request().Context(), &req2)
	if err != nil {
		return utils.HandleError(c, constans.ErrInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, model.JsonResponse{
		Status:  constans.ResponseStatusSuccess,
		Message: "User registered",
		Data:    res.User,
	})

}

func (h *UserHandler) Login(c echo.Context) error {
	loginRequest := new(model.LoginRequest)

	// Bind the login request struct
	err := c.Bind(&loginRequest)
	if err != nil {
		return utils.HandleError(c, constans.ErrBadRequest, "Invalid request body")
	}

	// Validate the login request struct
	err = config.Validator.Struct(loginRequest)
	if err != nil {
		// Format the validation errors
		errors := utils.FormatValidationErrors(err)
		return utils.HandleValidationError(c, errors)
	}

	req := pb.LoginRequest{
		LoginInput: &pb.LoginInput{
			Email:    loginRequest.Email,
			Password: loginRequest.Password,
		},
	}

	// Do the gRPC call
	res, err := h.userService.Login(c.Request().Context(), &req)
	if err != nil {
		return utils.HandleError(c, constans.ErrNotFound, "Invalid email or password")
	}

	// Generate JWT token
	token, err := middlewares.GenerateToken(res.User.Id, res.User.Email, res.User.Role)
	if err != nil {
		return utils.HandleError(c, constans.ErrInternalServerError, "Failed to generate token")
	}

	userResponse := &model.User{
		ID:      res.User.Id,
		Name:    res.User.Name,
		Email:   res.User.Email,
		Address: res.User.Address,
		Role:    res.User.Role,
		Phone:   res.User.Phone,
		Age:     res.User.Age,
		Token:   token,
	}

	return c.JSON(http.StatusOK, model.JsonResponse{
		Status:  constans.ResponseStatusSuccess,
		Message: "User logged in",
		Data:    userResponse,
	})

}

func (h *UserHandler) ForgotPassword(c echo.Context) error {
	forgotPasswordRequest := new(pb.ForgotPasswordRequest)
	// Bind and validate the request
	if err := c.Bind(forgotPasswordRequest); err != nil {
		return utils.HandleError(c, constans.ErrBadRequest, "Invalid request body")
	}
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()
	user, err := h.userService.GetUserByEmail(c.Request().Context(), &pb.GetUserByEmailRequest{Email: forgotPasswordRequest.Email})
	if err != nil {
		grpcErr, ok := status.FromError(err)
		if ok && grpcErr.Code() == codes.NotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": grpcErr.Message()})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to find user"})
	}
	// Call gRPC service to find the user
	_, err = h.userService.ForgotPassword(ctx, forgotPasswordRequest)
	if err != nil {
		grpcErr, ok := status.FromError(err)
		if ok {
			switch grpcErr.Code() {
			case codes.NotFound:
				return c.JSON(http.StatusNotFound, map[string]string{"error": grpcErr.Message()})
			default:
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": grpcErr.Message()})
			}
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	//generate token
	token, err := middlewares.GenerateToken(user.User.Id, user.User.Email, user.User.Role)

	emailReq := &pb.SendForgotPasswordEmailRequest{
		Email:      forgotPasswordRequest.Email,
		UserId:     user.User.Id,
		ResetToken: token,
		ResetLink:  fmt.Sprintf("http://localhost:8080/user/reset-password?token=%s", token),
	}

	emailResp, err := h.emailService.SendForgotPasswordEmail(ctx, emailReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to send reset password email"})
	}

	if !emailResp.Success {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to send reset password email"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    emailResp.Meta.Code,
			"status":  emailResp.Meta.Status,
			"message": emailResp.Meta.Message,
		},
		"success": emailResp.Success,
	})
}

func (h *UserHandler) NewPassword(c echo.Context) error {
	newPasswordRequest := new(pb.NewPasswordRequest)

	token := c.QueryParam("token")
	if token == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Token is missing"})
	}

	claims, err := middlewares.VerifyToken(token)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
	}
	// Bind the reset password request struct
	err = c.Bind(newPasswordRequest)
	if err != nil {
		return utils.HandleError(c, constans.ErrBadRequest, "Invalid request body")
	}

	// Validate the reset password request struct
	err = config.Validator.Struct(newPasswordRequest)
	if err != nil {
		// Format the validation errors
		errors := utils.FormatValidationErrors(err)
		return utils.HandleValidationError(c, errors)
	}

	// set user id
	newPasswordRequest.UserId = claims.UserID

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// Do the gRPC call
	resp, err := h.userService.NewPassword(ctx, newPasswordRequest)
	if err != nil {
		utils.HandleError(c, constans.ErrBadRequest, err.Error())
		grpcErr, ok := status.FromError(err)
		if ok {
			switch grpcErr.Code() {
			case codes.InvalidArgument:
				return c.JSON(http.StatusBadRequest, map[string]string{"error": grpcErr.Message()})
			case codes.NotFound:
				return c.JSON(http.StatusNotFound, map[string]string{"error": grpcErr.Message()})
			case codes.Internal:
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": grpcErr.Message()})
			default:
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unexpected error"})
			}
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unexpected error"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    resp.Meta.Code,
			"status":  resp.Meta.Status,
			"message": resp.Meta.Message,
		},
	})
}

func (h *UserHandler) UpdateUserProfile(c echo.Context) error {
	// Bind the request body into the userProfile struct
	userProfile := new(model.UserProfileRequest)
	err := c.Bind(userProfile)
	if err != nil {
		return utils.HandleError(c, constans.ErrBadRequest, "Invalid request body")
	}

	// Validate the userProfile struct
	err = config.Validator.Struct(userProfile)
	if err != nil {
		errors := utils.FormatValidationErrors(err)
		return utils.HandleValidationError(c, errors)
	}

	// Extract user information from JWT token
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middlewares.JWTCustomClaims)
	userId := claims.UserID
	role := claims.Role
	log.Printf("UserID: %s, Role: %s", claims.UserID, claims.Role)
	if claims.Role == "" {
		return utils.HandleError(c, constans.ErrInternalServerError, "Role is missing in token")
	}

	req := pb.UpdateUserProfileRequest{
		UserId:    userId,
		Name:      userProfile.Name,
		Email:     userProfile.Email,
		Address:   userProfile.Address,
		Phone:     userProfile.Phone,
		Age:       uint32(userProfile.Age),
		Role:      role,
		UpdatedAt: timestamppb.New(time.Now()),
	}

	// Make the gRPC call to update the user profile
	res, err := h.userService.UpdateUserProfile(c.Request().Context(), &req)
	if err != nil {
		return utils.HandleError(c, constans.ErrInternalServerError, err.Error())
	}

	// Return the response to the client
	return c.JSON(http.StatusOK, model.JsonResponse{
		Status:  constans.ResponseStatusSuccess,
		Message: "User profile updated",
		Data:    res.UserProfile,
	})
}

func (h *UserHandler) GetUserProfile(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middlewares.JWTCustomClaims)

	userId := claims.UserID

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	userProfileResponse, err := h.userService.GetUserProfile(ctx, &pb.GetUserProfileRequest{
		UserId: userId,
	})
	if err != nil {
		st, _ := status.FromError(err)
		return c.JSON(http.StatusNotFound, map[string]string{"message": st.Message()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    userProfileResponse.Meta.Code,
			"status":  userProfileResponse.Meta.Status,
			"message": userProfileResponse.Meta.Message,
		},
		"user_profile": map[string]interface{}{
			"id":         userProfileResponse.UserProfile.Id,
			"name":       userProfileResponse.UserProfile.Name,
			"email":      userProfileResponse.UserProfile.Email,
			"phone":      userProfileResponse.UserProfile.Phone,
			"age":        userProfileResponse.UserProfile.Age,
			"address":    userProfileResponse.UserProfile.Address,
			"created_at": userProfileResponse.UserProfile.CreatedAt.AsTime().String(),
			"updated_at": userProfileResponse.UserProfile.UpdatedAt.AsTime().String(),
		},
	})
}
