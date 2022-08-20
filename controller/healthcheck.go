package controller

import "github.com/labstack/echo/v4"

func HealthCheck(c echo.Context) error {
	c.Logger().Info("Sa")
	return c.JSON(200, "OK")
}
