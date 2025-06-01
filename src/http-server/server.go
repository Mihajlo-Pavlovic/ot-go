package server

import (
	"encoding/json"
	"fmt"
	"http"
)

// Get request JSON structure
type GetRequest struct {
	ID              string `json:"id"`
	ContentType     string `json:"contentType"`
	IncludeMetadata bool   `json:"includeMetadata"`
	HashFunctionId  string `json:"hashFunctionId"`
	ParanetUAL      string `json:"paranetUAL"`
	SubjectUAL      string `json:"subjectUAL"`
}

func InitServer(port string) {
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

	fmt.Printf("Received /get request: %+v\n", req)

	// Send response with status and same id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "received",
		"id":     req.ID,
	})
}
