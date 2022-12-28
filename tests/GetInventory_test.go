package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/subramanyam-searce/product-catalog-go/helpers"
)

func GetInventoryViaAPI(t *testing.T) []map[string]any {
	response, err := http.Get(URL + "/inventory")
	helpers.HandleTestError("httpGetError", err, t)

	response_json := []map[string]any{}
	err = json.NewDecoder(response.Body).Decode(&response_json)
	helpers.HandleTestError("jsonDecodingError", err, t)

	RestoreDBTestingState(t)
	return response_json
}

func TestGetInventory(t *testing.T) {
	expected_response := `[{"product_id":3,"product_name":"Jeans","quantity":100}]`
	response := GetInventoryViaAPI(t)
	response_json, err := json.Marshal(response)
	helpers.HandleTestError("jsonMarshalError", err, t)

	if expected_response != string(response_json) {
		t.Errorf("Expected: %v, Got: %v", expected_response, response_json)
	}
}
