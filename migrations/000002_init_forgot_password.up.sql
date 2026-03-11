CREATE TABLE IF NOT EXISTS forgot_password (
    id SERIAL PRIMARY KEY,
    email VARCHAR(100),
    code INT
);