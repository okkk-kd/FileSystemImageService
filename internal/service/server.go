package service

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"tagesTestTask/config"
	"tagesTestTask/pkg/logger"
)

type server struct {
	gSrv   *grpc.Server
	ctx    *context.Context
	cfg    *config.Config
	logger logger.Logger
	pgDB   *sqlx.DB
}

type Server interface {
	RunServer() (err error)
}

func NewServer(
	ctx *context.Context,
	cfg *config.Config,
	logger logger.Logger,
	database *sqlx.DB,
) (obj Server, err error) {
	return &server{
		gSrv:   grpc.NewServer(),
		ctx:    ctx,
		cfg:    cfg,
		logger: logger,
		pgDB:   database,
	}, err
}

func (s *server) RunServer() (err error) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.cfg.Host, s.cfg.Port))
	if err := s.gSrv.Serve(lis); err != nil {
		fmt.Println("Error starting Server: ", err.Error())
		return
	}

	newService, err := NewService()
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	//controllers
	controllers, err := newService.InitService(s.cfg, s.logger, s.pgDB)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	err = controllers.Catalog.InitCatalog(s.gSrv)
	if err != nil {
		return
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	return
}
