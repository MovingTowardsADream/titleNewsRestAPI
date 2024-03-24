package main

import (
	"log"
	titleNews "titleNewsRestApi"
	"titleNewsRestApi/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	server := new(titleNews.Server)
	if err := server.Run("8000", handlers.InitHandler()); err != nil {
		log.Fatal("Error server")
	}
}
