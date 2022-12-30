package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/subramanyam-searce/product-catalog-go/constants/field_constraints"
	"github.com/subramanyam-searce/product-catalog-go/constants/responses"
	"github.com/subramanyam-searce/product-catalog-go/db/services"
	"github.com/subramanyam-searce/product-catalog-go/handlers/validators"
	"github.com/subramanyam-searce/product-catalog-go/helpers"
	"github.com/subramanyam-searce/product-catalog-go/typedefs"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	product_id := helpers.ParseMuxVarToInt(r, "id")

	product, err := services.GetProduct(product_id)
	helpers.HandleError("GetProductQueryHelperError", err)

	if product != nil {
		helpers.SendResponse(product, w)
	} else {
		helpers.SendJSONResponse(responses.ProductNotFound, w)
	}
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	page_no_str := r.URL.Query().Get("page")
	items_per_page_str := r.URL.Query().Get("items_per_page")

	if page_no_str == "" {
		page_no_str = "1"
	}

	if items_per_page_str == "" {
		items_per_page_str = "20"
	}

	page_no, err := strconv.Atoi(page_no_str)
	helpers.HandleError("stronvAtoiError", err)
	if err != nil {
		helpers.SendJSONResponse(responses.InvalidPageNo, w)
		return
	}

	items_per_page, err := strconv.Atoi(items_per_page_str)
	helpers.HandleError("stronvAtoiError", err)
	if err != nil {
		helpers.SendJSONResponse(responses.InvalidItemsPerPage, w)
		return
	}

	products, err := services.GetProducts(page_no, items_per_page)
	helpers.HandleError("GetProductsError", err)

	if err != nil {
		helpers.SendJSONResponse(err.Error(), w)
	} else {
		helpers.SendResponse(products, w)
	}
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := helpers.ParseMuxVarToInt(r, "id")

	response := services.DeleteProduct(id)

	helpers.SendJSONResponse(response, w)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	request_body, err := validators.ValidateRequestBody(r.Body, field_constraints.AddProduct)
	helpers.HandleError("validationError", err)
	if err != nil {
		helpers.SendJSONResponse(err.Error(), w)
		return
	}

	product := typedefs.Product{}
	err = json.Unmarshal(request_body, &product)
	helpers.HandleError("jsonUnmarshalError", err)

	response := services.AddProduct(product)

	helpers.SendJSONResponse(response, w)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var err error
	var input_json map[string]any
	id := helpers.ParseMuxVarToInt(r, "id")

	request_body, err := validators.ValidateRequestBody(r.Body, field_constraints.UpdateProduct)
	helpers.HandleError("validationError", err)
	if err != nil {
		helpers.SendJSONResponse(err.Error(), w)
		return
	}

	err = json.Unmarshal(request_body, &input_json)
	helpers.HandleError("decodingError", err)

	response := services.UpdateProduct(id, input_json)

	helpers.SendJSONResponse(response, w)
}
