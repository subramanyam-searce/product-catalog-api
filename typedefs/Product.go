package typedefs

type Product struct {
	Product_ID    int               `json:"product_id"`
	Name          string            `json:"name"`
	Specification map[string]string `json:"specification"`
	SKU           string            `json:"sku"`
	CategoryID    int               `json:"category_id"`
	Price         float32           `json:"price"`
}

type CategoryProduct struct {
	Product
	CategoryName string `json:"category_name"`
}
