package types

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// Datetime a thin wrapper around `time.Time` because sqlite3 doesn't support timestamp types
type Datetime struct {
	t time.Time
}

const format = "2006-01-02 15:04:05"

func NewDatetime(t time.Time) Datetime {
	return Datetime{
		t: t,
	}
}

func (dt *Datetime) Time() time.Time {
	return dt.t
}

func (dt *Datetime) Scan(value any) error {
	if value == nil {
		dt.t = time.Time{}
		return nil
	}

	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("invalid value type %T: %v", value, value)
	}

	t, err := time.Parse(format, str)
	if err != nil {
		return err
	}

	dt.t = t
	return nil
}

func (dt Datetime) Value() (driver.Value, error) {
	return dt.t.Format(format), nil
}
