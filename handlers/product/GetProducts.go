package handlers_product

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/subramanyam-searce/product-catalog-go/helpers"
	"github.com/subramanyam-searce/product-catalog-go/typedefs"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	products := []typedefs.Product{}
	vars := mux.Vars(r)
	page_no_str := vars["page_no"]
	page_no, err := strconv.Atoi(page_no_str)
	helpers.HandleError("strconvError", err)

	n := 5
	start_index_no := ((page_no - 1) * n)

	rows, err := helpers.RunQuery("SELECT * FROM product")
	helpers.HandleError("runQueryError:", err)

	for rows.Next() {
		newProduct := typedefs.Product{}
		spec_json := ""
		rows.Scan(&newProduct.Product_ID, &newProduct.Name, &spec_json, &newProduct.SKU, &newProduct.CategoryID, newProduct.Price)
		json.Unmarshal([]byte(spec_json), &newProduct.Specification)
		products = append(products, newProduct)
	}

	min_products := []map[string]any{}

	for _, v := range products {
		new_min_product := map[string]any{}
		new_min_product["name"] = v.Name
		new_min_product["product_id"] = v.Product_ID
		new_min_product["price"] = v.Price
		min_products = append(min_products, new_min_product)
	}

	end_index_no := int(math.Min(float64(start_index_no+n), float64(len(min_products))))

	if start_index_no <= (len(min_products)-1) && start_index_no >= 0 {
		helpers.SendResponse(min_products[start_index_no:end_index_no], w)
	} else {
		max_pages := ((end_index_no - 1) / n) + 1
		if end_index_no == 0 {
			max_pages = 0
		}
		helpers.SendResponse(map[string]string{"message": "Page not found. Max page number is " + fmt.Sprint(max_pages)}, w)
	}
}
