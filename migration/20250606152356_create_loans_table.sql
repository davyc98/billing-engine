-- +goose Up
-- SQL to create the loan_schedule table for MySQL.
CREATE TABLE loan (
    id                     INT AUTO_INCREMENT PRIMARY KEY, -- Use INT AUTO_INCREMENT PRIMARY KEY (MySQL)
    customer_id                 VARCHAR(255) NOT NULL,
    loan_amount                 NUMERIC(18, 2) NOT NULL,
    interest_rate               NUMERIC(5, 4) NOT NULL, -- e.g., 0.1000 for 10%
    loan_term_weeks             INT NOT NULL,
    start_date                  DATE NOT NULL,
    end_date                    DATE NOT NULL,
    total_payable_amount        NUMERIC(18, 2) NOT NULL,
    weekly_payment_amount       NUMERIC(18, 2) NOT NULL,
    current_outstanding_balance NUMERIC(18, 2) NOT NULL,
    status                      VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL
);


-- +goose Down
-- SQL to drop the LOAN table, reverting the migration.

DROP TABLE IF EXISTS loan;