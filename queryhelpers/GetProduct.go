package queryhelpers

import (
	"encoding/json"

	"github.com/subramanyam-searce/product-catalog-go/helpers"
	"github.com/subramanyam-searce/product-catalog-go/typedefs"
)

func GetProduct(id int) (*typedefs.CategoryProduct, error) {
	var product *typedefs.CategoryProduct = nil
	var err error
	rows, err := helpers.RunQuery("SELECT * FROM product p INNER JOIN category c ON p.category_id=c.category_id WHERE product_id=$1", id)
	helpers.HandleError("runQueryError", err)
	if err != nil {
		return nil, err
	}

	var spec_byte_slice []byte

	if rows.Next() {
		product = &typedefs.CategoryProduct{}
		err = rows.Scan(&product.Product_ID, &product.Name, &spec_byte_slice, &product.SKU, &product.CategoryID, &product.Price, &product.CategoryID, &product.CategoryName)
		json.Unmarshal(spec_byte_slice, &product.Specification)
	}
	helpers.HandleError("rowsScanError", err)
	return product, err
}
