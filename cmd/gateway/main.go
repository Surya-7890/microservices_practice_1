package main

import (
	"context"
	"net/http"

	"github.com/Surya-7890/book_store/gateway/config"
	_kafka "github.com/Surya-7890/book_store/gateway/kafka"
	"github.com/Surya-7890/book_store/gateway/redis"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

var (
	Redis *redis.Redis
	Kafka *_kafka.KafkaWriters
	ctx   = context.Background()
	App   *config.Application
)

func init() {
	App = config.LoadConfig()
	_kafka.CreateTopics(&App.Kafka)
	Redis = redis.ConnectToRedis(App.Redis)
}

func main() {
	mw := Middleware{
		Key: App.JWT_SECRET,
	}
	go func(cfg *config.KafkaConfig) {
		Kafka = _kafka.CreateWriters(cfg)
	}(&App.Kafka)
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

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
