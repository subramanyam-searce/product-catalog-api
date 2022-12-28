package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/subramanyam-searce/product-catalog-go/helpers"
)

func GetCartViaAPI(ref string, t *testing.T) any {
	response, err := http.Get(URL + "/cart/get?ref=" + ref)
	helpers.HandleTestError("httpGetError", err, t)

	var response_json any
	err = json.NewDecoder(response.Body).Decode(&response_json)
	helpers.HandleTestError("jsonDecodingError", err, t)

	return response_json
}

func TestGetCart(t *testing.T) {
	test_cases := []map[string]any{
		//Valid Cart reference
		{"ref": "4d4d8297-7663-451d-b79e-49a545728552", "expected_response": `{"created_at":"2022-12-28T13:52:51.016582Z","items":[{"product_id":3,"quantity":10}],"ref":"4d4d8297-7663-451d-b79e-49a545728552","total_cart_value":259.9}`},

		//Invalid cart Reference
		{"ref": "abcd", "expected_response": `{"message":"Cart Reference is Invalid"}`},
	}

	for _, v := range test_cases {
		response := GetCartViaAPI(v["ref"].(string), t)

		json_response, err := json.Marshal(response)
		helpers.HandleTestError("jsonMarshalError", err, t)

		if string(json_response) != v["expected_response"].(string) {
			t.Errorf("Expected Response: %v, Got Response: %v", v["expected_response"].(string), string(json_response))
		}
	}
}
