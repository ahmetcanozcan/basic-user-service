package main

import (
	"app/common/config"
	"app/common/db"
	"app/common/herrors"
	"app/controller"
	"app/repository"
	"app/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.uber.org/fx"
)

func main() {

	app := fx.New(
		fx.Provide(
			// Configurations
			config.NewFlags,
			config.ReadConfiguration,
			config.SplitConfigs,

			// Database
			db.NewMongo,

			// Repositories
			repository.NewUserRepository,

			// Services
			service.NewUserService,

			// Controllers
			controller.NewUserController,

			newLogger,
			newEcho,
		),
		fx.Invoke(startServer),
	)

	app.Run()
}

func newEcho(logger *log.Logger) *echo.Echo {
	e := echo.New()
	e.Logger = logger

	e.HTTPErrorHandler = herrors.ErrorHandler

	return e
}

func newLogger() *log.Logger {
	l := log.New("user-service")

	return l
}

func startServer(e *echo.Echo, ct *controller.UserController) {
	e.GET("/health_check", controller.HealthCheck)

	controller.MakeUserHandler(e, ct)

	e.Start(":8080")
}
