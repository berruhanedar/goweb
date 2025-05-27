package app

import (
	"github.com/berruhanedar/goweb/internal/config"
	"github.com/berruhanedar/goweb/internal/routes"
	"github.com/berruhanedar/goweb/middleware"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
)

func Start() {
	e := echo.New()

	e.Validator = &Validator{validator: validator.New()}

	config.LoadConfig()
	config.GetMongoClient()

	logger := logrus.New()
	e.Use(middleware.LogrusMiddleware(logger))

	routes.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":" + config.Cfg.Port))
}

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
