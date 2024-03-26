package main

import (
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	news "titleNewsRestApi"
	"titleNewsRestApi/pkg/handler"
	"titleNewsRestApi/pkg/repository"
	"titleNewsRestApi/pkg/service"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	db, err := repository.NewPostgresDB(repository.ConfigDB{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "qwerty",
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		logrus.Fatalf("Failed initialization database")
	}

	repos := repository.NewRepository(db)
	serv := service.NewService(repos)
	handlers := handler.NewHandler(serv)

	server := new(news.Server)
	if err := server.Run("8000", handlers.InitHandler()); err != nil {
		logrus.Fatal("Error server")
	}
}
