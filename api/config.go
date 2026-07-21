// api/config.go
package api

import (
	"encoding/json"
	"net/http"
	"os"
)

// Config for Vercel serverless functions
var (
	Version     = "1.0.0"
	Environment = getEnv("VERCEL_ENV", "development")
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// ErrorResponse for consistent error handling
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{
		Error:   http.StatusText(code),
		Message: message,
		Code:    code,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}
