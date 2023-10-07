package util

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application
type Config struct {
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBSource            string        `mapstructure:"DB_SOURCE"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	// DBDriver            string `mapstructure:"dbDriver"`
	// DBSource            string `mapstructure:"dbSource"`
	// ServerAddress       string `mapstructure:"serverAddress"`
	// Pippo               string `mapstructure:"someOther"`
	// TokenSymmetricKey   string `mapstructure:"tokenSymmetricKey"`
	// AccessTokenDuration string `mapstructure:"accessTokenDuration"`
}

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	viper.SetEnvPrefix("SB")

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
