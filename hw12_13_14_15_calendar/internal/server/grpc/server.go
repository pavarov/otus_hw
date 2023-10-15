package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/config"
	server2 "github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/server"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/server/pb"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
)

type server struct {
	cfg               config.GrpcServerConfig
	grpcServerService pb.EventServiceServer
	grpcServer        *grpc.Server
}

func NewServer(
	cfg config.GrpcServerConfig,
	grpcServerService pb.EventServiceServer,
) server2.Interface {
	return &server{
		cfg:               cfg,
		grpcServerService: grpcServerService,
	}
}

func InterceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}

func (s *server) Run() error {
	port := strconv.Itoa(s.cfg.Port)
	lsn, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
	}
	logg := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{}))
	s.grpcServer = grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(InterceptorLogger(logg), opts...),
		),
	)

	pb.RegisterEventServiceServer(s.grpcServer, s.grpcServerService)

	logg.Info(fmt.Sprintf("grpc server started on %s", lsn.Addr().String()))
	if err := s.grpcServer.Serve(lsn); err != nil {
		log.Fatal(err)
	}
	return nil
}

func (s *server) Stop() error {
	if s.grpcServer != nil {
		s.grpcServer.GracefulStop()
	}
	return nil
}
