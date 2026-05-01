package numberer_test

import (
	"testing"

	"github.com/boundedinfinity/canonical-model/go/idiomatic/numberer"
	"github.com/stretchr/testify/assert"
)

func Test_Fraction_Reduce(t *testing.T) {
	testCases := []struct {
		name     string
		input    numberer.FractionNumber
		expected numberer.FractionNumber
	}{
		{
			name:     "reduce 2/4 to 1/2",
			input:    numberer.FractionNumber{Numerator: 2, Denominator: 4},
			expected: numberer.FractionNumber{Numerator: 1, Denominator: 2},
		},
		{
			name:     "reduce 10/5 to 2/1",
			input:    numberer.FractionNumber{Numerator: 10, Denominator: 5},
			expected: numberer.FractionNumber{Numerator: 2, Denominator: 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := tc.input.Reduce()
			assert.Equal(tt, tc.expected.Numerator, actual.Numerator)
			assert.Equal(tt, tc.expected.Denominator, actual.Denominator)
		})
	}
}
