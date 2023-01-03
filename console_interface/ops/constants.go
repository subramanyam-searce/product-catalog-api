package ops

import (
	console_typedefs "github.com/subramanyam-searce/product-catalog-go/console_interface/typedefs"
)

var InvalidChoice string = "Invalid Choice: Please enter a valid choice"
var ChoiceInput string = "Enter your choice: "
var EnterPositiveValue string = "Please Enter a positive Value"
var GetProductID string = "Product ID: "
var GetCategoryID string = "Category ID: "
var GetCategoryName string = "Category Name: "
var GetItemsPerPage string = "Items Per Page: "
var GetQuantity string = "Quantity: "
var GetCartReference string = "Cart Reference"
var GetPageNo string = "Page No: "
var SizeOfOutputStructFields int = 15
var Divider = "-----------------------"
var Colon string = ": "
var EnterValidJson string = "Invalid JSON: Please enter a valid JSON string"
var EnterFieldNameToUpdate string = "Field"
var EnterTheValueToUpdate string = "New Value"
var FieldNotFoundForUpdation string = "Field not Found"
var EnterValidFloat string = "Please enter a valid float/number"
var CartItemsDisplay string = "Cart Items: "
var NewCartCreatedDisplay string = "New cart was created. Cart Reference: %v\n"
var GetCartItemNumberDisplay string = "How many products do you want to add?"

var HomePageOptions []console_typedefs.HomePageOption = []console_typedefs.HomePageOption{
	{DisplayName: "Quit", Handler: func() {}},
	{DisplayName: "Get Product", Handler: GetProduct},
	{DisplayName: "Get Products", Handler: GetProducts},
	{DisplayName: "Add Product", Handler: AddProduct},
	{DisplayName: "Delete Product", Handler: DeleteProduct},
	{DisplayName: "Update Product", Handler: UpdateProduct},

	{DisplayName: "Get Category", Handler: GetCategory},
	{DisplayName: "Get Categories", Handler: GetCategories},
	{DisplayName: "Add Category", Handler: AddCategory},
	{DisplayName: "Delete Category", Handler: DeleteCategory},
	{DisplayName: "Update Category", Handler: UpdateCategory},

	{DisplayName: "Get Inventory", Handler: GetInventory},
	{DisplayName: "Update Inventory", Handler: UpdateInventory},

	{DisplayName: "Get Cart", Handler: GetCart},
	{DisplayName: "Add Item to Cart", Handler: AddItemToCart},
	{DisplayName: "Add Items to Cart", Handler: AddItemsToCart},
	{DisplayName: "Remove Item to Cart", Handler: RemoveItemFromCart},
}
