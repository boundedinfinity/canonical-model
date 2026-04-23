package professional

import "github.com/boundedinfinity/canonical_model/idiomatic/ider"

type Profession struct {
	Id          ider.Id `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
}
