package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/subramanyam-searce/product-catalog-go/constants/responses"
	"github.com/subramanyam-searce/product-catalog-go/db/services"
	"github.com/subramanyam-searce/product-catalog-go/helpers"
)

var ExpectedGot string = "%v: Expected: %v, Got: %v"

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(responses.ErrorLoadingEnvFile + err.Error())
	}
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