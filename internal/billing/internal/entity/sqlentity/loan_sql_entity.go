package sqlentity

import (
	"database/sql/driver"
	"errors"
	"time"

	"github.com/shopspring/decimal"
)

type Loan struct {
	ID                        uint64
	CustomerID                uint64
	LoanAmount                decimal.Decimal
	InterestRate              decimal.Decimal
	LoanTermWeeks             int64
	Status                    LoanStatus
	StartDate                 string
	EndDate                   string
	TotalPayableAmount        decimal.Decimal
	WeeklyPaymentAmount       decimal.Decimal
	CurrentOutstandingBalance decimal.Decimal
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
}

func (l Loan) Columns() []any {
	return []any{
		"id",
		"customer_id",
		"loan_amount",
		"interest_rate",
		"status",
		"start_date",
		"end_date",
		"total_payable_amount",
		"weekly_payment_amount",
		"current_oustanding_balance",
		"created_at",
		"updated_at",
	}
}

func (l Loan) StringColumns() []string {
	vals := make([]string, len(l.Columns()))
	for i, col := range l.Columns() {
		c, ok := col.(string)
		if ok {
			vals[i] = c
		}
	}

	return vals
}

func (l *Loan) Values() []any {
	return []any{
		&l.ID,
		&l.CustomerID,
		&l.LoanAmount,
		&l.InterestRate,
		&l.Status,
		&l.LoanTermWeeks,
		&l.StartDate,
		&l.EndDate,
		&l.TotalPayableAmount,
		&l.WeeklyPaymentAmount,
		&l.CurrentOutstandingBalance,
		&l.CreatedAt,
		&l.UpdatedAt,
	}
}

func (l Loan) DriverValues() []driver.Value {
	vals := make([]driver.Value, len(l.Values()))
	for i, v := range l.Values() {
		vals[i] = v
	}

	return vals
}

type Loans []Loan

func (l Loans) IsEmpty() bool {
	return l.Len() == 0
}

func (l Loans) Len() int {
	return len(l)
}

func (l Loans) First() Loan {
	if l.IsEmpty() {
		return Loan{}
	}

	return l[0]
}

type LoanStatus int

const (
	UnknownStatus LoanStatus = iota
	Active
	Inactive
	Complete
)

func (ls LoanStatus) String() string {
	return [...]string{"UNKNOWN", "ACTIVE", "INACTIVE", "COMPLETED"}[ls]
}

func (ls LoanStatus) Value() (driver.Value, error) {
	return ls.String(), nil
}

func (ls LoanStatus) getMap() map[string]LoanStatus {
	return map[string]LoanStatus{
		"UNKNOWN":   UnknownStatus,
		"ACTIVE":    Active,
		"INACTIVE":  Inactive,
		"COMPLETED": Complete,
	}
}

func (ls *LoanStatus) Scan(value any) error {
	b, ok := value.([]byte)
	if ok {
		val := ls.getMap()[string(b)]

		*ls = val

		return nil
	}

	return errors.New("failed to scan loan status")
}

type UpdateLoan struct {
	CurrentOutstandingBalance decimal.Decimal
	UpdatedAt                 time.Time
}

func (l UpdateLoan) Columns() []any {
	return []any{
		"current_outstanding_balance",
		"updated_at",
	}
}

func (l UpdateLoan) StringColumns() []string {
	vals := make([]string, len(l.Columns()))
	for i, col := range l.Columns() {
		c, ok := col.(string)
		if ok {
			vals[i] = c
		}
	}

	return vals
}

func (l *UpdateLoan) Values() []any {
	return []any{
		&l.CurrentOutstandingBalance,
		&l.UpdatedAt,
	}
}

func (l *UpdateLoan) DriverValues() []driver.Value {
	vals := make([]driver.Value, len(l.Values()))
	for i, v := range l.Values() {
		vals[i] = v
	}

	return vals
}

func (a *UpdateLoan) MappedValues() map[string]driver.Value {
	vals := make(map[string]driver.Value)
	cols := a.StringColumns()
	for i, col := range cols {
		vals[col] = a.DriverValues()[i]
	}

	return vals
}
