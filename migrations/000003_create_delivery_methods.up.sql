CREATE TABLE IF NOT EXISTS delivery_methods (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    price BIGINT DEFAULT 0
);

INSERT INTO delivery_methods (name, price) VALUES 
('Dine in', 0),
('Door Delivery', 10000),
('Pick up', 0);
