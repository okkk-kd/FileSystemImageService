package main

import (
	"context"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"os/signal"
	"syscall"
	"tagesTestTask/config"
	"tagesTestTask/internal/service"
	"tagesTestTask/pkg/logger"
	"tagesTestTask/pkg/storage/postgres"
)

func main() {
	log.Default().Println("Service started")

	ctx := context.Background()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Default().Fatalf("Loading config: %s", err.Error())
		return
	}

	loggerDriver := logger.NewLogger(&cfg)
	err = loggerDriver.InitLogger()
	if err != nil {
		log.Default().Fatalf("Loading logger: %s", err.Error())
		return
	}

	psqlDB, err := postgres.InitPsqlDB(ctx, &cfg)
	if err != nil {
		log.Default().Fatalf("PostgreSQL init error: %s", err.Error())
		return
	}

	defer func(psqlDB *sqlx.DB) {
		err = psqlDB.Close()
		if err != nil {
			log.Default().Println(err.Error())
		} else {
			log.Default().Println("PostgreSQL closed properly")
		}
	}(psqlDB)

	s, err := service.NewServer(&ctx, &cfg, loggerDriver, psqlDB)
	if err != nil {
		return
	}

	err = s.RunServer()
	if err != nil {
		return
	}
	log.Default().Println("Service has been started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	if err != nil {
		log.Default().Fatalf(err.Error())
	} else {
		log.Default().Println("Fiber server exited properly")
	}
}
