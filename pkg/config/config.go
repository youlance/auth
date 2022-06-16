package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	RedisAddr           string        `mapstructure:"REDIS_ADDR"`
	RedisPassword       string        `mapstructure:"REDIS_PASSWORD"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN"`
	VerifyServerAddr    string        `mapstructure:"VERIFY_SERVER_ADDR"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
