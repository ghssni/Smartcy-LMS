package handler

import (
	"gateway-service/config"
	"gateway-service/constans"
	"gateway-service/model"
	"gateway-service/pb"
	"gateway-service/server/middleware"
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
	user := new(model.User)

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

	createdAt := time.Now()
	updatedAt := time.Now()

	req := pb.RegisterRequest{
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
	res, err := h.userService.Register(c.Request().Context(), &req)
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
		return utils.HandleError(c, constans.ErrInternalServerError, err.Error())
	}

	// Generate JWT token
	token, expiredAt, err := middleware.GenerateToken(utils.StringToUint(res.User.Id), res.User.Email, res.User.Role)
	if err != nil {
		return utils.HandleError(c, constans.ErrInternalServerError, err.Error())
	}

	jwtResponse := model.JWTResponse{
		Token:   token,
		Expires: expiredAt.Format(time.RFC3339),
	}

	userResponse := &model.User{
		ID:      res.User.Id,
		Name:    res.User.Name,
		Email:   res.User.Email,
		Address: res.User.Address,
		Role:    res.User.Role,
		Phone:   res.User.Phone,
		Age:     res.User.Age,
		Token:   jwtResponse,
	}

	return c.JSON(http.StatusOK, model.JsonResponse{
		Status:  constans.ResponseStatusSuccess,
		Message: "User logged in",
		Data:    userResponse,
	})

}
