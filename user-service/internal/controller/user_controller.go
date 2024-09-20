package controller

import (
	"errors"
	"github.com/ghssni/Smartcy-LMS/pkg"
	"github.com/ghssni/Smartcy-LMS/user-service/internal/models"
	"github.com/ghssni/Smartcy-LMS/user-service/internal/service"
	"github.com/go-playground/validator/v10"
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

	//return pkg.ResponseJson(ctx, http.StatusOK, response, "user login successfully")
	return pkg.ResponseJson(ctx, http.StatusOK, response, "user login successfully")
}
