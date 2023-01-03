package typedefs

import (
	"net/http"

	"github.com/subramanyam-searce/product-catalog-go/helpers"
)

type Product struct {
	Product_ID    int            `json:"product_id"`
	Name          string         `json:"name"`
	Specification map[string]any `json:"specification"`
	SKU           string         `json:"sku"`
	CategoryID    int            `json:"category_id"`
	Price         float64        `json:"price"`
}

type ShortProduct struct {
	Product_ID int     `json:"product_id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
}

type Category struct {
	CategoryID int    `json:"category_id"`
	Name       string `json:"name"`
}

type InventoryItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type ProductInventory struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
	ProductName string `json:"product_name"`
}

type CartItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type Cart struct {
	Ref       string     `json:"ref"`
	CreatedAt string     `json:"created_at"`
	Items     []CartItem `json:"items"`
	CartValue float64    `json:"total_cart_value"`
}

func (c *Cart) EvaluateCartValue() {
	for _, v := range c.Items {
		rows, err := helpers.RunQuery("SELECT price FROM product WHERE product_id=$1", v.ProductID)
		helpers.HandleError("runQueryError", err)

		var price float64
		rows.Next()
		err = rows.Scan(&price)
		helpers.HandleError("rowsScanError", err)

		c.CartValue += float64(price) * float64(v.Quantity)
	}
}

type Route struct {
	Path    string
	Handler func(http.ResponseWriter, *http.Request)
	Method  string
}
