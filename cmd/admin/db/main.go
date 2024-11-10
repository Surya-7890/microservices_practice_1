package db

import (
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToPostgres() *gorm.DB {
	postgres_uri := viper.GetString("postgres")
	db, err := gorm.Open(postgres.Open(postgres_uri))
	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&Admin{}); err != nil {
		panic(err)
	}
	return db
}
