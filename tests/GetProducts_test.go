package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/subramanyam-searce/product-catalog-go/helpers"
)

func GetProductsViaAPI(page_no int, t *testing.T) any {
	response, err := http.Get(URL + "/products/" + fmt.Sprint(page_no))
	helpers.HandleTestError("httpGetError", err, t)

	var v any

	err = json.NewDecoder(response.Body).Decode(&v)
	helpers.HandleTestError("jsonDecodingError", err, t)

	return v
}

func TestGetProducts(t *testing.T) {
	products := GetProductsViaAPI(1, t)

	_, ok := products.([]any)

	if !ok {
		t.Errorf("Expected a slice of products but got: " + fmt.Sprint(products))
	}

	products = GetProductsViaAPI(2, t)

	_, ok = products.(map[string]any)["message"]

	if !ok {
		t.Errorf("Expected an error of products but got: " + fmt.Sprint(products))
	}
}
