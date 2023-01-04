package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/subramanyam-searce/product-catalog-go/constants/queries"
	"github.com/subramanyam-searce/product-catalog-go/constants/responses"
	"github.com/subramanyam-searce/product-catalog-go/helpers"
	"github.com/subramanyam-searce/product-catalog-go/typedefs"
)

func CreateCart() (string, error) {
	ref := uuid.New().String()

	_, err := helpers.RunQuery(queries.InsertCartReference, ref, time.Now())
	helpers.HandleError("runQueryError:", err)
	if err != nil {
		return "", errors.New(responses.ErrorCreatingCart)
	}

	return ref, nil
}

func GetCart(ref string) (*typedefs.Cart, error) {
	cart := typedefs.Cart{}

	rows, err := helpers.RunQuery(queries.GetCartReference, ref)
	helpers.HandleError("runQueryError:", err)

	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, errors.New(responses.InvalidCart)
	}

	err = rows.Scan(&cart.Ref, &cart.CreatedAt)
	helpers.HandleError("rowsScanError", err)

	rows, err = helpers.RunQuery(queries.GetCartItems, ref)
	helpers.HandleError("runQueryError:", err)

	for rows.Next() {
		new_cart_item := typedefs.CartItem{}
		cart_ref_id := ""
		err = rows.Scan(&cart_ref_id, &new_cart_item.ProductID, &new_cart_item.Quantity)
		helpers.HandleError("rowsScanError", err)
		cart.Items = append(cart.Items, new_cart_item)
	}

	cart.EvaluateCartValue()

	return &cart, nil
}

func AddItemToCart(ref string, product_id int, quantity int) (string, string, bool) {
	var err error
	var isNewCart bool

	if quantity <= 0 {
		return ref, responses.QuantityNotPositive, isNewCart
	}
	
	if ref == "" {
		ref, err = CreateCart()
		isNewCart = true
		helpers.HandleError("createCartError", err)
		if err != nil {
			return ref, err.Error(), isNewCart
		}
	}

	_, err = GetCart(ref)
	helpers.HandleError("getCartError", err)

	if err != nil {
		return ref, err.Error(), isNewCart
	}

	response := UpdateInventory(product_id, -quantity)
	if response != responses.InventoryUpdatedSuccessfully {
		return ref, response, isNewCart
	}

	rows, err := helpers.RunQuery(queries.GetCartItemQuantity, ref, product_id)
	helpers.HandleError("runQueryError:", err)

	if err != nil {
		return ref, err.Error(), isNewCart
	}

	if rows.Next() {
		var db_quantity int
		err := rows.Scan(&db_quantity)
		helpers.HandleError("rowsScanError", err)

		_, err = helpers.RunQuery(queries.UpdateCartItemQuantity, db_quantity+quantity, ref, product_id)
		helpers.HandleError("runQueryError:", err)

		if err != nil {
			return ref, err.Error(), isNewCart
		}

	} else {
		_, err = helpers.RunQuery(queries.InsertCartItem, ref, product_id, quantity)
		helpers.HandleError("runQueryError:", err)

		if err != nil {
			return ref, err.Error(), isNewCart
		}
	}

	return ref, responses.ItemAddedToCart, isNewCart
}

func RemoveItemFromCart(ref string, product_id int, quantity int) string {

	_, err := GetCart(ref)
	helpers.HandleError("getCartError", err)
	if err != nil {
		return err.Error()
	}

	rows, err := helpers.RunQuery(queries.GetCartItemQuantity, ref, product_id)
	helpers.HandleError("runQueryError:", err)

	if !rows.Next() {
		return responses.ProductNotInCart
	}

	cart_item := typedefs.CartItem{ProductID: product_id}
	err = rows.Scan(&cart_item.Quantity)
	helpers.HandleError("rowsScanError", err)

	if cart_item.Quantity-quantity < 0 {
		return responses.CartQuantityLessThanRequiredRemoval
	}

	if cart_item.Quantity-quantity == 0 {
		_, err = helpers.RunQuery(queries.DeleteCartItem, ref, product_id)
		helpers.HandleError("runQueryError", err)

		_, err = helpers.RunQuery(queries.UpsertInventoryForCartItemRemoval, product_id, quantity)
		helpers.HandleError("runQueryError", err)

		return responses.CartItemsRemovedSuccessfully
	}

	if cart_item.Quantity-quantity > 0 {
		_, err = helpers.RunQuery(queries.UpdateCartItemQuantity, cart_item.Quantity-quantity, ref, product_id)
		helpers.HandleError("runQueryError", err)

		_, err = helpers.RunQuery(queries.UpsertInventoryForCartItemRemoval, product_id, quantity)
		helpers.HandleError("runQueryError", err)
	}

	return responses.CartItemsRemovedSuccessfully
}
