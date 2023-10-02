package logger

import (
	"os"
	"strings"

	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/config"
	"golang.org/x/exp/slog"
)

const LevelPanic = slog.Level(12)

var LevelNames = map[slog.Level]string{
	LevelPanic: "PANIC",
}

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
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				levelLabel, exists := LevelNames[level]
				if !exists {
					levelLabel = level.String()
				}
				a.Value = slog.StringValue(levelLabel)
			}
			return a
		},
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
	l.logger.Log(nil, LevelPanic, msg, args...)
}
