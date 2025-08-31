package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"ride-sharing/shared/env"
	"syscall"
	"time"
)

var (
	httpAddr = env.GetString("HTTP_ADDR", ":8081")
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /trip/preview", enableCORS(handleTripPreview))
	mux.HandleFunc("/ws/drivers", handleDriversWebSocket)
	mux.HandleFunc("/ws/riders", handleRidersWebSocket)

	srv := &http.Server{
		Addr:    httpAddr,
		Handler: mux,
	}

	errCh := make(chan error, 1)
	go func() {
		slog.Info("starting API Gateway server", "port", httpAddr)
		errCh <- srv.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-errCh:
		slog.Error("failed to start the server", "error", err)

	case sig := <-shutdown:
		slog.Warn("server is shutting down", "signal", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			slog.Error("couldn't stop the server gracefully", "error", err)
			srv.Close()
		}
	}
}
