
## API Reference

#### Get Products
Sends the list of all products in an array

```http
GET /products/${page_no}
```

#### Get Product
Sends a single product referenced by **${id}**

```http
GET /product/${id}
```

#### Add Product
Adds a product to the database

```http
POST /product/add
```
#### Request Body

JSON:

- `product_id` (int, required): ID of the product.
- `name` (string, required): The name of the product.
- `specification` (JSON, required): The Specifications of the product.
- `sku` (string, required): Stock Keeping Unit number of the product.
- `category_id` (int, required): The product's category ID. This needs to be present in the Category Table.
- `price` (float, required): The product's price

#### Delete Product
Deletes a product from the database referenced by the **${id}**

```http
DELETE /product/delete/${id}
```

#### Update a Product
Updates a product in the database referenced by **${id}**

**Note:** `product_id` cannot be updated.

```http
PUT /product/update/${id}
```

- `name` (string, optional): The name of the product.
- `specification` (JSON, optional): The Specifications of the product.
- `sku` (string, optional): Stock Keeping Unit number of the product.
- `category_id` (int, optional): The product's category ID. This needs to be present in the Category Table.
- `price` (float, optional): The product's price