package handlers_cart

import (
	"net/http"

	"github.com/subramanyam-searce/product-catalog-go/helpers"
	"github.com/subramanyam-searce/product-catalog-go/queryhelpers"
)

func AddToCart(w http.ResponseWriter, r *http.Request) {
	urlQuery := r.URL.Query()

	ref := urlQuery.Get("ref")
	quantity_str := urlQuery.Get("quantity")
	product_id := urlQuery.Get("product_id")

	response := queryhelpers.AddItemToCart(ref, quantity_str, product_id)

	helpers.SendResponse(response, w)
}
