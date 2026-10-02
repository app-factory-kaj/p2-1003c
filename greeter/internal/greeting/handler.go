package greeting

import (
	"encoding/json"
	"net/http"
)

// errorResponse is the JSON shape every error on this service returns.
type errorResponse struct {
	Error string `json:"error"`
}

// Handler serves GET /hello.
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(errorResponse{Error: "method not allowed"})
		return
	}

	name := r.URL.Query().Get("name")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(New(name))
}
