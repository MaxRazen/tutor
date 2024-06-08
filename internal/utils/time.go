package utils

import (
	"encoding/json"
	"strconv"
	"time"
)

// Alias to base time.Time type to use for JSON serialization
type JSONTime time.Time

// Marshaler interface implementation. Converts Unix timestamp to []byte
func (t JSONTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Unix())
}

// Parses Unix timestamp from string and returns time.Time with given timestamp
// The function might return 1970-01-01 00:00:00 when the given timestamp is empty string or invalid
func ParseTimestamp(ts string) time.Time {
	if ts == "" {
		return time.Unix(0, 0)
	}

	s, err := strconv.ParseInt(ts, 10, 64)

	if err != nil {
		s = 0
	}

	return time.Unix(s, 0)
}

// Parses datetime from string formatted in the format 2006-01-02 15:04:05
// The function might return 1970-01-01 00:00:00 when the given datetime is empty string or cannot be parsed
func ParseDatetime(dt string) time.Time {
	if dt == "" {
		return time.Unix(0, 0)
	}

	t, err := time.Parse(time.DateTime, dt)

	if err != nil {
		return time.Unix(0, 0)
	}

	return t
}
