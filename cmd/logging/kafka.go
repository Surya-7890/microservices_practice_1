package main

import (
	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
)

type KafkaConfig struct {
	Brokers []string
}

func createNewReader() *kafka.Reader {
	var kafkaConfig KafkaConfig
	if err := viper.UnmarshalKey("kafka", &kafkaConfig); err != nil {
		panic(err)
	}

	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: kafkaConfig.Brokers,
		Topic: "logging",
		GroupID: "logging",
	})
}