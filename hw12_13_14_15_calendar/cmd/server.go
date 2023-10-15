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
	grpc2 "github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/server/grpc"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/server/http/eventhandler"
	internalhttp "github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/server/http/middlewares"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/services"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/storage"
	memorystorage "github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/storage/sql"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/migrations"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func server(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "server calendar",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := serverRun(ctx)
			if err != nil {
				return err
			}
			return nil
		},
	}
}

func serverRun(ctx context.Context) error {
	appCtx, cancel := signal.NotifyContext(ctx, os.Interrupt)

	cfg := config.NewAppConfig()
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal("failed to load unmarshal configuration", "error", err)
	}

	logg := logger.New(cfg.LoggerConfig)

	var db storage.Interface
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

	service := services.NewEventService(db)
	grpcServerService := grpc2.NewService(service)
	GRPCServer := grpc2.NewServer(cfg.GrpcConfig, grpcServerService)

	HTTPServer := echo.New()
	HTTPServer.Use(internalhttp.LoggingMiddleware(logg))
	HTTPServer.HideBanner = true

	eventGroup := HTTPServer.Group("event")
	eventHandler := eventhandler.New(service, logg)
	eventGroup.GET("/list-by-date/:date", eventHandler.ListOnDate)
	eventGroup.GET("/list-by-week/:date", eventHandler.ListOnWeek)
	eventGroup.GET("/list-by-month/:date", eventHandler.ListOnMonth)
	eventGroup.POST("", eventHandler.Create)
	eventGroup.PUT("/:id", eventHandler.Update)
	eventGroup.DELETE("/:id", eventHandler.Delete)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		port := strconv.Itoa(cfg.ServerConfig.Port)
		if err := HTTPServer.Start(":" + port); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logg.Panic("failed to start http server", "error", err)
			cancel()
		}
		wg.Done()
	}()

	go func() {
		if err := GRPCServer.Run(); err != nil {
			logg.Panic("failed to start grpc server", "error", err)
			cancel()
		}
		wg.Done()
	}()

	go func() {
		<-appCtx.Done()
		_ = HTTPServer.Close()
		_ = GRPCServer.Stop()
	}()

	wg.Wait()
	logg.Info("shutdown!")
	return nil
}
