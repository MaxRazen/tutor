package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseTimestamp(t *testing.T) {
	loc, err := time.LoadLocation("UTC")

	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "1717848752",
			expected: "2024-06-08 12:12:32",
		},
		{
			input:    "2524622400",
			expected: "2050-01-01 04:00:00",
		},
		{
			input:    "",
			expected: "1970-01-01 00:00:00",
		},
		{
			input:    "0",
			expected: "1970-01-01 00:00:00",
		},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.expected, ParseTimestamp(tc.input).In(loc).Format(time.DateTime))
	}
}

func TestParseDatetime(t *testing.T) {
	testCases := []struct {
		input    string
		expected int64
	}{
		{
			input:    "2024-06-08 12:12:32",
			expected: 1717848752,
		},
		{
			input:    "2050-01-01 04:00:00",
			expected: 2524622400,
		},
		{
			input:    "1970-01-01 00:00:00",
			expected: 0,
		},
		{
			input:    "1970-01-01 00:00:00",
			expected: 0,
		},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.expected, ParseDatetime(tc.input).Unix())
	}
}
