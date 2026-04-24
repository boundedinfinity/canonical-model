package metric_test

import (
	"testing"

	"github.com/boundedinfinity/canonical_model/go/idiomatic/measurement/metric"
	"github.com/stretchr/testify/assert"
)

func Test_prefix_name_and_symbol(t *testing.T) {
	tcs := []struct {
		name           string
		input          metric.Prefix
		expectedSymbol string
		expectedName   string
		expectedFactor float64
	}{
		{
			name:           "kilo",
			input:          metric.Prefixes.Kilo,
			expectedSymbol: "k",
			expectedName:   "Kilo",
			expectedFactor: 1e3,
		},
		{
			name:           "unit",
			input:          metric.Prefixes.Unit,
			expectedSymbol: "",
			expectedName:   "",
			expectedFactor: 1,
		},
		{
			name:           "unknown",
			input:          metric.Prefixes.Unknown,
			expectedSymbol: "",
			expectedName:   "",
			expectedFactor: 0,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			assert.Equal(t, tc.expectedSymbol, metric.Prefixes.Symbol(tc.input))
			assert.Equal(t, tc.expectedName, metric.Prefixes.Name(tc.input))
			assert.Equal(t, tc.expectedFactor, metric.Prefixes.Factor(tc.input))
		})
	}
}
