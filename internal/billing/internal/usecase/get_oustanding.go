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
		LoanID uint64 `json:"loan_id"     validate:"required"`
	}

	GetOustandingOutput struct {
		LoanID                 uint64          `json:"loan_id"`
		TotalOutstandingAmount decimal.Decimal `json:"total_outstanding_amount"`
		TotalOustandingWeeks   int             `json:"total_outstanding_week"`
	}
)
