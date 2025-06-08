package interactor

import (
	"context"

	"github.com/davyc98/billing-engine/internal/billing/internal/entity/sqlentity"
	"github.com/davyc98/billing-engine/internal/billing/internal/usecase"
	"go.uber.org/zap"
)

type (
	IsDelinquent struct {
		store  GetAndUpdateLoanStore
		logger *zap.SugaredLogger
	}
)

func NewIsDelinquent(
	store GetAndUpdateLoanStore,
	logger *zap.SugaredLogger,
) *IsDelinquent {
	return &IsDelinquent{
		store:  store,
		logger: logger,
	}
}

func (c *IsDelinquent) Execute(
	ctx context.Context,
	in usecase.IsDelinquentInput,
) (*usecase.IsDelinquentOutput, error) {
	var loanSchedules sqlentity.LoanSchedules
	var err error
	if loanSchedules, err = c.store.GetLoanSchedule(ctx, in.LoanID); err != nil {
		c.logger.Errorw("failed to get loan schedule", "error", err)

		return nil, err
	}

	if len(loanSchedules) < 2 {
		return &usecase.IsDelinquentOutput{
			IsDelinquent: false,
		}, nil
	}

	// Iterate through the sorted overdue weeks to find consecutive ones
	for i := range len(loanSchedules) - 1 {
		if loanSchedules[i+1].WeekNumber == loanSchedules[i].WeekNumber+1 {
			return &usecase.IsDelinquentOutput{
				IsDelinquent: true,
			}, nil
		}
	}

	return &usecase.IsDelinquentOutput{
		IsDelinquent: false,
	}, nil
}
