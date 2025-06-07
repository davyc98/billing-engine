-- +goose Up
-- SQL to create the customer table for MySQL.

CREATE TABLE customers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    email       VARCHAR(255) UNIQUE NOT NULL, -- Email should be unique
    phone       VARCHAR(50),
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Optional: Add an index on email for faster lookups if frequently queried by email
CREATE INDEX idx_customer_email ON customers (email);

-- +goose Down
-- SQL to drop the customer table, reverting the migration.

DROP TABLE IF EXISTS customers;
