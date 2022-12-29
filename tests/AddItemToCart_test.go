package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/subramanyam-searce/product-catalog-go/helpers"
)

func AddToCartViaAPI(ref string, product_id int, quantity int, t *testing.T) map[string]string {
	request_url := fmt.Sprintf("%v/additemtocart?ref=%v&product_id=%v&quantity=%v", URL, ref, product_id, quantity)
	response, err := http.Post(request_url, "application/json", nil)
	helpers.HandleTestError("httpPostError", err, t)

	json_response := map[string]string{}
	json.NewDecoder(response.Body).Decode(&json_response)

	return json_response
}

func TestAddToCart(t *testing.T) {
	test_cases := []map[string]any{
		//valid reference, valid product_id and quantity available in inventory, add new item to cart
		{"ref": "4d4d8297-7663-451d-b79e-49a545728552", "product_id": 3, "quantity": 1, "expected_response": "Item was added to the cart"},

		//same as above, but increasing quantity of existing cart_item
		{"ref": "4d4d8297-7663-451d-b79e-49a545728552", "product_id": 3, "quantity": 3, "expected_response": "Item was added to the cart"},

		//quantity more than inventory stock
		{"ref": "4d4d8297-7663-451d-b79e-49a545728552", "product_id": 3, "quantity": 200, "expected_response": "The Required Quantity is more than the Available Inventory Quantity: 96"},

		//Invalid Cart reference
		{"ref": "4abcd", "product_id": 3, "quantity": 1, "expected_response": "Invalid cart_reference"},

		//invalid product_id
		{"ref": "4d4d8297-7663-451d-b79e-49a545728552", "product_id": 100, "quantity": 1, "expected_response": "Product id is invalid"},
	}

	for _, v := range test_cases {
		response := AddToCartViaAPI(v["ref"].(string), v["product_id"].(int), v["quantity"].(int), t)

		if response["message"] != v["expected_response"] {
			t.Errorf("Expected: %v, Got: %v", v["expected_response"], response["message"])
		}
	}

	RestoreDBTestingState(t)
}
