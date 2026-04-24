package fiction

import "github.com/boundedinfinity/canonical_model/go/idiomatic/ider"

type Series struct {
	Id          ider.Id `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
}
