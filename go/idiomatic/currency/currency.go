package currency

import "github.com/boundedinfinity/canonical_model/go/idiomatic/specification/iso/iso4217"

type Currency struct {
	Iso iso4217.Currency `json:"iso"`
}
