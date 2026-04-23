package fraction

import (
	"fmt"
	"strconv"

	"github.com/boundedinfinity/go-commoner/idiomatic"
	"github.com/boundedinfinity/go-commoner/idiomatic/mather"
	"github.com/boundedinfinity/go-commoner/idiomatic/numberer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

// ----------------------------------------------------------------------------------------------------
// Constructors
// ----------------------------------------------------------------------------------------------------

func New[T ~int](numerator, denominator T) Fraction[T] {
	return Fraction[T]{
		Numerator:   numerator,
		Denominator: denominator,
	}
}

// FromFloat creates a Fraction from a floating point number.
func FromFloat[T ~int, F idiomatic.Float](n F) Fraction[T] {
	denominator := mather.FractionalPartPlace(n)
	numerator := n * F(denominator)

	return Fraction[T]{
		Numerator:   T(numerator),
		Denominator: T(denominator),
	}
}

// ----------------------------------------------------------------------------------------------------
// Type
// ----------------------------------------------------------------------------------------------------

type Fraction[T ~int] struct {
	Numerator   T
	Denominator T
}

func (t Fraction[T]) String() string {
	return fmt.Sprintf("%d/%d", t.Numerator, t.Denominator)
}

func (t Fraction[T]) Float() float64 {
	return float64(t.Numerator / t.Denominator)
}

func (t Fraction[T]) Copy() Fraction[T] {
	return Fraction[T]{
		Numerator:   t.Numerator,
		Denominator: t.Denominator,
	}
}

func (t Fraction[T]) Reduce() Fraction[T] {
	gcd := numberer.GreatestCommonFactor(t.Numerator, t.Denominator)

	return Fraction[T]{
		Numerator:   t.Numerator / gcd,
		Denominator: t.Denominator / gcd,
	}
}

func (t Fraction[T]) IsZero() bool {
	return IsZero(t)
}

func (t Fraction[T]) Enumerate(l, h T) []Fraction[T] {
	var items []Fraction[T]
	item := t.Reduce()
	l = mather.Max(l, item.Denominator)

	for i := l; i <= h; i <<= 1 {
		items = append(items, item)
		item = New(item.Numerator*2, item.Denominator*2)
	}

	return items
}

// ----------------------------------------------------------------------------------------------------
// Helpers
// ----------------------------------------------------------------------------------------------------

func IsZero[T ~int](elem Fraction[T]) bool {
	var zero Fraction[T]
	return elem == zero
}

func Component[T idiomatic.Float](n T) int {
	s := fmt.Sprintf("%v", n)
	comps := stringer.Split(s, ".")
	var i int

	if len(comps) > 1 {
		s = comps[1]
		x, err := strconv.Atoi(s)

		if err != nil {
			panic(err)
		}

		i = x
	} else {
		i = 0
	}

	return i
}

func Magnitude[T idiomatic.Float](n T) int {
	s := fmt.Sprintf("%v", n)
	comps := stringer.Split(s, ".")
	var size int

	if len(comps) > 1 {
		s = stringer.Split(s, ".")[1]
		size = len(s)
	} else {
		s = "0"
		size = 0
	}

	return size
}
