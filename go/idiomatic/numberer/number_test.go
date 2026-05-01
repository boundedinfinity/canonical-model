package numberer_test

import (
	"testing"

	"github.com/boundedinfinity/canonical-model/go/idiomatic/numberer"
	"github.com/stretchr/testify/assert"
)

func Test_Fraction_Reciprocal(t *testing.T) {
	testCases := []struct {
		name     string
		input    numberer.Number
		expected numberer.Number
	}{
		{
			name:     "reciprocal 2/4 to 4/2",
			input:    numberer.Fraction(2, 4),
			expected: numberer.Fraction(4, 2),
		},
		{
			name:     "reciprocal 10/5 to 5/10",
			input:    numberer.Fraction(10, 5),
			expected: numberer.Fraction(5, 10),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			assert.IsType(tt, &numberer.FractionNumber{}, tc.input)
			assert.IsType(tt, &numberer.FractionNumber{}, tc.expected)

			input := tc.input.(*numberer.FractionNumber)
			expected := tc.expected.(*numberer.FractionNumber)

			actual := input.Reciprocal()
			assert.IsType(tt, &numberer.FractionNumber{}, actual)

			assert.Equal(tt, expected.Numerator, actual.(*numberer.FractionNumber).Numerator)
			assert.Equal(tt, expected.Denominator, actual.(*numberer.FractionNumber).Denominator)
		})
	}
}
