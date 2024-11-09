package main

import (
	"context"
	"fmt"
	"reflect"
	"sync"

	"github.com/Surya-7890/book_store/logging/config"
	_kafka "github.com/Surya-7890/book_store/logging/kafka"
	"github.com/segmentio/kafka-go"
)

func init() {
	config.LoadConfig()
}

func main() {
	readers := _kafka.CreateReaders()
	fmt.Println("readers", readers)
	ctx := context.Background()

	wg := &sync.WaitGroup{}

	_value := reflect.ValueOf(*readers)
	for i := 0; i < _value.NumField(); i++ {
		reader := _value.Field(i).Interface().(*kafka.Reader)

		wg.Add(1)
		go func(reader *kafka.Reader, ctx context.Context, wg *sync.WaitGroup) {
			defer wg.Done()
			for {
				msg, err := reader.ReadMessage(ctx)
				if err != nil {
					fmt.Println(err.Error())
					continue
				}
				fmt.Println("["+reader.Config().Topic+"]:", string(msg.Value))
			}
		}(reader, ctx, wg)
	}

	wg.Wait()
}
