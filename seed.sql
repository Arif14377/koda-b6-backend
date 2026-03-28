-- DROP TABLES IN REVERSE ORDER TO HANDLE CONSTRAINTS
DROP TABLE IF EXISTS product_category CASCADE;
DROP TABLE IF EXISTS delivery_methods CASCADE;
DROP TABLE IF EXISTS categories CASCADE;
DROP TABLE IF EXISTS transaction_product CASCADE;
DROP TABLE IF EXISTS transactions CASCADE;
DROP TABLE IF EXISTS reviews CASCADE;
DROP TABLE IF EXISTS cart CASCADE;
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS roles CASCADE;
DROP TABLE IF EXISTS product_images CASCADE;
DROP TABLE IF EXISTS product_size CASCADE;
DROP TABLE IF EXISTS product_variant CASCADE;
DROP TABLE IF EXISTS products CASCADE;

-- CREATE TABLES (consistent with migrations/000001_init_db.up.sql)
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(200) NOT NULL,
    quantity INT,
    price BIGINT,
    rating INT DEFAULT 0,
    old_price BIGINT,
    is_flash_sale BOOLEAN DEFAULT FALSE
);

CREATE TABLE product_variant (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    product_id INT NOT NULL,
    add_price BIGINT,
    CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES products(id)
);

CREATE TABLE product_size (
    id SERIAL PRIMARY KEY,
    name VARCHAR(20) NOT NULL,
    product_id INT NOT NULL,
    add_price BIGINT,
    CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES products(id)
);

CREATE TABLE product_images (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    path VARCHAR(500),
    CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES products(id)
);

CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(10) CHECK (name IN('admin', 'user'))
);

CREATE TABLE users (
    id VARCHAR(36) PRIMARY KEY,
    full_name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password TEXT,
    address VARCHAR(200),
    phone VARCHAR(20),
    picture VARCHAR(500),
    role_id INT,
    created_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_roles FOREIGN KEY (role_id) REFERENCES roles(id)
);

CREATE TABLE cart (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(36),
    product_id INT,
    quantity INT NOT NULL DEFAULT 1,
    size_id INT,
    variant_id INT,
    created_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_products FOREIGN KEY(product_id) REFERENCES products(id),
    CONSTRAINT fk_size FOREIGN KEY(size_id) REFERENCES product_size(id),
    CONSTRAINT fk_variant FOREIGN KEY(variant_id) REFERENCES product_variant(id)
);

CREATE TABLE reviews (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(36),
    messages VARCHAR,
    rating INT CHECK (rating BETWEEN 1 AND 5),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_user_review FOREIGN KEY(user_id) REFERENCES users(id)
);

CREATE TABLE transactions (
    id BIGSERIAL PRIMARY KEY,
    user_id VARCHAR(36),
    trx_code VARCHAR(50) UNIQUE NOT NULL,
    delivery_method VARCHAR(60) NOT NULL,
    full_name VARCHAR(100),
    email VARCHAR(100),
    address VARCHAR(200),
    sub_total BIGINT,
    tax BIGINT,
    total BIGINT,
    date TIMESTAMP DEFAULT NOW(),
    status VARCHAR(20),
    payment_method VARCHAR(100),
    CONSTRAINT fk_user_transaction FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE transaction_product (
    id SERIAL PRIMARY KEY,
    product_id INT,
    transaction_id INT,
    quantity INT,
    size_id INT,
    variant_id INT,
    price INT,
    CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES products(id),
    CONSTRAINT fk_transaction FOREIGN KEY(transaction_id) REFERENCES transactions(id),
    CONSTRAINT fk_size FOREIGN KEY(size_id) REFERENCES product_size(id),
    CONSTRAINT fk_variant FOREIGN KEY(variant_id) REFERENCES product_variant(id)
);

CREATE TABLE delivery_methods (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    price BIGINT DEFAULT 0
);

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(80)
);

CREATE TABLE product_category (
    id SERIAL PRIMARY KEY,
    product_id INT,
    category_id INT,
    CONSTRAINT product FOREIGN KEY(product_id) REFERENCES products(id),
    CONSTRAINT category FOREIGN KEY(category_id) REFERENCES categories(id)
);

-- FILL DATA
-- 1. Products
INSERT INTO products(name, description, quantity, price, rating, old_price, is_flash_sale) VALUES
    ('Espresso', 'Kopi hitam pekat dengan crema tebal', 50, 18000, 5, 20000, false),
    ('Americano', 'Espresso dengan tambahan air panas', 45, 20000, 4, 25000, false),
    ('Cappuccino', 'Espresso dengan steamed milk dan foam lembut', 40, 25000, 5, 30000, true),
    ('Caffe Latte', 'Espresso dengan susu creamy', 42, 25000, 4, 0, false),
    ('Caramel Latte', 'Latte dengan sirup karamel manis', 35, 28000, 5, 35000, false),
    ('Vanilla Latte', 'Latte dengan aroma vanilla lembut', 30, 28000, 4, 0, false),
    ('Hazelnut Latte', 'Latte dengan rasa hazelnut gurih', 28, 28000, 4, 0, false),
    ('Mocha', 'Perpaduan espresso, cokelat, dan susu', 32, 30000, 5, 35000, true),
    ('Affogato', 'Espresso dengan scoop es krim vanilla', 20, 32000, 5, 0, false),
    ('Cold Brew', 'Kopi seduh dingin dengan rasa smooth', 25, 27000, 4, 30000, false),
    ('Matcha Latte', 'Teh hijau Jepang dengan susu creamy', 30, 28000, 5, 0, false),
    ('Chocolate', 'Minuman cokelat hangat premium', 35, 26000, 4, 0, false),
    ('Taro Latte', 'Minuman taro manis dan creamy', 25, 27000, 4, 0, false),
    ('Red Velvet Latte', 'Minuman red velvet lembut', 20, 29000, 4, 0, false),
    ('Thai Tea', 'Teh Thailand dengan susu kental manis', 30, 24000, 4, 0, false),
    ('Lemon Tea', 'Teh segar dengan perasan lemon', 40, 20000, 4, 0, false),
    ('Peach Tea', 'Teh dengan aroma dan rasa peach', 22, 23000, 4, 0, false),
    ('Mineral Water', 'Air mineral botol dingin', 60, 10000, 5, 0, false),
    ('Croissant Butter', 'Pastry renyah dengan aroma butter', 20, 22000, 5, 25000, false),
    ('Chocolate Croissant', 'Croissant isi cokelat lumer', 18, 25000, 5, 30000, true),
    ('Chicken Sandwich', 'Roti isi ayam dan sayur segar', 15, 30000, 4, 0, false),
    ('Beef Burger', 'Burger daging sapi dengan saus spesial', 12, 35000, 5, 45000, true),
    ('French Fries', 'Kentang goreng crispy', 25, 20000, 4, 0, false),
    ('Spaghetti Bolognese', 'Pasta dengan saus daging tomat', 14, 38000, 4, 0, false),
    ('Chicken Wrap', 'Tortilla isi ayam dan sayuran', 16, 32000, 4, 0, false),
    ('Cheese Cake Slice', 'Potongan cheesecake lembut', 10, 30000, 5, 0, false),
    ('Extra Shot Espresso', 'Tambahan satu shot espresso', 100, 8000, 5, 0, false),
    ('Syrup Caramel', 'Tambahan sirup karamel', 80, 5000, 4, 0, false),
    ('Syrup Vanilla', 'Tambahan sirup vanilla', 80, 5000, 4, 0, false),
    ('Whipped Cream', 'Tambahan whipped cream lembut', 70, 7000, 4, 0, false);

-- 2. Variants
INSERT INTO product_variant(product_id, name, add_price) VALUES
    (1, 'ice', 500), (1, 'extra shot', 8000),
    (2, 'ice', 1000), (2, 'less sugar', 0),
    (3, 'ice', 1000), (3, 'soy milk', 3000),
    (4, 'ice', 1000), (4, 'extra sugar', 1000),
    (5, 'ice', 1500), (5, 'soy milk', 3000),
    (8, 'ice', 1500), (8, 'extra topping', 5000),
    (11, 'ice', 1000), (11, 'soy milk', 3000),
    (15, 'less sugar', 0), (15, 'extra sugar', 1000),
    (22, 'extra cheese', 4000), (22, 'double meat', 8000),
    (24, 'extra sauce', 3000), (24, 'extra cheese', 4000);

-- 3. Sizes
INSERT INTO product_size(name, product_id, add_price) VALUES
    ('Regular', 1, 0), ('Medium', 1, 3000), ('Large', 1, 5000),
    ('Regular', 2, 0), ('Medium', 2, 4000), ('Large', 2, 6000),
    ('Regular', 3, 0), ('Medium', 3, 5000), ('Large', 3, 7000),
    ('Regular', 4, 0), ('Medium', 4, 5000), ('Large', 4, 7000),
    ('Regular', 5, 0), ('Medium', 5, 6000), ('Large', 5, 8000),
    ('Regular', 8, 0), ('Medium', 8, 6000), ('Large', 8, 9000),
    ('Regular', 11, 0), ('Medium', 11, 5000);

-- 4. Roles
INSERT INTO roles(name) VALUES ('admin'), ('user');

-- 5. Users | password for all user = 1234
INSERT INTO users(id, full_name, email, password, address, phone, picture, role_id) VALUES
    ('550e8400-e29b-41d4-a716-446655440001', 'Arif Rahman', 'arif.rahman@gmail.com', '$argon2id$v=19$m=65536,t=3,p=4$lJpFSh+Xs4r1BKw/AqJgMw$VVoW/s4D5IOki95/z+XmoRFSxnG/PtTAYuz4aPFwXXs', 'Jl. Kapuk No. 3 Tangsel', '081234560001', 'https://images.unsplash.com/photo-1500648767791-00dcc994a43e?q=80&w=200&h=200&auto=format&fit=crop', 1),
    ('550e8400-e29b-41d4-a716-446655440002', 'Dewi Lestari', 'dewi.lestari@gmail.com', '$argon2id$v=19$m=65536,t=3,p=4$lJpFSh+Xs4r1BKw/AqJgMw$VVoW/s4D5IOki95/z+XmoRFSxnG/PtTAYuz4aPFwXXs', 'Jl. Kenanga No. 8 Bandung', '081234560002', 'https://images.unsplash.com/photo-1494790108377-be9c29b29330?q=80&w=200&h=200&auto=format&fit=crop', 2),
    ('550e8400-e29b-41d4-a716-446655440003', 'Rizky Saputra', 'rizky.saputra@gmail.com', '$argon2id$v=19$m=65536,t=3,p=4$lJpFSh+Xs4r1BKw/AqJgMw$VVoW/s4D5IOki95/z+XmoRFSxnG/PtTAYuz4aPFwXXs', 'Jl. Mawar No. 21 Surabaya', '081234560003', 'https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?q=80&w=200&h=200&auto=format&fit=crop', 2),
    ('550e8400-e29b-41d4-a716-446655440004', 'Sinta Maharani', 'sinta.maharani@gmail.com', '$argon2id$v=19$m=65536,t=3,p=4$lJpFSh+Xs4r1BKw/AqJgMw$VVoW/s4D5IOki95/z+XmoRFSxnG/PtTAYuz4aPFwXXs', 'Jl. Anggrek No. 5 Yogyakarta', '081234560004', 'https://images.unsplash.com/photo-1438761681033-6461ffad8d80?q=80&w=200&h=200&auto=format&fit=crop', 2),
    ('550e8400-e29b-41d4-a716-446655440005', 'Fajar Nugroho', 'fajar.nugroho@gmail.com', '$argon2id$v=19$m=65536,t=3,p=4$lJpFSh+Xs4r1BKw/AqJgMw$VVoW/s4D5IOki95/z+XmoRFSxnG/PtTAYuz4aPFwXXs', 'Jl. Flamboyan No. 17 Semarang', '081234560005', 'https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?q=80&w=200&h=200&auto=format&fit=crop', 2),
    ('550e8400-e29b-41d4-a716-446655440006', 'Nadia Putri', 'nadia.putri@gmail.com', '$argon2id$v=19$m=65536,t=3,p=4$lJpFSh+Xs4r1BKw/AqJgMw$VVoW/s4D5IOki95/z+XmoRFSxnG/PtTAYuz4aPFwXXs', 'Jl. Dahlia No. 9 Medan', '081234560006', 'https://images.unsplash.com/photo-1544005313-94ddf0286df2?q=80&w=200&h=200&auto=format&fit=crop', 2),
    ('550e8400-e29b-41d4-a716-446655440007', 'Bima Kurniawan', 'bima.kurniawan@gmail.com', '$argon2id$v=19$m=65536,t=3,p=4$lJpFSh+Xs4r1BKw/AqJgMw$VVoW/s4D5IOki95/z+XmoRFSxnG/PtTAYuz4aPFwXXs', 'Jl. Cempaka No. 14 Bekasi', '081234560007', 'https://images.unsplash.com/photo-1506794778202-cad84cf45f1d?q=80&w=200&h=200&auto=format&fit=crop', 2),
    ('550e8400-e29b-41d4-a716-446655440008', 'Laras Wulandari', 'laras.wulandari@gmail.com', '$argon2id$v=19$m=65536,t=3,p=4$lJpFSh+Xs4r1BKw/AqJgMw$VVoW/s4D5IOki95/z+XmoRFSxnG/PtTAYuz4aPFwXXs', 'Jl. Teratai No. 3 Depok', '081234560008', 'https://images.unsplash.com/photo-1517841905240-472988babdf9?q=80&w=200&h=200&auto=format&fit=crop', 2),
    ('550e8400-e29b-41d4-a716-446655440009', 'Andika Ramadhan', 'andika.ramadhan@gmail.com', '$argon2id$v=19$m=65536,t=3,p=4$lJpFSh+Xs4r1BKw/AqJgMw$VVoW/s4D5IOki95/z+XmoRFSxnG/PtTAYuz4aPFwXXs', 'Jl. Bougenville No. 11 Tangerang', '081234560009', 'https://images.unsplash.com/photo-1539571696357-5a69c17a67c6?q=80&w=200&h=200&auto=format&fit=crop', 2),
    ('550e8400-e29b-41d4-a716-446655440010', 'Maya Oktaviani', 'maya.oktaviani@gmail.com', '$argon2id$v=19$m=65536,t=3,p=4$lJpFSh+Xs4r1BKw/AqJgMw$VVoW/s4D5IOki95/z+XmoRFSxnG/PtTAYuz4aPFwXXs', 'Jl. Sakura No. 6 Bogor', '081234560010', 'https://images.unsplash.com/photo-1488426862026-3ee34a7d66df?q=80&w=200&h=200&auto=format&fit=crop', 2);

-- 6. Delivery Methods
INSERT INTO delivery_methods (name, price) VALUES 
('Dine in', 0),
('Door Delivery', 10000),
('Pick up', 0);

-- 7. Categories
INSERT INTO categories(name) VALUES ('Favourite Product'), ('Coffee'), ('Non-Coffee'), ('Foods'), ('Add-On');

-- 7. Product Category mapping
INSERT INTO product_category(product_id, category_id) VALUES
    (1, 2), (2, 2), (3, 2), (3, 1), (4, 2), (5, 2), (6, 2), (6, 1), (7, 2), (8, 2), (9, 2), (10, 2),
    (11, 3), (12, 3), (13, 3), (14, 3), (15, 3), (16, 3), (17, 3), (18, 3),
    (19, 4), (20, 4), (21, 4), (22, 4), (23, 4), (24, 4), (25, 4), (26, 4),
    (27, 5), (28, 5), (29, 5), (30, 5);

-- 8. Reviews (pointing to actual users)
INSERT INTO reviews(user_id, messages, rating) VALUES
    ('550e8400-e29b-41d4-a716-446655440010', 'Kopinya enak dan aromanya kuat, pasti order lagi.', 5),
    ('550e8400-e29b-41d4-a716-446655440009', 'Tempatnya nyaman tapi minumannya agak lama datangnya.', 4),
    ('550e8400-e29b-41d4-a716-446655440008', 'Rasa matchanya kurang terasa, mungkin bisa ditingkatkan.', 3),
    ('550e8400-e29b-41d4-a716-446655440007', 'Pelayanannya ramah dan cepat, sangat puas.', 5),
    ('550e8400-e29b-41d4-a716-446655440006', 'Harga sedikit mahal tapi kualitas sesuai.', 4),
    ('550e8400-e29b-41d4-a716-446655440005', 'Kopi terlalu pahit untuk selera saya.', 2),
    ('550e8400-e29b-41d4-a716-446655440004', 'Makanannya enak tapi minuman kurang dingin.', 3),
    ('550e8400-e29b-41d4-a716-446655440003', 'Pelayanan kurang responsif saat ramai.', 2),
    ('550e8400-e29b-41d4-a716-446655440002', 'Suka banget sama caramel lattenya, recommended!', 5),
    ('550e8400-e29b-41d4-a716-446655440001', 'Overall oke, tapi tempat parkir agak sempit.', 4);

-- 9. Product Images
INSERT INTO product_images(product_id, path) VALUES
    (1, 'https://images.unsplash.com/photo-1510707577719-ae7c14805e3a?q=80&w=600&h=600&auto=format&fit=crop'),
    (1, 'https://images.unsplash.com/photo-1579992357154-faf4bde95b3d?q=80&w=600&h=600&auto=format&fit=crop'),
    (1, 'https://images.unsplash.com/photo-1565538183181-79282740a618?q=80&w=600&h=600&auto=format&fit=crop'),
    (2, 'https://images.unsplash.com/photo-1551033406-611cf9a28f67?q=80&w=600&h=600&auto=format&fit=crop'),
    (2, 'https://images.unsplash.com/photo-1557006021-b85faa2bc5e2?q=80&w=600&h=600&auto=format&fit=crop'),
    (2, 'https://images.unsplash.com/photo-1580915411954-282cb1b0d780?q=80&w=600&h=600&auto=format&fit=crop'),
    (3, 'https://images.unsplash.com/photo-1572442388796-11668a67e53d?q=80&w=600&h=600&auto=format&fit=crop'),
    (3, 'https://images.unsplash.com/photo-1534778101976-62847782c213?q=80&w=600&h=600&auto=format&fit=crop'),
    (3, 'https://images.unsplash.com/photo-1551033406-611cf9a28f67?q=80&w=600&h=600&auto=format&fit=crop'),
    (4, 'https://images.unsplash.com/photo-1461023058943-07fcbe16d735?q=80&w=600&h=600&auto=format&fit=crop'),
    (4, 'https://images.unsplash.com/photo-1541167760496-162955ed8a9f?q=80&w=600&h=600&auto=format&fit=crop'),
    (4, 'https://images.unsplash.com/photo-1459755486867-b55449bb39ff?q=80&w=600&h=600&auto=format&fit=crop'),
    (5, 'https://images.unsplash.com/photo-1495474472287-4d71bcdd2085?q=80&w=600&h=600&auto=format&fit=crop'),
    (5, 'https://images.unsplash.com/photo-1572286258217-40142c1c6a70?q=80&w=600&h=600&auto=format&fit=crop'),
    (5, 'https://images.unsplash.com/photo-1599398054066-846f28917f38?q=80&w=600&h=600&auto=format&fit=crop'),
    (6, 'https://images.unsplash.com/photo-1514432324607-a09d9b4aefdd?q=80&w=600&h=600&auto=format&fit=crop'),
    (6, 'https://images.unsplash.com/photo-1544145945-f904253d0c71?q=80&w=600&h=600&auto=format&fit=crop'),
    (6, 'https://images.unsplash.com/photo-1570968865863-dc4ac048be6b?q=80&w=600&h=600&auto=format&fit=crop'),
    (7, 'https://images.unsplash.com/photo-1494314671902-399b18174975?q=80&w=600&h=600&auto=format&fit=crop'),
    (7, 'https://images.unsplash.com/photo-1512568400610-62da28bc8a13?q=80&w=600&h=600&auto=format&fit=crop'),
    (7, 'https://images.unsplash.com/photo-1497633762265-9d179a990aa6?q=80&w=600&h=600&auto=format&fit=crop'),
    (8, 'https://images.unsplash.com/photo-1559496417-e7f25cb247f3?q=80&w=600&h=600&auto=format&fit=crop'),
    (8, 'https://images.unsplash.com/photo-1517701550927-30cf4ba1dba5?q=80&w=600&h=600&auto=format&fit=crop'),
    (8, 'https://images.unsplash.com/photo-1553909489-cd47e0907d3f?q=80&w=600&h=600&auto=format&fit=crop'),
    (9, 'https://images.unsplash.com/photo-1511920170033-f8396924c348?q=80&w=600&h=600&auto=format&fit=crop'),
    (9, 'https://images.unsplash.com/photo-1594631252845-29fc4586c552?q=80&w=600&h=600&auto=format&fit=crop'),
    (9, 'https://images.unsplash.com/photo-1447078806655-40579c2520d6?q=80&w=600&h=600&auto=format&fit=crop'),
    (10, 'https://images.unsplash.com/photo-1512568400610-62da28bc8a13?q=80&w=600&h=600&auto=format&fit=crop'),
    (10, 'https://images.unsplash.com/photo-1461023058943-07fcbe16d735?q=80&w=600&h=600&auto=format&fit=crop'),
    (10, 'https://images.unsplash.com/photo-1495474472287-4d71bcdd2085?q=80&w=600&h=600&auto=format&fit=crop'),
    (11, 'https://images.unsplash.com/photo-1515824918246-a8a556f3c27f?q=80&w=600&h=600&auto=format&fit=crop'),
    (11, 'https://images.unsplash.com/photo-1536256263959-770b48d82b0a?q=80&w=600&h=600&auto=format&fit=crop'),
    (11, 'https://images.unsplash.com/photo-1596791011531-10c0e5a60e6e?q=80&w=600&h=600&auto=format&fit=crop'),
    (12, 'https://images.unsplash.com/photo-1445116572660-236099ec97a0?q=80&w=600&h=600&auto=format&fit=crop'),
    (12, 'https://images.unsplash.com/photo-1544787219-7f47ccb76574?q=80&w=600&h=600&auto=format&fit=crop'),
    (12, 'https://images.unsplash.com/photo-1511381939415-e44015466834?q=80&w=600&h=600&auto=format&fit=crop'),
    (13, 'https://images.unsplash.com/photo-1534706936160-d5ee67737249?q=80&w=600&h=600&auto=format&fit=crop'),
    (13, 'https://images.unsplash.com/photo-1627834377411-8da5f4f09de8?q=80&w=600&h=600&auto=format&fit=crop'),
    (13, 'https://images.unsplash.com/photo-1579888945649-2f1585aaabbb?q=80&w=600&h=600&auto=format&fit=crop'),
    (14, 'https://images.unsplash.com/photo-1578314675249-a6910f80cc4e?q=80&w=600&h=600&auto=format&fit=crop'),
    (14, 'https://images.unsplash.com/photo-1579306194872-64d3b7bac4c2?q=80&w=600&h=600&auto=format&fit=crop'),
    (14, 'https://images.unsplash.com/photo-1576618148400-f54bed99fcfd?q=80&w=600&h=600&auto=format&fit=crop'),
    (15, 'https://images.unsplash.com/photo-1558160074-4d7d8bdf4256?q=80&w=600&h=600&auto=format&fit=crop'),
    (15, 'https://images.unsplash.com/photo-1594631252845-29fc4586c552?q=80&w=600&h=600&auto=format&fit=crop'),
    (15, 'https://images.unsplash.com/photo-1571328003758-4a3921661709?q=80&w=600&h=600&auto=format&fit=crop'),
    (16, 'https://images.unsplash.com/photo-1513558161293-cdaf765ed2fd?q=80&w=600&h=600&auto=format&fit=crop'),
    (16, 'https://images.unsplash.com/photo-1556679343-c7306c1976bc?q=80&w=600&h=600&auto=format&fit=crop'),
    (16, 'https://images.unsplash.com/photo-1543508282-5c1f427f023f?q=80&w=600&h=600&auto=format&fit=crop'),
    (17, 'https://images.unsplash.com/photo-1525193612562-0ec53b0e5d7c?q=80&w=600&h=600&auto=format&fit=crop'),
    (17, 'https://images.unsplash.com/photo-1597318181409-cf64d0b5d8a2?q=80&w=600&h=600&auto=format&fit=crop'),
    (17, 'https://images.unsplash.com/photo-1556679343-c7306c1976bc?q=80&w=600&h=600&auto=format&fit=crop'),
    (18, 'https://images.unsplash.com/photo-1523362628742-0c2673ee5202?q=80&w=600&h=600&auto=format&fit=crop'),
    (18, 'https://images.unsplash.com/photo-1550507992-eb63ffee0847?q=80&w=600&h=600&auto=format&fit=crop'),
    (18, 'https://images.unsplash.com/photo-1559839914-17aae19cea9e?q=80&w=600&h=600&auto=format&fit=crop'),
    (19, 'https://images.unsplash.com/photo-1555507036-ab1f4038808a?q=80&w=600&h=600&auto=format&fit=crop'),
    (19, 'https://images.unsplash.com/photo-1549437254-4796332a6859?q=80&w=600&h=600&auto=format&fit=crop'),
    (19, 'https://images.unsplash.com/photo-1530610476181-d83430b64dcd?q=80&w=600&h=600&auto=format&fit=crop'),
    (20, 'https://images.unsplash.com/photo-1509440159596-0249088772ff?q=80&w=600&h=600&auto=format&fit=crop'),
    (20, 'https://images.unsplash.com/photo-1530610476181-d83430b64dcd?q=80&w=600&h=600&auto=format&fit=crop'),
    (20, 'https://images.unsplash.com/photo-1551024601-bec78aea704b?q=80&w=600&h=600&auto=format&fit=crop'),
    (21, 'https://images.unsplash.com/photo-1509440159596-0249088772ff?q=80&w=600&h=600&auto=format&fit=crop'),
    (21, 'https://images.unsplash.com/photo-1528735602780-2552fd46c7af?q=80&w=600&h=600&auto=format&fit=crop'),
    (21, 'https://images.unsplash.com/photo-1550507992-eb63ffee0847?q=80&w=600&h=600&auto=format&fit=crop'),
    (22, 'https://images.unsplash.com/photo-1568901346375-23c9450c58cd?q=80&w=600&h=600&auto=format&fit=crop'),
    (22, 'https://images.unsplash.com/photo-1550547660-d9450f859349?q=80&w=600&h=600&auto=format&fit=crop'),
    (22, 'https://images.unsplash.com/photo-1571091718767-18b5b1457add?q=80&w=600&h=600&auto=format&fit=crop'),
    (23, 'https://images.unsplash.com/photo-1573088693243-9874a625e3bb?q=80&w=600&h=600&auto=format&fit=crop'),
    (23, 'https://images.unsplash.com/photo-1518013431117-eb1465fd5752?q=80&w=600&h=600&auto=format&fit=crop'),
    (23, 'https://images.unsplash.com/photo-1630384060421-cb20d0e0649d?q=80&w=600&h=600&auto=format&fit=crop'),
    (24, 'https://images.unsplash.com/photo-1551183053-bf91a1d81141?q=80&w=600&h=600&auto=format&fit=crop'),
    (24, 'https://images.unsplash.com/photo-1572441713132-c542fc4fe282?q=80&w=600&h=600&auto=format&fit=crop'),
    (24, 'https://images.unsplash.com/photo-1622973536968-3ead9e780960?q=80&w=600&h=600&auto=format&fit=crop'),
    (25, 'https://images.unsplash.com/photo-1513104890138-7c749659a591?q=80&w=600&h=600&auto=format&fit=crop'),
    (25, 'https://images.unsplash.com/photo-1626700051175-6818013e1d4f?q=80&w=600&h=600&auto=format&fit=crop'),
    (25, 'https://images.unsplash.com/photo-1504674900247-0877df9cc836?q=80&w=600&h=600&auto=format&fit=crop'),
    (26, 'https://images.unsplash.com/photo-1551024601-bec78aea704b?q=80&w=600&h=600&auto=format&fit=crop'),
    (26, 'https://images.unsplash.com/photo-1533134242443-d4fd215305ad?q=80&w=600&h=600&auto=format&fit=crop'),
    (26, 'https://images.unsplash.com/photo-1508737804141-4c3b688e2546?q=80&w=600&h=600&auto=format&fit=crop'),
    (27, 'https://images.unsplash.com/photo-1510707577719-ae7c14805e3a?q=80&w=600&h=600&auto=format&fit=crop'),
    (27, 'https://images.unsplash.com/photo-1579992357154-faf4bde95b3d?q=80&w=600&h=600&auto=format&fit=crop'),
    (27, 'https://images.unsplash.com/photo-1565538183181-79282740a618?q=80&w=600&h=600&auto=format&fit=crop'),
    (28, 'https://images.unsplash.com/photo-1551024506-0bccd828d307?q=80&w=600&h=600&auto=format&fit=crop'),
    (28, 'https://images.unsplash.com/photo-1582845241727-46487e411b2b?q=80&w=600&h=600&auto=format&fit=crop'),
    (28, 'https://images.unsplash.com/photo-1555529731-118a5bb67af7?q=80&w=600&h=600&auto=format&fit=crop'),
    (29, 'https://images.unsplash.com/photo-1551024506-0bccd828d307?q=80&w=600&h=600&auto=format&fit=crop'),
    (29, 'https://images.unsplash.com/photo-1506368249639-73a05d6f6488?q=80&w=600&h=600&auto=format&fit=crop'),
    (29, 'https://images.unsplash.com/photo-1621236304198-6515855885ba?q=80&w=600&h=600&auto=format&fit=crop'),
    (30, 'https://images.unsplash.com/photo-1551024506-0bccd828d307?q=80&w=600&h=600&auto=format&fit=crop'),
    (30, 'https://images.unsplash.com/photo-1553909489-cd47e0907d3f?q=80&w=600&h=600&auto=format&fit=crop'),
    (30, 'https://images.unsplash.com/photo-1610450954843-05f42df22204?q=80&w=600&h=600&auto=format&fit=crop');