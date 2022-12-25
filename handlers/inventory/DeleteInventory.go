package handlers_inventory

import (
	"net/http"

	"github.com/subramanyam-searce/product-catalog-go/helpers"
	"github.com/subramanyam-searce/product-catalog-go/queryhelpers"
)

func DeleteInventory(w http.ResponseWriter, r *http.Request) {
	id := helpers.ParseMuxVarToInt(r, "id")

	inventoryItem, err := queryhelpers.GetInventory(id)

	helpers.HandleError("getInventoryError", err)

	if err != nil {
		helpers.SendResponse(map[string]string{"message": err.Error()}, w)
		return
	}

	if inventoryItem == nil {
		helpers.SendResponse(map[string]string{"message": "Inventory Item Not Found"}, w)
		return
	}

	query := "DELETE FROM inventory WHERE product_id=$1"
	_, err = helpers.RunQuery(query, id)
	helpers.HandleError("runQueryError", err)

	if err != nil {
		helpers.SendResponse(map[string]string{"message": err.Error()}, w)
	} else {
		helpers.SendResponse(map[string]string{"message": "Successfully deleted the Inventory Item"}, w)
	}
}
