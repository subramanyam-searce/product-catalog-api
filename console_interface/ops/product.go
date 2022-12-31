package ops

import (
	"fmt"

	"github.com/fatih/structs"
	"github.com/subramanyam-searce/product-catalog-go/constants/field_constraints"
	"github.com/subramanyam-searce/product-catalog-go/constants/responses"
	"github.com/subramanyam-searce/product-catalog-go/db/services"
	"github.com/subramanyam-searce/product-catalog-go/handlers/validators"
	"github.com/subramanyam-searce/product-catalog-go/typedefs"
)

func GetProduct() {
	product_id, err := GetPositiveFloatFromConsole(GetProductID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	product, err := services.GetProduct(int(product_id))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if product == nil {
		fmt.Println(responses.ProductNotFound)
		return
	}

	FormatPrintStruct(product)
}

func GetProducts() {
	items_per_page, err := GetPositiveFloatFromConsole(GetItemsPerPage)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	page_no, err := GetPositiveFloatFromConsole(GetPageNo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	products, err := services.GetProducts(int(page_no), int(items_per_page))
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range products {
		FormatPrintStruct(v)
	}
}

func AddProduct() {
	product := typedefs.Product{}
	product_struct := structs.New(&product)
	err := GetInputBody(product_struct, field_constraints.AddProduct)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	response := services.AddProduct(product)
	fmt.Println(response)
}

func DeleteProduct() {
	product_id, err := GetPositiveFloatFromConsole(GetProductID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	response := services.DeleteProduct(int(product_id))
	fmt.Println(response)
}

func UpdateProduct() {
	product_id, err := GetPositiveFloatFromConsole(GetProductID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	key, err := ScanField(EnterFieldNameToUpdate)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	field, ok := structs.New(typedefs.Product{}).FieldOk(key)
	if !ok {
		fmt.Println(FieldNotFoundForUpdation)
		return
	}

	key = field.Tag("json")

	val, err := ScanField(EnterTheValueToUpdate)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	to_update := map[string]any{key: val}

	err = validators.ValidateRequestBodyMap(to_update, field_constraints.UpdateProduct)
	if err != nil {
		fmt.Println(err)
		return
	}

	response := services.UpdateProduct(int(product_id), to_update)

	fmt.Println(response)
}
