# Product Catalog API - Golang

This is a product catalog API which has a list of products that can be added to the cart.
This works similar to how normal e-com website works, but in the form of API.

## Dependencies

Install Postgresql

```bash
sudo sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list'
wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -
sudo apt-get update
sudo apt-get -y install postgresql
```

Install the following packages using go get

```bash
go get github.com/gorilla/mux
go get github.com/google/uuid
go get github.com/lib/pq
```

## Environment Setup

Start the Postgresql server

```bash
sudo service postgresql start
```

Connect to the psql server

```bash
sudo -u postgres psql
```

Create a database user for the application

```bash
createuser service-pc-api
```

Create a database for the API

```bash
createdb -O service-pc-api product-catalog -h localhost -U service-pc-api
```

Initialize the Database schema

```bash
psql -h localhost -U service-pc-api -f sql_commands/init.sql
```

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`CONNECTION_STRING` - Postgres connection string/url

## Database Design

![alt text](https://imgur.com/p4EyQnJ.png)

## API Reference

### Product

<!--Get Products-->
<details>
<summary>Get Products</summary>
<br>

Sends the list of all products in an array

```http
GET /products
```

#### Query Parameters

| Field            | Type  | Description                     |
| :--------------- | :---- | :------------------------------ |
| `page`           | `int` | **Optional**. Default Value: 1  |
| `items_per_page` | `int` | **Optional**. Default Value: 20 |

</details>

<!--Get Product-->
<details>
<summary>Get Product</summary>
<br>

Sends a single product referenced by `id`

```http
GET /product/${id}
```

</details>

<!--Add Product-->
<details>
<summary>Add Product</summary>
<br>

Adds a product to the database

```http
POST /product
```

#### Request Body

| Field           | Type     | Description                                      |
| :-------------- | :------- | :----------------------------------------------- |
| `product_id`    | `int`    | **Required**. ID of the product.                 |
| `name`          | `int`    | **Required**. The name of the product.           |
| `specification` | `JSON`   | **Required**. The Specifications of the product. |
| `sku`           | `string` | **Required**. The product's category ID (FK).    |
| `category_id`   | `int`    | **Required**. The name of the product.           |
| `price`         | `float`  | **Required**. The product's price                |

</details>

<!--Delete Product-->
<details>
<summary>Delete Product</summary>
<br>

Deletes a product from the database referenced by the `id`

```http
DELETE /product/${id}
```

</details>

<!--Update Product-->
<details>
<summary>Update a Product</summary>
<br>

Updates a product in the database referenced by `id`

```http
PUT /product/${id}
```

#### Request Body

| Field           | Type     | Description                                      |
| :-------------- | :------- | :----------------------------------------------- |
| `product_id`    | `int`    | **Optional**. ID of the product.                 |
| `name`          | `int`    | **Optional**. The name of the product.           |
| `specification` | `JSON`   | **Optional**. The Specifications of the product. |
| `sku`           | `string` | **Optional**. The product's category ID (FK).    |
| `category_id`   | `int`    | **Optional**. The name of the product.           |
| `price`         | `float`  | **Optional**. The product's price                |

**Note:** `product_id` cannot be updated.

</details>

### Category

<!--Add Category-->
<details>
<summary>Add Category</summary>
<br>

Adds a new category to the database

```http
POST /category
```

#### Request Body

| Field         | Type     | Description                               |
| :------------ | :------- | :---------------------------------------- |
| `category_id` | `int`    | **Required**. Unique ID for the Category. |
| `name`        | `string` | **Required**. The name of the category.   |

</details>

<!--Get Categories-->
<details>
<summary>Get Categories</summary>
<br>

Get an array of all Categories

```http
GET /categories
```

</details>

<!--Delete Category-->
<details>
<summary>Delete Category</summary>
<br>

Deletes a category from the database referenced by `id`

```http
DELETE /category/${id}
```

</details>

<!--Update Category-->
<details>
<summary>Update Category</summary>
<br>

Update the category field referenced by `id`

```http
PUT /category/${id}
```

</details>

### Inventory

<!--Get Inventory-->
<details>
<summary>Get Inventory</summary>
<br>

View the current stock in inventory

```http
GET /inventory
```

</details>

<!--Update Inventory Item-->
<details>
<summary>Update Inventory Item</summary>
<br>

Add/Update the inventory.

```http
POST /inventory
```

#### Request Body: JSON

| Field        | Type  | Description                                             |
| :----------- | :---- | :------------------------------------------------------ |
| `product_id` | `int` | **Required**. Product ID to update/update in inventory. |
| `quantity`   | `int` | **Required**. The quantity to update/update             |

</details>

### Cart

<!--Add Item to Cart-->
<details>
<summary>Add Item to Cart</summary>
<br>

Add new item or increase existing quantity.

```http
POST /additemtocart
```

#### Query Parameters:

| Field | Type     | Description                                                         |
| :---- | :------- | :------------------------------------------------------------------ |
| `ref` | `string` | **Optional**. Cart Reference ID. New cart is created if not passed. |

#### Request Body:

| Field        | Type  | Description                                  |
| :----------- | :---- | :------------------------------------------- |
| `product_id` | `int` | **Required**. Product ID                     |
| `quantity`   | `int` | **Required**. Quantity of the product to add |

</details>

<!--Add items to cart-->
<details>
<summary>Add items to cart</summary>
<br>

Add new **items** or increase existing quantity.

```http
POST /additemstocart
```

#### Query Parameters:

| Field | Type     | Description                     |
| :---- | :------- | :------------------------------ |
| `ref` | `string` | **Required**. Cart Reference ID |

#### Request Body

Array of JSON of structure:

| Field        | Type  | Description                                  |
| :----------- | :---- | :------------------------------------------- |
| `product_id` | `int` | **Required**. Product ID                     |
| `quantity`   | `int` | **Required**. Quantity of the product to add |

</details>

<!--Get Cart-->
<details>
<summary>Get Cart</summary>
<br>

Get the cart referenced by `ref`

```http
GET /cart
```

#### Query Parameters:

| Field | Type     | Description                     |
| :---- | :------- | :------------------------------ |
| `ref` | `string` | **Required**. Cart Reference ID |

</details>

<!--Remove Item from Cart-->
<details>
<summary>Remove Item from Cart</summary>
<br>

Remove a Product ID from a cart referenced by

```http
DELETE /removeitemfromcart
```

#### Query Parameters:

| Field | Type     | Description                      |
| :---- | :------- | :------------------------------- |
| `ref` | `string` | **Required**. Cart Reference ID. |

#### Request Body:

| Field        | Type  | Description                                     |
| :----------- | :---- | :---------------------------------------------- |
| `product_id` | `int` | **Required**. Product ID                        |
| `quantity`   | `int` | **Required**. Quantity of the product to remove |

</details>

## Acknowledgements

- [Database Design](https://app.diagrams.net/#G1YuY3PY67Qg_d9O4dic71VcaWJdOD_obM)
- [API Demo Hosted Base URL](https://product-catalog-api.subramanyamr.repl.co)
