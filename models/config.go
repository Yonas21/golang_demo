package models

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	DB_TYPE     string `mapstructure:"DB_TYPE"`
	DB_HOST     string `mapstructure:"DB_HOST"`
	DB_PORT     int    `mapstructure:"DB_PORT"`
	DB_USER     string `mapstructure:"DB_USER"`
	DB_NAME     string `mapstructure:"DB_NAME"`
	DB_PREFIX   string `mapstructure:"DB_PREFIX"`
	DB_PASSWORD string `mapstructure:"DB_PASSWORD"`
}

func getDbConfig(path string) (config Config, err error) {

	// Set the file name of the configurations file
	// viper.SetConfigName(".env")

	viper.SetConfigFile(path)

	// Set the path to look for the configurations file
	// viper.AddConfigPath("../")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	// viper.SetConfigType("yml")

	err = viper.ReadInConfig()

	if err != nil {
		return Config{}, errors.WithMessage(err, "Error loading .env file")
	}

	err = viper.Unmarshal(&config)

	if err != nil {
		return Config{}, errors.WithMessage(err, "Invalid type assertion")
	}

	return config, nil
}
