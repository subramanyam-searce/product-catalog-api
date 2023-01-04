package services

import (
	"errors"

	"github.com/subramanyam-searce/product-catalog-go/constants/queries"
	"github.com/subramanyam-searce/product-catalog-go/constants/responses"
	"github.com/subramanyam-searce/product-catalog-go/helpers"
	"github.com/subramanyam-searce/product-catalog-go/typedefs"
)

func GetInventory() (*[]typedefs.ProductInventory, error) {
	inventory := []typedefs.ProductInventory{}

	rows, err := helpers.RunQuery(queries.GetInventory)
	helpers.HandleError("runQueryError", err)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		newInventoryItem := typedefs.ProductInventory{}
		err := rows.Scan(&newInventoryItem.ProductID, &newInventoryItem.ProductName, &newInventoryItem.Quantity)
		helpers.HandleError("rowsScanError", err)
		inventory = append(inventory, newInventoryItem)
	}

	return &inventory, err
}

func GetInventoryItem(product_id int) (*typedefs.InventoryItem, error) {
	inventory_item := typedefs.InventoryItem{}

	rows, err := helpers.RunQuery(queries.GetInventoryItem, product_id)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		err = rows.Scan(&inventory_item.ProductID, &inventory_item.Quantity)
		helpers.HandleError("rowsScanError", err)
		return &inventory_item, nil
	} else {
		return nil, errors.New(responses.ProductNotInInventory)
	}
}

func UpdateInventory(product_id int, quantity int) string {

	inventory_item, err := GetInventoryItem(product_id)
	helpers.HandleError("getInventoryItemError", err)
	if err != nil {
		return err.Error()
	}

	if inventory_item != nil {
		existing_quantity := inventory_item.Quantity

		_, err = helpers.RunQuery(queries.UpdateInventoryItem, existing_quantity+quantity, product_id)
		helpers.HandleError("runQueryError", err)

		if err != nil {
			return responses.InventoryQuantityLessThanRequiredRemoval
		}

		if existing_quantity+quantity == 0 {
			_, err = helpers.RunQuery(queries.DeleteInventoryItem, product_id)
			helpers.HandleError("runQueryError", err)

			if err != nil {
				return err.Error()
			}
		}
	} else {
		_, err = helpers.RunQuery(queries.InsertInventoryItem, product_id, quantity)
		helpers.HandleError("runQueryError", err)
		if err != nil {
			return responses.InventoryQuantityLessThanRequiredRemoval
		}
	}

	return responses.InventoryUpdatedSuccessfully
}
