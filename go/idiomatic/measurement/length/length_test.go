package length_test

import (
	"testing"

	"github.com/boundedinfinity/canonical-model/go/idiomatic/measurement/imperial"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/measurement/length"
	"github.com/stretchr/testify/assert"
)

func Test_Length_parse(t *testing.T) {
	tcs := []struct {
		name     string
		input    string
		expected length.Length
		err      error
	}{
		{
			name:     "1 mile",
			input:    "1 mile",
			expected: &length.ImperialLength{Amount: 1, Kind: imperial.Lengths.Mile},
			err:      nil,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			actual, err := length.Parse(tc.input)

			if tc.err != nil {
				assert.ErrorIs(t, err, tc.err)
				return
			} else {
				assert.InDelta(t, tc.expected, actual, 0.0001)
				assert.NoError(t, err)
			}
		})
	}
}

// func Test_Length_conversion(t *testing.T) {
// 	tcs := []struct {
// 		name     string
// 		input    float64
// 		expected float64
// 		from     length.Length
// 		to       length.Length
// 	}{
// 		{
// 			name:     "1 mile to feet",
// 			input:    1,
// 			expected: 5280,
// 			from:     lengths.,
// 			to:       imperial.Lengths.Foot,
// 		},

// 	}

// 	for _, tc := range tcs {
// 		t.Run(tc.name, func(tt *testing.T) {
// 			actual := imperial.Lengths.Convert(tc.input, tc.from, tc.to)
// 			assert.InDelta(t, tc.expected, actual, 0.0001)
// 		})
// 	}
// }
