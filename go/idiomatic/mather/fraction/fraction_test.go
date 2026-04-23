package fraction_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/mather/fraction"
	"github.com/stretchr/testify/assert"
)

func Test_Fraction_String(t *testing.T) {
	type TestInt int

	testCases := []struct {
		name     string
		input    fraction.Fraction[TestInt]
		expected string
	}{
		{
			name:     "1/2",
			input:    fraction.New[TestInt](1, 2),
			expected: "1/2",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := tc.input.String()
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

func Test_Fraction_Enumerate(t *testing.T) {
	testCases := []struct {
		name     string
		s, e     int
		input    fraction.Fraction[int]
		expected []fraction.Fraction[int]
	}{
		{
			name:  "4/32 x/2 -> x/64",
			input: fraction.New(4, 32),
			s:     2, e: 64,
			expected: []fraction.Fraction[int]{
				fraction.New(1, 8),
				fraction.New(2, 16),
				fraction.New(4, 32),
				fraction.New(8, 64),
			},
		},
		{
			name:  "1/2, 1/2 -> 32 / 64",
			input: fraction.New(1, 2),
			s:     2, e: 64,
			expected: []fraction.Fraction[int]{
				fraction.New(1, 2),
				fraction.New(2, 4),
				fraction.New(4, 8),
				fraction.New(8, 16),
				fraction.New(16, 32),
				fraction.New(32, 64),
			},
		},
		{
			name:  "8/16, 1/2 -> 32 / 64",
			input: fraction.New(8, 16),
			s:     2, e: 64,
			expected: []fraction.Fraction[int]{
				fraction.New(1, 2),
				fraction.New(2, 4),
				fraction.New(4, 8),
				fraction.New(8, 16),
				fraction.New(16, 32),
				fraction.New(32, 64),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := tc.input.Enumerate(tc.s, tc.e)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

func Test_Fraction_Component(t *testing.T) {
	testCases := []struct {
		name     string
		input    float64
		expected int
	}{
		{
			name:     "1/2",
			input:    1.0 / 2.0,
			expected: 5,
		},
		{
			name:     "3.14",
			input:    3.14,
			expected: 14,
		},
		{
			name:     "0.14",
			input:    0.14,
			expected: 14,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := fraction.Component(tc.input)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

func Test_Fraction_Magnitude(t *testing.T) {
	testCases := []struct {
		name     string
		input    float64
		expected int
	}{
		{
			name:     "1/2",
			input:    1.0 / 2.0,
			expected: 1,
		},
		{
			name:     "3.14",
			input:    3.14,
			expected: 2,
		},
		{
			name:     "0.14",
			input:    0.14,
			expected: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := fraction.Magnitude(tc.input)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

func Test_Fraction_FromFloat(t *testing.T) {
	testCases := []struct {
		name     string
		input    float64
		expected fraction.Fraction[int]
		reduced  fraction.Fraction[int]
	}{
		{
			name:  ".5",
			input: .5,
			expected: fraction.Fraction[int]{
				Numerator:   5,
				Denominator: 10,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   1,
				Denominator: 2,
			},
		},
		{
			name:  "1 / 32",
			input: float64(1) / float64(32),
			expected: fraction.Fraction[int]{
				Numerator:   3125,
				Denominator: 100000,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   1,
				Denominator: 32,
			},
		},
		{
			name:  "2 / 32",
			input: float64(2) / float64(32),
			expected: fraction.Fraction[int]{
				Numerator:   625,
				Denominator: 10000,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   1,
				Denominator: 16,
			},
		},
		{
			name:  "3 / 32",
			input: float64(3) / float64(32),
			expected: fraction.Fraction[int]{
				Numerator:   9375,
				Denominator: 100000,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   3,
				Denominator: 32,
			},
		},
		{
			name:  "4 / 32",
			input: float64(4) / float64(32),
			expected: fraction.Fraction[int]{
				Numerator:   125,
				Denominator: 1000,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   1,
				Denominator: 8,
			},
		},
		{
			name:  "5 / 32",
			input: float64(5) / float64(32),
			expected: fraction.Fraction[int]{
				Numerator:   15625,
				Denominator: 100000,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   5,
				Denominator: 32,
			},
		},
		{
			name:  "6 / 32",
			input: float64(6) / float64(32),
			expected: fraction.Fraction[int]{
				Numerator:   1875,
				Denominator: 10000,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   3,
				Denominator: 16,
			},
		},
		{
			name:  "7 / 32",
			input: float64(7) / float64(32),
			expected: fraction.Fraction[int]{
				Numerator:   21875,
				Denominator: 100000,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   7,
				Denominator: 32,
			},
		},
		{
			name:  "8 / 32",
			input: float64(8) / float64(32),
			expected: fraction.Fraction[int]{
				Numerator:   25,
				Denominator: 100,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   1,
				Denominator: 4,
			},
		},
		{
			name:  "9 / 32",
			input: float64(9) / float64(32),
			expected: fraction.Fraction[int]{
				Numerator:   28125,
				Denominator: 100000,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   9,
				Denominator: 32,
			},
		},
		{
			name:  "10 / 32",
			input: float64(10) / float64(32),
			expected: fraction.Fraction[int]{
				Numerator:   3125,
				Denominator: 10000,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   5,
				Denominator: 16,
			},
		},
		{
			name:  "11 / 32",
			input: float64(11) / float64(32),
			expected: fraction.Fraction[int]{
				Numerator:   34375,
				Denominator: 100000,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   11,
				Denominator: 32,
			},
		},
		{
			name:  "12 / 32",
			input: float64(12) / float64(32),
			expected: fraction.Fraction[int]{
				Numerator:   375,
				Denominator: 1000,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   3,
				Denominator: 8,
			},
		},
		{
			name:  "13 / 32",
			input: float64(13) / float64(32),
			expected: fraction.Fraction[int]{
				Numerator:   40625,
				Denominator: 100000,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   13,
				Denominator: 32,
			},
		},
		{
			name:  "14 / 32",
			input: float64(14) / float64(32),
			expected: fraction.Fraction[int]{
				Numerator:   4375,
				Denominator: 10000,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   7,
				Denominator: 16,
			},
		},
		{
			name:  "15 / 32",
			input: float64(15) / float64(32),
			expected: fraction.Fraction[int]{
				Numerator:   46875,
				Denominator: 100000,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   15,
				Denominator: 32,
			},
		},
		{
			name:  "16 / 32",
			input: float64(16) / float64(32),
			expected: fraction.Fraction[int]{
				Numerator:   5,
				Denominator: 10,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   1,
				Denominator: 2,
			},
		},
		{
			name:  "17 / 32",
			input: float64(17) / float64(32),
			expected: fraction.Fraction[int]{
				Numerator:   53125,
				Denominator: 100000,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   17,
				Denominator: 32,
			},
		},
		{
			name:  "18 / 32",
			input: float64(18) / float64(32),
			expected: fraction.Fraction[int]{
				Numerator:   5625,
				Denominator: 10000,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   9,
				Denominator: 16,
			},
		},
		{
			name:  "19 / 32",
			input: float64(19) / float64(32),
			expected: fraction.Fraction[int]{
				Numerator:   59375,
				Denominator: 100000,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   19,
				Denominator: 32,
			},
		},
		{
			name:  "20 / 32",
			input: float64(20) / float64(32),
			expected: fraction.Fraction[int]{
				Numerator:   625,
				Denominator: 1000,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   5,
				Denominator: 8,
			},
		},
		{
			name:  "21 / 32",
			input: float64(21) / float64(32),
			expected: fraction.Fraction[int]{
				Numerator:   65625,
				Denominator: 100000,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   21,
				Denominator: 32,
			},
		},
		{
			name:  "22 / 32",
			input: float64(22) / float64(32),
			expected: fraction.Fraction[int]{
				Numerator:   6875,
				Denominator: 10000,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   11,
				Denominator: 16,
			},
		},
		{
			name:  "23 / 32",
			input: float64(23) / float64(32),
			expected: fraction.Fraction[int]{
				Numerator:   71875,
				Denominator: 100000,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   23,
				Denominator: 32,
			},
		},
		{
			name:  "24 / 32",
			input: float64(24) / float64(32),
			expected: fraction.Fraction[int]{
				Numerator:   75,
				Denominator: 100,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   3,
				Denominator: 4,
			},
		},
		{
			name:  "25 / 32",
			input: float64(25) / float64(32),
			expected: fraction.Fraction[int]{
				Numerator:   78125,
				Denominator: 100000,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   25,
				Denominator: 32,
			},
		},
		{
			name:  "26 / 32",
			input: float64(26) / float64(32),
			expected: fraction.Fraction[int]{
				Numerator:   8125,
				Denominator: 10000,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   13,
				Denominator: 16,
			},
		},
		{
			name:  "27 / 32",
			input: float64(27) / float64(32),
			expected: fraction.Fraction[int]{
				Numerator:   84375,
				Denominator: 100000,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   27,
				Denominator: 32,
			},
		},
		{
			name:  "28 / 32",
			input: float64(28) / float64(32),
			expected: fraction.Fraction[int]{
				Numerator:   875,
				Denominator: 1000,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   7,
				Denominator: 8,
			},
		},
		{
			name:  "29 / 32",
			input: float64(29) / float64(32),
			expected: fraction.Fraction[int]{
				Numerator:   90625,
				Denominator: 100000,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   29,
				Denominator: 32,
			},
		},
		{
			name:  "30 / 32",
			input: float64(30) / float64(32),
			expected: fraction.Fraction[int]{
				Numerator:   9375,
				Denominator: 10000,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   15,
				Denominator: 16,
			},
		},
		{
			name:  "31 / 32",
			input: float64(31) / float64(32),
			expected: fraction.Fraction[int]{
				Numerator:   96875,
				Denominator: 100000,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   31,
				Denominator: 32,
			},
		},
		{
			name:  "32 / 32",
			input: float64(32) / float64(32),
			expected: fraction.Fraction[int]{
				Numerator:   1,
				Denominator: 1,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   1,
				Denominator: 1,
			},
		},
		{
			name:  "3.14159",
			input: 3.14159,
			expected: fraction.Fraction[int]{
				Numerator:   314159,
				Denominator: 100000,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   314159,
				Denominator: 100000,
			},
		},
		{
			name:  "2.71828",
			input: 2.71828,
			expected: fraction.Fraction[int]{
				Numerator:   271828,
				Denominator: 100000,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   67957,
				Denominator: 25000,
			},
		},
		{
			name:  "3.14",
			input: 3.14,
			expected: fraction.Fraction[int]{
				Numerator:   314,
				Denominator: 100,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   157,
				Denominator: 50,
			},
		},
		{
			name:  "6.124",
			input: 6.124,
			expected: fraction.Fraction[int]{
				Numerator:   6124,
				Denominator: 1000,
			},
			reduced: fraction.Fraction[int]{
				Numerator:   1531,
				Denominator: 250,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := fraction.FromFloat[int](tc.input)
			reduced := actual.Reduce()

			assert.Equal(tt, tc.expected, actual)
			assert.Equal(tt, tc.reduced, reduced)
		})
	}
}
