
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
github.com/gorilla/mux
github.com/google/uuid
github.com/lib/pq
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