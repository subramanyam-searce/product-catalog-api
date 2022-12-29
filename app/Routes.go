package app

import (
	handlers_cart "github.com/subramanyam-searce/product-catalog-go/handlers/cart"
	handlers_category "github.com/subramanyam-searce/product-catalog-go/handlers/category"
	handlers_inventory "github.com/subramanyam-searce/product-catalog-go/handlers/inventory"
	handlers_product "github.com/subramanyam-searce/product-catalog-go/handlers/product"
	"github.com/subramanyam-searce/product-catalog-go/typedefs"
)

var Routes []typedefs.Route = []typedefs.Route{
	{Path: "/product/add", Handler: handlers_product.AddProduct, Method: "POST"},
	{Path: "/products/{page_no:[0-9]+}", Handler: handlers_product.GetProducts, Method: "GET"},
	{Path: "/product/{id:[0-9]+}", Handler: handlers_product.GetProduct, Method: "GET"},
	{Path: "/product/delete/{id:[0-9]+}", Handler: handlers_product.DeleteProduct, Method: "DELETE"},
	{Path: "/product/update/{id:[0-9]+}", Handler: handlers_product.UpdateProduct, Method: "PUT"},

	{Path: "/category/add", Handler: handlers_category.AddCategory, Method: "POST"},
	{Path: "/categories", Handler: handlers_category.GetCategories, Method: "GET"},
	{Path: "/category/delete/{id:[0-9]+}", Handler: handlers_category.DeleteCategory, Method: "DELETE"},
	{Path: "/category/update/{id:[0-9]+}", Handler: handlers_category.UpdateCategory, Method: "PUT"},

	{Path: "/inventory/update", Handler: handlers_inventory.UpdateInventory, Method: "POST"},
	{Path: "/inventory", Handler: handlers_inventory.GetInventory, Method: "GET"},

	{Path: "/cart/create", Handler: handlers_cart.CreateCart, Method: "POST"},
	{Path: "/additemtocart", Handler: handlers_cart.AddToCart, Method: "POST"},
	{Path: "/additemstocart", Handler: handlers_cart.AddItemsToCart, Method: "POST"},
	{Path: "/cart", Handler: handlers_cart.GetCart, Method: "GET"},
	{Path: "/removeitemfromcart", Handler: handlers_cart.RemoveItemFromCart, Method: "DELETE"},
}
