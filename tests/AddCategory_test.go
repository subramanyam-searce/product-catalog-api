package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/subramanyam-searce/product-catalog-go/helpers"
	"github.com/subramanyam-searce/product-catalog-go/typedefs"
)

func AddCategoryViaAPI(category typedefs.Category, t *testing.T) map[string]string {

	json_product, err := json.Marshal(category)
	helpers.HandleError("jsonMarshalError", err)

	request_body := bytes.NewBuffer(json_product)

	res, err := http.Post(URL+"/category/add", "application/json", request_body)
	helpers.HandleError("httpPostRequestError", err)

	var json_response map[string]string
	json.NewDecoder(res.Body).Decode(&json_response)

	return json_response
}

func TestAddCategory(t *testing.T) {
	test_cases := []map[string]any{
		//Valid ID - Unique
		{"category": typedefs.Category{CategoryID: 4, Name: "Plastics"}, "expected_response": "Category added successfully"},
		//Invalid ID - Already Existing
		{"category": typedefs.Category{CategoryID: 1, Name: "Plastics"}, "expected_response": "pq: duplicate key value violates unique constraint \"category_pkey\""},
	}

	for _, v := range test_cases {
		response := AddCategoryViaAPI(v["category"].(typedefs.Category), t)

		if response["message"] != v["expected_response"] {
			t.Errorf("Expected: %v, Got: %v", v["expected_response"], response["message"])
		}
	}
}
