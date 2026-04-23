package length

import (
	"strconv"

	"github.com/boundedinfinity//model/measurement/imperial"
	"github.com/boundedinfinity/canonical_model/idiomatic/measurement/metric"
	"github.com/boundedinfinity/go-commoner/idiomatic/errorer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

type Length interface {
	GetAmount() float64
	Imperial(imperial.Length) Length
	Metric(metric.Prefix) Length
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	ErrLengthInvalid   = errorer.New("invalid length")
	errLengthInvalidFn = errorer.Func(ErrLengthInvalid)
)

func Parse(s string) (Length, error) {
	comps := stringer.Split(s, "")

	if len(comps) != 2 {
		return nil, nil
	}

	for i := range len(comps) {
		comps[i] = stringer.TrimSpace(comps[i])
	}

	var amount float64
	var err error

	if amount, err = strconv.ParseFloat(comps[0], 64); err != nil {
		return nil, errLengthInvalidFn("%s does not look like a valid number: %w", comps[0], err)
	}

	if kind, ok := imperial.Lengths.IsSymbol(comps[1]); ok {
		return imperialLength{Amount: amount, Kind: kind}, nil
	}

	if prefix, ok := metric.Lengths.IsSymbol(comps[1]); ok {
		return metricLength{Amount: amount, Prefix: prefix}, nil
	}

	return nil, errLengthInvalidFn("%s does not look like a valid metric or imperial unit", comps[1])
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	inchToMillimeters   = 25.4
	millimetersToInches = 1 / inchToMillimeters
)
