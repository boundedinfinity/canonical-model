package person

import "github.com/boundedinfinity/canonical_model/idiomatic/ider"

type NameType struct {
	Id          ider.Id `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
}
