package form

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
)

// https://www.irs.gov/forms-pubs/about-form-w-4
// https://www.irs.gov/pub/irs-pdf/fw4.pdf

type FormW4 struct {
	Id ider.Id `json:"id,omitempty"`
}
