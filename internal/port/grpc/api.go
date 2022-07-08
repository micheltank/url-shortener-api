package grpc

import (
	"database/sql"
	"fmt"
	"github.com/micheltank/url-shortener-api/internal/application"
	"github.com/micheltank/url-shortener-api/internal/infra/config"
	"github.com/micheltank/url-shortener-api/internal/infra/repository"
	"github.com/micheltank/url-shortener-api/internal/port/grpc/handler"
	"github.com/micheltank/url-shortener-api/internal/port/grpc/pb"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"sync"
)

type Api struct {
	Db       *sql.DB
	server   *grpc.Server
	listener net.Listener
}

func NewServer(config config.Config) (*Api, error) {
	//dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.DbConfig.User, config.DbConfig.Password, config.DbConfig.Host, config.DbConfig.Port, config.DbConfig.Name)
	//db, err := sql.Open("postgres", dataSourceName)
	//if err != nil {
	//	return nil, errors.Wrap(err, "failed to open postgres connection")
	//}

	// di
	urlRepository := repository.NewUrlCassandraRepository()
	urlService := application.NewService(urlRepository)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GrpcPort))
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute grpc listener")
	}
	grpcServer := grpc.NewServer()

	// handlers
	pb.RegisterUrlHandlerServer(grpcServer, handler.NewUrlHandler(urlService))

	return &Api{
		server:   grpcServer,
		listener: listener,
		//Db:       db,
	}, nil
}

func (api *Api) Run() <-chan error {
	out := make(chan error)
	go func() {
		if err := api.server.Serve(api.listener); err != nil {
			out <- errors.Wrap(err, "failed to listen and serve api")
		}
	}()
	return out
}

func (api *Api) Shutdown() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		api.server.GracefulStop()
	}()
	err := api.Db.Close()
	if err != nil {
		logrus.WithError(err).Error("Failed to close db")
	}
}
