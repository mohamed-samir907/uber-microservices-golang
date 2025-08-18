package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	handlers "ride-sharing/services/trip-service/internal/infrastructure/http"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"ride-sharing/shared/env"
	"syscall"
	"time"
)

var (
	httpAddr = env.GetString("HTTP_ADDR", ":8083")
)

func main() {
	inmemRepo := repository.NewInMemoryRepository()
	svc := service.NewTripService(inmemRepo)
	h := handlers.NewTripHandler(svc)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /preview", h.TripPreview)

	srv := &http.Server{
		Addr:    httpAddr,
		Handler: mux,
	}

	errCh := make(chan error, 1)
	go func() {
		errCh <- srv.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-errCh:
		slog.Error("Failed to start the server", "error", err)

	case sig := <-shutdown:
		slog.Warn("server is shutting down", "signal", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			slog.Error("failed to shutdown the server gracefully", "error", err)
			srv.Close()
		}
	}
}
