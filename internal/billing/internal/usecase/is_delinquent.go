package usecase

import "context"

type (
	IsDelinquent interface {
		Execute(ctx context.Context, in IsDelinquentInput) (*IsDelinquentOutput, error)
	}

	IsDelinquentInput struct {
		LoanID uint64 `json:"loan_id"`
	}
	IsDelinquentOutput struct {
		IsDelinquent bool `json:"is_delinquent"`
	}
)
