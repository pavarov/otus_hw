package logger

import (
	"os"
	"strings"

	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/config"
	"golang.org/x/exp/slog"
)

type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
	Panic(msg string, args ...any)
}

type SlogLogger struct {
	logger *slog.Logger
}

func New(cfg config.LoggerConfig) Logger {
	loggerLevel := slog.LevelDebug

	switch strings.ToLower(cfg.Level) {
	case "info":
		loggerLevel = slog.LevelInfo
	case "warn":
		loggerLevel = slog.LevelWarn
	case "error":
		loggerLevel = slog.LevelError
	}

	opts := &slog.HandlerOptions{
		Level: loggerLevel,
	}

	var h slog.Handler
	switch cfg.Format {
	case "json":
		h = slog.NewJSONHandler(os.Stdout, opts)
	default:
		h = slog.NewTextHandler(os.Stdout, opts)
	}
	logger := slog.New(h)

	return &SlogLogger{
		logger: logger,
	}
}

func (l *SlogLogger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

func (l *SlogLogger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

func (l *SlogLogger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

func (l *SlogLogger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}

func (l *SlogLogger) Panic(msg string, args ...any) {
	l.logger.Error(msg, args...)
}
