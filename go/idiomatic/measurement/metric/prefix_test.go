package metric_test

import (
	"testing"

	"github.com/boundedinfinity/canonical-model/go/idiomatic/measurement/metric"
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

func Test_Metric_Prefix_Convert(t *testing.T) {
	tcs := []struct {
		name     string
		input    float64
		expected float64
		from     metric.Prefix
		to       metric.Prefix
	}{
		{
			name:     "1 kilo to mega",
			input:    1,
			expected: 0.001,
			from:     metric.Prefixes.Kilo,
			to:       metric.Prefixes.Mega,
		},
		{
			name:     "1 mega to kilo",
			input:    1,
			expected: 1000,
			from:     metric.Prefixes.Mega,
			to:       metric.Prefixes.Kilo,
		},
		{
			name:     "1 unit to kilo",
			input:    1,
			expected: 0.001,
			from:     metric.Prefixes.Unit,
			to:       metric.Prefixes.Kilo,
		},
		{
			name:     "1 kilo to unit",
			input:    1,
			expected: 1000,
			from:     metric.Prefixes.Kilo,
			to:       metric.Prefixes.Unit,
		},
		{
			name:     "1 unit to milli",
			input:    1,
			expected: 1000,
			from:     metric.Prefixes.Unit,
			to:       metric.Prefixes.Milli,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			actual := metric.Prefixes.Convert(tc.input, tc.from, tc.to)
			assert.InDelta(t, tc.expected, actual, 0.0001)
		})
	}
}
