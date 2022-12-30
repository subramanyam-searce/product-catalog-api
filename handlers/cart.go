package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/subramanyam-searce/product-catalog-go/constants/field_constraints"
	"github.com/subramanyam-searce/product-catalog-go/constants/responses"
	"github.com/subramanyam-searce/product-catalog-go/db/services"
	"github.com/subramanyam-searce/product-catalog-go/handlers/validators"
	"github.com/subramanyam-searce/product-catalog-go/helpers"
)

type NewCartJSONResponse struct {
	Message       string `json:"message"`
	CartReference string `json:"ref"`
}

func AddItemsToCart(w http.ResponseWriter, r *http.Request) {
	response := []map[string]any{}
	request_body := []map[string]any{}

	ref := r.URL.Query().Get("ref")

	if ref == "" {
		helpers.SendJSONResponse(responses.InvalidCart, w)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&request_body)
	if err != nil {
		helpers.SendJSONResponse(responses.BadRequestBody, w)
		return
	}

	for _, v := range request_body {
		new_response_item := map[string]any{}
		response = append(response, new_response_item)

		product_id := v["product_id"]
		quantity := v["quantity"]
		new_response_item["product_id"] = product_id
		new_response_item["quantity"] = quantity

		err := validators.ValidateRequestBodyMap(v, field_constraints.AddItemToCart)
		helpers.HandleError("validationError", err)

		if err != nil {
			new_response_item["message"] = err.Error()
			continue
		}

		ref, new_response_item["message"], _ = services.AddItemToCart(ref, int(product_id.(float64)), int(quantity.(float64)))
	}

	helpers.SendResponse(response, w)
}

func AddToCart(w http.ResponseWriter, r *http.Request) {
	urlQuery := r.URL.Query()

	ref := urlQuery.Get("ref")

	request_body, err := validators.ValidateRequestBody(r.Body, field_constraints.AddItemToCart)
	helpers.HandleError("validationError", err)

	if err != nil {
		helpers.SendJSONResponse(err.Error(), w)
		return
	}

	input_json := map[string]int{}
	err = json.Unmarshal(request_body, &input_json)
	helpers.HandleError("jsonUnmarshalError", err)

	ref, response, isNewCart := services.AddItemToCart(ref, input_json["product_id"], input_json["quantity"])

	if isNewCart {
		helpers.SendResponse(NewCartJSONResponse{Message: response, CartReference: ref}, w)
		return
	}

	helpers.SendJSONResponse(response, w)
}

func GetCart(w http.ResponseWriter, r *http.Request) {
	ref := r.URL.Query().Get("ref")

	cart, err := services.GetCart(ref)
	helpers.HandleError("getCartError", err)

	if err != nil {
		helpers.SendJSONResponse(err.Error(), w)
		return
	}

	helpers.SendResponse(cart, w)
}

func RemoveItemFromCart(w http.ResponseWriter, r *http.Request) {
	ref := r.URL.Query().Get("ref")

	request_body, err := validators.ValidateRequestBody(r.Body, field_constraints.RemoveItemFromCart)
	helpers.HandleError("validationError", err)

	if err != nil {
		helpers.SendJSONResponse(err.Error(), w)
		return
	}

	input_json := map[string]any{}
	err = json.Unmarshal(request_body, &input_json)
	helpers.HandleError("jsonUnmarshalError", err)

	response := services.RemoveItemFromCart(ref, int(input_json["product_id"].(float64)), int(input_json["quantity"].(float64)))
	helpers.SendJSONResponse(response, w)
}
