package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLastSegment(t *testing.T) {
	testCases := []struct {
		input    string
		sep      string
		expected string
	}{
		{
			input:    "path/to/target",
			sep:      "/",
			expected: "target",
		},
		{
			input:    "https://example.com/category/auto?ref=affiliate",
			sep:      "/",
			expected: "auto?ref=affiliate",
		},
		{
			input:    "App::class::test",
			sep:      "::",
			expected: "test",
		},
		{
			input:    "/path/to/target",
			sep:      ",",
			expected: "/path/to/target",
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, GetLastSegment(tc.input, tc.sep))
	}
}
