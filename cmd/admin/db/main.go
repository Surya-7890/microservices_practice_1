package db

import (
	"fmt"

	"github.com/Surya-7890/book_store/admin/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getPostgresConnectionString(cfg *config.DBConfig) (string, error) {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.SSLMode,
	), nil
}

func ConnectToPostgres(cfg *config.DBConfig) *gorm.DB {
	postgres_uri, err := getPostgresConnectionString(cfg)
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
