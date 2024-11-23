package kafka

import (
	"fmt"
	"reflect"
	"time"

	"github.com/Surya-7890/book_store/user/config"
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

/* returns a set of writers for logging purposes */
func CreateWriters(cfg *config.KafkaConfig) *config.KafkaWriters {
	if err := checkKafkaConn(cfg.Brokers[0]); err != nil {
		panic(err)
	}

	writers := cfg.Writers
	_type := reflect.TypeOf(writers)
	_value := reflect.ValueOf(writers)

	_return := &config.KafkaWriters{}
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
