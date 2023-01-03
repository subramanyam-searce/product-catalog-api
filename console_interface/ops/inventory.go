package ops

import (
	"fmt"
	"strconv"

	"github.com/subramanyam-searce/product-catalog-go/db/services"
)

func GetInventory() {
	inventory, err := services.GetInventory()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range *inventory {
		FormatPrintStruct(v)
	}
}

func UpdateInventory() {
	product_id, err := GetPositiveFloatFromConsole(GetProductID)
	if err != nil {
		fmt.Println(err)
		return
	}

	quantity_str, err := ScanField(GetQuantity)
	if err != nil {
		fmt.Println(err)
		return
	}

	quantity, err := strconv.Atoi(quantity_str)
	if err != nil {
		fmt.Println(EnterValidFloat)
		return
	}

	response := services.UpdateInventory(int(product_id), quantity)
	fmt.Println(response)
}