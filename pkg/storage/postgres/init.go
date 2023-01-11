package postgres

import (
	"context"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"log"
	"tagesTestTask/config"
	"time"

	"github.com/jmoiron/sqlx"
)

func InitPsqlDB(ctx context.Context, cfg *config.Config) (*sqlx.DB, error) {
	log.Println(fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.DBName,
	))
	//log.Println(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
	//	cfg.Postgres.Host,
	//	cfg.Postgres.Port,
	//	cfg.Postgres.User,
	//	cfg.Postgres.Password,
	//	cfg.Postgres.DBName,
	//	cfg.Postgres.SSLMode,
	//))
	connectionURL := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.DBName,
	)
	//connectionURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
	//	cfg.Postgres.Host,
	//	cfg.Postgres.Port,
	//	cfg.Postgres.User,
	//	cfg.Postgres.Password,
	//	cfg.Postgres.DBName,
	//	cfg.Postgres.SSLMode,
	//)

	database, err := sqlx.ConnectContext(ctx, cfg.Postgres.PGDriver, connectionURL)
	if err != nil {
		return nil, err
	}

	database.SetMaxOpenConns(cfg.Postgres.Settings.MaxOpenConns)
	database.SetConnMaxLifetime(cfg.Postgres.Settings.ConnMaxLifetime * time.Second)
	database.SetMaxIdleConns(cfg.Postgres.Settings.MaxIdleConns)
	database.SetConnMaxIdleTime(cfg.Postgres.Settings.ConnMaxIdleTime * time.Second)

	if err = database.Ping(); err != nil {
		return nil, err
	}

	return database, nil
}
