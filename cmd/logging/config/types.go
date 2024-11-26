package config

import "github.com/segmentio/kafka-go"

type KafkaReaders struct {
	Error   *kafka.Reader
	Info    *kafka.Reader
	Warning *kafka.Reader
}

type KafkaConfigReaders struct {
	Error   string `mapstructure:"error"`
	Info    string `mapstructure:"info"`
	Warning string `mapstructure:"warning"`
}

type KafkaConfig struct {
	Brokers []string           `mapstructure:"brokers"`
	Readers KafkaConfigReaders `mapstructure:"readers"`
}

type Application struct {
	KafkaConfig KafkaConfig `mapstructure:"kafka"`
	Kafka       *KafkaReaders
}
