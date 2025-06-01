package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ot-go/service"
)

// Get request JSON structure
type GetRequest struct {
	ID              string `json:"id"`
	ContentType     string `json:"contentType"`
	IncludeMetadata bool   `json:"includeMetadata"`
}

func InitServer(port string) error {
	http.HandleFunc("/get", handleGet)
	return http.ListenAndServe(port, nil)
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	// Make sure this is POST request
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req GetRequest
	// Validate JSON structure
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := service.UALSvc.Validate(req.ID); err != nil {
		http.Error(w, fmt.Sprintf("Invalid UAL: %v", err), http.StatusBadRequest)
		return
	}

	fmt.Printf("Received /get request: %+v\n", req)

	// Send response with status and same id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "received",
		"id":     req.ID,
	})
}
