package usecase

import (
	"context"

	"github.com/shopspring/decimal"
)

type (
	MakePayment interface {
		Execute(ctx context.Context, in PaymentInput) error
	}

	PaymentInput struct {
		LoanID        uint64 `json:"loan_id"     validate:"required"`
		PaymentAmount string `json:"payment_amount"  validate:"required"`
	}

	PaymentOutput struct {
		ID              uint64          `json:"id"`
		LoanID          uint64          `json:"loan_id"`
		WeekNumber      int             `json:"week_number"`
		DueDate         string          `json:"due_date"`
		ScheduledAmount decimal.Decimal `json:"scheduled_amount"`
		PaidAmount      decimal.Decimal `json:"paid_amount"`
		PaymentStatus   string          `json:"payment_status"`
	}
)
