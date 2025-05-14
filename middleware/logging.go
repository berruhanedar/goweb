package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func LogrusMiddleware(logger *logrus.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			err := next(c)

			logger.WithFields(logrus.Fields{
				"method":  c.Request().Method,
				"path":    c.Request().URL.Path,
				"status":  c.Response().Status,
				"latency": time.Since(start),
				"ip":      c.RealIP(),
			}).Info("Request completed")

			return err
		}
	}
}
