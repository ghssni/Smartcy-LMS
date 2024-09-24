package controller

import (
	"errors"
	"github.com/ghssni/Smartcy-LMS/pkg"
	"github.com/ghssni/Smartcy-LMS/user-service/internal/models"
	"github.com/ghssni/Smartcy-LMS/user-service/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) RegisterUser(ctx echo.Context) error {
	var req models.RegisterInput
	if err := ctx.Bind(&req); err != nil {
		return pkg.ResponseJson(ctx, http.StatusBadRequest, nil, "invalid request format : "+err.Error())
	}

	// validate request
	if err := req.Validate(); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			formatterErrors := pkg.FormatValidationError(&req, validationErrors)
			return pkg.ResponseJson(ctx, http.StatusBadRequest, nil, formatterErrors)
		}
		return pkg.ResponseJson(ctx, http.StatusBadRequest, nil, err.Error())
	}

	user, err := c.userService.RegisterUser(ctx.Request().Context(), req)
	if err != nil {
		return pkg.ResponseJson(ctx, http.StatusInternalServerError, nil, err.Error())
	}

	response := map[string]interface{}{
		"name":       user.Name,
		"email":      user.Email,
		"address":    user.Address,
		"role":       user.Role,
		"phone":      user.Phone,
		"age":        user.Age,
		"created_at": time.Now().Format("2006-01-02 15:04:05"),
		"updated_at": time.Now().Format("2006-01-02 15:04:05"),
	}

	return pkg.ResponseJson(ctx, http.StatusCreated, response, "user created successfully")
}

func (c *UserController) LoginUser(ctx echo.Context) error {
	var req models.LoginInput
	if err := ctx.Bind(&req); err != nil {
		return pkg.ResponseJson(ctx, http.StatusBadRequest, nil, "invalid request format : "+err.Error())
	}

	// validate request
	if err := req.Validate(); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			formatterErrors := pkg.FormatValidationError(&req, validationErrors)
			return pkg.ResponseJson(ctx, http.StatusBadRequest, nil, formatterErrors)
		}
		return pkg.ResponseJson(ctx, http.StatusBadRequest, nil, err.Error())
	}

	user, err := c.userService.LoginUser(ctx.Request().Context(), &req)
	if err != nil {
		return pkg.ResponseJson(ctx, http.StatusNotFound, nil, "invalid Email or password")
	}

	response := models.UserResponse{
		Name:      user.Name,
		Email:     user.Email,
		Address:   user.Address,
		Age:       user.Age,
		Role:      user.Role,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	response.UserToken.Token = user.Token

	return pkg.ResponseJson(ctx, http.StatusOK, response, "user login successfully")
}

// FindUserByEmail find user by email
func (c *UserController) FindUserByEmail(ctx echo.Context) error {
	email := ctx.Param("email")
	user, err := c.userService.FindUserByEmail(ctx.Request().Context(), email)
	if err != nil {
		return pkg.ResponseJson(ctx, http.StatusNotFound, nil, "user not found")
	}

	response := models.UserResponse{
		Email: user.Email,
	}

	return pkg.ResponseJson(ctx, http.StatusOK, response, "user found successfully")
}

func (c *UserController) RegisterActivity(ctx echo.Context) error {
	var req models.UserActivityLog
	if err := ctx.Bind(&req); err != nil {
		return pkg.ResponseJson(ctx, http.StatusBadRequest, nil, "invalid request format: "+err.Error())
	}

	// Extract the JWT token from the context
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	// Get the user ID from the token claims
	id, ok := claims["id"].(string)
	if !ok {
		return pkg.ResponseJson(ctx, http.StatusUnauthorized, nil, "invalid token: user ID not found")
	}

	// Set the UserID in the request from the token
	req.UserID = id

	// Call the service to register the activity
	err := c.userService.RegisterActivity(ctx.Request().Context(), req.UserID, req.ActivityType)
	if err != nil {
		return pkg.ResponseJson(ctx, http.StatusInternalServerError, nil, "failed to register activity: "+err.Error())
	}

	// Return a successful response
	return pkg.ResponseJson(ctx, http.StatusCreated, nil, "activity registered successfully")
}
