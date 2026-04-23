package rational

import (
	"strconv"

	"github.com/boundedinfinity/go-commoner/idiomatic"
	"github.com/boundedinfinity/go-commoner/idiomatic/mather/fraction"
)

// ----------------------------------------------------------------------------------------------------
// Constructors
// ----------------------------------------------------------------------------------------------------

func New[T ~int](whole, numerator, denominator T) Rational[T] {
	return Rational[T]{
		Whole:    whole,
		Fraction: fraction.New[T](numerator, denominator),
	}
}

func FromString[T ~int, S ~string](s S) (Rational[T], error) {
	n, err := strconv.ParseFloat(string(s), 64)

	if err != nil {
		return Rational[T]{}, err
	}

	return FromFloat[T](n), nil
}

// ----------------------------------------------------------------------------------------------------
// Type
// ----------------------------------------------------------------------------------------------------

type Rational[T ~int] struct {
	Whole    T
	Fraction fraction.Fraction[T]
}

func (t Rational[T]) String() string {
	var result string

	if t.Whole != 0 {
		result = strconv.Itoa(int(t.Whole))
	}

	if !t.Fraction.IsZero() {
		if result != "" {
			result += " "
		}

		result += t.Fraction.String()
	}

	return result
}

func (t Rational[T]) Copy() Rational[T] {
	return Rational[T]{
		Whole:    t.Whole,
		Fraction: t.Fraction.Copy(),
	}
}

func (t Rational[T]) Float() float64 {
	return float64(t.Whole) + t.Fraction.Float()
}

func (t Rational[T]) Reduce() Rational[T] {
	return Rational[T]{
		Whole:    t.Whole,
		Fraction: t.Fraction.Reduce(),
	}
}

func (t Rational[T]) Mixed() Rational[T] {
	if t.Fraction.Numerator < t.Fraction.Denominator {
		return t.Copy()
	}

	whole := T(t.Fraction.Numerator / t.Fraction.Denominator)
	numerator := t.Fraction.Numerator % t.Fraction.Denominator

	return Rational[T]{
		Whole: t.Whole + whole,
		Fraction: fraction.Fraction[T]{
			Numerator:   numerator,
			Denominator: t.Fraction.Denominator,
		},
	}
}

func (t Rational[T]) Improper() Rational[T] {
	if t.Whole == 0 {
		return t.Copy()
	}

	return Rational[T]{
		Whole: 0,
		Fraction: fraction.Fraction[T]{
			Numerator:   t.Fraction.Denominator*t.Whole + t.Fraction.Numerator,
			Denominator: t.Fraction.Denominator,
		},
	}
}

// ----------------------------------------------------------------------------------------------------
// Helpers
// ----------------------------------------------------------------------------------------------------

func IsZero[T ~int](elem Rational[T]) bool {
	var zero Rational[T]
	return elem == zero
}

func MustString[T ~int, S ~string](s S) Rational[T] {
	n, err := strconv.ParseFloat(string(s), 64)

	if err != nil {
		panic(err)
	}

	return FromFloat[T](n)
}

func FromFloat[T ~int, F idiomatic.Float](n F) Rational[T] {
	return Rational[T]{
		Whole:    T(n),
		Fraction: fraction.FromFloat[T](n),
	}
}

func Whole[T idiomatic.Float](x T) int {
	return int(x)
}

func FractionComponent[T idiomatic.Float](x T) int {
	return fraction.Component(x)
}

func Fraction[T ~int, F idiomatic.Float](x F) fraction.Fraction[T] {
	return fraction.FromFloat[T](x)
}
