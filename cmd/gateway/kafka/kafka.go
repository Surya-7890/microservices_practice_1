package kafka

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
)

type KafkaWriterConfig struct {
	Error   string `yaml:"error"`
	Info    string `yaml:"info"`
	Warning string `yaml:"warning"`
}

type KafkaConfig struct {
	Brokers []string          `yaml:"brokers"`
	Address string            `yaml:"address"`
	Writers KafkaWriterConfig `yaml:"writers"`
}

type KafkaWriters struct {
	Error   *kafka.Writer `yaml:"error"`
	Info    *kafka.Writer `yaml:"info"`
	Warning *kafka.Writer `yaml:"warning"`
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

/* creates different kafka configs for logging purposes */
func createKafkaTopics(conn *kafka.Conn, topic string) {
	err := conn.CreateTopics(kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     1,
		ReplicationFactor: 1,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("created kafka topic:", topic)
}

func CreateTopics() {
	conn := connectToKafka()
	if conn == nil {
		panic("error while connecting to kafka")
	}
	defer conn.Close()
	fmt.Println("connected to kafka")

	// types and values of the topics fetched from config.yml file
	writers := kafkaConfig.Writers

	_type := reflect.TypeOf(writers)
	_value := reflect.ValueOf(writers)

	for i := 0; i < _type.NumField(); i++ {
		topic := _value.Field(i).String()
		createKafkaTopics(conn, topic)
	}
}

/* returns a set of writers for logging purposes */
func CreateWriters() *KafkaWriters {
	writers := kafkaConfig.Writers
	_type := reflect.TypeOf(writers)
	_value := reflect.ValueOf(writers)

	_return := &KafkaWriters{}
	_elem := reflect.ValueOf(_return).Elem()

	for i := 0; i < _type.NumField(); i++ {
		topic := _value.Field(i).String()
		writer := createNewWriter(topic)

		field := _elem.FieldByName(_type.Field(i).Name)
		if field.IsValid() && field.CanSet() {
			field.Set(reflect.ValueOf(writer))
		}
	}
	return _return
}

func createNewWriter(topic string) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers: kafkaConfig.Brokers,
		Topic:   topic,
	})
}
