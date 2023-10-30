package measurement

import "github.com/boundedinfinity/go-commoner/idiomatic/math/rational"

type Dimensions struct {
	Width    rational.Rational `json:"width,omitempty"`
	Heigth   rational.Rational `json:"heigth,omitempty"`
	Depth    rational.Rational `json:"depth,omitempty"`
	Length   rational.Rational `json:"length,omitempty"`
	Radius   rational.Rational `json:"radius,omitempty"`
	Diameter rational.Rational `json:"diameter,omitempty"`
}
