package csv

import (
	"time"
)

const dateFormat = "2006-01-02"

// Date for custom format as YYYY-MM-DD
type Date struct {
	time.Time
}

// MarshalCSV overrides custom format
func (d Date) MarshalCSV() ([]byte, error) {
	var b [len(dateFormat)]byte
	return d.AppendFormat(b[:0], dateFormat), nil
}

// UnmarshalCSV overrides custom format
func (d *Date) UnmarshalCSV(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	t, err := time.Parse(dateFormat, string(data))
	if err != nil {
		return err
	}
	*d = Date{Time: t}
	return nil
}
