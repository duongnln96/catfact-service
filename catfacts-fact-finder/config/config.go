package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config struct
type Config struct {
	LocalPort     int    `mapstructure:"LOCAL_PORT"`
	LocalProtocal string `mapstructure:"LOCAL_PROTOCOL"`
	OfflineMode   bool   `mapstructure:"FINDFACT_OFFLINE"`
}

var values Config

func init() {
	config := viper.New()
	config.SetConfigName("config")
	config.AddConfigPath("./config/")
	config.AutomaticEnv()

	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Error while reading config: %s", err)
	}

	if err := config.Unmarshal(&values); err != nil {
		log.Fatalf("Error while parsing config: %s", err)
	}
}

// GetConfig returns config instance
func GetConfig() *Config {
	return &values
}
