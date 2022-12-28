package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/subramanyam-searce/product-catalog-go/helpers"
	"github.com/subramanyam-searce/product-catalog-go/typedefs"
)

func UpdateCategoryViaAPI(category typedefs.Category, t *testing.T) map[string]string {
	update_json_req_body_map := map[string]any{
		"name": category.Name,
	}
	json_product, err := json.Marshal(update_json_req_body_map)
	helpers.HandleTestError("jsonMarshalError", err, t)

	request_body := bytes.NewBuffer(json_product)
	req, err := http.NewRequest("PUT", fmt.Sprintf("%v/category/update/%v", URL, category.CategoryID), request_body)
	helpers.HandleTestError("httpNewRequestError", err, t)

	res, err := http.DefaultClient.Do(req)
	helpers.HandleTestError("httpDefaultClientDoError", err, t)

	var v map[string]string
	json.NewDecoder(res.Body).Decode(&v)

	RestoreDBTestingState(t)
	return v
}

func TestUpdateCategory(t *testing.T) {
	test_cases := []map[string]any{
		//Valid Category ID
		{"category": typedefs.Category{CategoryID: 1, Name: "Hoodies"}, "expected_response": "Successfully Updated"},

		//Invalid Category ID
		{"category": typedefs.Category{CategoryID: 125, Name: "Hoodies"}, "expected_response": "Category Not Found"},
	}

	for _, v := range test_cases {
		response := UpdateCategoryViaAPI(v["category"].(typedefs.Category), t)

		if response["message"] != v["expected_response"] {
			t.Errorf("Expected %v, Got: %v", v["expected_response"], response["message"])
		}
	}
}
