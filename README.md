## Setup

- Install goose -> go install github.com/pressly/goose/v3/cmd/goose@latest
- Copy .env.example to .env
- Fill in the .env file
- Create Database with credential same like env, using mysql
- Run migration files below
- Use postman collection to test out endpoints

## Migrations

- goose -dir migration/ mysql "user:password@tcp(localhost:3306)/loan_db?parseTime=true" up

## Seed

- goose -no-versioning -dir seed/ mysql "user:password@tcp(localhost:3306)/loan_db?parseTime=true" up

## ERD

+-------------------+       +-------------------+       +-----------------------+
|     CUSTOMER      |       |      LOAN         |       |    LOAN_SCHEDULE      |
+-------------------+       +-------------------+       +-----------------------+
| * customer_id (PK)|-------| * loan_id (PK)    |-------| * schedule_id (PK)    |
|   name            |       |   customer_id (FK)|       |   loan_id (FK)        |
|   email           |       |   loan_amount     |       |   week_number         |
|   phone           |       |   interest_rate   |       |   due_date            |
|   address         |       |   loan_term_weeks |       |   scheduled_amount    |
|   ...             |       |   start_date      |       |   paid_amount         |
+-------------------+       |   end_date        |       |   payment_status      |
                            |   total_payable   |       |   payment_date        |
                            |   weekly_payment  |       +-----------------------+
                            |   current_outstanding |
                            |   status          |
                            +-------------------+