package rational_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/mather/rational"
	"github.com/stretchr/testify/assert"
)

func Test_Rational_String(t *testing.T) {
	testCases := []struct {
		name     string
		input    rational.Rational[int]
		expected string
	}{
		{
			name:     "1 1/2",
			input:    rational.New[int](1, 1, 2),
			expected: "1 1/2",
		},
		{
			name:     "1 3/2",
			input:    rational.New[int](1, 3, 2),
			expected: "1 3/2",
		},
		{
			name:     "1 3/2 -> 2 1/2",
			input:    rational.New[int](1, 3, 2).Mixed(),
			expected: "2 1/2",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := tc.input.String()
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

func Test_Rational_Whole(t *testing.T) {
	assert.Equal(t, 3, rational.Whole(3.14))
}

func Test_Rational_Reduce(t *testing.T) {
	testCases := []struct {
		name     string
		input    rational.Rational[int]
		expected rational.Rational[int]
	}{
		{
			name:     "3.75",
			input:    rational.New[int](3, 75, 100),
			expected: rational.New[int](3, 3, 4),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := tc.input.Reduce()
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

func Test_Rational_Mixed(t *testing.T) {
	testCases := []struct {
		name     string
		input    rational.Rational[int]
		expected rational.Rational[int]
	}{
		{
			name:     "14/3",
			input:    rational.New[int](0, 14, 3),
			expected: rational.New[int](4, 2, 3),
		},
		{
			name:     "2 14/3",
			input:    rational.New[int](2, 14, 3),
			expected: rational.New[int](6, 2, 3),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := tc.input.Mixed()
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

func Test_Rational_Improper(t *testing.T) {
	testCases := []struct {
		name     string
		input    rational.Rational[int]
		expected rational.Rational[int]
	}{
		{
			name:     "2 8/32",
			input:    rational.New[int](2, 8, 32).Reduce(),
			expected: rational.New[int](0, 9, 4),
		},
		{
			name:     "1 3/4",
			input:    rational.New[int](1, 3, 4),
			expected: rational.New[int](0, 7, 4),
		},
		{
			name:     "3/4",
			input:    rational.New[int](0, 3, 4),
			expected: rational.New[int](0, 3, 4),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := tc.input.Improper()
			assert.Equal(tt, tc.expected, actual)
		})
	}
}
