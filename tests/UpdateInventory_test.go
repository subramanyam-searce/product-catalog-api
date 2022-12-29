package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/subramanyam-searce/product-catalog-go/helpers"
)

func UpdateInventoryViaAPI(product_id int, quantity int) map[string]string {
	inventory_item := map[string]int{
		"product_id": product_id,
		"quantity":   quantity,
	}
	json_product, err := json.Marshal(inventory_item)
	helpers.HandleError("jsonMarshalError", err)

	request_body := bytes.NewBuffer(json_product)

	res, err := http.Post(URL+"/inventory/update", "application/json", request_body)
	helpers.HandleError("httpPostRequestError", err)

	var json_response map[string]string
	json.NewDecoder(res.Body).Decode(&json_response)

	return json_response
}

func TestUpdateInventory(t *testing.T) {
	test_cases := []map[string]any{
		{"product_id": 1, "quantity": 10, "expected_response": "Inventory Items added successfully"},                                                                          //Valid Product ID
		{"product_id": 100, "quantity": 10, "expected_response": "pq: insert or update on table \"inventory\" violates foreign key constraint \"inventory_product_id_fkey\""}, //Invalid Product ID
		{"product_id": 1, "quantity": -110, "expected_response": "pq: new row for relation \"inventory\" violates check constraint \"inventory_quantity_check\""},             //Negative Quantity for existing product
		{"product_id": 2, "quantity": -110, "expected_response": "pq: new row for relation \"inventory\" violates check constraint \"inventory_quantity_check\""},             //Negative Quantity for new product addition
	}

	for _, v := range test_cases {
		response := UpdateInventoryViaAPI(v["product_id"].(int), v["quantity"].(int))

		if response["message"] != v["expected_response"] {
			t.Errorf("Expected: %v, Got: %v", v["expected_response"], response["message"])
		}
	}
}
