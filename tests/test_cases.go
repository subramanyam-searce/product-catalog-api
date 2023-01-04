package tests

import (
	"errors"

	"github.com/subramanyam-searce/product-catalog-go/constants/responses"
	"github.com/subramanyam-searce/product-catalog-go/typedefs"
)

type GetProductTestCase struct {
	ProductID int
	Error error
}

type GetProductsTestCase struct {
	PageNo int
	ItemsPerPage int
	Error error // empty string if no error occurs and products array gets returned
}

type AddProductTestCase struct {
	Product typedefs.Product
	Response string //response returned by the function
}

type DeleteProductTestCase struct {
	ProductID int
	Response string
}

type UpdateProductTestCase struct {
	ProductID int
	ToUpdate map[string]any
	Response string
}





type GetCategoryTestCase struct {
	CategoryID int
	Error error
}

type AddCategoryTestCase struct {
	Category typedefs.Category
	Response string
}

type DeleteCategoryTestCase struct {
	CategoryID int
	Response string
}

type UpdateCategoryTestCase struct {
	CategoryID int
	Name string
	Response string
}





type GetInventoryItemTestCase struct {
	ProductID int
	ExpectedQuantity int
	Error error
}

type UpdateInventoryTestCase struct {
	ProductID int
	QuantityToUpdate int
	Response string
}




type GetCartTestCase struct {
	CartReference string
	Error error
}

type AddItemToCartTestCase struct {
	CartReference string
	ProductID int
	Quantity int
	Response string
	IsNewCart bool
}

type RemoveItemFromCartTestCase struct {
	CartReference string
	ProductID int
	Quantity int
	Response string
}





var GetProductTestCases []GetProductTestCase = []GetProductTestCase{
	{ProductID: 1},
	{ProductID: 2},
	{ProductID: 3},
	{ProductID: 4},
	{ProductID: 5},
	{ProductID: 6},
	{ProductID: 7},

	{ProductID: 56},
	{ProductID: 100},
	{ProductID: 10},
	{ProductID: 1000000},

	{ProductID: 0, Error: errors.New(responses.ProductIDNotPositive)},
	{ProductID: -1, Error: errors.New(responses.ProductIDNotPositive)},
	{ProductID: -2, Error: errors.New(responses.ProductIDNotPositive)},
	{ProductID: -3, Error: errors.New(responses.ProductIDNotPositive)},
	{ProductID: -101, Error: errors.New(responses.ProductIDNotPositive)},
	{ProductID: -1060, Error: errors.New(responses.ProductIDNotPositive)},
}

var GetProductsTestCases []GetProductsTestCase = []GetProductsTestCase {
	{PageNo: -10, ItemsPerPage: 0, Error: errors.New(responses.InvalidPageNo)},
	{PageNo: 0, ItemsPerPage: 1, Error: errors.New(responses.InvalidPageNo)},
	{PageNo: 1, ItemsPerPage: 0, Error: errors.New(responses.InvalidItemsPerPage)},
	{PageNo: -3, ItemsPerPage: 1, Error: errors.New(responses.InvalidPageNo)},
	{PageNo: 1, ItemsPerPage: -5, Error: errors.New(responses.InvalidItemsPerPage)},

	{PageNo: 1, ItemsPerPage: 1},
	{PageNo: 2, ItemsPerPage: 1},
	{PageNo: 3, ItemsPerPage: 1},
	{PageNo: 4, ItemsPerPage: 1},
	{PageNo: 5, ItemsPerPage: 1},
	{PageNo: 6, ItemsPerPage: 1},
	{PageNo: 7, ItemsPerPage: 1},
	{PageNo: 8, ItemsPerPage: 1, Error: errors.New(responses.ProductsOutOfRange)},
	{PageNo: 9, ItemsPerPage: 1, Error: errors.New(responses.ProductsOutOfRange)},

	{PageNo: 1, ItemsPerPage: 1000},

	{PageNo: 5, ItemsPerPage: 10, Error: errors.New(responses.ProductsOutOfRange)},
	{PageNo: 5, ItemsPerPage: 10, Error: errors.New(responses.ProductsOutOfRange)},
}

var AddProductTestCases []AddProductTestCase = []AddProductTestCase {
	{Product: typedefs.Product{
		Product_ID: 101,
		Name: "Money Purse",
		Specification: map[string]any{"color": "black"},
		SKU: "665589",
		CategoryID: 5,
		Price: 10.99,
	}, Response: responses.ProductAddedSuccessfully,},
	{Product: typedefs.Product{
		Product_ID: -101,
		Name: "Money Purse",
		Specification: map[string]any{"color": "black"},
		SKU: "665589",
		CategoryID: 5,
		Price: 10.99,
	}, Response: responses.ProductIDNotPositive,},
	{Product: typedefs.Product{
		Product_ID: 101,
		Name: "Money Purse",
		Specification: map[string]any{"color": "black"},
		SKU: "665589",
		CategoryID: -5,
		Price: 10.99,
	}, Response: responses.CategoryIDNotPositive,},
	{Product: typedefs.Product{
		Product_ID: 101,
		Name: "Money Purse",
		Specification: map[string]any{"color": "black"},
		SKU: "665589",
		CategoryID: 5,
		Price: -10.99,
	}, Response: responses.PriceNotPositive,},

	{Product: typedefs.Product{
		Product_ID: 101,
		Name: "Money Purse",
		Specification: map[string]any{"color": "black"},
		SKU: "665589",
		CategoryID: 5,
		Price: 10.99,
	}, Response: responses.ProductAlreadyExists,},
	{Product: typedefs.Product{
		Product_ID: 102,
		Name: "Money Purse",
		Specification: map[string]any{"color": "black"},
		SKU: "665589",
		CategoryID: 100,
		Price: 10.99,
	}, Response: responses.InvalidCategoryIDForProductFKConstraint,},
	{Product: typedefs.Product{
		Product_ID: 1,
		Name: "Money Purse",
		Specification: map[string]any{"color": "black"},
		SKU: "665589",
		CategoryID: 5,
		Price: 10.99,
	}, Response: responses.ProductAlreadyExists,},
	{Product: typedefs.Product{
		Product_ID: 2,
		Name: "Money Purse",
		Specification: map[string]any{"color": "black"},
		SKU: "665589",
		CategoryID: -1,
		Price: 10.99,
	}, Response: responses.CategoryIDNotPositive,},

	{Product: typedefs.Product{
		Product_ID: 5,
		Name: "Money Purse",
		Specification: map[string]any{"color": "black"},
		SKU: "665589",
		CategoryID: 100,
		Price: 10.99,
	}, Response: responses.ProductAlreadyExists,},
}

var DeleteProductTestCases []DeleteProductTestCase = []DeleteProductTestCase {
	{ProductID: 1, Response: responses.ProductSuccessfullyDeleted},
	{ProductID: 3, Response: responses.ProductSuccessfullyDeleted},
	{ProductID: 5, Response: responses.ProductSuccessfullyDeleted},
	{ProductID: 4, Response: responses.ProductSuccessfullyDeleted},

	{ProductID: -1, Response: responses.ProductIDNotPositive},
	{ProductID: 0, Response: responses.ProductIDNotPositive},
	{ProductID: -100, Response: responses.ProductIDNotPositive},

	{ProductID: 1, Response: responses.ProductNotFound},
	{ProductID: 3, Response: responses.ProductNotFound},
	{ProductID: 100, Response: responses.ProductNotFound},
}

var UpdateProductTestCases []UpdateProductTestCase = []UpdateProductTestCase {
	{ProductID: 1, ToUpdate: map[string]any{
		"name": "TWS",
	}, Response: responses.ProductUpdatedSuccessfully,},
	{ProductID: 1034, ToUpdate: map[string]any{
		"name": "TWS",
	}, Response: responses.ProductNotFound,},
	{ProductID: -1, ToUpdate: map[string]any{
		"name": "TWS",
	}, Response: responses.ProductIDNotPositive,},
	{ProductID: 3, ToUpdate: map[string]any{
		"name": "TWS",
		"price": float64(-1),
	}, Response: responses.PriceNotPositive},
	{ProductID: 3, ToUpdate: map[string]any{
		"name": "TWS",
		"price": float64(0),
		"category_id": 23,
	}, Response: responses.InvalidCategoryIDForProductFKConstraint,},
	{ProductID: 3, ToUpdate: map[string]any{
		"name": "TWS",
		"price": float64(-1),
		"category_id": -3,
	}, Response: responses.PriceNotPositive,},
	{ProductID: 3, ToUpdate: map[string]any{
		"name": "TWS",
		"price": float64(-1),
		"specfici": "",
		"category_id": -3,
	}, Response: responses.PriceNotPositive,},
	{ProductID: 3, ToUpdate: map[string]any{
		"name": "TWS",
		"price": float64(10.99),
		"category_id": 1,
	}, Response: responses.ProductUpdatedSuccessfully,},
}





var GetCategoryTestCases []GetCategoryTestCase = []GetCategoryTestCase {
	{CategoryID: 1},
	{CategoryID: 2},
	{CategoryID: 3},
	{CategoryID: 4},
	{CategoryID: 5},

	{CategoryID: 100, Error: errors.New(responses.InvalidCategoryID)},
	{CategoryID: 101, Error: errors.New(responses.InvalidCategoryID)},
	{CategoryID: 1010, Error: errors.New(responses.InvalidCategoryID)},

	{CategoryID: 0, Error: errors.New(responses.CategoryIDNegative)},
	{CategoryID: -1, Error: errors.New(responses.CategoryIDNegative)},
	{CategoryID: -2, Error: errors.New(responses.CategoryIDNegative)},
	{CategoryID: -3, Error: errors.New(responses.CategoryIDNegative)},
	{CategoryID: -100, Error: errors.New(responses.CategoryIDNegative)},
}

var AddCategoryTestCases []AddCategoryTestCase = []AddCategoryTestCase {
	{Category: typedefs.Category{
		CategoryID: 6,
		Name: "Test Category",
	}, Response: responses.CategoryAddedSuccessfully,},
	{Category: typedefs.Category{
		CategoryID: 100,
		Name: "Test Category",
	}, Response: responses.CategoryAddedSuccessfully,},

	{Category: typedefs.Category{
		CategoryID: 123,
		Name: "",
	}, Response: responses.CategoryNameLengthError,},
	{Category: typedefs.Category{
		CategoryID: 123,
		Name: "M",
	}, Response: responses.CategoryNameLengthError,},
	{Category: typedefs.Category{
		CategoryID: 123,
		Name: "Ma",
	}, Response: responses.CategoryNameLengthError,},
	{Category: typedefs.Category{
		CategoryID: 123,
		Name: "Max",
	}, Response: responses.CategoryAddedSuccessfully,},

	{Category: typedefs.Category{
		CategoryID: 1,
		Name: "Test Category",
	}, Response: responses.CategoryAlreadyExist,},
	{Category: typedefs.Category{
		CategoryID: 2,
		Name: "Test Category",
	}, Response: responses.CategoryAlreadyExist,},
	{Category: typedefs.Category{
		CategoryID: 4,
		Name: "Test Category",
	}, Response: responses.CategoryAlreadyExist,},

	{Category: typedefs.Category{
		CategoryID: 0,
		Name: "Test Category",
	}, Response: responses.CategoryIDNegative,},
	{Category: typedefs.Category{
		CategoryID: -1,
		Name: "Test Category",
	}, Response: responses.CategoryIDNegative,},
	{Category: typedefs.Category{
		CategoryID: -2,
		Name: "Test Category",
	}, Response: responses.CategoryIDNegative,},
	{Category: typedefs.Category{
		CategoryID: -100,
		Name: "Test Category",
	}, Response: responses.CategoryIDNegative,},
}

var DeleteCategoryTestCases []DeleteCategoryTestCase = []DeleteCategoryTestCase {
	{CategoryID: 1, Response: responses.CategoryIDUsedByProduct},
	{CategoryID: 2, Response: responses.CategoryIDUsedByProduct},
	{CategoryID: 4, Response: responses.CategoryIDUsedByProduct},

	{CategoryID: 3, Response: responses.CategorySuccessfullyDeleted},
	{CategoryID: 5, Response: responses.CategorySuccessfullyDeleted},

	{CategoryID: 0, Response: responses.CategoryIDNegative},
	{CategoryID: -1, Response: responses.CategoryIDNegative},
	{CategoryID: -100, Response: responses.CategoryIDNegative},
	{CategoryID: -45, Response: responses.CategoryIDNegative},

	{CategoryID: 3, Response: responses.InvalidCategoryID},
	{CategoryID: 5, Response: responses.InvalidCategoryID},
	{CategoryID: 100, Response: responses.InvalidCategoryID},
	{CategoryID: 64, Response: responses.InvalidCategoryID},
}

var UpdateCategoryTestCases []UpdateCategoryTestCase = []UpdateCategoryTestCase {
	{CategoryID: 1, Name: "Test", Response: responses.CategoryUpdatedSuccessfully},
	{CategoryID: 2, Name: "Test", Response: responses.CategoryUpdatedSuccessfully},
	{CategoryID: 4, Name: "", Response: responses.CategoryNameLengthError},
	{CategoryID: 4, Name: "a", Response: responses.CategoryNameLengthError},
	{CategoryID: 4, Name: "as", Response: responses.CategoryNameLengthError},
	{CategoryID: 4, Name: "Max", Response: responses.CategoryUpdatedSuccessfully},

	{CategoryID: 11, Name: "Test", Response: responses.InvalidCategoryID},
	{CategoryID: 1001, Name: "Test", Response: responses.InvalidCategoryID},
	{CategoryID: 253, Name: "Test", Response: responses.InvalidCategoryID},
	{CategoryID: 9865, Name: "Test", Response: responses.InvalidCategoryID},

	{CategoryID: 0, Name: "Test", Response: responses.CategoryIDNegative},
	{CategoryID: -1, Name: "Test", Response: responses.CategoryIDNegative},
	{CategoryID: -2, Name: "Test", Response: responses.CategoryIDNegative},
	{CategoryID: -5, Name: "Test", Response: responses.CategoryIDNegative},
	{CategoryID: -99, Name: "Test", Response: responses.CategoryIDNegative},
	{CategoryID: -1000, Name: "Test", Response: responses.CategoryIDNegative},
}





var GetInventoryItemTestCases []GetInventoryItemTestCase = []GetInventoryItemTestCase {
	{ProductID: 1, ExpectedQuantity: 100},
	{ProductID: 2, ExpectedQuantity: 50},
	{ProductID: 3, ExpectedQuantity: 20},
	{ProductID: 4, ExpectedQuantity: 40},
	{ProductID: 6, ExpectedQuantity: 5},

	{ProductID: 5, ExpectedQuantity: 0, Error: errors.New(responses.ProductNotInInventory)},
	{ProductID: 1001, ExpectedQuantity: 5, Error: errors.New(responses.ProductNotInInventory)},
	{ProductID: 103, ExpectedQuantity: 5, Error: errors.New(responses.ProductNotInInventory)},
}

var UpdateInventoryTestCases []UpdateInventoryTestCase = []UpdateInventoryTestCase {
	{ProductID: 1, QuantityToUpdate: 10, Response: responses.InventoryUpdatedSuccessfully},
	{ProductID: 1, QuantityToUpdate: 200, Response: responses.InventoryUpdatedSuccessfully},
	{ProductID: 2, QuantityToUpdate: 2, Response: responses.InventoryUpdatedSuccessfully},
	{ProductID: 3, QuantityToUpdate: -10, Response: responses.InventoryUpdatedSuccessfully},
	{ProductID: 3, QuantityToUpdate: -10, Response: responses.InventoryUpdatedSuccessfully},

	{ProductID: 3, QuantityToUpdate: -10, Response: responses.ProductNotInInventory},
	{ProductID: 100, QuantityToUpdate: -10, Response: responses.ProductNotInInventory},
	{ProductID: -2, QuantityToUpdate: -10, Response: responses.ProductNotInInventory},
	{ProductID: -102, QuantityToUpdate: -10, Response: responses.ProductNotInInventory},

	{ProductID: 1, QuantityToUpdate: -10000000, Response: responses.InventoryQuantityLessThanRequiredRemoval},
	{ProductID: 2, QuantityToUpdate: -100000002, Response: responses.InventoryQuantityLessThanRequiredRemoval},
}



var GetCartTestCases []GetCartTestCase = []GetCartTestCase {
	{CartReference: "4d4d8297-7663-451d-b79e-49a545728552"},

	{CartReference: "asdasdasdasdasd", Error: errors.New(responses.InvalidCart)},
	{CartReference: "123123", Error: errors.New(responses.InvalidCart)},
	{CartReference: "4d4d8297-7663-451d-b79e-1111111111111", Error: errors.New(responses.InvalidCart)},
}

var AddItemToCartTestCases []AddItemToCartTestCase = []AddItemToCartTestCase {
	{CartReference: "4d4d8297-7663-451d-b79e-49a545728552", ProductID: 1, Quantity: 10, Response: responses.ItemAddedToCart},
	{CartReference: "4d4d8297-7663-451d-b79e-49a545728552", ProductID: 1, Quantity: 10, Response: responses.ItemAddedToCart},
	{CartReference: "asdasd", ProductID: 1, Quantity: 10, Response: responses.InvalidCart},
	{CartReference: "", ProductID: 1, Quantity: 10, Response: responses.ItemAddedToCart, IsNewCart: true},

	{CartReference: "4d4d8297-7663-451d-b79e-49a545728552", ProductID: 1, Quantity: 1000, Response: responses.InventoryQuantityLessThanRequiredRemoval},
	{CartReference: "4d4d8297-7663-451d-b79e-49a545728552", ProductID: -1, Quantity: 1000, Response: responses.ProductNotInInventory},
	{CartReference: "4d4d8297-7663-451d-b79e-49a545728552", ProductID: 1, Quantity: -1, Response: responses.QuantityNotPositive},

	{CartReference: "4d4d8297-7663-451d-b79e-111111111111", ProductID: 1, Quantity: 1, Response: responses.InvalidCart},
	{CartReference: "4d4d8297-7663-451d-b79e-49a545728552", ProductID: 1, Quantity: 0, Response: responses.QuantityNotPositive},
	{CartReference: "4d4d8297-7663-451d-b79e-49a545728552", ProductID: 10, Quantity: 0, Response: responses.QuantityNotPositive},
	{CartReference: "4d4d8297-7663-451d-b79e-49a545728552", ProductID: 10, Quantity: 1, Response: responses.ProductNotInInventory},
}

var RemoveItemFromCartTestCases []RemoveItemFromCartTestCase = []RemoveItemFromCartTestCase {
	{CartReference: "4d4d8297-7663-451d-b79e-49a545728552", ProductID: 3, Quantity: 5, Response: responses.CartItemsRemovedSuccessfully},
	{CartReference: "4d4d8297-7663-451d-b79e-49a545728552", ProductID: 3, Quantity: 6, Response: responses.CartQuantityLessThanRequiredRemoval},
	{CartReference: "4d4d8297-7663-451d-b79e-49a545728552", ProductID: 3, Quantity: 5, Response: responses.CartItemsRemovedSuccessfully},
	{CartReference: "4d4d8297-7663-451d-b79e-111111111111", ProductID: 3, Quantity: 5, Response: responses.InvalidCart},
	{CartReference: "", ProductID: 3, Quantity: 5, Response: responses.InvalidCart},
	{CartReference: "4d4d8297-7663-451d-b79e-49a545728552", ProductID: 2, Quantity: 5, Response: responses.ProductNotInCart},
	{CartReference: "4d4d8297-7663-451d-b79e-49a545728552", ProductID: -2, Quantity: 5, Response: responses.ProductNotInCart},
}