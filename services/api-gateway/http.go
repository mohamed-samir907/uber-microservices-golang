package main

import (
	"io"
	"net/http"
)

func handleTripPreview(w http.ResponseWriter, r *http.Request) {
	// var req previewTripRequest
	// if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	// 	http.Error(w, "Invalid request body", http.StatusBadRequest)
	// 	return
	// }
	// defer r.Body.Close()

	// // validation
	// if req.UserId == "" {
	// 	http.Error(w, "User ID is required", http.StatusBadRequest)
	// 	return
	// }

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
