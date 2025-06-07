package sqlentity

import (
	"database/sql/driver"
	"time"
)

type User struct {
	ID        uint64
	Name      string
	Email     string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u User) Columns() []any {
	return []any{
		"id",
		"name",
		"email",
		"phone",
		"created_at",
		"updated_at",
	}
}

func (u User) StringColumns() []string {
	vals := make([]string, len(u.Columns()))
	for i, col := range u.Columns() {
		c, ok := col.(string)
		if ok {
			vals[i] = c
		}
	}

	return vals
}

func (u *User) Values() []any {
	return []any{
		u.ID,
		u.Name,
		u.Email,
		u.Phone,
		u.CreatedAt,
		u.UpdatedAt,
	}
}

func (u *User) DriverValues() []driver.Value {
	vals := make([]driver.Value, len(u.Values()))
	for i, v := range u.Values() {
		vals[i] = v
	}

	return vals
}
