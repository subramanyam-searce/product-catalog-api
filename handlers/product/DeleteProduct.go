package handlers_product

import (
	"net/http"

	"github.com/subramanyam-searce/product-catalog-go/helpers"
	"github.com/subramanyam-searce/product-catalog-go/queryhelpers"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := helpers.ParseMuxVarToInt(r, "id")

	var err error
	product, err := queryhelpers.GetProduct(id)
	helpers.HandleError("queryHelperGetProductError", err)

	if product != nil {
		_, err := helpers.RunQuery("DELETE FROM product WHERE product_id=$1", id)
		helpers.HandleError("runQueryError:", err)

		if err == nil {
			helpers.SendResponse(map[string]string{"message": "Successfully Deleted"}, w)
		}
	} else {
		helpers.SendResponse(map[string]string{"message": "Product not found"}, w)
	}

}
