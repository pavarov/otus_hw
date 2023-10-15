package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/logger"
)

func LoggingMiddleware(log logger.Logger) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(reqLoggerConfig(log))
}

func reqLoggerConfig(log logger.Logger) middleware.RequestLoggerConfig {
	return middleware.RequestLoggerConfig{
		LogRemoteIP: true,
		LogProtocol: true,
		LogHost:     true,
		LogMethod:   true,
		LogURI:      true,
		LogStatus:   true,
		LogError:    true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				log.Info(
					"request",
					"client IP", v.RemoteIP,
					"method", v.Method,
					"uri", v.URI,
					"protocol", v.Protocol,
					"status", v.Status,
					"latency", v.Latency,
					"user_agent", v.UserAgent,
				)
			} else {
				log.Info(
					"failed request",
					"client IP", v.RemoteIP,
					"method", v.Method,
					"uri", v.URI,
					"protocol", v.Protocol,
					"status", v.Status,
					"latency", v.Latency,
					"user_agent", v.UserAgent,
					"error", v.Error,
				)
			}
			return nil
		},
	}
}
