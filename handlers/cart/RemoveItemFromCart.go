package handlers_cart

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/subramanyam-searce/product-catalog-go/helpers"
	"github.com/subramanyam-searce/product-catalog-go/typedefs"
)

func RemoveItemFromCart(w http.ResponseWriter, r *http.Request) {
	urlQuery := r.URL.Query()

	ref := urlQuery.Get("ref")
	product_id := urlQuery.Get("product_id")
	quantity_str := urlQuery.Get("quantity")

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

	rows, err = helpers.RunQuery("SELECT product_id, quantity FROM cart_item WHERE ref=$1 AND product_id=$2;", ref, product_id)
	helpers.HandleError("runQueryError:", err)

	if !rows.Next() {
		helpers.SendResponse(map[string]string{"message": "Product is not found in your cart"}, w)
		return
	}

	cart_item := typedefs.CartItem{}
	err = rows.Scan(&cart_item.ProductID, &cart_item.Quantity)
	helpers.HandleError("rowsScanError", err)

	if quantity_str == "" {
		quantity_str = fmt.Sprint(cart_item.Quantity)
	}

	quantity, err := strconv.Atoi(quantity_str)
	helpers.HandleError("strconvError", err)

	if cart_item.Quantity-quantity < 0 {
		helpers.SendResponse(map[string]string{"message": fmt.Sprintf("Cannot remove %v from available cart quantity %v", quantity, cart_item.Quantity)}, w)
		return
	}

	if cart_item.Quantity-quantity == 0 {
		_, err = helpers.RunQuery("DELETE FROM cart_item WHERE ref=$1 AND product_id=$2", ref, product_id)
		helpers.HandleError("runQueryError", err)

		_, err = helpers.RunQuery("INSERT INTO inventory VALUES($1, $2) ON CONFLICT ON CONSTRAINT inventory_pkey DO UPDATE SET quantity=(SELECT quantity FROM inventory WHERE product_id=$1)+$2", product_id, quantity)
		helpers.HandleError("runQueryError", err)

		helpers.SendResponse(map[string]string{"message": "All quantity of this product was removed from your cart"}, w)
		return
	}

	if cart_item.Quantity-quantity > 0 {
		_, err = helpers.RunQuery("UPDATE cart_item SET quantity=$1 WHERE ref=$2 AND product_id=$3", cart_item.Quantity-quantity, ref, product_id)
		helpers.HandleError("runQueryError", err)

		_, err = helpers.RunQuery("INSERT INTO inventory VALUES($1, $2) ON CONFLICT ON CONSTRAINT inventory_pkey DO UPDATE SET quantity=(SELECT quantity FROM inventory WHERE product_id=$1)+$2", product_id, quantity)
		helpers.HandleError("runQueryError", err)

		helpers.SendResponse(map[string]string{"message": "Successfully removed"}, w)
		return
	}
}
