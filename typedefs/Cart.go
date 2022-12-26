package typedefs

import "github.com/subramanyam-searce/product-catalog-go/helpers"

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
