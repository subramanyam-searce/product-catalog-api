package ops

import (
	"fmt"

	"github.com/subramanyam-searce/product-catalog-go/constants/responses"
	"github.com/subramanyam-searce/product-catalog-go/db/services"
	"github.com/subramanyam-searce/product-catalog-go/typedefs"
)

func GetCategoryFromConsole() (*typedefs.Category, error) {
	category := typedefs.Category{}

	category_id, err := GetPositiveFloatFromConsole(GetCategoryID)
	if err != nil {
		return	nil, err
	}

	category.CategoryID = int(category_id)

	category_name, err := ScanField(GetCategoryName)
	if err != nil {
		return nil, err
	}

	category.Name = category_name

	return &category, nil
}

func GetCategory() {
	category_id, err := GetPositiveFloatFromConsole(GetCategoryID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	category, err := services.GetCategory(int(category_id))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if category == nil {
		fmt.Println(responses.InvalidCategoryID)
		return
	}

	FormatPrintStruct(category)
}

func GetCategories() {
	categories, err := services.GetCategories()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, v := range *categories {
		FormatPrintStruct(v)
	}
}

func AddCategory() {
	category, err := GetCategoryFromConsole()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	response := services.AddCategory(*category)
	fmt.Println(response)
}

func DeleteCategory() {
	category_id, err := GetPositiveFloatFromConsole(GetCategoryID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	response := services.DeleteCategory(int(category_id))
	fmt.Println(response)
}

func UpdateCategory() {
	category, err := GetCategoryFromConsole()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	response := services.UpdateCategory(category.CategoryID, category.Name)
	fmt.Println(response)
}