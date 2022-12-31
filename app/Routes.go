package app

import (
	"github.com/subramanyam-searce/product-catalog-go/handlers"
	"github.com/subramanyam-searce/product-catalog-go/typedefs"
)

var Routes []typedefs.Route = []typedefs.Route{
	{Path: "/", Handler: handlers.Home, Method: "GET"},
	{Path: "/product", Handler: handlers.AddProduct, Method: "POST"},
	{Path: "/products", Handler: handlers.GetProducts, Method: "GET"},
	{Path: "/product/{id:[0-9]+}", Handler: handlers.GetProduct, Method: "GET"},
	{Path: "/product/{id:[0-9]+}", Handler: handlers.DeleteProduct, Method: "DELETE"},
	{Path: "/product/{id:[0-9]+}", Handler: handlers.UpdateProduct, Method: "PUT"},

	{Path: "/category", Handler: handlers.AddCategory, Method: "POST"},
	{Path: "/categories", Handler: handlers.GetCategories, Method: "GET"},
	{Path: "/category/{id:[0-9]+}", Handler: handlers.GetCategory, Method: "GET"},
	{Path: "/category/{id:[0-9]+}", Handler: handlers.DeleteCategory, Method: "DELETE"},
	{Path: "/category/{id:[0-9]+}", Handler: handlers.UpdateCategory, Method: "PUT"},

	{Path: "/inventory", Handler: handlers.UpdateInventory, Method: "POST"},
	{Path: "/inventory", Handler: handlers.GetInventory, Method: "GET"},

	{Path: "/additemtocart", Handler: handlers.AddToCart, Method: "POST"},
	{Path: "/additemstocart", Handler: handlers.AddItemsToCart, Method: "POST"},
	{Path: "/cart", Handler: handlers.GetCart, Method: "GET"},
	{Path: "/removeitemfromcart", Handler: handlers.RemoveItemFromCart, Method: "DELETE"},
}
