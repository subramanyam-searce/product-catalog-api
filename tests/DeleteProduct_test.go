package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/subramanyam-searce/product-catalog-go/helpers"
)

func TestDeleteProduct(t *testing.T) {
	//Valid Product ID
	product_id := 1
	DeleteProductViaAPI(product_id, "Successfully Deleted", t)

	//Invalid Product Id
	product_id = 100
	DeleteProductViaAPI(product_id, "Product not found", t)

}

func DeleteProductViaAPI(product_id int, expected_response string, t *testing.T) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%v/product/delete/%v", URL, product_id), nil)
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
