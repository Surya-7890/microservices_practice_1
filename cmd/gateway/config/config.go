package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func LoadConfig() *Application {
	viper.SetConfigFile("config.yml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	app := &Application{}

	err = viper.Unmarshal(app)
	if err != nil {
		panic(err)
	}

	fmt.Println(app.KafkaConfig)

	return app
}
