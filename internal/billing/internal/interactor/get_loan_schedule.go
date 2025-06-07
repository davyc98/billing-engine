package interactor

import (
	"context"

	"github.com/davyc98/billing-engine/internal/billing/internal/entity/sqlentity"
	"github.com/davyc98/billing-engine/internal/billing/internal/usecase"
	"go.uber.org/zap"
)

type (
	GetAndUpdateLoanStore interface {
		GetLoanSchedule(ctx context.Context, loanID uint64) (*sqlentity.LoanSchedule, error)
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
	var loanSchedule *sqlentity.LoanSchedule
	var err error
	if loanSchedule, err = c.store.GetLoanSchedule(ctx, in.LoanID); err != nil {
		c.logger.Errorw("failed to insert loan", "error", err)

		return nil, err
	}

	return &usecase.GetOustandingOutput{
		ID:              loanSchedule.ID,
		LoanID:          loanSchedule.LoanID,
		WeekNumber:      loanSchedule.WeekNumber,
		DueDate:         loanSchedule.DueDate,
		ScheduledAmount: loanSchedule.ScheduledAmount,
		PaidAmount:      loanSchedule.PaidAmount,
		PaymentStatus:   loanSchedule.PaymentStatus.String(),
	}, nil
}
