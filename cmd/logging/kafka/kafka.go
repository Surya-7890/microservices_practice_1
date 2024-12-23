package kafka

import (
	"fmt"
	"reflect"
	"time"

	"github.com/Surya-7890/book_store/logging/config"
	"github.com/segmentio/kafka-go"
)

func checkKafkaConn(broker string) error {

	for i := 0; i < 10; i++ {
		conn, err := kafka.Dial("tcp", broker)
		if err != nil {
			time.Sleep(time.Second * 5)
			continue
		}
		defer conn.Close()
		return nil
	}

	return fmt.Errorf("couldnt connect to kafka")
}

func CreateReaders(cfg *config.KafkaConfig) *config.KafkaReaders {
	if err := checkKafkaConn(cfg.Brokers[0]); err != nil {
		panic(err)
	}

	_type := reflect.TypeOf(cfg.Readers)
	_value := reflect.ValueOf(cfg.Readers)

	_return := &config.KafkaReaders{}

	_elem := reflect.ValueOf(_return).Elem()

	for i := 0; i < _type.NumField(); i++ {
		topic := _value.Field(i).String()
		reader := createNewReader(cfg, topic)

		field := _elem.FieldByName(_type.Field(i).Name)
		if field.IsValid() && field.CanSet() {
			field.Set(reflect.ValueOf(reader))
		}
	}

	return _return
}

func createNewReader(cfg *config.KafkaConfig, topic string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: cfg.Brokers,
		Topic:   topic,
		GroupID: "logging",
	})
}
