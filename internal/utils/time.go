package utils

import (
	"strconv"
	"time"
)

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
