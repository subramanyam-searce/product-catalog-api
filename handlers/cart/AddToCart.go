package handlers_cart

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/subramanyam-searce/product-catalog-go/helpers"
	"github.com/subramanyam-searce/product-catalog-go/typedefs"
)

func AddToCart(w http.ResponseWriter, r *http.Request) {
	urlQuery := r.URL.Query()

	ref := urlQuery.Get("ref")
	quantity_str := urlQuery.Get("quantity")
	product_id := urlQuery.Get("product_id")

	if ref == "" || quantity_str == "" || product_id == "" {
		helpers.SendResponse(map[string]string{"message": "ref / quantity / product_id missing in the url"}, w)
		return
	}

	quantity, err := strconv.Atoi(quantity_str)
	helpers.HandleError("strconvError:", err)

	rows, err := helpers.RunQuery("SELECT * FROM cart_reference WHERE ref=$1;", ref)
	helpers.HandleError("runQueryError:", err)

	if !rows.Next() {
		helpers.SendResponse(map[string]string{"message": "Invalid cart_reference"}, w)
		return
	}

	rows, err = helpers.RunQuery("SELECT p.product_id, i.quantity, p.name FROM product p LEFT JOIN inventory i ON p.product_id=i.product_id WHERE p.product_id=$1", product_id)
	helpers.HandleError("runQueryError", err)

	if !rows.Next() {
		helpers.SendResponse(map[string]string{"message": "Product id is invalid"}, w)
		return
	}

	inventory_item := typedefs.ProductInventory{}
	err = rows.Scan(&inventory_item.ProductID, &inventory_item.Quantity, &inventory_item.ProductName)
	helpers.HandleError("rowsScanError", err)

	if inventory_item.Quantity-quantity < 0 {
		helpers.SendResponse(map[string]string{"message": "Inventory Quantity is less than the required quantity: " + fmt.Sprint(inventory_item.Quantity)}, w)
		return
	}

	_, err = helpers.RunQuery("UPDATE inventory SET quantity=$1 WHERE product_id=$2", inventory_item.Quantity-quantity, product_id)
	helpers.HandleError("runQueryError:", err)

	rows, err = helpers.RunQuery("SELECT quantity FROM cart_item WHERE ref=$1 AND product_id=$2", ref, product_id)
	helpers.HandleError("runQueryError:", err)

	if rows.Next() {
		var db_quantity int
		rows.Scan(&db_quantity)

		_, err = helpers.RunQuery("UPDATE cart_item SET quantity=$1 WHERE ref=$2 AND product_id=$3", db_quantity+quantity, ref, product_id)
		helpers.HandleError("runQueryError:", err)

	} else {
		_, err = helpers.RunQuery("INSERT INTO cart_item VALUES($1, $2, $3);", ref, product_id, quantity)
		helpers.HandleError("runQueryError:", err)
	}

	if err != nil {
		helpers.SendResponse(map[string]string{"message": err.Error()}, w)
		return
	}

	helpers.SendResponse(map[string]string{"message": "Item was added to the cart"}, w)
}
