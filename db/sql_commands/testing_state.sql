DELETE FROM cart_item;
DELETE FROM inventory;
DELETE FROM product;
DELETE FROM cart_reference;
DELETE FROM category;

INSERT INTO category VALUES(1, 'Clothes');
INSERT INTO category VALUES(2, 'Electronics');
INSERT INTO category VALUES(3, 'Perfumes');

INSERT INTO product VALUES(1, 'Hoodie', '{"color": "black", "gender": "male"}', 65544, 1, 10.99);
INSERT INTO product VALUES(2, 'TWS', '{"color": "white", "ANC": "true"}', 65546, 2, 14.99);
INSERT INTO product VALUES(3, 'Jeans', '{"color": "Jean Blue"}', 65547, 1, 25.99);

INSERT INTO inventory VALUES(3, 100);

INSERT INTO cart_item VALUES('4d4d8297-7663-451d-b79e-49a545728552', 3, 10);

INSERT INTO cart_reference VALUES('4d4d8297-7663-451d-b79e-49a545728552', '2022-12-28 13:52:51.016582');