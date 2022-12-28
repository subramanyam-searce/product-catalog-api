package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/subramanyam-searce/product-catalog-go/helpers"
)

func DeleteInventoryItemViaAPI(product_id int, t *testing.T) map[string]string {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%v/inventory/delete/%v", URL, product_id), nil)
	helpers.HandleTestError("httpNewRequestError", err, t)

	res, err := http.DefaultClient.Do(req)
	helpers.HandleTestError("httpDefaultClientDoError", err, t)

	var v map[string]string
	json.NewDecoder(res.Body).Decode(&v)

	return v
}

func TestDeleteInventory(t *testing.T) {
	test_cases := []map[string]any{
		{"product_id": 3, "expected_response": "Successfully deleted the Inventory Item"}, //valid existing inventory product_id
		{"product_id": 345, "expected_response": "Inventory Item Not Found"},              //Invalid product_id
	}

	for _, v := range test_cases {
		response := DeleteInventoryItemViaAPI(v["product_id"].(int), t)

		if response["message"] != v["expected_response"] {
			t.Errorf("Expected: %v, Got: %v", v["expected_response"], response["message"])
		}
	}
}
