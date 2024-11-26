package config

import (
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

type KafkaWriters struct {
	Error   *kafka.Writer `mapstructure:"error"`
	Info    *kafka.Writer `mapstructure:"info"`
	Warning *kafka.Writer `mapstructure:"warning"`
}

type KafkaWriterConfig struct {
	Error   string `mapstructure:"error"`
	Info    string `mapstructure:"info"`
	Warning string `mapstructure:"warning"`
}

type KafkaConfig struct {
	Brokers []string          `mapstructure:"brokers"`
	Address string            `mapstructure:"address"`
	Writers KafkaWriterConfig `mapstructure:"writers"`
}

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
}

type Application struct {
	DatabaseConfig DBConfig    `mapstructure:"postgres"`
	KafkaConfig    KafkaConfig `mapstructure:"kafka"`
	Port           string      `mapstructure:"port"`
	DB             *gorm.DB
	Kafka          *KafkaWriters
}
