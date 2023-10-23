package usa

import "github.com/boundedinfinity/schema/idiomatic/id"

// https://www.irs.gov/pub/irs-pdf/fw4.pdf

type FormW4 struct {
	Id id.Id `json:"id,omitempty"`
}
