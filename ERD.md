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