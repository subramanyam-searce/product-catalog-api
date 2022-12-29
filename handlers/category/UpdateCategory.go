package handlers_category

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/subramanyam-searce/product-catalog-go/helpers"
	"github.com/subramanyam-searce/product-catalog-go/queryhelpers"
)

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	var err error
	id := helpers.ParseMuxVarToInt(r, "id")
	json_input := map[string]any{}

	json.NewDecoder(r.Body).Decode(&json_input)

	category, err := queryhelpers.GetCategory(id)
	helpers.HandleError("getCategoryError", err)

	if category == nil {
		helpers.SendResponse(map[string]string{"message": "Category Not Found"}, w)
		return
	}

	for k, v := range json_input {
		if k == "category_id" {
			helpers.SendResponse(map[string]string{"message": "Category ID cannot be updated"}, w)
			return
		} else {
			err = queryhelpers.UpdateTableField("category", fmt.Sprintf("WHERE category_id=%v", id), k, fmt.Sprintf("%v", v))
			helpers.HandleError("updateTableFieldError", err)
		}
	}

	if err != nil {
		helpers.SendResponse(map[string]string{"message": err.Error()}, w)
	} else {
		helpers.SendResponse(map[string]string{"message": "Successfully Updated"}, w)
	}
}
