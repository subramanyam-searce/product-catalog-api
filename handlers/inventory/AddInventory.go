package handlers_inventory

import (
	"encoding/json"
	"net/http"

	"github.com/subramanyam-searce/product-catalog-go/helpers"
	"github.com/subramanyam-searce/product-catalog-go/typedefs"
)

func AddInventory(w http.ResponseWriter, r *http.Request) {
	inventory_item := typedefs.Inventory{}
	json.NewDecoder(r.Body).Decode(&inventory_item)

	query := "INSERT INTO inventory VALUES($1, $2)"
	_, err := helpers.RunQuery(query, inventory_item.ProductID, inventory_item.Quantity)
	helpers.HandleError("runQueryError", err)

	if err != nil {
		helpers.SendResponse(map[string]string{"message": err.Error()}, w)
	} else {
		helpers.SendResponse(map[string]string{"message": "Inventory Items added successfully"}, w)
	}
}
