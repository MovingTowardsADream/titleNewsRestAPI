package main

import (
	_ "github.com/lib/pq"
	"log"
	news "titleNewsRestApi"
	"titleNewsRestApi/pkg/handler"
	"titleNewsRestApi/pkg/repository"
	"titleNewsRestApi/pkg/service"
)

func main() {

	db, err := repository.NewPostgresDB(repository.ConfigDB{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "qwerty",
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("Failed initialization database")
	}

	repos := repository.NewRepository(db)
	serv := service.NewService(repos)
	handlers := handler.NewHandler(serv)

	server := new(news.Server)
	if err := server.Run("8000", handlers.InitHandler()); err != nil {
		log.Fatal("Error server")
	}
}
