package physical_product

import "github.com/boundedinfinity/go-commoner/idiomatic/math/rational"

type ProductDimensions struct {
	Width  rational.Rational `json:"width,omitempty"`
	Heigth rational.Rational `json:"heigth,omitempty"`
	Depth  rational.Rational `json:"depth,omitempty"`
	Length rational.Rational `json:"length,omitempty"`
}
