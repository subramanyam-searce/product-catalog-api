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

func GetCategory(w http.ResponseWriter, r *http.Request) {
	id := helpers.ParseMuxVarToInt(r, "id")
	category, err := services.GetCategory(id)
	helpers.HandleError("GetCategoryError", err)
	if err != nil {
		helpers.SendJSONResponse(err.Error(), w)
		return
	}

	helpers.SendResponse(category, w)
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := services.GetCategories()

	if err != nil {
		helpers.SendJSONResponse(err.Error(), w)
		return
	}

	helpers.SendResponse(categories, w)
}

func AddCategory(w http.ResponseWriter, r *http.Request) {
	request_body, err := validators.ValidateRequestBody(r.Body, field_constraints.AddCategory)
	helpers.HandleError("validationError", err)
	if err != nil {
		helpers.SendJSONResponse(err.Error(), w)
		return
	}

	category := typedefs.Category{}
	err = json.Unmarshal(request_body, &category)
	helpers.HandleError("jsonUnmarshalError", err)

	response := services.AddCategory(category)

	helpers.SendJSONResponse(response, w)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id := helpers.ParseMuxVarToInt(r, "id")

	response := services.DeleteCategory(id)

	helpers.SendJSONResponse(response, w)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	var err error
	var input_json map[string]any
	id := helpers.ParseMuxVarToInt(r, "id")

	request_body, err := validators.ValidateRequestBody(r.Body, field_constraints.UpdateCategory)
	helpers.HandleError("validationError", err)
	if err != nil {
		helpers.SendJSONResponse(err.Error(), w)
		return
	}

	err = json.Unmarshal(request_body, &input_json)
	helpers.HandleError("decodingError", err)

	response := services.UpdateCategory(id, input_json)

	helpers.SendJSONResponse(response, w)
}
