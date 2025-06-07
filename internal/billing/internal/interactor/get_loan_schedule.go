package interactor

import (
	"context"
	"database/sql"
	"errors"

	"github.com/davyc98/billing-engine/internal/billing/internal/entity/sqlentity"
	"github.com/davyc98/billing-engine/internal/billing/internal/gateway"
	"github.com/davyc98/billing-engine/internal/billing/internal/usecase"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

type (
	GetAndUpdateLoanStore interface {
		GetLoanSchedule(ctx context.Context, loanID uint64) (sqlentity.LoanSchedules, error)
		UpdateLoanSchedule(ctx context.Context, in sqlentity.UpdateEntity, opts ...gateway.UpdateLoanScheduleOption) error
		UpdateLoan(ctx context.Context, in sqlentity.UpdateEntity, opts ...gateway.UpdateLoanOption) error
		GetLoan(ctx context.Context, loanID uint64) (sqlentity.Loans, error)
	}

	GetOutstandingLoan struct {
		store  GetAndUpdateLoanStore
		logger *zap.SugaredLogger
	}
)

func NewGetOustandingLoan(
	store GetAndUpdateLoanStore,
	logger *zap.SugaredLogger,
) *GetOutstandingLoan {
	return &GetOutstandingLoan{
		store:  store,
		logger: logger,
	}
}

func (c *GetOutstandingLoan) Execute(
	ctx context.Context,
	in usecase.GetOutstandingInput,
) (*usecase.GetOustandingOutput, error) {
	var loanSchedule sqlentity.LoanSchedules
	var err error
	if loanSchedule, err = c.store.GetLoanSchedule(ctx, in.LoanID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &usecase.GetOustandingOutput{
				LoanID:                 in.LoanID,
				TotalOutstandingAmount: decimal.Zero,
			}, nil
		}
		c.logger.Errorw("failed to get loan schedule", "error", err)

		return nil, err
	}
	var total decimal.Decimal
	var totalOutstandingWeek int
	for _, v := range loanSchedule {
		total = v.ScheduledAmount.Add(total)
		totalOutstandingWeek++
	}

	return &usecase.GetOustandingOutput{
		LoanID:                 in.LoanID,
		TotalOutstandingAmount: total,
		TotalOustandingWeeks:   totalOutstandingWeek,
	}, nil
}
