package main

import (
	"context"
	"errors"
	"level_2/develop/dev11/handler"
	"level_2/develop/dev11/repository"
	"level_2/develop/dev11/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	chanOs := make(chan os.Signal)
	signal.Notify(chanOs, syscall.SIGINT, syscall.SIGTERM)

	server := http.Server{Addr: ":8080"}
	storage := repository.NewStorage()
	repositories := repository.NewRepository(storage)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)
	handlers.InitRoutes()

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalln("can not start server")
		}
	}()

	<-chanOs
	cancel()
	log.Println("Shutting Down")

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalln("Can not shutting down...")
	}
}
