package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/subramanyam-searce/product-catalog-go/helpers"
)

func DeleteCategoryViaAPI(category_id int, expected_response string, t *testing.T) map[string]string {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%v/category/delete/%v", URL, category_id), nil)
	helpers.HandleTestError("httpNewRequestError", err, t)

	res, err := http.DefaultClient.Do(req)
	helpers.HandleTestError("httpDefaultClientDoError", err, t)

	var v map[string]string
	json.NewDecoder(res.Body).Decode(&v)

	RestoreDBTestingState(t)
	return v
}

func TestDeleteCategories(t *testing.T) {
	test_cases := []map[string]any{
		//existing category not used as foreign key by product table
		{"category_id": 3, "expected_response": "Successfully deleted the category"},

		//existing category which is being used as foreign key by product table
		{"category_id": 2, "expected_response": "pq: update or delete on table \"category\" violates foreign key constraint \"product_category_id_fkey\" on table \"product\""},

		//invalid category id
		{"category_id": 500, "expected_response": "Category Not Found"},
	}

	for _, v := range test_cases {
		response := DeleteCategoryViaAPI(v["category_id"].(int), v["expected_response"].(string), t)

		if response["message"] != v["expected_response"] {
			t.Errorf("expected: %v, Got: %v", v["expected_response"], response["message"])
		}
	}
}
