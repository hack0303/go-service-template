package handler

import (
	"net/http"
)

// HealthCheckResponse represents the health check response structure
type HealthCheckResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

// HealthCheck godoc
// @Summary Health check endpoint
// @Description Returns the health status of the service
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} HTTPResponse{data=HealthCheckResponse}
// @Router /health [get]
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := NewSuccessResponse(HealthCheckResponse{
		Status:  "healthy",
		Version: "1.0.0",
	})
	WriteResponse(w, http.StatusOK, response)
}
