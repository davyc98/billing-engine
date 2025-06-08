package gateway

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/davyc98/billing-engine/internal/billing/internal/entity/sqlentity"
	"github.com/davyc98/billing-engine/internal/pkg/pkgsql"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"

	"go.uber.org/zap"
)

type LoanSQLGateway struct {
	db           pkgsql.SQL
	logger       *zap.SugaredLogger
	queryBuilder pkgsql.GoquBuilder

	loanTableName         string
	loanScheduleTableName string
	customersTableName    string
}

func NewLoanSQLGateway(
	db *sql.DB,
	logger *zap.SugaredLogger,
	queryBuilder pkgsql.GoquBuilder,
) *LoanSQLGateway {
	return &LoanSQLGateway{
		db:           db,
		logger:       logger,
		queryBuilder: queryBuilder,

		loanTableName:         "loan",
		loanScheduleTableName: "loan_schedule",
		customersTableName:    "customers",
	}
}

func (r *LoanSQLGateway) GetLoanSchedule(ctx context.Context, loanID uint64) (sqlentity.LoanSchedules, error) {
	var res sqlentity.LoanSchedule
	currDate := time.Now().Format("2006-01-02")
	query := r.queryBuilder.Select(res.Columns()...).From(r.loanScheduleTableName).Where(goqu.Ex{
		"loan_id":        loanID,
		"payment_status": goqu.Op{"neq": "PAID"},
	}, goqu.C("due_date").Lte(currDate))

	sql, _, err := query.ToSQL()
	if err != nil {
		r.logger.Errorw("failed to build query", "error", err)

		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, sql)
	if err != nil {
		r.logger.Errorw("failed to execute query", "error", err)

		return nil, err
	}

	var loanSchs sqlentity.LoanSchedules
	for rows.Next() {
		err := rows.Scan(res.Values()...)
		if err != nil {
			r.logger.Errorw("failed to scan row", "error", err)

			return nil, err
		}

		loanSchs = append(loanSchs, res)
	}

	return loanSchs, nil
}

func (r *LoanSQLGateway) GetLoan(ctx context.Context, loanID uint64) (sqlentity.Loans, error) {
	var res sqlentity.Loan
	query := r.queryBuilder.Select(res.Columns()...).From(r.loanTableName).Where(goqu.Ex{
		"id": loanID,
	})

	sql, _, err := query.ToSQL()
	if err != nil {
		r.logger.Errorw("failed to build query", "error", err)

		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, sql)
	if err != nil {
		r.logger.Errorw("failed to execute query", "error", err)

		return nil, err
	}

	var loans sqlentity.Loans
	for rows.Next() {
		err := rows.Scan(res.Values()...)
		if err != nil {
			r.logger.Errorw("failed to scan row", "error", err)

			return nil, err
		}

		loans = append(loans, res)
	}

	return loans, nil
}

type UpdateLoanOption func(*goqu.UpdateDataset) *goqu.UpdateDataset

func UpdateLoanWithLoanIDFilter(loanID uint64) UpdateLoanOption {
	return func(query *goqu.UpdateDataset) *goqu.UpdateDataset {
		return query.Where(goqu.Ex{
			"id": loanID,
		})
	}
}

func (r *LoanSQLGateway) UpdateLoan(
	ctx context.Context,
	in sqlentity.UpdateEntity,
	opts ...UpdateLoanOption,
) error {
	query := r.queryBuilder.Update(r.loanTableName).Set(in.MappedValues())

	for _, opt := range opts {
		query = opt(query)
	}

	sql, _, err := query.ToSQL()
	if err != nil {
		r.logger.Errorw("failed to build query", "error", err)

		return err
	}

	res, err := r.db.ExecContext(ctx, sql)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		r.logger.Errorw("failed to execute query", "error", err)

		return err
	}

	row, err := res.RowsAffected()
	if err != nil {
		r.logger.Errorw("failed to get last insert id", "error", err)

		return err
	}

	if row == 0 {
		return fmt.Errorf("loan not found")
	}

	return nil
}

type UpdateLoanScheduleOption func(*goqu.UpdateDataset) *goqu.UpdateDataset

func UpdateLoanScheduleWithLoanIDAndIDFilter(id, loanID uint64) UpdateLoanScheduleOption {
	return func(query *goqu.UpdateDataset) *goqu.UpdateDataset {
		return query.Where(goqu.Ex{
			"id":      id,
			"loan_id": loanID,
		})
	}
}

func (r *LoanSQLGateway) UpdateLoanSchedule(
	ctx context.Context,
	in sqlentity.UpdateEntity,
	opts ...UpdateLoanScheduleOption,
) error {
	query := r.queryBuilder.Update(r.loanScheduleTableName).Set(in.MappedValues())

	for _, opt := range opts {
		query = opt(query)
	}

	sql, _, err := query.ToSQL()
	if err != nil {
		r.logger.Errorw("failed to build query", "error", err)

		return err
	}

	res, err := r.db.ExecContext(ctx, sql)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		r.logger.Errorw("failed to execute query", "error", err)

		return err
	}

	row, err := res.RowsAffected()
	if err != nil {
		r.logger.Errorw("failed to get last insert id", "error", err)

		return err
	}

	if row == 0 {
		return fmt.Errorf("loan not found")
	}

	return nil
}
