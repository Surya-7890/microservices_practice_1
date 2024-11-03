package main

import (
	"net/http"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func main() {
	gw := gwruntime.NewServeMux()
	setup(gw)

	mux := http.NewServeMux()
	mux.Handle("/", gw)

	s := http.Server{
		Handler: mux,
		Addr: ":10000",
	}

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}