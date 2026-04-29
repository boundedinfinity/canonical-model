package details

import "github.com/boundedinfinity/canonical-model/go/idiomatic/ider"

type ContentWarning struct {
	Id          ider.Id `json:"id"`
	Rating      string  `json:"rating"`
	Description string  `json:"description"`
}
