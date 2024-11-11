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
	DB DBConfig `mapstructure:"postgres"`
}

func getPostgresConnectionString() (string, error) {
	cfg := &Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		return "", nil
	}
	fmt.Println(cfg)
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.DBName,
		cfg.DB.SSLMode,
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
	if err := db.AutoMigrate(&Book{}); err != nil {
		panic(err)
	}
	return db
}
