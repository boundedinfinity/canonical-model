package percent

import (
	"github.com/boundedinfinity/go-commoner/idiomatic"
	"github.com/boundedinfinity/go-commoner/idiomatic/mather/fraction"
)

func FromFloat[T ~int, F idiomatic.Float](v F) Percent[T] {
	return Percent[T]{
		fraction: fraction.FromFloat[T](v),
	}
}

type Percent[T ~int] struct {
	fraction fraction.Fraction[T]
	size     int
}

func (this Percent[T]) Percentage() float64 {
	return this.fraction.Float()
}

func (this Percent[T]) Decimal() float64 {
	return this.fraction.Float() / 100
}
