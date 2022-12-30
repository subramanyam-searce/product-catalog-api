package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/subramanyam-searce/product-catalog-go/constants/field_constraints"
	"github.com/subramanyam-searce/product-catalog-go/db/services"
	"github.com/subramanyam-searce/product-catalog-go/handlers/validators"
	"github.com/subramanyam-searce/product-catalog-go/helpers"
	"github.com/subramanyam-searce/product-catalog-go/typedefs"
)

func GetInventory(w http.ResponseWriter, r *http.Request) {
	inventory, err := services.GetInventory()
	helpers.HandleError("GetInventoryError", err)

	if err != nil {
		helpers.SendJSONResponse(err.Error(), w)
		return
	}

	helpers.SendResponse(inventory, w)
}

func UpdateInventory(w http.ResponseWriter, r *http.Request) {
	request_body, err := validators.ValidateRequestBody(r.Body, field_constraints.UpdateInventory)
	helpers.HandleError("validationError", err)
	if err != nil {
		helpers.SendJSONResponse(err.Error(), w)
		return
	}

	inventory_item := typedefs.InventoryItem{}
	err = json.Unmarshal(request_body, &inventory_item)
	helpers.HandleError("jsonUnmarshalError", err)

	response := services.UpdateInventory(inventory_item.ProductID, inventory_item.Quantity)

	helpers.SendJSONResponse(response, w)
}
