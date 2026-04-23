package business_structure

// https://www.irs.gov/businesses/partnerships

import "github.com/boundedinfinity/canonical_model/idiomatic/ider"

var _ BusinessStructure = &Partnership{}

type Partnership struct {
	Id   ider.Id `json:"id"`
	Name string  `json:"name"`
}

func (this Partnership) GetName() string {
	return this.Name
}

func (_ Partnership) GetKind() Kind {
	return Kinds.Partnership
}
