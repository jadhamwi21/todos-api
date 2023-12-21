package config

import "github.com/spf13/viper"

func SetupConfig() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
}
