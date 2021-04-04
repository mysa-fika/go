package storage

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// JSONB represents a PostgreSQL's JSONB value:
// https://www.postgresql.org/docs/9.6/static/datatype-json.html. JSONB
// satisfies sqlbuilder.ScannerValuer.
type JSONB struct {
	V interface{}
}

// MarshalJSON encodes the wrapper value as JSON.
func (j JSONB) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.V)
}

// UnmarshalJSON decodes the given JSON into the wrapped value.
func (j *JSONB) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	j.V = v
	return nil
}

// Scan satisfies the sql.Scanner interface.
func (j *JSONB) Scan(src interface{}) error {
	if src == nil {
		j.V = nil
		return nil
	}

	b, ok := src.([]byte)
	if !ok {
		return errors.New("Scan source was not []bytes")
	}

	if err := json.Unmarshal(b, &j.V); err != nil {
		return err
	}
	return nil
}

// Value satisfies the driver.Valuer interface.
func (j JSONB) Value() (driver.Value, error) {
	// See https://github.com/lib/pq/issues/528#issuecomment-257197239 on why are
	// we returning string instead of []byte.
	if j.V == nil {
		return nil, nil
	}
	if v, ok := j.V.(json.RawMessage); ok {
		return string(v), nil
	}
	b, err := json.Marshal(j.V)
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

// JSONBValue takes an interface and provides a driver.Value that can be
// stored as a JSONB column.
func JSONBValue(i interface{}) (driver.Value, error) {
	v := JSONB{i}
	return v.Value()
}

// ScanJSONB decodes a JSON byte stream into the passed dst value.
func ScanJSONB(dst interface{}, src interface{}) error {
	v := JSONB{dst}
	return v.Scan(src)
}
