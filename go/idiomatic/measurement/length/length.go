package length

import (
	"strconv"

	"github.com/boundedinfinity/canonical-model/go/idiomatic/measurement/imperial"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/measurement/metric"
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
	s = stringer.TrimSpace(s)
	s = stringer.CompactSpace(s)
	comps := stringer.Split(s, " ")

	if len(comps) != 2 {
		return nil, errLengthInvalidFn("%s does not look like a valid length", s)
	}

	var amount float64
	var err error

	if amount, err = strconv.ParseFloat(comps[0], 64); err != nil {
		return nil, errLengthInvalidFn("%s does not look like a valid number: %w", comps[0], err)
	}

	if kind, ok := imperial.Lengths.ParseOk(comps[1]); ok {
		return ImperialLength{Amount: amount, Kind: kind}, nil
	}

	if prefix, ok := metric.Lengths.ParseOk(comps[1]); ok {
		return metricLength{Amount: amount, Prefix: prefix}, nil
	}

	return nil, errLengthInvalidFn("%s does not look like a valid metric or imperial length", comps[1])
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	inchesToMillimeters = 25.4
	millimetersToInches = 1 / inchesToMillimeters
)
