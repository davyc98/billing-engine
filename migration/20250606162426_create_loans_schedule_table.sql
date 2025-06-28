-- +goose Up
CREATE TABLE loan_schedule (
    id         INT AUTO_INCREMENT PRIMARY KEY, -- Use INT AUTO_INCREMENT PRIMARY KEY (MySQL)
    loan_id             INT NOT NULL,
    week_number         INT NOT NULL,
    due_date            DATE NOT NULL,
    scheduled_amount    NUMERIC(18, 2) NOT NULL,
    paid_amount         NUMERIC(18, 2) DEFAULT 0.00 NOT NULL,
    payment_status      VARCHAR(50) NOT NULL,
    payment_date        DATE, -- NULLable, as payment might not have occurred yet

    -- Foreign Key constraint
    CONSTRAINT fk_loan_id
        FOREIGN KEY (loan_id)
        REFERENCES loan (id)
        ON DELETE CASCADE -- If a loan is deleted, its schedule entries are also deleted.
);

CREATE INDEX idx_loan_schedule_loan_id ON loan_schedule (loan_id);

-- Optional: Add a unique constraint for loan_id and week_number to prevent duplicate schedule entries
-- for the same week within a loan.
CREATE UNIQUE INDEX uix_loan_schedule_loan_week ON loan_schedule (loan_id, week_number);

-- +goose Down
-- SQL to drop the loan_schedule table, reverting the migration.

DROP TABLE IF EXISTS loan_schedule;