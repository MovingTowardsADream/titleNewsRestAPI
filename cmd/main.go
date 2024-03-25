package main

import (
	"log"
	news "titleNewsRestApi"
	"titleNewsRestApi/pkg/handler"
	"titleNewsRestApi/pkg/repository"
	"titleNewsRestApi/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	serv := service.NewService(repos)
	handlers := handler.NewHandler(serv)

	server := new(news.Server)
	if err := server.Run("8000", handlers.InitHandler()); err != nil {
		log.Fatal("Error server")
	}
}
