package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Bilal-Ahmed4/to-do-app/internal/config"
	postgres "github.com/Bilal-Ahmed4/to-do-app/internal/database"
	"github.com/Bilal-Ahmed4/to-do-app/internal/handlers"
	_ "github.com/lib/pq"
)

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

	router.HandleFunc("POST /todos", handlers.CreateNewTodo(pool))

	// create api to fetch all the todos
	router.HandleFunc("GET /todos", handlers.GetTodos(pool))

	// http.ListenAndServe(":8080", router) // here you will provide the port and the mux object
	// we can also use an alternative for this the above basically auto create the &http.Server and
	// gin router.run use the http.ListenAndServe under the hood
	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("Unable to start server: ", err)
	}
}
