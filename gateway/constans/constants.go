package constans

import (
	"gateway-service/utils"
	"net/http"
)

const (
	ResponseStatusFailed  = "Failed"
	ResponseStatusSuccess = "Success"
)

var (
	ErrNotFound            = utils.NewAPIError(http.StatusNotFound, "Resource not found", nil)
	ErrBadRequest          = utils.NewAPIError(http.StatusBadRequest, "Invalid request data", nil)
	ErrInternalServerError = utils.NewAPIError(http.StatusInternalServerError, "Internal Server Error", "We have encountered an error, please try again later")
	ErrUnauthorized        = utils.NewAPIError(http.StatusUnauthorized, "Unauthorized access", nil)
	ErrConflict            = utils.NewAPIError(http.StatusConflict, "Resource already exists", nil)
	ErrForbidden           = utils.NewAPIError(http.StatusForbidden, "Forbidden access", nil)
)

type UserRoleEnum string

const (
	UserRoleReader UserRoleEnum = "reader"
	UserRoleAdmin  UserRoleEnum = "admin"
)
