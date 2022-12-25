package typedefs

type Inventory struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type ProductInventory struct {
	Inventory
	ProductName string `json:"product_name"`
}
