package herrors

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Generic Errors
func NewInternalServer(cause error) error {
	return echo.NewHTTPError(http.StatusInternalServerError, "server error").SetInternal(cause)
}

func NewBadRequest(cause error) error {
	return echo.NewHTTPError(http.StatusBadRequest, "bad request").SetInternal(cause)
}

// domain errors
var (
	ErrExistingUser = echo.NewHTTPError(http.StatusForbidden, "User with that email already exists")
	ErrUserNotFound = echo.NewHTTPError(http.StatusNotFound, "User with that id does not exist")
)
