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
		err_ := Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.DB_ERROR),
			Value: []byte(err.Error()),
		})
		if err_ != nil {
			fmt.Println(err_)
		}
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(postgres_uri))
	if err != nil {
		err_ := Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.DB_ERROR),
			Value: []byte(err.Error()),
		})
		if err_ != nil {
			fmt.Println(err_)
		}
		panic(err)
	}
	if err := db.AutoMigrate(&Book{}); err != nil {
		err_ := Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.DB_ERROR),
			Value: []byte(err.Error()),
		})
		if err_ != nil {
			fmt.Println(err_)
		}
		panic(err)
	}

	err = Kafka.Info.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(utils.DB_INFO),
		Value: []byte("[books-service]: connected to postgres"),
	})
	if err != nil {
		fmt.Println(err)
	}
	return db
}
