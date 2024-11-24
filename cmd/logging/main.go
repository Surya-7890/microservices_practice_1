package main

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"strings"
	"sync"

	"github.com/Surya-7890/book_store/logging/config"
	_kafka "github.com/Surya-7890/book_store/logging/kafka"
	"github.com/segmentio/kafka-go"
)

var LogFilesMap = make(map[string]*os.File)

var (
	App *config.Application
)

func init() {
	App = config.LoadConfig()
}

func handleLogs(reader *kafka.Reader, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Println("["+reader.Config().Topic+"]:", string(msg.Value))
	}
}

func main() {
	readers := _kafka.CreateReaders(&App.KafkaConfig)
	_kafka.CreateLogFiles(App.KafkaConfig, LogFilesMap)

	ctx := context.Background()

	wg := &sync.WaitGroup{}

	_value := reflect.ValueOf(*readers)
	for i := 0; i < _value.NumField(); i++ {
		reader := _value.Field(i).Interface().(*kafka.Reader)

		wg.Add(1)
		go func(reader *kafka.Reader, ctx context.Context, wg *sync.WaitGroup) {
			defer wg.Done()
			topic := reader.Config().Topic
			for {
				msg, err := reader.ReadMessage(ctx)
				if err != nil {
					fmt.Println(err.Error())
					continue
				}
				fmt.Println(msg.Value)
				if file, ok := LogFilesMap[strings.Split(topic, "-")[0]]; ok || file != nil {
					file.WriteString("[" + topic + "]: " + string(msg.Value) + "\n")
				}
			}
		}(reader, ctx, wg)
	}

	wg.Wait()
}
