package numberer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/numberer"
	"github.com/stretchr/testify/assert"
)

func Test_Integer_GreatestCommonFactor(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "10, 0",
			input:    []int{10, 0},
			expected: 0,
		},
		{
			name:     "-10, 0",
			input:    []int{-10, 0},
			expected: 0,
		},
		{
			name:     "0, -10",
			input:    []int{0, -10},
			expected: 0,
		},
		{
			name:     "3, 5",
			input:    []int{3, 5},
			expected: 1,
		},
		{
			name:     "-3, 5",
			input:    []int{-3, 5},
			expected: 1,
		},
		{
			name:     "3, -5",
			input:    []int{3, -5},
			expected: 1,
		},
		{
			name:     "2, 4",
			input:    []int{2, 4},
			expected: 2,
		},
		{
			name:     "-2, 4",
			input:    []int{-2, 4},
			expected: 2,
		},
		{
			name:     "2, -4",
			input:    []int{2, -4},
			expected: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := numberer.GreatestCommonFactor(tc.input[0], tc.input[1])
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

func Test_Integer_AllFactors(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected []int
	}{
		{
			name:     "0",
			input:    0,
			expected: nil,
		},
		{
			name:     "10",
			input:    10,
			expected: []int{1, 2, 5, 10},
		},
		{
			name:     "-10",
			input:    -10,
			expected: []int{1, 2, 5, 10},
		},
		{
			name:     "13",
			input:    13,
			expected: []int{1, 13},
		},
		{
			name:     "-13",
			input:    -13,
			expected: []int{1, 13},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := numberer.EnumerateFactors(tc.input)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}
