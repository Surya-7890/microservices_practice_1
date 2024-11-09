package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Surya-7890/book_store/gateway/config"
	_kafka "github.com/Surya-7890/book_store/gateway/kafka"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/segmentio/kafka-go"
)

func init() {
	config.LoadConfig()
	_kafka.CreateTopics()
}

func main() {
	gw := gwruntime.NewServeMux()
	setup(gw)

	mux := http.NewServeMux()
	mux.Handle("/", gw)

	writer := _kafka.CreateWriters()
	go func() {
		for {
			err := writer.Info.WriteMessages(context.Background(), kafka.Message{
				Key:   []byte("sample"),
				Value: []byte("sample testing"),
			})
			if err != nil {
				fmt.Println("error while writing", err.Error())
				continue
			}
			time.Sleep(5 * time.Second)
		}
	}()

	s := http.Server{
		Handler: mux,
		Addr:    ":10000",
	}

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
