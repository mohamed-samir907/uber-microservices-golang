package main

import (
	"io"
	"net/http"
)

func handleTripPreview(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Post("http://trip-service:8083/preview", "application/json", r.Body)
	if err != nil {
		http.Error(w, "Failed to call trip service", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	w.WriteHeader(resp.StatusCode)

	w.Write(body)
}
