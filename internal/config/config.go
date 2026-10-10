package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	RabbitMQ RabbitMQConfig `mapstructure:"rabbitmq"`
}

type RabbitMQConfig struct {
	URL string `mapstructure:"url"`
}

func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Warning: .env file not found, using environment variables")
	}

	bindEnvVars()

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	if cfg.RabbitMQ.URL == "" {
		return nil, fmt.Errorf("RABBITMQ_URL is required")
	}

	return &cfg, nil
}

func bindEnvVars() {
	viper.BindEnv("rabbitmq.url", "RABBITMQ_URL")
}
