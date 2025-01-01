package types

import (
	"database/sql/driver"
	"encoding/json"
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
	switch t.Month() {
	case time.February, time.March, time.April, time.May, time.June, time.July:
		return Semester{
			year:   t.Year() - 1,
			period: Spring,
		}
	case time.August, time.September, time.October, time.November, time.December, time.January:
		return Semester{
			year:   t.Year(),
			period: Fall,
		}
	default:
		panic("impossible case: sum type in golang when")
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

func (sem Semester) String() string {
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

	*sem = _sem
	return nil
}

func (sem Semester) Value() (driver.Value, error) {
	return sem.String(), nil
}

func (sem *Semester) UnmarshalJSON(b []byte) error {
	var data string
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}

	_sem, err := ParseSemester(data)
	if err != nil {
		return err
	}

	*sem = _sem
	return nil
}

func (sem Semester) MarshalJSON() ([]byte, error) {
	return json.Marshal(sem.String())
}
