package typedefs

type CartItem struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type Cart struct {
	Ref       string     `json:"ref"`
	CreatedAt string     `json:"created_at"`
	Items     []CartItem `json:"items"`
}
