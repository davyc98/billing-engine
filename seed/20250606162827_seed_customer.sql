-- +goose Up
-- Seed data for the CUSTOMER table (MySQL)

INSERT INTO customers (id, name, email, phone) VALUES
(1, 'Alice Smith', 'alice.smith@example.com', '123-456-7890'),
(2, 'Bob Johnson', 'bob.j@example.com', '987-654-3210'),
(3, 'Charlie Brown', 'charlie.b@example.com', '555-123-4567');


-- +goose Down
-- SQL to revert the customer seed data.
-- Deletes the three customer records inserted by the 'Up' migration.

DELETE FROM customers WHERE id IN (1, 2, 3);