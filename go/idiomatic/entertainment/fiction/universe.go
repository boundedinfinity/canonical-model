package fiction

import "github.com/boundedinfinity/canonical_model/go/idiomatic/ider"

type Universe struct {
	Id          ider.Id `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
}
