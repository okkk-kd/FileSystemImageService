package config

import (
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Server
	Postgres
	Logger
}

type Server struct {
	Host                        string `validate:"required"`
	Port                        string `validate:"required"`
	ShowUnknownErrorsInResponse bool
	APIRateLimiter              struct {
		Max        int           `validate:"required,min=1"`
		Expiration time.Duration `validate:"required,min=1"`
	}
}

type Postgres struct {
	Host     string `validate:"required"`
	Port     string `validate:"required"`
	User     string `validate:"required"`
	Password string `validate:"required"`
	DBName   string `validate:"required"`
	SSLMode  string `validate:"required"`
	PGDriver string `validate:"required"`
	Settings struct {
		MaxOpenConns    int           `validate:"required,min=1"`
		ConnMaxLifetime time.Duration `validate:"required,min=1"`
		MaxIdleConns    int           `validate:"required,min=1"`
		ConnMaxIdleTime time.Duration `validate:"required,min=1"`
	}
}

type Logger struct {
	Level zerolog.Level `validate:"required"`
	File  string
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./config/")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&c)
	if err != nil {
		return
	}
	return
}
