package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/subramanyam-searce/product-catalog-go/helpers"
)

func GetCategoriesViaAPI(t *testing.T) any {
	response, err := http.Get(URL + "/categories")
	helpers.HandleTestError("httpGetError", err, t)

	var v any

	err = json.NewDecoder(response.Body).Decode(&v)
	helpers.HandleTestError("jsonDecodingError", err, t)

	return v
}

func TestGetCategories(t *testing.T) {
	categories := GetProductsViaAPI(1, t)

	_, ok := categories.([]any)

	if !ok {
		t.Errorf("Expected a slice of categories but got: " + fmt.Sprint(categories))
	}

	categories = GetProductsViaAPI(2, t)

	_, ok = categories.(map[string]any)["message"]

	if !ok {
		t.Errorf("Expected an error of categories but got: " + fmt.Sprint(categories))
	}
}
