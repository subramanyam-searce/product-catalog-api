package services

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/subramanyam-searce/product-catalog-go/constants/queries"
	"github.com/subramanyam-searce/product-catalog-go/constants/responses"
	"github.com/subramanyam-searce/product-catalog-go/helpers"
	"github.com/subramanyam-searce/product-catalog-go/typedefs"
)

func GetProduct(id int) (*typedefs.Product, error) {
	var product *typedefs.Product = nil
	var err error

	if id <= 0 {
		return nil, errors.New(responses.ProductIDNotPositive)
	}

	rows, err := helpers.RunQuery(queries.GetProduct, id)
	if err != nil {
		return nil, err
	}

	var spec_byte_slice []byte

	if rows.Next() {
		product = &typedefs.Product{}
		err = rows.Scan(&product.Product_ID, &product.Name, &spec_byte_slice, &product.SKU, &product.CategoryID, &product.Price)
		if err != nil {
			return nil, err
		}	

		err := json.Unmarshal(spec_byte_slice, &product.Specification)
		if err != nil {
			return nil, err
		}
	}

	return product, nil
}

func GetProducts(page_no int, items_per_page int) ([]typedefs.ShortProduct, error) {
	products := []typedefs.Product{}

	if page_no <= 0 {
		return nil, errors.New(responses.InvalidPageNo)
	}

	if items_per_page <= 0 {
		return nil, errors.New(responses.InvalidItemsPerPage)
	}

	rows, err := helpers.RunQuery(queries.GetAllProducts)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		newProduct := typedefs.Product{}
		spec_json := ""
		err := rows.Scan(&newProduct.Product_ID, &newProduct.Name, &spec_json, &newProduct.SKU, &newProduct.CategoryID, &newProduct.Price)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal([]byte(spec_json), &newProduct.Specification)
		if err != nil {
			return nil, err
		}

		products = append(products, newProduct)
	}

	min_products := []typedefs.ShortProduct{}

	for _, v := range products {
		new_min_product := typedefs.ShortProduct{Product_ID: v.Product_ID, Name: v.Name, Price: v.Price}
		min_products = append(min_products, new_min_product)
	}

	start_index, end_index, err := helpers.Paginate(page_no, len(min_products), items_per_page)

	if err != nil {
		return nil, err
	} else {
		return min_products[start_index:end_index], nil
	}
}

func DeleteProduct(product_id int) string {
	var err error
	product, err := GetProduct(product_id)
	helpers.HandleError("queryHelperGetProductError", err)
	if err != nil {
		return err.Error()
	}

	_, err = helpers.RunQuery(queries.DeleteInventoryItem, product_id)
	if err != nil {
		return err.Error()
	}

	_, err = helpers.RunQuery(queries.DeleteCartItemsWithProductID, product_id)
	if err != nil {
		return err.Error()
	}

	if product != nil {
		_, err := helpers.RunQuery(queries.DeleteProduct, product_id)

		if err == nil {
			return responses.ProductSuccessfullyDeleted
		} else {
			return err.Error()
		}
	} else {
		return responses.ProductNotFound
	}
}

func AddProduct(product typedefs.Product) string {
	existing_product, err := GetProduct(product.Product_ID)
	helpers.HandleError("getProductError", err)
	if err != nil {
		return err.Error()
	}

	if product.Price <= 0 {
		return responses.PriceNotPositive
	}

	if product.CategoryID <= 0 {
		return responses.CategoryIDNotPositive
	}

	if existing_product != nil {
		return responses.ProductAlreadyExists
	}

	spec_json_str, err := json.Marshal(product.Specification)
	helpers.HandleError("jsonMarshalError", err)
	if err != nil {
		return responses.BadRequestBody
	}

	_, err = helpers.RunQuery(queries.InsertProduct, product.Product_ID,
		product.Name, spec_json_str, product.SKU, product.CategoryID, product.Price)
	if err != nil {
		return responses.InvalidCategoryIDForProductFKConstraint
	}

	return responses.ProductAddedSuccessfully
}

func updateProductTableField(product_id int, fieldName string, val string) error {
	query := "UPDATE product SET " + fieldName + "=$1 WHERE product_id=$2;"
	_, err := helpers.RunQuery(query, val, fmt.Sprint(product_id))

	return err
}

func UpdateProduct(product_id int, to_update map[string]any) string {
	if len(to_update) == 0 {
		return responses.EmptyInputJson
	}

	if to_update["price"] != nil && to_update["price"].(float64) < 0 {
		return responses.PriceNotPositive
	}

	product, err := GetProduct(product_id)
	helpers.HandleError("getProductError", err)
	if err != nil {
		return err.Error()
	}

	if product == nil {
		return responses.ProductNotFound
	}

	for k, v := range to_update {
		if k == "product_id" {
			return responses.ProductIDCannotBeUpdated
		}
		err = updateProductTableField(product_id, k, fmt.Sprint(v))
		helpers.HandleError("updateTableFieldError", err)
		if err != nil {
			return responses.InvalidCategoryIDForProductFKConstraint
		}
	}

	return responses.ProductUpdatedSuccessfully
}
