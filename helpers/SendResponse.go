package helpers

import (
	"encoding/json"
	"net/http"
)

func SendResponse(v any, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
