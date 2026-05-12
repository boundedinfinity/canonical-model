package payment_card

import "github.com/boundedinfinity/go-commoner/idiomatic/stringer"

type Number string

func (this Number) String() string {
	return string(this)
}

func (this Number) Last4Digits() string {
	return stringer.TruncateEnd(this, 4, "")
}
