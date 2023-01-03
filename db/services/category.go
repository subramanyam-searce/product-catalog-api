package services

import (
	"errors"
	"fmt"

	"github.com/subramanyam-searce/product-catalog-go/constants/queries"
	"github.com/subramanyam-searce/product-catalog-go/constants/responses"
	"github.com/subramanyam-searce/product-catalog-go/helpers"
	"github.com/subramanyam-searce/product-catalog-go/typedefs"
)

func GetCategories() (*[]typedefs.Category, error) {
	categories := []typedefs.Category{}

	rows, err := helpers.RunQuery(queries.GetAllCategories)
	helpers.HandleError("runQueryError", err)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		category := typedefs.Category{}
		err := rows.Scan(&category.CategoryID, &category.Name)
		helpers.HandleError("rowsScanError", err)
		categories = append(categories, category)
	}

	return &categories, nil
}

func GetCategory(category_id int) (*typedefs.Category, error) {
	var category *typedefs.Category

	if category_id <= 0 {
		return nil, errors.New(responses.CategoryIDNegative)
	}

	rows, err := helpers.RunQuery(queries.GetCategory, category_id)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		category = &typedefs.Category{}
		rows.Scan(&category.CategoryID, &category.Name)
		return category, nil
	}

	return nil, errors.New(responses.InvalidCategoryID)
}

func DeleteCategory(category_id int) string {

	if category_id <= 0 {
		return responses.CategoryIDNegative
	}

	category, err := GetCategory(category_id)
	helpers.HandleError("getCategoryError", err)
	if err != nil {
		return err.Error()
	}

	if category == nil {
		return responses.InvalidCategoryID
	}

	_, err = helpers.RunQuery(queries.DeleteCategory, category_id)
	helpers.HandleError("runQueryError", err)
	if err != nil {
		return responses.CategoryIDUsedByProduct
	}

	return responses.CategorySuccessfullyDeleted
}

func AddCategory(category typedefs.Category) string {

	existing_category, err := GetCategory(category.CategoryID)
	helpers.HandleError("getCategoryError", err)
	if category.CategoryID <= 0 {
		return responses.CategoryIDNegative
	}

	if len(category.Name) < 3 {
		return responses.CategoryNameLengthError
	}

	if existing_category != nil {
		return responses.CategoryAlreadyExist
	}

	_, err = helpers.RunQuery(queries.AddCategory, category.CategoryID, category.Name)
	helpers.HandleError("runQueryError", err)
	if err != nil {
		return err.Error()
	}

	return responses.CategoryAddedSuccessfully
}

func updateCategoryTableField(category_id int, fieldName string, val string) error {
	query := "UPDATE category SET " + fieldName + "=$1 WHERE category_id=$2;"
	_, err := helpers.RunQuery(query, val, fmt.Sprint(category_id))

	return err
}

func UpdateCategory(category_id int, name string) string {

	category, err := GetCategory(category_id)
	helpers.HandleError("getCategoryError", err)
	if err != nil {
		return err.Error()
	}

	if category == nil {
		return responses.InvalidCategoryID
	}

	if len(name) < 3 {
		return responses.CategoryNameLengthError
	}

	err = updateCategoryTableField(category_id, "name", name)
	helpers.HandleError("updateTableFieldError", err)
	if err != nil {
		return err.Error()
	}

	return responses.CategoryUpdatedSuccessfully
}
