package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/subramanyam-searce/product-catalog-go/helpers"
	"github.com/subramanyam-searce/product-catalog-go/typedefs"
)

func AddProductViaAPI(product typedefs.Product) string {
	json_product, err := json.Marshal(product)
	helpers.HandleError("jsonMarshalError", err)

	request_body := bytes.NewBuffer(json_product)

	res, err := http.Post(URL+"/product/add", "application/json", request_body)
	helpers.HandleError("httpPostRequestError", err)

	var json_response map[string]string
	json.NewDecoder(res.Body).Decode(&json_response)

	response_message := json_response["message"]
	return response_message
}

func checkAddProductResponse(t *testing.T, expected_response string, got_response string) {
	if got_response != expected_response {
		t.Errorf("Expected Response: %v, Got Response: %v", expected_response, got_response)
	}
	RestoreDBTestingState(t)
}

func TestAddProduct(t *testing.T) {
	product := typedefs.Product{
		Product_ID: 294,
		Name:       "Shorts",
		Specification: map[string]any{
			"color": "Brown",
			"size":  "40",
		},
		SKU:        "65548",
		CategoryID: 1,
		Price:      9.99,
	}

	got_response := AddProductViaAPI(product)
	expected_response := "Product Added Successfully"
	checkAddProductResponse(t, expected_response, got_response)

	product.Product_ID = 1
	got_response = AddProductViaAPI(product)
	expected_response = "pq: duplicate key value violates unique constraint \"product_pkey\""
	checkAddProductResponse(t, expected_response, got_response)

	product.Product_ID = 494
	product.CategoryID = 200
	got_response = AddProductViaAPI(product)
	expected_response = "pq: insert or update on table \"product\" violates foreign key constraint \"product_category_id_fkey\""
	checkAddProductResponse(t, expected_response, got_response)
}
