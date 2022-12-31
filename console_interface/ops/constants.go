package ops

import (
	console_typedefs "github.com/subramanyam-searce/product-catalog-go/console_interface/typedefs"
)

var InvalidChoice string = "Invalid Choice: Please enter a valid choice"
var ChoiceInput string = "Enter your choice: "
var EnterPositiveValue string = "Please Enter a positive Value"
var GetProductID string = "Product ID: "
var GetItemsPerPage string = "Items Per Page: "
var GetPageNo string = "Page No: "
var SizeOfOutputStructFields int = 15
var Divider = "-----------------------"
var Colon string = ": "
var EnterValidJson string = "Invalid JSON: Please enter a valid JSON string"
var EnterFieldNameToUpdate string = "Field"
var EnterTheValueToUpdate string = "New Value"
var FieldNotFoundForUpdation string = "Field not Found"

var HomePageOptions []console_typedefs.HomePageOption = []console_typedefs.HomePageOption{
	{DisplayName: "Quit", Handler: func() {}},
	{DisplayName: "Get Product", Handler: GetProduct},
	{DisplayName: "Get Products", Handler: GetProducts},
	{DisplayName: "Add Product", Handler: AddProduct},
	{DisplayName: "Delete Product", Handler: DeleteProduct},
	{DisplayName: "Update Product", Handler: UpdateProduct},
}
