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

// HealthCheckHandler godoc
// @Summary Health check
// @Description Check the health of the server
// @Tags health
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /health [get]
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{Status: "OK", Go: runtime.Version()}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}