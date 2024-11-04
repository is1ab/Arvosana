package types

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

type NullDatetime struct {
	Datetime Datetime
	Valid    bool
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
	return n.Datetime, nil
}

func (n NullDatetime) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.Datetime)
	}
	return json.Marshal(nil)
}

type NullSemester struct {
	Semester Semester
	Valid    bool
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

func (n NullSemester) Value() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.Semester)
	}
	return json.Marshal(nil)
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
