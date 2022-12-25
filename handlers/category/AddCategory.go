package handlers_category

import (
	"encoding/json"
	"net/http"

	"github.com/subramanyam-searce/product-catalog-go/helpers"
	"github.com/subramanyam-searce/product-catalog-go/typedefs"
)

func AddCategory(w http.ResponseWriter, r *http.Request) {
	category := typedefs.Category{}
	json.NewDecoder(r.Body).Decode(&category)

	query := "INSERT INTO category VALUES($1, $2)"
	_, err := helpers.RunQuery(query, category.CategoryID, category.Name)
	helpers.HandleError("runQueryError", err)

	if err != nil {
		helpers.SendResponse(map[string]string{"message": err.Error()}, w)
	} else {
		helpers.SendResponse(map[string]string{"message": "Category added successfully"}, w)
	}
}
