package main

import (
	"log/slog"
	"net/http"
	"ride-sharing/services/trip-service/internal/infrastructure/handlers"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"ride-sharing/shared/env"
)

var (
	httpAddr = env.GetString("HTTP_ADDR", ":8083")
)

func main() {
	inmemRepo := repository.NewInMemoryRepository()
	svc := service.NewTripService(inmemRepo)
	h := handlers.NewTripHandler(svc)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /preview", h.CreateTrip)

	srv := &http.Server{
		Addr:    httpAddr,
		Handler: mux,
	}

	if err := srv.ListenAndServe(); err != nil {
		slog.Error("Failed to start server", "error", err)
		return
	}
}
