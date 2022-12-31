package queries

var InsertProduct string = "INSERT INTO product VALUES($1, $2, $3, $4, $5, $6)"
var GetProduct string = "SELECT * FROM product WHERE product_id=$1"
var GetAllProducts string = "SELECT * FROM product ORDER BY product_id ASC"
var DeleteProduct string = "DELETE FROM product WHERE product_id=$1"

var GetAllCategories string = "SELECT * FROM category ORDER BY category_id ASC"
var GetCategory string = "SELECT * FROM category WHERE category_id=$1"
var DeleteCategory string = "DELETE FROM category WHERE category_id=$1"
var AddCategory string = "INSERT INTO category VALUES($1, $2)"

var GetInventory string = "SELECT p.product_id, p.name, i.quantity FROM product p INNER JOIN inventory i ON p.product_id=i.product_id"
var GetInventoryItem string = "SELECT * FROM inventory WHERE product_id=$1"
var UpdateInventoryItem string = "UPDATE inventory SET quantity=$1 WHERE product_id=$2"
var DeleteInventoryItem string = "DELETE FROM inventory WHERE product_id=$1"
var InsertInventoryItem string = "INSERT INTO inventory VALUES($1, $2)"

var InsertCartReference string = "INSERT INTO cart_reference VALUES($1, $2);"
var GetCartReference string = "SELECT * FROM cart_reference WHERE ref=$1;"
var GetCartItems string = "SELECT * FROM cart_item WHERE ref=$1"
var GetCartItemQuantity string = "SELECT quantity FROM cart_item WHERE ref=$1 AND product_id=$2"
var UpdateCartItemQuantity string = "UPDATE cart_item SET quantity=$1 WHERE ref=$2 AND product_id=$3"
var InsertCartItem string = "INSERT INTO cart_item VALUES($1, $2, $3);"
var DeleteCartItem string = "DELETE FROM cart_item WHERE ref=$1 AND product_id=$2"
var DeleteCartItemsWithProductID string = "DELETE FROM cart_item WHERE product_id=$1"
var UpsertInventoryForCartItemRemoval string = "INSERT INTO inventory VALUES($1, $2) ON CONFLICT ON CONSTRAINT inventory_pkey DO UPDATE SET quantity=(SELECT quantity FROM inventory WHERE product_id=$1)+$2"
