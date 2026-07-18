package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Bilal-Ahmed4/student-api/internal/config"
	"github.com/Bilal-Ahmed4/student-api/internal/http/handlers/student"
	"github.com/Bilal-Ahmed4/student-api/internal/storage/sqlite"
)

func main() {
	//we have to load the config first

	config := config.MustLoad()

	//database
	storage, err := sqlite.New(config)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("storage intialized", slog.String("env", config.Env), slog.String("version", "1.0.0"))

	//setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New(storage))
	router.HandleFunc("GET /api/students/{id}", student.GetById(storage))
	// setup the server

	server := http.Server{
		Addr:    config.Addr, 
		Handler: router,
	}

	fmt.Println("serverstarted at", config.HTTPServer.Addr)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM) //this means if any signal comes from these notify us on done chan

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start the server")
		}
	}()

	<-done

	// runs after the signal from the keyboard
	slog.Info("Shutting down the server")
	// if the shutdown doesnt happened for some reason so we give a time to the server shutdown using the context
	// if not shutdown in 5 sec than throw error server.shutdown expect an var of type of context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown((ctx)); err != nil {
		slog.Info("failed to shutdown the server", slog.String("error", err.Error()))
	}
	slog.Info("Server Shutdown succesfully")

}
