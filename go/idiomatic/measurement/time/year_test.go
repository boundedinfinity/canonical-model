package time

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Year_String(t *testing.T) {
	testCases := []struct {
		name     string
		input    Year
		expected string
	}{
		{
			name:     "positive year",
			input:    Year(2024),
			expected: "2024",
		},
		{
			name:     "zero year",
			input:    Year(0),
			expected: "0",
		},
		{
			name:     "negative year",
			input:    Year(-44),
			expected: "-44",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := tc.input.String()
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

func Test_Year_IsLeap(t *testing.T) {
	testCases := []struct {
		name     string
		input    Year
		expected bool
	}{
		{
			name:     "divisible by 400",
			input:    Year(2000),
			expected: true,
		},
		{
			name:     "divisible by 100 but not 400",
			input:    Year(1900),
			expected: false,
		},
		{
			name:     "divisible by 4 but not 100",
			input:    Year(2024),
			expected: true,
		},
		{
			name:     "not divisible by 4",
			input:    Year(2023),
			expected: false,
		},
		{
			name:     "zero year",
			input:    Year(0),
			expected: true,
		},
		{
			name:     "negative leap year",
			input:    Year(-4),
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := tc.input.IsLeap()
			assert.Equal(tt, tc.expected, actual)
		})
	}
}
