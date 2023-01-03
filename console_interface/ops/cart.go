package ops

import (
	"fmt"

	"github.com/subramanyam-searce/product-catalog-go/db/services"
)

func GetCart() {
	ref, err := ScanField(GetCartReference)
	if err != nil {
		fmt.Println(err)
		return
	}

	cart, err := services.GetCart(ref)
	if err != nil {
		fmt.Println(err)
		return
	}

	FormatPrintStruct(cart)
	fmt.Println(CartItemsDisplay)
	for _, v := range cart.Items {
		FormatPrintStruct(v)
	}
}

func AddItemToCartHelper(ref *string) {
	product_id, err := GetPositiveFloatFromConsole(GetProductID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	quantity, err := GetPositiveFloatFromConsole(GetQuantity)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ref_NewCart, response, isNewCart:= services.AddItemToCart(*ref, int(product_id), int(quantity))
	*ref = ref_NewCart

	if isNewCart {
		fmt.Printf(NewCartCreatedDisplay, *ref)
	}

	fmt.Println(response)
}

func AddItemToCart() {
	ref_str, err := ScanField(GetCartReference)
	ref := &ref_str

	if err != nil {
		fmt.Println(err)
		return
	}

	AddItemToCartHelper(ref)
}

func AddItemsToCart() {
	ref, err := ScanField(GetCartReference)
	if err != nil {
		fmt.Println(err)
		return
	}

	number_of_cart_items, err := GetPositiveFloatFromConsole(GetCartItemNumberDisplay)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < int(number_of_cart_items); i++ {
		AddItemToCartHelper(&ref)
	}
}

func RemoveItemFromCart() {
	ref, err := ScanField(GetCartReference)
	if err != nil {
		fmt.Println(err)
		return
	}

	product_id, err := GetPositiveFloatFromConsole(GetProductID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	quantity, err := GetPositiveFloatFromConsole(GetQuantity)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	response := services.RemoveItemFromCart(ref, int(product_id), int(quantity))
	fmt.Println(response)
}