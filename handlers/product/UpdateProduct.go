package handlers_product

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/subramanyam-searce/product-catalog-go/helpers"
	"github.com/subramanyam-searce/product-catalog-go/queryhelpers"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var err error
	var input_json map[string]any
	id := helpers.ParseMuxVarToInt(r, "id")

	err = json.NewDecoder(r.Body).Decode(&input_json)
	helpers.HandleError("decodingError", err)

	product, err := queryhelpers.GetProduct(id)

	if product == nil {
		helpers.SendResponse(map[string]string{"message": "Product not found"}, w)
		return
	}

	for k, v := range input_json {
		if k != "product_id" {
			if k == "specification" {
				v, err = json.Marshal(v)
				v = string(v.([]uint8))
				helpers.HandleError("jsonMarshalError", err)
			}
			err = queryhelpers.UpdateTableField("product", fmt.Sprintf("WHERE product_id=%v", id), k, fmt.Sprint(v))
			helpers.HandleError("updateTableFieldError", err)

			if err != nil {
				helpers.SendResponse(map[string]string{"message": err.Error()}, w)
				return
			}
		} else {
			err = errors.New("product_id cannot be updated")
			helpers.SendResponse(map[string]string{"message": err.Error()}, w)
			return
		}
	}

	if err == nil {
		helpers.SendResponse(map[string]string{"message": "Update done successfully"}, w)
	}
}
