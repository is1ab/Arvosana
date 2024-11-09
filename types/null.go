package types

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type NullDatetime struct {
	Datetime Datetime
	Valid    bool
}

var InvalidNullDatetime = NullDatetime{
	Datetime: Datetime{},
	Valid:    false,
}

func (n *NullDatetime) Scan(value any) error {
	if value == nil {
		n.Datetime = Datetime{}
		n.Valid = false
		return nil
	}
	n.Valid = true
	return n.Datetime.Scan(value)
}

func (n NullDatetime) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Datetime.Value()
}

func (n NullDatetime) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.Datetime)
	}
	return json.Marshal(nil)
}

// ParseNullDateTime treats `int64(0)` as null
func ParseNullDateTime(value any) (NullDatetime, error) {
	if value == nil {
		return InvalidNullDatetime, nil
	}

	i, ok := value.(int64)
	if !ok {
		return InvalidNullDatetime, fmt.Errorf("invalid value type %T: %v", value, value)
	}

	if i == 0 {
		return InvalidNullDatetime, nil
	}

	dt := time.Unix(i, 0)

	return NullDatetime{
		Datetime: NewDatetime(dt),
		Valid:    true,
	}, nil
}

type NullSemester struct {
	Semester Semester
	Valid    bool
}

var InvalidNullSemester = NullSemester{
	Semester: Semester{},
	Valid:    false,
}

func (n *NullSemester) Scan(value any) error {
	if value == nil {
		n.Semester = Semester{}
		n.Valid = false
		return nil
	}
	n.Valid = true
	return n.Semester.Scan(value)
}

func (n NullSemester) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Semester.Value()
}

// ParseNullSemester treats `""` as null
func ParseNullSemester(value any) (NullSemester, error) {
	if value == nil {
		return InvalidNullSemester, nil
	}

	str, ok := value.(string)
	if !ok {
		return InvalidNullSemester, fmt.Errorf("invalid value type %T: %v", value, value)
	}

	if str == "" {
		return InvalidNullSemester, nil
	}

	sem, err := ParseSemester(str)
	if err != nil {
		return InvalidNullSemester, err
	}

	return NullSemester{
		Semester: sem,
		Valid:    true,
	}, nil
}

type NullString struct {
	sql.NullString
}

func (nf NullString) MarshalJSON() ([]byte, error) {
	if nf.Valid {
		return json.Marshal(nf.String)
	}
	return json.Marshal(nil)
}

var InvalidNullString = NullString{
	NullString: sql.NullString{
		String: "",
		Valid:  false,
	},
}

// ParseNullString treats `""` as null
func ParseNullString(value any) (NullString, error) {
	if value == nil {
		return InvalidNullString, nil
	}

	str, ok := value.(string)
	if !ok {
		return InvalidNullString, fmt.Errorf("invalid value type %T: %v", value, value)
	}

	if str == "" {
		return InvalidNullString, nil
	}

	return NullString{
		NullString: sql.NullString{
			String: str,
			Valid:  true,
		},
	}, nil
}

type NullFloat64 struct {
	sql.NullFloat64
}

func (nf NullFloat64) MarshalJSON() ([]byte, error) {
	if nf.Valid {
		return json.Marshal(nf.Float64)
	}
	return json.Marshal(nil)
}
