package config

import "github.com/spf13/viper"

type Config struct {
	Host string `validate:"required"`
	Port string `validate:"required"`
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
