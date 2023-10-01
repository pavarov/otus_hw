package main

import (
	"context"

	_ "github.com/lib/pq"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/cmd"
)

func main() {
	ctx := context.Background()
	cmd.Execute(ctx)
}

// func main() {
//	flag.Parse()
//
//	if flag.Arg(0) == "version" {
//		printVersion()
//		return
//	}
//
//	config := NewConfig()
//	logg := logger.New(config.Logger.Level)
//
//	storage := memorystorage.New()
//	calendar := app.New(logg, storage)
//
//	server := internalhttp.NewServer(logg, calendar)
//
//	ctx, cancel := signal.NotifyContext(context.Background(),
//		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
//	defer cancel()
//
//	go func() {
//		<-ctx.Done()
//
//		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
//		defer cancel()
//
//		if err := server.Stop(ctx); err != nil {
//			logg.Error("failed to stop http server: " + err.Error())
//		}
//	}()
//
//	logg.Info("calendar is running...")
//
//	if err := server.Start(ctx); err != nil {
//		logg.Error("failed to start http server: " + err.Error())
//		cancel()
//		os.Exit(1) //nolint:gocritic
//	}
//}
