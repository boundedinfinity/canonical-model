package business_structure

// https://www.irs.gov/businesses/small-businesses-self-employed/forming-a-corporation

import "github.com/boundedinfinity/canonical_model/go/idiomatic/ider"

var _ BusinessStructure = &Corporation{}

type Corporation struct {
	Id   ider.Id `json:"id"`
	Name string  `json:"name"`
}

func (this Corporation) GetName() string {
	return this.Name
}

func (_ Corporation) GetKind() Kind {
	return Kinds.Corporation
}
