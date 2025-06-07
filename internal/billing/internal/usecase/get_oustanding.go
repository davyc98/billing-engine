package usecase

import (
	"context"

	"github.com/shopspring/decimal"
)

type (
	GetOutstanding interface {
		Execute(ctx context.Context, in GetOutstandingInput) (*GetOustandingOutput, error)
	}

	GetOutstandingInput struct {
		LoanID     uint64 `json:"loan_id"     validate:"required"`
		CustomerID uint64 `json:"customer_id"`
	}

	GetOustandingOutput struct {
		ID              uint64
		LoanID          uint64
		WeekNumber      int
		DueDate         string
		ScheduledAmount decimal.Decimal
		PaidAmount      decimal.Decimal
		PaymentStatus   string
	}
)
