package handlers_cart

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/subramanyam-searce/product-catalog-go/helpers"
	"github.com/subramanyam-searce/product-catalog-go/queryhelpers"
)

func AddItemsToCart(w http.ResponseWriter, r *http.Request) {
	response := []map[string]any{}
	request_body := []map[string]int{}

	ref := r.URL.Query().Get("ref")

	err := json.NewDecoder(r.Body).Decode(&request_body)
	if err != nil {
		return
	}

	for _, v := range request_body {
		new_response_item := map[string]any{}
		product_id := v["product_id"]
		quantity := v["quantity"]
		new_response_item["product_id"] = product_id
		new_response_item["quantity"] = quantity

		new_response_item["message"] = queryhelpers.AddItemToCart(ref, fmt.Sprint(quantity), fmt.Sprint(product_id))["message"]
		response = append(response, new_response_item)
	}

	helpers.SendResponse(response, w)
}
