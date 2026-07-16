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
)

func main() {
	//we have to load the config first
	config := config.MustLoad()

	//setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /api/students", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
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
	ctx, cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	if err:=server.Shutdown((ctx)); err!=nil{
		slog.Info("failed to shutdown the server",slog.String("error",err.Error()))
	}
	slog.Info("Server Shutdown succesfully")
	
}
