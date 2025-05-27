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

			entry := logger.WithFields(logrus.Fields{
				"method":  c.Request().Method,
				"path":    c.Request().URL.Path,
				"status":  c.Response().Status,
				"latency": time.Since(start),
				"ip":      c.RealIP(),
			})

			if err != nil || c.Response().Status >= 400 {
				entry.Error("Request completed with error or client/server error")
			} else {
				entry.Info("Request completed successfully")
			}

			return err
		}
	}
}
