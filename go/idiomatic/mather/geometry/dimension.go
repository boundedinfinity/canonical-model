package geometry

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/numberer"
)

type Dimension2d struct {
	Height numberer.Number
	Width  numberer.Number
}

func NewDimension2d(height, width numberer.Number) Dimension2d {
	return Dimension2d{
		Height: height,
		Width:  width,
	}
}
