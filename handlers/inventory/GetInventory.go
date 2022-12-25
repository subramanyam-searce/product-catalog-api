package handlers_inventory

import (
	"net/http"

	"github.com/subramanyam-searce/product-catalog-go/helpers"
	"github.com/subramanyam-searce/product-catalog-go/typedefs"
)

func GetInventory(w http.ResponseWriter, r *http.Request) {
	inventory := []typedefs.ProductInventory{}

	query := "SELECT p.product_id, p.name, i.quantity FROM product p INNER JOIN inventory i ON p.product_id=i.product_id"
	rows, err := helpers.RunQuery(query)
	helpers.HandleError("runQueryError", err)

	for rows.Next() {
		newInventoryItem := typedefs.ProductInventory{}
		rows.Scan(&newInventoryItem.ProductID, &newInventoryItem.ProductName, &newInventoryItem.Quantity)
		inventory = append(inventory, newInventoryItem)
	}

	if err != nil {
		helpers.SendResponse(map[string]string{"message": err.Error()}, w)
	} else {
		helpers.SendResponse(inventory, w)
	}
}
