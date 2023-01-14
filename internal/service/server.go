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
	"tagesTestTask/pkg/diskStorage"
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

	//utils
	storage, err := diskStorage.NewDiskStorage(s.pgDB, 2)
	if err != nil {
		return err
	}

	err = controllers.Catalog.InitCatalog(s.pgDB, storage, s.gSrv)
	if err != nil {
		return
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.cfg.Server.Host, s.cfg.Server.Port))
	if err != nil {
		fmt.Println("Error listening: ", err.Error())
		return err
	}
	go func() {
		if err := s.gSrv.Serve(lis); err != nil {
			fmt.Println("Error starting Server: ", err.Error())
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	return
}
