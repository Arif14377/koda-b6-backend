CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(200) NOT NULL,
    quantity INT,
    price BIGINT
);

CREATE TABLE IF NOT EXISTS product_variant (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    product_id INT NOT NULL,
    add_price BIGINT,
    CONSTRAINT fk_product
        FOREIGN KEY(product_id) REFERENCES products(id)
);

CREATE TABLE IF NOT EXISTS product_size (
    id SERIAL PRIMARY KEY,
    name VARCHAR(20) NOT NULL,
    product_id INT,
    add_price BIGINT,
    constraint fk_product
        foreign key(product_id)
        references products(id)
);

CREATE TABLE IF NOT EXISTS product_images (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    path VARCHAR(500),
    CONSTRAINT fk_product
        FOREIGN KEY(product_id)
        REFERENCES products(id)
);

CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(10) CHECK (name IN('admin', 'user'))
);

CREATE TABLE IF NOT EXISTS users (
    ID VARCHAR(36) PRIMARY KEY,
    full_name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100),
    address VARCHAR(200),
    phone VARCHAR(20),
    picture VARCHAR(500),
    role_id INT,
    created_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_roles
        FOREIGN KEY (role_id) REFERENCES roles(id)
);

CREATE TABLE IF NOT EXISTS cart (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(36),
    product_id int,
    constraint fk_user
        foreign key(user_id)
        references users(id),
    constraint fk_products
        foreign key(product_id)
        references products(id)
);


CREATE TABLE IF NOT EXISTS reviews (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(36),
    messages VARCHAR,
    rating INT CHECK (rating BETWEEN 1 AND 5),
    CONSTRAINT users
        foreign key(user_id)
        references users(id)
);

CREATE TABLE IF NOT EXISTS transactions (
    id BIGSERIAL PRIMARY KEY,
    trx_code VARCHAR(50) UNIQUE NOT NULL,
    delivery_method VARCHAR(60) NOT NULL,
    full_name VARCHAR(100),
    email VARCHAR(100),
    address VARCHAR(200),
    sub_total BIGINT,
    tax BIGINT,
    total BIGINT,
    date TIMESTAMP DEFAULT now(),
    status VARCHAR(20),
    payment_method VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS transaction_product (
    id SERIAL PRIMARY KEY,
    product_id INT,
    transaction_id INT,
    quantity INT,
    CONSTRAINT fk_product
        FOREIGN KEY(product_id)
        REFERENCES products(id),
    CONSTRAINT fk_transaction
        FOREIGN KEY(transaction_id)
        REFERENCES transactions(id)
);

CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(80)
);

CREATE TABLE IF NOT EXISTS roduct_category (
    id SERIAL PRIMARY KEY,
    product_id INT,
    category_id INT,
    CONSTRAINT product
        FOREIGN KEY(product_id)
        REFERENCES products(id),
    CONSTRAINT category
        FOREIGN KEY(category_id)
        REFERENCES categories(id)
);