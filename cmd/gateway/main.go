package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Surya-7890/book_store/gateway/config"
	_kafka "github.com/Surya-7890/book_store/gateway/kafka"
	"github.com/Surya-7890/book_store/gateway/redis"
	"github.com/Surya-7890/book_store/gateway/utils"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/segmentio/kafka-go"
)

var (
	Redis *redis.Redis
	ctx   = context.Background()
	App   *config.Application
)

func init() {
	App = config.LoadConfig()
	_kafka.CreateTopics(&App.KafkaConfig)
	Redis = redis.ConnectToRedis(App.Redis)
	App.Kafka = _kafka.CreateWriters(&App.KafkaConfig)
}

func main() {
	mw := Middleware{
		Key:   App.JWT_SECRET,
		Kafka: App.Kafka,
	}

	gw := gwruntime.NewServeMux([]gwruntime.ServeMuxOption{
		gwruntime.WithMetadata(mw.requestInterceptor),
		gwruntime.WithForwardResponseOption(mw.responseInterceptor),
	}...)
	setup(gw, App)

	mux := http.NewServeMux()
	mux.Handle("/", gw)

	s := http.Server{
		Handler: mux,
		Addr:    App.Port,
	}

	err := App.Kafka.Info.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(utils.SERVER_INFO),
		Value: []byte("[gateway-service]: running server... on port: " + App.Port),
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	err = s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
