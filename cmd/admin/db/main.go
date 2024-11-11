package db

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
}

type Config struct {
	Database DBConfig `mapstructure:"postgres"`
}

func getPostgresConnectionString() (string, error) {
	cfg := &Config{}
	fmt.Println(viper.GetString("postgres.user"))
	err := viper.Unmarshal(cfg)
	if err != nil {
		fmt.Println("error: ", err)
		return "", err
	}
	fmt.Println("cfg: ", cfg.Database)
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.SSLMode,
	), nil
}

func ConnectToPostgres() *gorm.DB {
	postgres_uri, err := getPostgresConnectionString()
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(postgres_uri))
	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&Admin{}); err != nil {
		panic(err)
	}
	return db
}
