package main

import (
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
)

type KafkaConfig struct {
	Brokers []string
	Address string
}

var kafkaConfig = &KafkaConfig{}

func connectToKafka() *kafka.Conn {
	if err := viper.UnmarshalKey("kafka", kafkaConfig); err != nil {
		log.Fatalf("Error unmarshalling kafka config: %v", err)
	}
	for i := 0; i < 10; i++ {
		conn, err := kafka.Dial("tcp", kafkaConfig.Address)
		if err == nil {
			return conn
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}

func createKafkaTopics(topics []string) {
	conn := connectToKafka()
	if conn == nil {
		panic("error while connecting to kafka")
	}
	defer conn.Close()
	fmt.Println("connected to kafka")

	for _, topic := range topics {
		err := conn.CreateTopics(kafka.TopicConfig{
			Topic: topic,
			NumPartitions: 1,
			ReplicationFactor: 1,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println("created kafka topic:", topic)
	}
}

func createNewWriter() *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers: kafkaConfig.Brokers,
		Topic: "logging",
	})
}