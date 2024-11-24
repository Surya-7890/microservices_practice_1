package config

import "github.com/segmentio/kafka-go"

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

type Service struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type Application struct {
	KafkaConfig KafkaConfig `mapstructure:"kafka"`
	Port        string      `mapstructure:"port"`
	JWT_SECRET  string      `mapstructure:"jwt_key"`
	User        Service     `mapstructure:"user"`
	Admin       Service     `mapstructure:"admin"`
	Books       Service     `mapstructure:"books"`
	Redis       string      `mapstructure:"redis"`
	Kafka       *KafkaWriters
}
