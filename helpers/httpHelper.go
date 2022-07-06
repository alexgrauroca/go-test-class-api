package helpers

import (
	"encoding/json"
	"net/http"
)

func HttpJsonResponse(w http.ResponseWriter, v any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(v)
}
