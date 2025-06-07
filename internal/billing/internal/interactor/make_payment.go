package interactor

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/davyc98/billing-engine/internal/billing/internal/entity/sqlentity"
	"github.com/davyc98/billing-engine/internal/billing/internal/gateway"
	"github.com/davyc98/billing-engine/internal/billing/internal/usecase"
	"github.com/davyc98/billing-engine/internal/pkg/pkgerror"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

type (
	MakePayment struct {
		store  GetAndUpdateLoanStore
		logger *zap.SugaredLogger
	}
)

func NewMakePayment(
	store GetAndUpdateLoanStore,
	logger *zap.SugaredLogger,
) *MakePayment {
	return &MakePayment{
		store:  store,
		logger: logger,
	}
}

func (c *MakePayment) Execute(
	ctx context.Context,
	in usecase.PaymentInput,
) error {
	// Handle missed payment, and two payment in row
	var loanSchedules sqlentity.LoanSchedules
	var err error
	if loanSchedules, err = c.store.GetLoanSchedule(ctx, in.LoanID); err != nil {
		c.logger.Errorw("failed to get loan schedule", "error", err)

		return err
	}

	paymentAmount, err := decimal.NewFromString(in.PaymentAmount)
	if err != nil {
		c.logger.Errorw("failed to parse payment amount", "error", err)
		return pkgerror.ServerErrorFrom(err)
	}
	totalPayment := paymentAmount
	firstLoanS := loanSchedules.First()
	if !paymentAmount.Equal(firstLoanS.ScheduledAmount) && paymentAmount.LessThan(firstLoanS.ScheduledAmount) {
		err := errors.New("invalid amount")
		c.logger.Errorw("invalid amount", "error", err)
		return err
	}
	for _, v := range loanSchedules {
		if paymentAmount.LessThan(v.ScheduledAmount) {
			break
		}

		if err := c.store.UpdateLoanSchedule(ctx, &sqlentity.UpdateLoanSchedule{
			PaidAmount:    v.ScheduledAmount,
			PaymentStatus: sqlentity.Paid,
			PaymentDate: sql.NullString{
				String: time.Now().Format("2006-01-02"),
				Valid:  true,
			},
		}, gateway.UpdateLoanScheduleWithLoanIDAndIDFilter(v.ID, in.LoanID)); err != nil {
			c.logger.Errorw("failed to update loan", "error", err)

			return pkgerror.ServerErrorFrom(err)
		}

		paymentAmount = paymentAmount.Sub(v.ScheduledAmount)
	}

	var loans sqlentity.Loans
	if loans, err = c.store.GetLoan(ctx, in.LoanID); err != nil {
		c.logger.Errorw("failed to get loan", "error", err)

		return pkgerror.ServerErrorFrom(err)
	}

	fLoan := loans.First()
	if err := c.store.UpdateLoan(ctx, &sqlentity.UpdateLoan{
		CurrentOutstandingBalance: fLoan.CurrentOutstandingBalance.Sub(totalPayment),
		UpdatedAt:                 time.Now(),
	}, gateway.UpdateLoanWithLoanIDFilter(in.LoanID)); err != nil {
		c.logger.Errorw("failed to update loan", "error", err)

		return pkgerror.ServerErrorFrom(err)
	}

	return nil
}
