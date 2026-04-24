package country

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/specification/iso/iso3166"
)

type Country struct {
	Iso iso3166.Country `json:"iso"`
}
