package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
}

func New() (*Config, error) {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	dbUser := viper.Get("APP_DB_USERNAME").(string)
	dbPassword := viper.Get("APP_DB_PASSWORD").(string)
	dbName := viper.Get("APP_DB_NAME").(string)
	config := &Config{
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBName:     dbName,
	}
	return config, nil
}
