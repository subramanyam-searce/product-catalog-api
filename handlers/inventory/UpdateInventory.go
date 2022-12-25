package handlers_inventory

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/subramanyam-searce/product-catalog-go/helpers"
	"github.com/subramanyam-searce/product-catalog-go/queryhelpers"
)

func UpdateInventory(w http.ResponseWriter, r *http.Request) {
	var err error
	id := helpers.ParseMuxVarToInt(r, "id")
	json_input := map[string]any{}

	json.NewDecoder(r.Body).Decode(&json_input)

	inventory_item, err := queryhelpers.GetInventory(id)
	helpers.HandleError("getInventoryError", err)
	if inventory_item == nil {
		helpers.SendResponse(map[string]string{"message": "Inventory Item Not Found"}, w)
		return
	}

	for k, v := range json_input {
		if k == "product_id" {
			helpers.SendResponse(map[string]string{"message": "Inventory Item cannot be updated"}, w)
			return
		} else {
			err = queryhelpers.UpdateTableField("inventory", fmt.Sprintf("WHERE product_id=%v", id), k, fmt.Sprintf("%v", v))
			helpers.HandleError("updateTableFieldError", err)
			break
		}
	}

	if err != nil {
		helpers.SendResponse(map[string]string{"message": err.Error()}, w)
	} else {
		helpers.SendResponse(map[string]string{"message": "Successfully Updated"}, w)
	}
}
