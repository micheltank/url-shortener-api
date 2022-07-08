package main

import (
	"github.com/micheltank/url-shortener-api/internal/infra/config"
	"github.com/micheltank/url-shortener-api/internal/port/grpc"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	appConfig, err := config.NewConfig()
	if err != nil {
		logrus.WithError(err).Fatalf("failed to read config")
	}
	setLogLevel(appConfig)

	err = run(appConfig)
	if err != nil {
		logrus.WithError(err).Fatal("failed running application")

		return
	}
}

func setLogLevel(appConfig config.Config) {
	logLevel, err := logrus.ParseLevel(appConfig.LogLevel)
	if err != nil {
		logrus.WithError(err).Warnf("failed to parse log level, using default 'info'")

		return
	}
	logrus.SetLevel(logLevel)
}

func run(appConfig config.Config) error {
	logrus.Info("Starting application")

	// GRPC Server
	grpcApiServer, err := grpc.NewServer(appConfig)
	if err != nil {
		return errors.Wrap(err, "failed to initialize grpcApiServer")
	}
	grpcApiErr := grpcApiServer.Run()
	logrus.Infof("Running grpc server on port %d", appConfig.GrpcPort)
	defer grpcApiServer.Shutdown()

	// Shutdown
	quit := notifyShutdown()
	select {
	case err := <-grpcApiErr:
		return errors.Wrap(err, "failed while running restApiServer")
	case <-quit:
		logrus.Info("Gracefully shutdown")
		return nil
	}
}

func notifyShutdown() chan os.Signal {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	return quit
}
