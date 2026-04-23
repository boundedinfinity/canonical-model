package listing_site

import "github.com/boundedinfinity/canonical_model/idiomatic/ider"

type ListingSite struct {
	Id   ider.Id `json:"id"`
	Name string  `json:"name"`
}
