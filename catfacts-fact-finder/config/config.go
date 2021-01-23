package config

import (
	"fmt"
	"log"

	"github.com/duongnln96/catfact-service/catfacts-fact-finder/factfinder"
	"github.com/spf13/viper"
)

// Config class
type Config struct {
	ServiceInfo factfinder.ServiceInfo `mapstructure:"metadata"`
	CoreConfig  factfinder.CoreConfig  `mapstructure:"core-config"`
}

var values Config

// GetServiceInfo method
func (cfg *Config) GetServiceInfo() string {
	return fmt.Sprintf("Service %+v", cfg.ServiceInfo)
}

func init() {
	config := viper.New()
	config.SetConfigName("config")
	config.AddConfigPath("./config/")

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
