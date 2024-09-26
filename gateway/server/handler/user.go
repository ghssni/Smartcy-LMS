package handler

import (
	"gateway-service/config"
	"gateway-service/constans"
	"gateway-service/model"
	"gateway-service/pb"
	"gateway-service/server/middlewares"
	"gateway-service/utils"
	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
	"time"
)

type UserHandler struct {
	// Add the user service client here
	userService pb.UserServiceClient
}

func NewUserHandler(userService pb.UserServiceClient) *UserHandler {
	return &UserHandler{
		userService: userService,
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
