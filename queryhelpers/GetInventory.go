package queryhelpers

import (
	"github.com/subramanyam-searce/product-catalog-go/helpers"
	"github.com/subramanyam-searce/product-catalog-go/typedefs"
)

func GetInventory(id int) (*typedefs.Inventory, error) {
	query := "SELECT * FROM inventory WHERE product_id=$1;"
	var inventoryItem *typedefs.Inventory = nil

	rows, err := helpers.RunQuery(query, id)

	if rows.Next() {
		inventoryItem = &typedefs.Inventory{}
		rows.Scan(&inventoryItem.ProductID, &inventoryItem.Quantity)
	}

	return inventoryItem, err
}
