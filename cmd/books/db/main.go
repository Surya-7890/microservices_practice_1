package db

import (
	"context"
	"fmt"

	"github.com/Surya-7890/book_store/books/config"
	"github.com/Surya-7890/book_store/books/utils"
	"github.com/segmentio/kafka-go"
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

func ConnectToPostgres(Kafka *config.KafkaWriters, cfg *config.DBConfig) *gorm.DB {
	postgres_uri, err := getPostgresConnectionString(cfg)
	if err != nil {
		Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.DB_ERROR),
			Value: []byte(err.Error()),
		})
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(postgres_uri))
	if err != nil {
		Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.DB_ERROR),
			Value: []byte(err.Error()),
		})
		panic(err)
	}
	if err := db.AutoMigrate(&Book{}); err != nil {
		Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.DB_ERROR),
			Value: []byte(err.Error()),
		})
		panic(err)
	}

	err = Kafka.Info.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(utils.DB_INFO),
		Value: []byte("connected to postgres successfully"),
	})

	if err != nil {
		fmt.Println("err db:", err.Error())
	}

	return db
}
