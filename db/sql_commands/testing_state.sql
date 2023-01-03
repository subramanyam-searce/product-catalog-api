DELETE FROM cart_item;
DELETE FROM inventory;
DELETE FROM product;
DELETE FROM cart_reference;
DELETE FROM category;

INSERT INTO category VALUES(1, 'Clothes');
INSERT INTO category VALUES(2, 'Electronics');
INSERT INTO category VALUES(3, 'Plastics');
INSERT INTO category VALUES(4, 'Wooden Products');
INSERT INTO category VALUES(5, 'Luxury');

INSERT INTO product VALUES(1, 'Hoodie', '{"color": "black", "gender": "male"}', 65544, 1, 10.99);
INSERT INTO product VALUES(2, 'TWS', '{"color": "white", "ANC": "true"}', 65546, 2, 14.99);
INSERT INTO product VALUES(3, 'Jeans', '{"color": "Jean Blue"}', 65547, 1, 25.99);
INSERT INTO product VALUES(4, 'Formal Shirt', '{"color": "White", "pattern": "striped", "full_sleeves": "true"}', 65548, 1, 15.99);
INSERT INTO product VALUES(5, 'Pen Stand', '{"Height":"15cm"}', 65549, 4, 5.99);
INSERT INTO product VALUES(6, 'Keyboard', '{"type":"Gaming"}', 65550, 2, 30.99);
INSERT INTO product VALUES(7, 'Mouse', '{"type":"Gaming"}', 65551, 2, 24.99);

INSERT INTO inventory VALUES(1, 100);
INSERT INTO inventory VALUES(2, 50);
INSERT INTO inventory VALUES(3, 20);
INSERT INTO inventory VALUES(4, 40);
INSERT INTO inventory VALUES(6, 5);

INSERT INTO cart_item VALUES('4d4d8297-7663-451d-b79e-49a545728552', 3, 10);

INSERT INTO cart_reference VALUES('4d4d8297-7663-451d-b79e-49a545728552', '2022-12-28 13:52:51.016582');