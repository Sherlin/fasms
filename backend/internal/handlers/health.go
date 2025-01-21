package handlers

import (
	"encoding/json"
	"net/http"
	"runtime"
)

// HealthResponse represents the response structure for the health check.
type HealthResponse struct {
	Status string `json:"status"`
	Go string `json:"go"`
}

// HealthCheckHandler handles the health check endpoint.
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{Status: "OK", Go: runtime.Version()}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}