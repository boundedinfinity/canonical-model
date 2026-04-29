package imperial_test

import (
	"testing"

	"github.com/boundedinfinity/canonical-model/go/idiomatic/measurement/imperial"
	"github.com/stretchr/testify/assert"
)

func Test_Imperial_Convert(t *testing.T) {
	tcs := []struct {
		name     string
		input    float64
		expected float64
		from     imperial.Length
		to       imperial.Length
	}{
		{
			name:     "1 mile to feet",
			input:    1,
			expected: 5280,
			from:     imperial.Lengths.Mile,
			to:       imperial.Lengths.Foot,
		},
		{
			name:     "1 mile to nautical mile",
			input:    1,
			expected: 0.868976,
			from:     imperial.Lengths.Mile,
			to:       imperial.Lengths.NauticalMile,
		},
		{
			name:     "24 inches to feet",
			input:    24,
			expected: 2,
			from:     imperial.Lengths.Inch,
			to:       imperial.Lengths.Foot,
		},
		{
			name:     "foot to inches",
			input:    1,
			expected: 12,
			from:     imperial.Lengths.Foot,
			to:       imperial.Lengths.Inch,
		},
		{
			name:     "2 feet to inches",
			input:    2,
			expected: 24,
			from:     imperial.Lengths.Foot,
			to:       imperial.Lengths.Inch,
		},
		{
			name:     "1 foot to feet",
			input:    1,
			expected: 1,
			from:     imperial.Lengths.Foot,
			to:       imperial.Lengths.Foot,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			actual := imperial.Lengths.Convert(tc.input, tc.from, tc.to)
			assert.InDelta(t, tc.expected, actual, 0.0001)
		})
	}
}
