package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config from ENV
type Config struct {
	LocalPort          int    `mapstructure:"LOCAL_PORT"`
	LocalProtocol      string `mapstructure:"LOCAL_PROTOCOL"`
	FactFinderHost     string `mapstructure:"FACT_FINDER_HOST"`
	FactFinderPort     int    `mapstructure:"FACT_FINDER_PORT"`
	FactFinderProtocol string `mapstructure:"FACT_FINDER_PROTOCOL"`
	FactFinderURI      string `mapstructure:"FACT_FINDER_URI"`
}

var values Config

func init() {
	config := viper.New()
	config.SetConfigName("config")
	config.AddConfigPath("./config")
	config.SetConfigType("env")
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
