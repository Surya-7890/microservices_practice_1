package config

import "github.com/spf13/viper"

func LoadConfig() *Application {
	viper.SetConfigFile("config.yml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	app := &Application{}
	err = viper.Unmarshal(app)
	if err != nil {
		panic(err)
	}

	return app
}
