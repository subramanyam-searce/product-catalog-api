package responses

var ProjectRoot string = "/root/learning/Golang/rest-api-projects/product-catalogue/"

var ErrorLoadingEnvFile string = "An Error occurred while loading .env file: "
var FailedToRestoreDB string = "Failed to restore DB"

var ProductNotFound string = "Product not found"
var ProductSuccessfullyDeleted string = "Product was successfully deleted from the database"
var ProductsOutOfRange string = "Page number is out of range"
var InvalidPageNo string = "Invalid Page Number. Please enter a Valid integer"
var InvalidItemsPerPage string = "Invalid Items per Page. Please enter a Valid integer"
var ProductAlreadyExists string = "Product with this ID already exists"
var BadRequestBody string = "Bad Request Body"
var ProductAddedSuccessfully = "Product was added Successfully"
var MissingFields string = "Missing Fields: "
var InvalidDataTypeForFieldInJSONBody string = "Invalid Datatype for one/more fields in the Request body: "
var FollowingFieldsCannotBeNegative string = "The Following Fields cannot be Negative: "
var InvalidFieldsInRequestBody string = "Invalid Field detected in Request Body: "
var InvalidCategoryIDForProductFKConstraint string = "Invalid CategoryID"
var ProductIDCannotBeUpdated string = "ProductID cannot be Updated"
var ProductUpdatedSuccessfully string = "Product was updated Successfully"
var EmptyInputJson string = "The Input JSON is empty in the request body"
var InvalidCategoryID string = "Invalid Category ID"
var CategoryIDNegative string = "Category ID cannot be Negative"
var CategorySuccessfullyDeleted string = "Category was Successfully deleted from the database"
var CategoryIDUsedByProduct string = "This Category is being used in the product table and cannot be deleted"
var CategoryAlreadyExist string = "Category with this ID already exists"
var CategoryAddedSuccessfully string = "Category was added Successfully"
var CategoryIDCannotBeUpdated string = "CategoryID cannot be Updated"
var CategoryUpdatedSuccessfully string = "Category was Updated Successfully"
var ProductNotInInventory string = "Product is not in the Inventory"
var InventoryQuantityLessThanRequiredRemoval string = "Inventory quantity is less than the amount you are trying to remove"
var InventoryUpdatedSuccessfully = "Inventory items was updated Successfully"
var ErrorCreatingCart string = "An Error occurred while creating your cart. Please try again"
var InvalidCart string = "Cart Reference is Invalid"
var ReqQuantityMoreThanStock string = "The Required Quantity is more than the Available Inventory Quantity: "
var ItemAddedToCart string = "Item was added to the cart Sucessfully"
var ProductNotInCart string = "Product is not present in your Cart"
var CartQuantityLessThanRequiredRemoval string = "Cart quantity is less than the amount you are trying to remove"
var CartItemsRemovedSuccessfully string = "Cart Item(s) were removed Successfully"

var ProductIDNotPositive string = "Product ID cannot be Zero or Negative"
var PriceNotPositive string = "Price cannot be Zero or Negative"
var CategoryIDNotPositive string = "Category ID cannot be Zero or Negative"
var QuantityNotPositive string = "Quantity cannot be Zero or Negative"
var ProductNotAddedToDatabase string = "Product was not Added to the database but got a Success response"
var CategoryNotAddedToDatabase string = "Category was not Added to the database but got a Success response"
var ProductNotDeletedFromDatabase string = "Product was not Deleted from the database but got a Success response"
var CategoryNotDeletedFromDatabase string = "Category was not Deleted from the database but got a Success response"
var CategoryNameLengthError string = "Category Name's length needs to be atleast 3 Characters"
var QuantityNotUpdateInInventoryProperly string = "Quantity was not updated in Inventory Properly."
var CartItemNotAddedSuccessfully string = "Cart Item was not added successfully in database"

var ServerStarted string = "Server running on PORT 8080"
