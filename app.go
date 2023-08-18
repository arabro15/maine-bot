package main

import (
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"maineBot/config"
	"maineBot/pkg/client/postgresql"
	"net"
	"net/http"
	"time"
)

func main() {
	router := httprouter.New()
	cfg := config.GetConfig()
	_, err := postgresql.NewClient(context.TODO(), 3, cfg.Storage) //first param postgreSQLClient
	if err != nil {
		log.Fatalf("%v", err)
	}
	//repository := repository2.NewUserRepository(postgreSQLClient)
	//create := usecase.CreateUserUseCase.Execute()
	//controller := controller.UserController{
	//
	//}
	// Не знаю что тут делать
	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
	listener, listenErr := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))

	if listenErr != nil {
		log.Fatal(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.Serve(listener))
}
