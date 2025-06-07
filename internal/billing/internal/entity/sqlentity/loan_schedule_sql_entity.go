package sqlentity

import (
	"database/sql"
	"database/sql/driver"
	"errors"

	"github.com/shopspring/decimal"
)

type LoanSchedule struct {
	ID              uint64
	LoanID          uint64
	WeekNumber      int
	DueDate         string
	ScheduledAmount decimal.Decimal
	PaidAmount      decimal.Decimal
	PaymentStatus   PaymentStatus
	PaymentDate     sql.NullString
}

func (l LoanSchedule) Columns() []any {
	return []any{
		"id",
		"loan_id",
		"week_number",
		"due_date",
		"scheduled_amount",
		"paid_amount",
		"payment_status",
		"payment_date",
	}
}

func (l LoanSchedule) StringColumns() []string {
	vals := make([]string, len(l.Columns()))
	for i, col := range l.Columns() {
		c, ok := col.(string)
		if ok {
			vals[i] = c
		}
	}

	return vals
}

func (l *LoanSchedule) Values() []any {
	return []any{
		&l.ID,
		&l.LoanID,
		&l.WeekNumber,
		&l.DueDate,
		&l.ScheduledAmount,
		&l.PaidAmount,
		&l.PaymentStatus,
		&l.PaymentDate,
	}
}

func (l *LoanSchedule) DriverValues() []driver.Value {
	vals := make([]driver.Value, len(l.Values()))
	for i, v := range l.Values() {
		vals[i] = v
	}

	return vals
}

func (a *LoanSchedule) MappedValues() map[string]driver.Value {
	vals := make(map[string]driver.Value)
	cols := a.StringColumns()
	for i, col := range cols {
		vals[col] = a.DriverValues()[i]
	}

	return vals
}

type PaymentStatus int

const (
	UnknownPaymentStatus PaymentStatus = iota
	Due
	Overdue
	Paid
)

func (ls PaymentStatus) String() string {
	return [...]string{"UNKNOWN", "DUE", "OVERDUE", "PAID"}[ls]
}

func (ls PaymentStatus) Value() (driver.Value, error) {
	return ls.String(), nil
}

func (ls PaymentStatus) getMap() map[string]PaymentStatus {
	return map[string]PaymentStatus{
		"UNKNOWN": UnknownPaymentStatus,
		"DUE":     Due,
		"OVERDUE": Overdue,
		"PAID":    Paid,
	}
}

func (ls *PaymentStatus) Scan(value any) error {
	b, ok := value.([]byte)
	if ok {
		val := ls.getMap()[string(b)]

		*ls = val

		return nil
	}

	return errors.New("failed to scan loan status")
}

type LoanSchedules []LoanSchedule

func (l LoanSchedules) IsEmpty() bool {
	return l.Len() == 0
}

func (l LoanSchedules) Len() int {
	return len(l)
}

func (l LoanSchedules) First() LoanSchedule {
	if l.IsEmpty() {
		return LoanSchedule{}
	}

	return l[0]
}

type UpdateLoanSchedule struct {
	PaidAmount    decimal.Decimal
	PaymentStatus PaymentStatus
	PaymentDate   sql.NullString
}

func (l UpdateLoanSchedule) Columns() []any {
	return []any{
		"paid_amount",
		"payment_status",
		"payment_date",
	}
}

func (l UpdateLoanSchedule) StringColumns() []string {
	vals := make([]string, len(l.Columns()))
	for i, col := range l.Columns() {
		c, ok := col.(string)
		if ok {
			vals[i] = c
		}
	}

	return vals
}

func (l *UpdateLoanSchedule) Values() []any {
	return []any{
		&l.PaidAmount,
		&l.PaymentStatus,
		&l.PaymentDate,
	}
}

func (l *UpdateLoanSchedule) DriverValues() []driver.Value {
	vals := make([]driver.Value, len(l.Values()))
	for i, v := range l.Values() {
		vals[i] = v
	}

	return vals
}

func (a *UpdateLoanSchedule) MappedValues() map[string]driver.Value {
	vals := make(map[string]driver.Value)
	cols := a.StringColumns()
	for i, col := range cols {
		vals[col] = a.DriverValues()[i]
	}

	return vals
}
