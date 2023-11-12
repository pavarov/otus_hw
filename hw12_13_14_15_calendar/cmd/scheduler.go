package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/config"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/logger"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/rabbit"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/rabbit/models"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/services"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/storage"
	memorystorage "github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/storage/sql"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/migrations"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func schedulerCommand(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "calendar_scheduler",
		Short: "notification scheduler",
		RunE: func(cmd *cobra.Command, args []string) error {
			return schedulerRun(ctx)
		},
	}
}

func schedulerRun(ctx context.Context) error {
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

	producer := rabbit.NewProducer(cfg.RabbitConfig)
	err := producer.Connect()
	if err != nil {
		logg.Panic(fmt.Sprintf("failed to create rabbit connection. Error: %s", err))
	}
	defer func() {
		if err := producer.Close(); err != nil {
			logg.Panic(fmt.Sprintf("failed to close producer connection. Error: %s", err))
		}
	}()

	service := services.NewEventService(db)

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	go schedulerWorker(ctx, service, producer, logg)
	go cleanerWorker(ctx, service, logg)

	<-ctx.Done()
	return nil
}

func schedulerWorker(
	ctx context.Context,
	eventProvider services.EventServiceInterface,
	producer *rabbit.Producer,
	logger logger.Logger,
) {
	for {
	start:
		select {
		case <-ctx.Done():
			logger.Info("producer: got DONE from context")
			return
		case <-time.Tick(time.Second * 3):
			events, err := eventProvider.ListToNotify(ctx)
			if err != nil {
				logger.Error("producer: failed to get list for notify. Error: ", err)
				break start
			}

			for _, event := range events {
				notification := models.Notification{
					ID:     event.ID,
					Title:  event.Title,
					Start:  event.Start,
					UserID: event.UserID,
				}

				data, errM := json.Marshal(notification)
				if errM != nil {
					logger.Error(fmt.Sprintf("producer: failed to marshal event. Error: %s", errM))
				}

				errP := producer.Publish(ctx, data)
				if errP != nil {
					logger.Error(fmt.Sprintf("producer: failed to publish notification. Error: %s", errM))
				}

				logger.Debug(fmt.Sprintf("producer: published notification. Data: %s", data))
			}
		}
	}
}

func cleanerWorker(
	ctx context.Context,
	eventProvider services.EventServiceInterface,
	logger logger.Logger,
) {
	for {
	start:
		select {
		case <-ctx.Done():
			return
		case <-time.Tick(time.Second * 5):
			err := eventProvider.RemoveOld(ctx, time.Now().AddDate(-1, 0, 0))
			if err != nil {
				logger.Error(fmt.Sprintf("failed to remove old notifications. Error: %s", err))
				break start
			}
		}
	}
}
