package federal

import "github.com/boundedinfinity/canonical-model/go/idiomatic/ider"

type EntityModel struct {
	Id   ider.Id `json:"id"`
	Name string  `json:"name"`
}
