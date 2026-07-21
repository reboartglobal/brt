// api/health.go
package api

import (
	"encoding/json"
	"net/http"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
	Version string `json:"version"`
	Env     string `json:"environment"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	response := HealthResponse{
		Status:  "ok",
		Service: "rgn-website",
		Version: "1.0.0",
		Env:     getEnv("VERCEL_ENV", "development"),
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
