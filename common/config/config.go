package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	var err error
	viper.SetConfigName(".env.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("databases fatal error config file: %w", err))
	}
}

func GetString(key string) string {
	return viper.GetString(key)
}
