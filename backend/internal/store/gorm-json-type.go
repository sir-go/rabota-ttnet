package store

import (
	"bytes"
	"database/sql/driver"
	"errors"
)

// JSON implements a database model interface for storing user profile as a JSON
type JSON []byte

// Value returns a record as a string
func (j *JSON) Value() (driver.Value, error) {
	if j.IsNull() {
		return nil, nil
	}
	return string(*j), nil
}

// Scan converts a value to a byte slice
func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		return errors.New("invalid Scan Source")
	}
	*j = append((*j)[0:0], s...)
	return nil
}

// MarshalJSON presents a record as a byte slice
func (j *JSON) MarshalJSON() ([]byte, error) {
	if j == nil {
		return []byte("null"), nil
	}
	return *j, nil
}

// UnmarshalJSON stores the given byte slice as a JSON object
func (j *JSON) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("null point exception")
	}
	*j = append((*j)[0:0], data...)
	return nil
}

// IsNull checks if the record is empty
func (j *JSON) IsNull() bool {
	return len(*j) == 0 || string(*j) == "null"
}

// Equals makes JSON objects comparable
func (j *JSON) Equals(j1 JSON) bool {
	return bytes.Equal(*j, j1)
}
