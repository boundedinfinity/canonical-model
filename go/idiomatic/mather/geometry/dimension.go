package geometry

import (
	"github.com/boundedinfinity/go-commoner/idiomatic"
)

type Dimension2d[T idiomatic.Number] struct {
	Height T
	Width  T
}

func NewDimension2d[T idiomatic.Number](height, width T) Dimension2d[T] {
	return Dimension2d[T]{
		Height: height,
		Width:  width,
	}
}
