// api/contact.go
package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type ContactRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Message string `json:"message"`
}

type ContactResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    ContactRequest `json:"data"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Only allow POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	// Read body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	
	var req ContactRequest
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	
	// Validate
	if req.Name == "" || req.Email == "" || req.Message == "" {
		http.Error(w, "Name, email, and message are required", http.StatusBadRequest)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	
	response := ContactResponse{
		Success: true,
		Message: "Pesan berhasil dikirim! Tim RGN akan menghubungi Anda segera.",
		Data:    req,
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}