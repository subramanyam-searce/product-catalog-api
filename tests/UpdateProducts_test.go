package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/subramanyam-searce/product-catalog-go/helpers"
)

func TestUpdateProducts(t *testing.T) {
	//valid update operation
	product_id := 1
	update_json_req_body_map := map[string]any{
		"name": "Sweater",
		"specification": map[string]any{
			"color":  "blue",
			"gender": "female",
		},
		"sku":   62244,
		"price": 11.99,
	}

	CheckUpdateEndpoint(product_id, update_json_req_body_map, "Update done successfully", t)

	//trying to update product_id
	update_json_req_body_map["product_id"] = 2
	CheckUpdateEndpoint(product_id, update_json_req_body_map, "product_id cannot be updated", t)

	//invalid product_id
	delete(update_json_req_body_map, "product_id")
	product_id = 100
	CheckUpdateEndpoint(product_id, update_json_req_body_map, "Product not found", t)

}

func CheckUpdateEndpoint(product_id int, update_json_req_body_map map[string]any, expected_response string, t *testing.T) {
	json_product, err := json.Marshal(update_json_req_body_map)
	helpers.HandleTestError("jsonMarshalError", err, t)

	request_body := bytes.NewBuffer(json_product)
	req, err := http.NewRequest("PUT", fmt.Sprintf("%v/product/update/%v", URL, product_id), request_body)
	helpers.HandleTestError("httpNewRequestError", err, t)

	res, err := http.DefaultClient.Do(req)
	helpers.HandleTestError("httpDefaultClientDoError", err, t)

	var v map[string]string
	json.NewDecoder(res.Body).Decode(&v)

	if expected_response != v["message"] {
		t.Errorf("Expected: %v, Got: %v", expected_response, v["message"])
	}

	RestoreDBTestingState(t)
}
