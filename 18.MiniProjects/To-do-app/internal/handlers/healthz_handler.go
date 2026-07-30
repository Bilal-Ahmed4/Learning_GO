package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/Bilal-Ahmed4/to-do-app/internal/response"
	"github.com/jackc/pgx/v5/pgxpool"
)

func HealthzHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.WriteJson(w, http.StatusOK, map[string]string{
			"status": "alive",
		})
	}
}

func ReadyzHandler(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		if err := pool.Ping(ctx); err != nil {

		}

	}
}
