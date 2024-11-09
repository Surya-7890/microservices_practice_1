package kafka

import (
	"reflect"

	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
)

type KafkaReaders struct {
	Error   *kafka.Reader
	Info    *kafka.Reader
	Warning *kafka.Reader
}

type KafkaConfigReaders struct {
	Error   string `yaml:"error"`
	Info    string `yaml:"info"`
	Warning string `yaml:"warning"`
}

type KafkaConfig struct {
	Brokers []string           `yaml:"brokers"`
	Readers KafkaConfigReaders `yaml:"readers"`
}

var kafkaConfig = &KafkaConfig{}

func CreateReaders() *KafkaReaders {
	if err := viper.UnmarshalKey("kafka", &kafkaConfig); err != nil {
		panic(err)
	}

	_type := reflect.TypeOf(kafkaConfig.Readers)
	_value := reflect.ValueOf(kafkaConfig.Readers)

	_return := &KafkaReaders{}

	_elem := reflect.ValueOf(_return).Elem()

	for i := 0; i < _type.NumField(); i++ {
		topic := _value.Field(i).String()
		reader := createNewReader(topic)

		field := _elem.FieldByName(_type.Field(i).Name)
		if field.IsValid() && field.CanSet() {
			field.Set(reflect.ValueOf(reader))
		}
	}

	return _return
}

func createNewReader(topic string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: kafkaConfig.Brokers,
		Topic:   topic,
		GroupID: "logging",
	})
}
