package db

import (
	"context"
	"fmt"
	"os"

	"github.com/Surya-7890/book_store/user/config"
	"github.com/Surya-7890/book_store/user/utils"
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
			fmt.Println(err_.Error())
		}
		os.Exit(1)
	}
	db, err := gorm.Open(postgres.Open(postgres_uri))
	if err != nil {
		err_ := Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.DB_ERROR),
			Value: []byte(err.Error()),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		os.Exit(1)
	}
	if err := db.AutoMigrate(&User{}); err != nil {
		err_ := Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.DB_ERROR),
			Value: []byte(err.Error()),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		os.Exit(1)
	}
	err_ := Kafka.Info.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(utils.DB_INFO),
		Value: []byte("[user-service]: connected to postgres"),
	})
	if err_ != nil {
		fmt.Println(err_.Error())
	}
	return db
}
