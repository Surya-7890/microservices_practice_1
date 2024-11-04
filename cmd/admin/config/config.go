package config

import "github.com/spf13/viper"

func LoadConfig() {
	viper.SetConfigFile("config.yml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}