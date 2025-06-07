## Setup

- Install goose -> go install github.com/pressly/goose/v3/cmd/goose@latest
- Copy .env.example to .env
- Fill in the .env file

## Migrations

- goose -dir migration/ mysql "user:password@tcp(localhost:3306)/loan_db?parseTime=true" up

## Seed

- goose -no-versioning -dir seed/ mysql "user:password@tcp(localhost:3306)/loan_db?parseTime=true" up

