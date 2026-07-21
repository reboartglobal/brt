// api/packages.go
package api

import (
	"encoding/json"
	"net/http"
)

type Package struct {
	Name        string   `json:"name"`
	Price       string   `json:"price"`
	Speed       string   `json:"speed"`
	Features    []string `json:"features"`
	IsPopular   bool     `json:"isPopular"`
	Description string   `json:"description"`
}

func PackagesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	packages := []Package{
		{
			Name:      "Start",
			Price:     "Rp199.000",
			Speed:     "50 Mbps",
			Features:  []string{"Unlimited kuota", "1 perangkat", "Dukungan standar"},
			IsPopular: false,
			Description: "Cocok untuk penggunaan dasar",
		},
		{
			Name:      "Pro",
			Price:     "Rp349.000",
			Speed:     "200 Mbps",
			Features:  []string{"Unlimited kuota", "5 perangkat", "Dukungan prioritas", "VPN & keamanan"},
			IsPopular: true,
			Description: "Paling laris untuk rumah & kerja",
		},
		{
			Name:      "Ultimate",
			Price:     "Rp599.000",
			Speed:     "1 Gbps",
			Features:  []string{"Unlimited kuota", "10 perangkat", "Dukungan 24/7 prioritas", "VPN premium", "IP publik statis"},
			IsPopular: false,
			Description: "Untuk pengguna power & bisnis",
		},
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(packages)
}