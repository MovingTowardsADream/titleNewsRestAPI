package main

import (
	"log"
	titleNews "titleNewsRestApi"
	"titleNewsRestApi/pkg/handler"
	"titleNewsRestApi/pkg/repository"
	"titleNewsRestApi/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	serv := service.NewService(repos)
	handlers := handler.NewHandler(serv)

	server := new(titleNews.Server)
	if err := server.Run("8000", handlers.InitHandler()); err != nil {
		log.Fatal("Error server")
	}
}
