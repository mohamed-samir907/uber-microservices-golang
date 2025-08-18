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

	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello from API Gateway!"))
	})

	http.ListenAndServe(httpAddr, nil)
}
