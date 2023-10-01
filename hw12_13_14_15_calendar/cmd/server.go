package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/config"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/server/http"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/storage"
	memorystorage "github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/storage/sql"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/migrations"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/exp/slog"
)

func httpServer(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "httpServer",
		Short: "http server calendar",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := httpServerRun(ctx)
			if err != nil {
				return err
			}
			return nil
		},
	}
}

func httpServerRun(ctx context.Context) error {
	appCtx, cancel := signal.NotifyContext(ctx, os.Interrupt)

	cfg := config.NewAppConfig()
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal("failed to load unmarshal configuration", "error", err)
	}

	logg := logger.New(cfg.LoggerConfig)

	var db storage.StoreInterface
	switch cfg.DBConfig.Type {
	case "db":
		sqlClient, err := sqlstorage.NewClient(cfg.DBConfig)
		if err != nil {
			return fmt.Errorf("failed to create db client: %w", err)
		}
		db = sqlstorage.New(sqlClient)
		if err := migrations.MigrationRun(sqlClient.Connection().DB); err != nil {
			return err
		}
	default:
		db = memorystorage.New()
	}

	server := echo.New()
	server.Use(internalhttp.LoggingMiddleware(logg))
	server.HideBanner = true
	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	list, err := db.List(ctx)
	if err != nil {
		return err
	}
	fmt.Sprintln(list)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		port := strconv.Itoa(cfg.ServerConfig.Port)
		if err := server.Start(":" + port); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("failed to start server", "error", err)
			cancel()
		}

		wg.Done()
	}()

	go func() {
		<-appCtx.Done()
		_ = server.Close()
	}()

	wg.Wait()
	logg.Info("shutdown!")
	return nil
}
