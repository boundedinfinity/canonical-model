package ratio

import "github.com/boundedinfinity/go-commoner/idiomatic/mather/fraction"

type Ratio[T ~int] struct {
	fraction.Fraction[T]
}
