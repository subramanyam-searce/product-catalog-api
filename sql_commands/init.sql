CREATE TABLE category (
  category_id int,
  name varchar,
  PRIMARY KEY(category_id)
);

CREATE TABLE product (
  product_id int,
  name varchar,
  specification JSON,
  sku varchar,
  category_id int,
  price float,
  PRIMARY KEY(product_id),
  FOREIGN KEY(category_id) REFERENCES category(category_id)
);
 
 CREATE TABLE inventory (
   product_id int,
   quantity int CHECK(quantity >= 0),
   PRIMARY KEY(product_id),
   FOREIGN KEY (product_id) REFERENCES product(product_id)
  );
  
  CREATE TABLE cart_reference (
    ref varchar,
    created_at timestamp,
    PRIMARY KEY(ref)
  );
   
   CREATE TABLE cart_item (
     ref varchar,
     product_id int,
     quantity int,
     PRIMARY KEY(ref, product_id),
     FOREIGN KEY(product_id) REFERENCES product(product_id)
   );

   GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO "service-pc-api"; 