package ider_test

import (
	"testing"

	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
	"github.com/stretchr/testify/assert"
)

func Test_prefix_name_and_symbol(t *testing.T) {
	tcs := []struct {
		name     string
		input    ider.Id
		expected ider.Id
	}{
		{
			name:     "is zero",
			input:    ider.Zero,
			expected: ider.Zero,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			assert.Equal(t, tc.expected, tc.input)
		})
	}
}
