package numberer

import (
	"strconv"
	"strings"

	"github.com/boundedinfinity/go-commoner/errorer"
)

var (
	ErrNumber   = errorer.New("number error")
	errNumberFn = errorer.Func(ErrNumber)
)

type Number interface {
	IsZero() bool
	ToFloat() Number
	ToFraction() Number
	ToMixed() Number
	ToImproper() Number
	IsProper() bool
	IsImproper() bool
	Reciprocal() Number
	Float() float64
}

func New(n string) (Number, error) {
	var number Number
	var err error
	errUnknownFormat := errNumberFn("unknown number format '%s'", n)

	n = strings.TrimSpace(n)
	containsPeriod := strings.Contains(n, ".")
	containsSlash := strings.Contains(n, "/")
	containsSpace := strings.Contains(n, " ")

	parse := func(x string) (int, error) {
		if i, err := strconv.Atoi(x); err != nil {
			return i, errNumberFn("cannot parse integer '%s' from '%s'", x, n)
		} else {
			return i, nil
		}
	}

	switch {
	case containsPeriod && !containsSlash && !containsSpace:
		if f, err := strconv.ParseFloat(n, 64); err != nil {
			return nil, errNumberFn("cannot parse float '%s': %w", n, err)
		} else {
			number = &FloatNumber{value: f}
		}
	case !containsPeriod && containsSlash && containsSpace:
		mixed := MixedNumber{}

		numberParts := strings.Split(n, " ")
		if len(numberParts) != 2 {
			return nil, errUnknownFormat
		}

		if mixed.Whole, err = parse(numberParts[0]); err != nil {
			return nil, err
		}

		fractionsParts := strings.Split(numberParts[1], "/")
		if len(fractionsParts) != 2 {
			return nil, errUnknownFormat
		}

		if mixed.Fraction.Numerator, err = parse(fractionsParts[0]); err != nil {
			return nil, err
		}

		if mixed.Fraction.Denominator, err = parse(fractionsParts[1]); err != nil {
			return nil, err
		}

		number = &mixed
	case !containsPeriod && containsSlash && !containsSpace:
		fraction := FractionNumber{}

		fractionsParts := strings.Split(n, "/")
		if len(fractionsParts) != 2 {
			return nil, errUnknownFormat
		}

		if fraction.Numerator, err = parse(fractionsParts[0]); err != nil {
			return nil, err
		}
		if fraction.Denominator, err = parse(fractionsParts[1]); err != nil {
			return nil, err
		}

		number = &fraction
	default:
		return nil, errUnknownFormat
	}

	return number, nil
}
