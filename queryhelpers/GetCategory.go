package queryhelpers

import (
	"github.com/subramanyam-searce/product-catalog-go/helpers"
	"github.com/subramanyam-searce/product-catalog-go/typedefs"
)

func GetCategory(id int) (*typedefs.Category, error) {
	query := "SELECT * FROM category WHERE category_id=$1;"
	var category *typedefs.Category = nil

	rows, err := helpers.RunQuery(query, id)

	if rows.Next() {
		category = &typedefs.Category{}
		rows.Scan(&category.CategoryID, &category.Name)
	}

	return category, err
}
