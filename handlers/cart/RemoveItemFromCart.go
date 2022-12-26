package handlers_cart

import (
	"net/http"

	"github.com/subramanyam-searce/product-catalog-go/helpers"
)

func RemoveItemFromCart(w http.ResponseWriter, r *http.Request) {
	urlQuery := r.URL.Query()

	ref := urlQuery.Get("ref")
	product_id := urlQuery.Get("product_id")

	if ref == "" || product_id == "" {
		helpers.SendResponse(map[string]string{"message": "ref / product_id missing in the url"}, w)
		return
	}

	rows, err := helpers.RunQuery("SELECT * FROM cart_reference WHERE ref=$1;", ref)
	helpers.HandleError("runQueryError:", err)

	if !rows.Next() {
		helpers.SendResponse(map[string]string{"message": "Invalid cart_reference"}, w)
		return
	}

	result, err := helpers.ConnectToDB().Exec("DELETE FROM cart_item WHERE ref=$1 AND product_id=$2", ref, product_id)
	helpers.HandleError("runQueryError", err)

	rows_affected, err := result.RowsAffected()
	helpers.HandleError("rowsAffectedError", err)

	if rows_affected != 0 {
		if err != nil {
			helpers.SendResponse(map[string]string{"message": err.Error()}, w)
		} else {
			helpers.SendResponse(map[string]string{"message": "Cart item deleted successfully"}, w)
		}
	} else {
		helpers.SendResponse(map[string]string{"message": "Product is not found in your cart"}, w)
	}
}
