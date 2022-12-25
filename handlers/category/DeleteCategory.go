package handlers_category

import (
	"net/http"

	"github.com/subramanyam-searce/product-catalog-go/helpers"
	"github.com/subramanyam-searce/product-catalog-go/queryhelpers"
)

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id := helpers.ParseMuxVarToInt(r, "id")
	category, err := queryhelpers.GetCategory(id)
	helpers.HandleError("getCategoryError", err)
	if err != nil {
		helpers.SendResponse(map[string]string{"message": err.Error()}, w)
		return
	}

	if category == nil {
		helpers.SendResponse(map[string]string{"message": "Category Not Found"}, w)
		return
	}

	query := "DELETE FROM category WHERE category_id=$1"
	_, err = helpers.RunQuery(query, id)
	helpers.HandleError("runQueryError", err)

	if err != nil {
		helpers.SendResponse(map[string]string{"message": err.Error()}, w)
	} else {
		helpers.SendResponse(map[string]string{"message": "Successfully deleted the category"}, w)
	}
}
