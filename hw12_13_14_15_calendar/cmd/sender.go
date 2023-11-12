package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/config"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/logger"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/rabbit"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func senderCommand(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "calendar_sender",
		Short: "notification sender",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := senderRun(ctx)
			if err != nil {
				return err
			}
			return nil
		},
	}
}

func senderRun(ctx context.Context) error {
	cfg := config.NewAppConfig()
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal("failed to load unmarshal configuration", "error", err)
	}

	logg := logger.New(cfg.LoggerConfig)

	consumer := rabbit.NewConsumer(cfg.RabbitConfig)
	err := consumer.Connect()
	if err != nil {
		logg.Panic(fmt.Sprintf("failed to create rabbit connection. Error: %s", err))
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			logg.Panic(fmt.Sprintf("failed to close consumer connection. Error: %s", err))
		}
	}()

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	go senderWorker(ctx, consumer, logg)

	<-ctx.Done()

	return nil
}

func senderWorker(
	ctx context.Context,
	consumer *rabbit.Consumer,
	logger logger.Logger,
) {
	for {
	start:
		select {
		case <-ctx.Done():
			logger.Info("consumer: got DONE from context")
			return
		case <-time.Tick(time.Second * 3):
			events, err := consumer.Consume()
			if err != nil {
				logger.Error("consumer: failed to get list for notify. Error: ", err)
				break start
			}

			for e := range events {
				logger.Info("sender: new notification.", "data", string(e.Body))
			}
		}
	}
}
