package main

import (
	"fmt"
	"net/http"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func main() {
	gw := gwruntime.NewServeMux()
	setup(gw)

	mux := http.NewServeMux()
	mux.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("swagger")
		http.ServeFile(w, r, "./swagger.json")
	})
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