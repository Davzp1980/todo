package main

import (
	"log"

	"todo"

	"todo/handler"
	"todo/repository"

	"todo/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRouters()); err != nil {
		log.Fatalf("Error occured while running http server %s", err.Error())
	}

}
