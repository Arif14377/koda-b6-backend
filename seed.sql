select * from products;
select * from product_variant;
select * from product_size;
select * from users;
select * from cart;
select * from reviews;
select * from transactions;
select * from transaction_product;
select * from categories; -- baru 5
select * from product_category;
select * from product_images;

create table products (
    id serial primary key,
    name varchar(30),
    description varchar(200),
    quantity int,
    price int
);

CREATE TABLE product_variant (
    id serial PRIMARY KEY,
    name varchar(20) not null,
    product_id int not null,
    add_price int,
    constraint fk_product
        foreign key(product_id)
        references products(id)
);

create table product_size (
    id serial PRIMARY KEY,
    name varchar(20) not null,
    product_id int not null,
    add_price int,
    constraint fk_product
        foreign key(product_id)
        references products(id)
);

create table product_images (
    id serial PRIMARY KEY,
    product_id int not null,
    path varchar(200),
    constraint fk_product
        foreign key(product_id)
        references products(id)
);

create table users (
    id serial PRIMARY KEY,
    full_name varchar(80),
    email varchar(25) UNIQUE NOT NULL,
    password text,
    address varchar(80),
    phone varchar(16),
    picture varchar(120),
    created_at TIMESTAMP DEFAULT now()
);

create table cart (
    id serial PRIMARY KEY,
    user_id int,
    product_id int,
    constraint fk_user
        foreign key(user_id)
        references users(id),
    constraint fk_products
        foreign key(product_id)
        references products(id)
);

create table reviews (
    id serial PRIMARY KEY,
    user_id int,
    messages varchar(600),
    rating int,
    constraint users
        foreign key(user_id)
        references users(id)
);

create table transactions (
    id bigserial PRIMARY KEY,
    trx_code varchar(25) UNIQUE,
    delivery_method varchar(60),
    full_name varchar(80),
    email varchar(25),
    address varchar(80),
    sub_total int,
    tax int,
    total int,
    date TIMESTAMP DEFAULT now(),
    status varchar(15),
    payment_method varchar(20)
);

create table transaction_product (
    id serial PRIMARY KEY,
    product_id int,
    transaction_id int,
    quantity int,
    constraint fk_product
        foreign key(product_id)
        references products(id),
    constraint fk_transaction
        foreign key(transaction_id)
        references transactions(id)
);

create table categories (
    id serial primary key,
    name varchar(80)
);

create table product_category (
    id serial primary key,
    product_id int,
    category_id int,
    constraint product
        foreign key(product_id)
        references products(id),
    constraint category
        foreign key(category_id)
        references categories(id)
);


-- Fill products
insert into products(name, description, quantity, price) values
    ('Espresso', 'Kopi hitam pekat dengan crema tebal', 50, 18000),
    ('Americano', 'Espresso dengan tambahan air panas', 45, 20000),
    ('Cappuccino', 'Espresso dengan steamed milk dan foam lembut', 40, 25000),
    ('Caffe Latte', 'Espresso dengan susu creamy', 42, 25000),
    ('Caramel Latte', 'Latte dengan sirup karamel manis', 35, 28000),
    ('Vanilla Latte', 'Latte dengan aroma vanilla lembut', 30, 28000),
    ('Hazelnut Latte', 'Latte dengan rasa hazelnut gurih', 28, 28000),
    ('Mocha', 'Perpaduan espresso, cokelat, dan susu', 32, 30000),
    ('Affogato', 'Espresso dengan scoop es krim vanilla', 20, 32000),
    ('Cold Brew', 'Kopi seduh dingin dengan rasa smooth', 25, 27000),
    ('Matcha Latte', 'Teh hijau Jepang dengan susu creamy', 30, 28000),
    ('Chocolate', 'Minuman cokelat hangat premium', 35, 26000),
    ('Taro Latte', 'Minuman taro manis dan creamy', 25, 27000),
    ('Red Velvet Latte', 'Minuman red velvet lembut', 20, 29000),
    ('Thai Tea', 'Teh Thailand dengan susu kental manis', 30, 24000),
    ('Lemon Tea', 'Teh segar dengan perasan lemon', 40, 20000),
    ('Peach Tea', 'Teh dengan aroma dan rasa peach', 22, 23000),
    ('Mineral Water', 'Air mineral botol dingin', 60, 10000),
    ('Croissant Butter', 'Pastry renyah dengan aroma butter', 20, 22000),
    ('Chocolate Croissant', 'Croissant isi cokelat lumer', 18, 25000),
    ('Chicken Sandwich', 'Roti isi ayam dan sayur segar', 15, 30000),
    ('Beef Burger', 'Burger daging sapi dengan saus spesial', 12, 35000),
    ('French Fries', 'Kentang goreng crispy', 25, 20000),
    ('Spaghetti Bolognese', 'Pasta dengan saus daging tomat', 14, 38000),
    ('Chicken Wrap', 'Tortilla isi ayam dan sayuran', 16, 32000),
    ('Cheese Cake Slice', 'Potongan cheesecake lembut', 10, 30000),
    ('Extra Shot Espresso', 'Tambahan satu shot espresso', 100, 8000),
    ('Syrup Caramel', 'Tambahan sirup karamel', 80, 5000),
    ('Syrup Vanilla', 'Tambahan sirup vanilla', 80, 5000),
    ('Whipped Cream', 'Tambahan whipped cream lembut', 70, 7000);

-- fill product_variant
insert into product_variant(product_id, name, add_price) values
-- Espresso (1)
(1, 'ice', 500),
(1, 'extra shot', 8000),

-- Americano (2)
(2, 'ice', 1000),
(2, 'less sugar', 0),

-- Cappuccino (3)
(3, 'ice', 1000),
(3, 'soy milk', 3000),

-- Caffe Latte (4)
(4, 'ice', 1000),
(4, 'extra sugar', 1000),

-- Caramel Latte (5)
(5, 'ice', 1500),
(5, 'soy milk', 3000),

-- Mocha (8)
(8, 'ice', 1500),
(8, 'extra topping', 5000),

-- Matcha Latte (11)
(11, 'ice', 1000),
(11, 'soy milk', 3000),

-- Thai Tea (15)
(15, 'less sugar', 0),
(15, 'extra sugar', 1000),

-- Beef Burger (22)
(22, 'extra cheese', 4000),
(22, 'double meat', 8000),

-- Spaghetti Bolognese (24)
(24, 'extra sauce', 3000),
(24, 'extra cheese', 4000);
    

-- fill product_size
insert into product_size(name, product_id, add_price) values
-- Espresso (1)
('Regular', 1, 0),
('Medium', 1, 3000),
('Large', 1, 5000),

-- Americano (2)
('Regular', 2, 0),
('Medium', 2, 4000),
('Large', 2, 6000),

-- Cappuccino (3)
('Regular', 3, 0),
('Medium', 3, 5000),
('Large', 3, 7000),

-- Caffe Latte (4)
('Regular', 4, 0),
('Medium', 4, 5000),
('Large', 4, 7000),

-- Caramel Latte (5)
('Regular', 5, 0),
('Medium', 5, 6000),
('Large', 5, 8000),

-- Mocha (8)
('Regular', 8, 0),
('Medium', 8, 6000),
('Large', 8, 9000),

-- Matcha Latte (11)
('Regular', 11, 0),
('Medium', 11, 5000);


select * from users;

-- fill roles
INSERT INTO roles(name) VALUES
('admin'),
('user');

-- fill users
INSERT INTO users(id, full_name, email, password, address, phone, picture, role_id) VALUES
('550e8400-e29b-41d4-a716-446655440001', 'Arif Rahman', 'arif.rahman@gmail.com', '1234', 'Jl. Kapuk No. 3 Tangsel', '081234560001', 'https://placehold.co/200x200', 1),
('550e8400-e29b-41d4-a716-446655440002', 'Dewi Lestari', 'dewi.lestari@gmail.com', '1234', 'Jl. Kenanga No. 8 Bandung', '081234560002', 'https://placehold.co/200x200', 2),
('550e8400-e29b-41d4-a716-446655440003', 'Rizky Saputra', 'rizky.saputra@gmail.com', '1234', 'Jl. Mawar No. 21 Surabaya', '081234560003', 'https://placehold.co/200x200', 2),
('550e8400-e29b-41d4-a716-446655440004', 'Sinta Maharani', 'sinta.maharani@gmail.com', '1234', 'Jl. Anggrek No. 5 Yogyakarta', '081234560004', 'https://placehold.co/200x200', 2),
('550e8400-e29b-41d4-a716-446655440005', 'Fajar Nugroho', 'fajar.nugroho@gmail.com', '1234', 'Jl. Flamboyan No. 17 Semarang', '081234560005', 'https://placehold.co/200x200', 2),
('550e8400-e29b-41d4-a716-446655440006', 'Nadia Putri', 'nadia.putri@gmail.com', '1234', 'Jl. Dahlia No. 9 Medan', '081234560006', 'https://placehold.co/200x200', 2),
('550e8400-e29b-41d4-a716-446655440007', 'Bima Kurniawan', 'bima.kurniawan@gmail.com', '1234', 'Jl. Cempaka No. 14 Bekasi', '081234560007', 'https://placehold.co/200x200', 2),
('550e8400-e29b-41d4-a716-446655440008', 'Laras Wulandari', 'laras.wulandari@gmail.com', '1234', 'Jl. Teratai No. 3 Depok', '081234560008', 'https://placehold.co/200x200', 2),
('550e8400-e29b-41d4-a716-446655440009', 'Andika Ramadhan', 'andika.ramadhan@gmail.com', '1234', 'Jl. Bougenville No. 11 Tangerang', '081234560009', 'https://placehold.co/200x200', 2),
('550e8400-e29b-41d4-a716-446655440010', 'Maya Oktaviani', 'maya.oktaviani@gmail.com', '1234', 'Jl. Sakura No. 6 Bogor', '081234560010', 'https://placehold.co/200x200', 2);

-- fill cart;
insert into cart(user_id, product_id) values
    (1, 3),
    (1, 7),
    (1, 12),
    (2, 5),
    (2, 14),
    (2, 21),
    (3, 1),
    (3, 9),
    (3, 18),
    (4, 2),
    (4, 11),
    (4, 25),
    (5, 4),
    (5, 16),
    (5, 30),
    (6, 6),
    (6, 13),
    (6, 22),
    (7, 8),
    (7, 19),
    (7, 27),
    (8, 10),
    (8, 15),
    (8, 24),
    (9, 17),
    (9, 20),
    (9, 28),
    (10, 23),
    (10, 26),
    (10, 29);

-- fill reviews
insert into reviews(user_id, messages, rating) values
    (1, 'Kopinya enak dan aromanya kuat, pasti order lagi.', 5),
    (2, 'Tempatnya nyaman tapi minumannya agak lama datangnya.', 4),
    (3, 'Rasa matchanya kurang terasa, mungkin bisa ditingkatkan.', 3),
    (4, 'Pelayanannya ramah dan cepat, sangat puas.', 5),
    (5, 'Harga sedikit mahal tapi kualitas sesuai.', 4),
    (6, 'Kopi terlalu pahit untuk selera saya.', 2),
    (7, 'Makanannya enak tapi minuman kurang dingin.', 3),
    (8, 'Pelayanan kurang responsif saat ramai.', 2),
    (9, 'Suka banget sama caramel lattenya, recommended!', 5),
    (10, 'Overall oke, tapi tempat parkir agak sempit.', 4);

-- fill transactions
insert into transactions (trx_code, delivery_method, full_name, email, address, sub_total, tax, total, status, payment_method) values
('TRX-2026-02-0001', 'Dine In', 'Arif Pratama', 'arif.pratama@gmail.com', 'Jl. Melati No. 12 Jakarta', 37000, 3700, 40700, 'Finish Order', 'Bank BRI'),
('TRX-2026-02-0002', 'Pick Up', 'Dewi Lestari', 'dewi.lestari@gmail.com', 'Jl. Kenanga No. 8 Bandung', 24000, 2400, 26400, 'Finish Order', 'DANA'),
('TRX-2026-02-0003', 'Door Delivery', 'Rizky Saputra', 'rizky.saputra@gmail.com', 'Jl. Mawar No. 21 Surabaya', 94500, 9450, 103950, 'Sending Goods', 'Bank BCA'),
('TRX-2026-02-0004', 'Dine In', 'Sinta Maharani', 'sinta.maharani@gmail.com', 'Jl. Anggrek No. 5 Yogyakarta', 25000, 2500, 27500, 'Finish Order', 'Gopay'),
('TRX-2026-02-0005', 'Door Delivery', 'Fajar Nugroho', 'fajar.nugroho@gmail.com', 'Jl. Flamboyan No. 17 Semarang', 140000, 14000, 154000, 'On Progress', 'OVO'),
('TRX-2026-02-0006', 'Pick Up', 'Nadia Putri', 'nadia.putri@gmail.com', 'Jl. Dahlia No. 9 Medan', 34000, 3400, 37400, 'Finish Order', 'Paypal'),
('TRX-2026-02-0007', 'Dine In', 'Bima Kurniawan', 'bima.kurniawan@gmail.com', 'Jl. Cempaka No. 14 Bekasi', 114000, 11400, 125400, 'Sending Goods', 'Bank BRI'),
('TRX-2026-02-0008', 'Door Delivery', 'Laras Wulandari', 'laras.wulandari@gmail.com', 'Jl. Teratai No. 3 Depok', 68000, 6800, 74800, 'On Progress', 'DANA'),
('TRX-2026-02-0009', 'Pick Up', 'Andika Ramadhan', 'andika.ramadhan@gmail.com', 'Jl. Bougenville No. 11 Tangerang', 38500, 3850, 42350, 'Finish Order', 'Gopay'),
('TRX-2026-02-0010', 'Dine In', 'Maya Oktaviani', 'maya.oktaviani@gmail.com', 'Jl. Sakura No. 6 Bogor', 54000, 5400, 59400, 'Finish Order', 'Bank BCA'),
('TRX-2026-02-0011', 'Door Delivery', 'Arif Pratama', 'arif.pratama@gmail.com', 'Jl. Melati No. 12 Jakarta', 22000, 2200, 24200, 'On Progress', 'OVO'),
('TRX-2026-02-0012', 'Pick Up', 'Dewi Lestari', 'dewi.lestari@gmail.com', 'Jl. Kenanga No. 8 Bandung', 62000, 6200, 68200, 'Sending Goods', 'Paypal'),
('TRX-2026-02-0013', 'Dine In', 'Rizky Saputra', 'rizky.saputra@gmail.com', 'Jl. Mawar No. 21 Surabaya', 28500, 2850, 31350, 'Finish Order', 'Bank BRI'),
('TRX-2026-02-0014', 'Door Delivery', 'Sinta Maharani', 'sinta.maharani@gmail.com', 'Jl. Anggrek No. 5 Yogyakarta', 96000, 9600, 105600, 'On Progress', 'DANA');

-- fill transaction_product (perlu isi tabel transactions dulu)
insert into transaction_product(product_id, transaction_id, quantity) values
    (1, 1, 3),
    (2, 2, 1),
    (3, 3, 2),
    (4, 4, 2),
    (5, 5, 2),
    (6, 6, 1),
    (7, 6, 1),
    (1, 7, 4),
    (8, 8, 5),
    (8, 9, 3),
    (1, 9, 2),
    (2, 9, 6),
    (10, 10, 3),
    (11, 11, 2),
    (12, 12, 2);

insert into categories(name) values
        ('Favourite Product'),
        ('Coffee'),
        ('Non-Coffee'),
        ('Foods'),
        ('Add-On');

insert into product_category(product_id, category_id) values
    (1, 2),
    (2, 2),
    (3, 2), (3, 1),
    (4, 2),
    (5, 2),
    (6, 2), (6, 1),
    (7, 2),
    (8, 2),
    (9, 2),
    (10, 2),

    (11, 3),
    (12, 3),
    (13, 3),
    (14, 3),
    (15, 3),
    (16, 3),
    (17, 3),
    (18, 3),

    (19, 4), (19, 3),
    (20, 4), (20, 3),
    (21, 4), (21, 3),
    (22, 4), (22, 3),
    (23, 4), (23, 3),
    (24, 4), (24, 3),
    (25, 4), (25, 3),
    (26, 4), (26, 3),

    (27, 5),
    (28, 5),
    (29, 5),
    (30, 5);

-- fill product_images
insert into product_images(product_id, path) values
(1, 'https://placehold.co/600x600'),
(1, 'https://placehold.co/600x600'),
(1, 'https://placehold.co/600x600'),

(2, 'https://placehold.co/600x600'),
(2, 'https://placehold.co/600x600'),
(2, 'https://placehold.co/600x600'),

(3, 'https://placehold.co/600x600'),
(3, 'https://placehold.co/600x600'),
(3, 'https://placehold.co/600x600'),

(4, 'https://placehold.co/600x600'),
(4, 'https://placehold.co/600x600'),
(4, 'https://placehold.co/600x600'),

(5, 'https://placehold.co/600x600'),
(5, 'https://placehold.co/600x600'),
(5, 'https://placehold.co/600x600'),

(6, 'https://placehold.co/600x600'),
(6, 'https://placehold.co/600x600'),
(6, 'https://placehold.co/600x600'),

(7, 'https://placehold.co/600x600'),
(7, 'https://placehold.co/600x600'),
(7, 'https://placehold.co/600x600'),

(8, 'https://placehold.co/600x600'),
(8, 'https://placehold.co/600x600'),
(8, 'https://placehold.co/600x600'),

(9, 'https://placehold.co/600x600'),
(9, 'https://placehold.co/600x600'),
(9, 'https://placehold.co/600x600'),

(10, 'https://placehold.co/600x600'),
(10, 'https://placehold.co/600x600'),
(10, 'https://placehold.co/600x600'),

(11, 'https://placehold.co/600x600'),
(11, 'https://placehold.co/600x600'),
(11, 'https://placehold.co/600x600'),

(12, 'https://placehold.co/600x600'),
(12, 'https://placehold.co/600x600'),
(12, 'https://placehold.co/600x600'),

(13, 'https://placehold.co/600x600'),
(13, 'https://placehold.co/600x600'),
(13, 'https://placehold.co/600x600'),

(14, 'https://placehold.co/600x600'),
(14, 'https://placehold.co/600x600'),
(14, 'https://placehold.co/600x600'),

(15, 'https://placehold.co/600x600'),
(15, 'https://placehold.co/600x600'),
(15, 'https://placehold.co/600x600'),

(16, 'https://placehold.co/600x600'),
(16, 'https://placehold.co/600x600'),
(16, 'https://placehold.co/600x600'),

(17, 'https://placehold.co/600x600'),
(17, 'https://placehold.co/600x600'),
(17, 'https://placehold.co/600x600'),

(18, 'https://placehold.co/600x600'),
(18, 'https://placehold.co/600x600'),
(18, 'https://placehold.co/600x600'),

(19, 'https://placehold.co/600x600'),
(19, 'https://placehold.co/600x600'),
(19, 'https://placehold.co/600x600'),

(20, 'https://placehold.co/600x600'),
(20, 'https://placehold.co/600x600'),
(20, 'https://placehold.co/600x600'),

(21, 'https://placehold.co/600x600'),
(21, 'https://placehold.co/600x600'),
(21, 'https://placehold.co/600x600'),

(22, 'https://placehold.co/600x600'),
(22, 'https://placehold.co/600x600'),
(22, 'https://placehold.co/600x600'),

(23, 'https://placehold.co/600x600'),
(23, 'https://placehold.co/600x600'),
(23, 'https://placehold.co/600x600'),

(24, 'https://placehold.co/600x600'),
(24, 'https://placehold.co/600x600'),
(24, 'https://placehold.co/600x600'),

(25, 'https://placehold.co/600x600'),
(25, 'https://placehold.co/600x600'),
(25, 'https://placehold.co/600x600'),

(26, 'https://placehold.co/600x600'),
(26, 'https://placehold.co/600x600'),
(26, 'https://placehold.co/600x600'),

(27, 'https://placehold.co/600x600'),
(27, 'https://placehold.co/600x600'),
(27, 'https://placehold.co/600x600'),

(28, 'https://placehold.co/600x600'),
(28, 'https://placehold.co/600x600'),
(28, 'https://placehold.co/600x600'),

(29, 'https://placehold.co/600x600'),
(29, 'https://placehold.co/600x600'),
(29, 'https://placehold.co/600x600'),

(30, 'https://placehold.co/600x600'),
(30, 'https://placehold.co/600x600'),
(30, 'https://placehold.co/600x600');