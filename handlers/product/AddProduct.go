package handlers_product

import (
	"encoding/json"
	"net/http"

	"github.com/subramanyam-searce/product-catalog-go/helpers"
	"github.com/subramanyam-searce/product-catalog-go/typedefs"
)

func AddProduct(w http.ResponseWriter, r *http.Request) {
	product := typedefs.Product{}
	json.NewDecoder(r.Body).Decode(&product)

	query := "INSERT INTO product VALUES($1, $2, $3, $4, $5, $6)"
	spec_json_str, err := json.Marshal(product.Specification)
	helpers.HandleError("jsonMarshalError", err)

	_, err = helpers.RunQuery(query, product.Product_ID,
		product.Name, spec_json_str, product.SKU, product.CategoryID, product.Price)

	response := map[string]string{"message": ""}
	if err == nil {
		response["message"] = "Product Added Successfully"
	} else {
		response["message"] = err.Error()
	}
	helpers.SendResponse(response, w)

}
