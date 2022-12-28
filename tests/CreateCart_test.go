package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/subramanyam-searce/product-catalog-go/helpers"
)

func TestCreateCart(t *testing.T) {
	response, err := http.Post(URL+"/cart/create", "application/json", nil)
	helpers.HandleTestError("httpPostError", err, t)

	json_response := map[string]string{}
	json.NewDecoder(response.Body).Decode(&json_response)

	got_response, ok := json_response["ref"]

	if !ok {
		t.Errorf("Expected: %v, Got: %v", "{\"red\":\"<Unique ID here>\"}", got_response)
	}
}
