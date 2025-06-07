-- +goose Up
-- Seed data for the LOAN table (MySQL)
--
-- This migration inserts one new loan record with the start_date set to the current date (NOW()).
-- The loan_id is 101, associated with customer_id 1 (Alice Smith from previous seed data).
-- Loan details: Rp 5,000,000 loan, 10% flat interest, 50 weeks term.
-- Total Payable: Rp 5,500,000. Weekly Payment: Rp 110,000.

INSERT INTO loan (
    loan_id,
    customer_id,
    loan_amount,
    interest_rate,
    loan_term_weeks,
    start_date,
    end_date,
    total_payable_amount,
    weekly_payment_amount,
    current_outstanding_balance,
    status
) VALUES (
    101, -- New loan ID
    1,   -- Assuming customer_id 1 exists (Alice Smith)
    5000000.00,
    0.10,
    50,
    CURDATE(), -- Sets the start_date to the current date
    DATE_ADD(CURDATE(), INTERVAL 50 WEEK), -- Calculates end_date 50 weeks from now
    5500000.00,
    110000.00,
    5500000.00, -- Initial outstanding, will be updated by payments
    'ACTIVE'
);


-- +goose Down
-- SQL to revert the loan seed data.
-- Deletes the loan record inserted by the 'Up' migration.

DELETE FROM loan WHERE loan_id = 101;