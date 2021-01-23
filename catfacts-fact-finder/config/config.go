package config

import (
	"log"

	ff "github.com/duongnln96/catfact-service/catfacts-fact-finder/factfinder"
	"github.com/spf13/viper"
)

// Config class
type Config struct {
	Metadata ff.Meta       `mapstructure:"Metadata"`
	Cfg      ff.CoreConfig `mapstructure:"Config"`
}

var values Config

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
