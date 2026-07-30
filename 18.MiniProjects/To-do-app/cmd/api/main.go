package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Bilal-Ahmed4/to-do-app/internal/config"
	postgres "github.com/Bilal-Ahmed4/to-do-app/internal/database"
	"github.com/Bilal-Ahmed4/to-do-app/internal/handlers"
	"github.com/Bilal-Ahmed4/to-do-app/internal/middleware"
	_ "github.com/lib/pq"
)

// migrate create -ext sql -dir migrations -seq create_todos_table
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func main() {

	cfg, err := config.MustLoad()
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	//now we have cfg now we create a database pool
	pool, err := postgres.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	defer pool.Close()

	var router *http.ServeMux
	router = http.NewServeMux()

	// router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
	// 	// this is equivalent to the context.Json(200,gin.H{}) gin.H is map of type map[string]interface{}
	// 	writeJSON(w, http.StatusOK, map[string]interface{}{
	// 		"message":   "Todo API is running!",
	// 		"status":    "success",
	// 		"datatbase": "connected",
	// 	})
	// })

	router.HandleFunc("POST /todos", handlers.CreateNewTodoHandler(pool))

	// create api to fetch all the todos
	router.HandleFunc("GET /todos", handlers.GetTodosHandler(pool))

	router.HandleFunc("GET /todos/{id}", handlers.GetTodosByIdHandler(pool))

	router.HandleFunc("PUT /todos/{id}", handlers.UpdateTodoHandler(pool))

	router.HandleFunc("DELETE /todos/{id}", handlers.DeleteTodoHandler(pool))

	router.HandleFunc("POST /auth/registration", handlers.CreateUser(pool))

	router.HandleFunc("POST /auth/login", handlers.LoginHandler(pool, cfg))

	router.HandleFunc("/test", middleware.AuthMiddleware(cfg, handlers.TestHandler()))

	router.HandleFunc("GET /healthz", handlers.HealthzHandler())
	router.HandleFunc("GET /readyz", handlers.ReadyzHandler(pool))

	// http.ListenAndServe(":8080", router) // here you will provide the port and the mux object
	// we can also use an alternative for this the above basically auto create the &http.Server and
	// gin router.run use the http.ListenAndServe under the hood
	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	//graceful shutdown

	go func() {
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("Unable to start server: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit // this is blocking and will only run when signal.notify write into the quit

	log.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	// 4. Ask the server to stop gracefully
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("forced shutdown: %v", err)
	}

	log.Println("server stopped cleanly")

}
