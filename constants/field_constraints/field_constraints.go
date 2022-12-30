package field_constraints

import "github.com/subramanyam-searce/product-catalog-go/handlers/validators"

var AddProduct []validators.FieldConstraint = []validators.FieldConstraint{
	{FieldName: "product_id", DataType: "float64", IsRequired: true},
	{FieldName: "name", DataType: "string", IsRequired: true},
	{FieldName: "category_id", DataType: "float64", IsRequired: true},
	{FieldName: "specification", DataType: "map[string]interface {}", IsRequired: true},
	{FieldName: "sku", DataType: "string", IsRequired: true},
	{FieldName: "price", DataType: "float64", IsRequired: true},
}

var UpdateProduct []validators.FieldConstraint = []validators.FieldConstraint{
	{FieldName: "product_id", DataType: "float64"},
	{FieldName: "name", DataType: "string"},
	{FieldName: "category_id", DataType: "float64"},
	{FieldName: "specification", DataType: "map[string]interface {}"},
	{FieldName: "sku", DataType: "string"},
	{FieldName: "price", DataType: "float64"},
}

var AddCategory []validators.FieldConstraint = []validators.FieldConstraint{
	{FieldName: "category_id", DataType: "float64", IsRequired: true},
	{FieldName: "name", DataType: "string", IsRequired: true},
}

var UpdateCategory []validators.FieldConstraint = []validators.FieldConstraint{
	{FieldName: "category_id", DataType: "float64"},
	{FieldName: "name", DataType: "string"},
}

var UpdateInventory []validators.FieldConstraint = []validators.FieldConstraint{
	{FieldName: "product_id", DataType: "float64", IsRequired: true},
	{FieldName: "quantity", DataType: "float64", IsRequired: true, IsNegativeAllowed: true},
}

var AddItemToCart []validators.FieldConstraint = []validators.FieldConstraint{
	{FieldName: "product_id", DataType: "float64", IsRequired: true},
	{FieldName: "quantity", DataType: "float64", IsRequired: true},
}

var RemoveItemFromCart []validators.FieldConstraint = []validators.FieldConstraint{
	{FieldName: "product_id", DataType: "float64", IsRequired: true},
	{FieldName: "quantity", DataType: "float64"},
}
