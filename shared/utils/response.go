package utils

import (
	"encoding/json"
	"net/http"
	"ride-sharing/shared/contracts"
	"strconv"
)

// APIError creates a new API error with the given message and status code.
func APIError(errMsg string, status int) *contracts.APIError {
	return &contracts.APIError{
		Code:    strconv.Itoa(status),
		Message: errMsg,
	}
}

// ResponseJson sends a JSON response with the given status, data, and error.
// It sets the appropriate headers and marshals the response into JSON format.
func ResponseJson(w http.ResponseWriter, status int, data any, err *contracts.APIError) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := contracts.APIResponse{
		Data:  data,
		Error: err,
	}

	return json.NewEncoder(w).Encode(resp)
}
