package currency

import "github.com/boundedinfinity/canonical_model/idiomatic/model/specifications/iso/iso4217"

type Currency struct {
	Iso iso4217.Currency `json:"iso"`
}
