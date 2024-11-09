package main

import (
	"context"
	"fmt"

	"github.com/Surya-7890/book_store/logging/config"
)

func init() {
	config.LoadConfig()
}

func main() {
	reader := createNewReader()
	ctx := context.Background()
	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			fmt.Println("[logging]:", err.Error())
			continue
		}
		fmt.Println(string(msg.Value))
	}
}