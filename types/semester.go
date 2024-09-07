package types

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

type SemesterPeriod string

const (
	Spring SemesterPeriod = "s"
	Fall   SemesterPeriod = "f"
)

func ParseSemesterPeriod(value string) (SemesterPeriod, error) {
	switch SemesterPeriod(value) {
	case Spring, Fall:
		return SemesterPeriod(value), nil
	default:
		return "", fmt.Errorf("invalid SemesterPeriod %s", value)
	}
}

type Semester struct {
	year   int
	period SemesterPeriod
}

func TimeToSemester(t time.Time) Semester {
	if t.Month() >= time.September {
		return Semester{
			year:   t.Year(),
			period: Fall,
		}
	} else {
		return Semester{
			year:   t.Year(),
			period: Spring,
		}
	}
}

// ParseSemester value should follow `YYYY(s|f)` format
// e.g. 2024f parses into Semester{year: 2024, period: "fall"}
func ParseSemester(value string) (Semester, error) {
	if len(value) != 5 {
		return Semester{}, fmt.Errorf("invalid string %s: length should be 5", value)
	}

	year, err := strconv.Atoi(value[:4])
	if err != nil {
		return Semester{}, fmt.Errorf("invalid year %s: %w", value[:4], err)
	}

	period, err := ParseSemesterPeriod(string(value[4]))
	if err != nil {
		return Semester{}, err
	}

	return Semester{
		year:   year,
		period: period,
	}, nil
}

func (sem Semester) Format() string {
	return strconv.Itoa(sem.year) + string(sem.period)
}

func (sem *Semester) Scan(value any) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("invalid value type %T: %v", value, value)
	}

	_sem, err := ParseSemester(str)
	if err != nil {
		return err
	}

	sem.year = _sem.year
	sem.period = _sem.period
	return nil
}

func (sem Semester) Value() (driver.Value, error) {
	return sem.Format(), nil
}
