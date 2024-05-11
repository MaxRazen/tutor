package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInSlice(t *testing.T) {
	slice := []string{
		"one",
		"two",
		"six",
	}

	assert.Equal(t, true, InSlice("two", slice))
	assert.Equal(t, false, InSlice("seven", slice))
}
