package tests

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/subramanyam-searce/product-catalog-go/constants/responses"
	"github.com/subramanyam-searce/product-catalog-go/db/services"
	"github.com/subramanyam-searce/product-catalog-go/helpers"
)

var ExpectedGot string = "%v: Expected: %v, Got: %v"

func init() {
	connection_string := os.Getenv("TESTING_DB_CONNECTION_STRING")
	helpers.DB = helpers.EstablishDBConnection(connection_string)
}

func IsErrorSame(err1 error, err2 error) bool {
	var err1String string
	var err2String string
	
	if err1 == nil {
		err1String = "nil"
	} else {
		err1String = err1.Error()
	}

	if err2 == nil {
		err2String = "nil"
	} else {
		err2String = err2.Error()
	}

	return err1String == err2String
}

func TestGetProduct(t *testing.T) {
	RestoreDBToTestingState()
	
	for _, v := range GetProductTestCases {
		_, err := services.GetProduct(v.ProductID)

		if !IsErrorSame(err, v.Error) {
			t.Errorf(ExpectedGot, fmt.Sprint(v), v.Error, err)
			continue
		}	
	}
}

func TestGetProducts(t *testing.T) {
	err := RestoreDBToTestingState()
	if err != nil {
		t.Errorf(err.Error())
	}

	for _, v := range GetProductsTestCases {
		_, err := services.GetProducts(v.PageNo, v.ItemsPerPage)

		if !IsErrorSame(err, v.Error) {
			t.Errorf(ExpectedGot, fmt.Sprint(v), v.Error, err)
			continue
		}

	}
}

func TestAddProduct(t *testing.T) {
	RestoreDBToTestingState()

	for _, v := range AddProductTestCases {
		response := services.AddProduct(v.Product)

		if response != v.Response {
			t.Errorf(ExpectedGot, fmt.Sprint(v), v.Response, response)
		}

		if response == responses.ProductAddedSuccessfully {
			product, _ := services.GetProduct(v.Product.Product_ID)
			if product == nil {
				t.Errorf(ExpectedGot, fmt.Sprint(v), "", responses.ProductNotAddedToDatabase)
			}
		}
	}

	RestoreDBToTestingState()
}

func TestDeleteProduct(t *testing.T) {
	RestoreDBToTestingState()
	for _, v := range DeleteProductTestCases {
		response := services.DeleteProduct(v.ProductID)

		if response != v.Response {
			t.Errorf(ExpectedGot, fmt.Sprint(v), v.Response, response)
		}

		if response == responses.ProductSuccessfullyDeleted {
			product, _ := services.GetProduct(v.ProductID)
			if product != nil {
				t.Error(responses.ProductNotDeletedFromDatabase)
			}
		}
	}
	RestoreDBToTestingState()
}

func TestUpdateProduct(t *testing.T) {
	RestoreDBToTestingState()

	for _, v := range UpdateProductTestCases {
		response := services.UpdateProduct(v.ProductID, v.ToUpdate)
		if response != v.Response {
			t.Errorf(ExpectedGot, fmt.Sprint(v), v.Response, response)
		}
	}

	RestoreDBToTestingState()
}





func TestGetCategory(t *testing.T) {
	RestoreDBToTestingState()

	for _, v := range GetCategoryTestCases {
		_, err := services.GetCategory(v.CategoryID)
		if !IsErrorSame(err, v.Error) {
			t.Errorf(ExpectedGot, fmt.Sprint(v), v.Error, err)
			continue
		}	
	}
}

func TestGetCategories(t *testing.T) {
	_, err := services.GetCategories()
	if err != nil {
		t.Errorf(ExpectedGot, "", nil, err)
	}
}

func TestAddCategory(t *testing.T) {
	RestoreDBToTestingState()
	for _, v := range AddCategoryTestCases {
		response := services.AddCategory(v.Category)
		if response != v.Response {
			t.Errorf(ExpectedGot, fmt.Sprint(v), v.Response, response)
		}

		if response == responses.CategoryAddedSuccessfully {
			category, _ := services.GetCategory(v.Category.CategoryID)
			if category == nil {
				t.Errorf(ExpectedGot, fmt.Sprint(v), "", responses.CategoryNotAddedToDatabase)
			}
		}
	}
	RestoreDBToTestingState()
}

func TestDeleteCategory(t *testing.T) {
	RestoreDBToTestingState()
	for _, v := range DeleteCategoryTestCases {
		response := services.DeleteCategory(v.CategoryID)
		if response != v.Response {
			t.Errorf(ExpectedGot, fmt.Sprint(v), v.Response, response)
		}

		if response == responses.CategorySuccessfullyDeleted {
			category, _ := services.GetCategory(v.CategoryID)
			if category != nil {
				t.Error(responses.CategoryNotDeletedFromDatabase)
			}
		}
	}
	RestoreDBToTestingState()
}

func TestUpdateCategory(t *testing.T) {
	RestoreDBToTestingState()

	for _, v := range UpdateCategoryTestCases {
		response := services.UpdateCategory(v.CategoryID, v.Name)
		if response != v.Response {
			t.Errorf(ExpectedGot, fmt.Sprint(v), v.Response, response)
		}
	}

	RestoreDBToTestingState()
}




func TestGetInventory(t *testing.T) {
	RestoreDBToTestingState()
	_, err := services.GetInventory()

	if err != nil {
		t.Error(err)
	}
	RestoreDBToTestingState()
}

func TestGetInventoryItem(t *testing.T) {
	RestoreDBToTestingState()

	for _, v := range GetInventoryItemTestCases {
		inventory_item, err := services.GetInventoryItem(v.ProductID)

		if err == nil {
			if inventory_item.Quantity != v.ExpectedQuantity {
				t.Errorf(ExpectedGot, fmt.Sprint(v), v.ExpectedQuantity, inventory_item.Quantity)
			}
		}

		if !IsErrorSame(err, v.Error) {
			t.Errorf(ExpectedGot, fmt.Sprint(v), v.Error, err)
		}
	}

	RestoreDBToTestingState()
}

func TestUpdateInventory(t *testing.T) {
	RestoreDBToTestingState()

	for _, v := range UpdateInventoryTestCases {
		inventory_item_before_update, err := services.GetInventoryItem(v.ProductID)
		if err != nil {
			if !IsErrorSame(err, errors.New(v.Response)) {
				t.Errorf(ExpectedGot, fmt.Sprint(v), v.Response, err.Error())
				continue
			}
		}

		response := services.UpdateInventory(v.ProductID, v.QuantityToUpdate)
		if response != v.Response {
			t.Errorf(ExpectedGot, fmt.Sprint(v), v.Response, response)
			continue
		}

		if response == responses.InventoryUpdatedSuccessfully {
			inventory_item_after_update, err := services.GetInventoryItem(v.ProductID)
			fmt.Println(inventory_item_after_update, err)
			if inventory_item_before_update.Quantity + v.QuantityToUpdate == 0 {
				if !IsErrorSame(err, errors.New(responses.ProductNotInInventory)) {
					t.Errorf(responses.QuantityNotUpdateInInventoryProperly)
					t.Errorf(ExpectedGot, fmt.Sprint(v), v.Response, err)	
				}
				continue
			}

			if inventory_item_after_update.Quantity - inventory_item_before_update.Quantity != v.QuantityToUpdate {
				t.Errorf(responses.QuantityNotUpdateInInventoryProperly)
				t.Errorf(ExpectedGot, fmt.Sprint(v), (inventory_item_before_update.Quantity + v.QuantityToUpdate), inventory_item_after_update.Quantity)
			}
		}

	}

	RestoreDBToTestingState()
}




func TestGetCart(t *testing.T) {
	RestoreDBToTestingState()

	for _, v := range GetCartTestCases {
		_, err := services.GetCart(v.CartReference)
		if !IsErrorSame(err, v.Error) {
			t.Errorf(ExpectedGot, fmt.Sprint(v), v.Error, err)
		}
	}

	RestoreDBToTestingState()
}

func TestAddItemToCart(t *testing.T) {
	RestoreDBToTestingState()

	for _, v := range AddItemToCartTestCases {
		cart_before_adding, _ := services.GetCart(v.CartReference)

		ref, response, isNewCart := services.AddItemToCart(v.CartReference, v.ProductID, v.Quantity)
		v.CartReference = ref

		if response != v.Response {
			t.Errorf(ExpectedGot, fmt.Sprint(v), v.Response, response)
			continue
		}

		if response == responses.InvalidCart {
			continue
		}

		if isNewCart != v.IsNewCart {
			t.Errorf(ExpectedGot, fmt.Sprint(v), v.IsNewCart, isNewCart)
			continue
		}

		cart_after_adding, err := services.GetCart(v.CartReference)
		if err != nil {
			t.Error(v, err)
			continue
		}

		if cart_before_adding == nil {
			for _, vc := range cart_after_adding.Items {
				if vc.ProductID == v.ProductID {
					if vc.Quantity != v.Quantity {
						t.Errorf(responses.CartItemNotAddedSuccessfully)
						t.Errorf(ExpectedGot, fmt.Sprint(v), v.Quantity, vc.Quantity)
					}
				}
			}
		} else {
			current_quantity := 0
			for _, vc := range cart_before_adding.Items {
				if vc.ProductID == v.ProductID {
					current_quantity = vc.Quantity
				}
			}

			for _, vc := range cart_after_adding.Items {
				if vc.ProductID == v.ProductID {
					if response == responses.ItemAddedToCart && vc.Quantity != current_quantity + v.Quantity {
						t.Errorf(responses.CartItemNotAddedSuccessfully)
						t.Errorf(ExpectedGot, fmt.Sprint(v), current_quantity + v.Quantity, vc.Quantity)
					}
				}
			}
		}
	}

	RestoreDBToTestingState()
}

func TestRemoveItemFromCart(t *testing.T) {
	RestoreDBToTestingState()

	for _, v := range RemoveItemFromCartTestCases {
		response := services.RemoveItemFromCart(v.CartReference, v.ProductID, v.Quantity)

		if response != v.Response {
			t.Errorf(ExpectedGot, fmt.Sprint(v), v.Response, response)
		}
	}

	RestoreDBToTestingState()
}