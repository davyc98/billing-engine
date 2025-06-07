package sqlentity

import (
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
	PaymentDate     string
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
		l.ID,
		l.LoanID,
		l.WeekNumber,
		l.DueDate,
		l.ScheduledAmount,
		l.PaidAmount,
		l.PaymentStatus,
		l.PaymentDate,
	}
}

func (l *LoanSchedule) DriverValues() []driver.Value {
	vals := make([]driver.Value, len(l.Values()))
	for i, v := range l.Values() {
		vals[i] = v
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
