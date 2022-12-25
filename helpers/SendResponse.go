package helpers

import (
	"encoding/json"
	"net/http"
)

func SendResponse(v any, w http.ResponseWriter) {
	json.NewEncoder(w).Encode(v)
}
