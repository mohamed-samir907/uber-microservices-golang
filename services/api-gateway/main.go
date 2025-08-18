package main

import (
	"log/slog"
	"net/http"
	"ride-sharing/shared/env"
)

var (
	httpAddr = env.GetString("HTTP_ADDR", ":8081")
)

func main() {
	slog.Info("Starting API Gateway")

	mux := http.NewServeMux()

	mux.HandleFunc("POST /trip/preview", handleTripPreview)

	srv := &http.Server{
		Addr:    httpAddr,
		Handler: mux,
	}

	if err := srv.ListenAndServe(); err != nil {
		slog.Error("Failed to start API Gateway", "error", err)
	}
}
