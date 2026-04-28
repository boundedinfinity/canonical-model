package geometry

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/mather"
)

type Dimension2d[T mather.Number] struct {
	Height T
	Width  T
}

func NewDimension2d[T mather.Number](height, width T) Dimension2d[T] {
	return Dimension2d[T]{
		Height: height,
		Width:  width,
	}
}
