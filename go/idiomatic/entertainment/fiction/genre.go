package fiction

import "github.com/boundedinfinity/canonical-model/go/idiomatic/ider"

type Genre struct {
	Id          ider.Id `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
}
