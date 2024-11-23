package kafka

import (
	"fmt"
	"reflect"
	"time"

	"github.com/Surya-7890/book_store/gateway/config"
	"github.com/segmentio/kafka-go"
)

type KafkaWriters struct {
	Error   *kafka.Writer `mapstructure:"error"`
	Info    *kafka.Writer `mapstructure:"info"`
	Warning *kafka.Writer `mapstructure:"warning"`
}

func connectToKafka(cfg *config.KafkaConfig) *kafka.Conn {
	for i := 0; i < 10; i++ {
		conn, err := kafka.Dial("tcp", cfg.Address)
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

func CreateTopics(cfg *config.KafkaConfig) {
	conn := connectToKafka(cfg)
	if conn == nil {
		panic("error while connecting to kafka")
	}
	defer conn.Close()
	fmt.Println("connected to kafka")

	// types and values of the topics fetched from config.yml file
	writers := cfg.Writers

	_type := reflect.TypeOf(writers)
	_value := reflect.ValueOf(writers)

	for i := 0; i < _type.NumField(); i++ {
		topic := _value.Field(i).String()
		createKafkaTopics(conn, topic)
	}
}

/* returns a set of writers for logging purposes */
func CreateWriters(cfg *config.KafkaConfig) *KafkaWriters {
	writers := cfg.Writers
	_type := reflect.TypeOf(writers)
	_value := reflect.ValueOf(writers)

	_return := &KafkaWriters{}
	_elem := reflect.ValueOf(_return).Elem()

	for i := 0; i < _type.NumField(); i++ {
		topic := _value.Field(i).String()
		writer := createNewWriter(cfg, topic)

		field := _elem.FieldByName(_type.Field(i).Name)
		if field.IsValid() && field.CanSet() {
			field.Set(reflect.ValueOf(writer))
		}
	}
	return _return
}

func createNewWriter(cfg *config.KafkaConfig, topic string) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers: cfg.Brokers,
		Topic:   topic,
	})
}
