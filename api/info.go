// api/info.go
package api

import (
	"encoding/json"
	"net/http"
	"time"
)

type InfoResponse struct {
	Name      string    `json:"name"`
	Tagline   string    `json:"tagline"`
	Version   string    `json:"version"`
	Timestamp time.Time `json:"timestamp"`
	Status    string    `json:"status"`
	Features  []string  `json:"features"`
}

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	response := InfoResponse{
		Name:      "RGN - Unlimited Connection",
		Tagline:   "Unlimited Connection, Unlimited Possibilities",
		Version:   "1.0.0",
		Timestamp: time.Now(),
		Status:    "running",
		Features: []string{
			"Internet Fiber Optik",
			"Unlimited Kuota",
			"Kecepatan 1 Gbps",
			"Dukungan 24/7",
			"Keamanan Terjamin",
		},
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}