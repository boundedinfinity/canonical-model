package business_structure

// https://www.irs.gov/businesses/small-businesses-self-employed/s-corporations

import "github.com/boundedinfinity/canonical-model/go/idiomatic/ider"

var _ BusinessStructure = &SCorporation{}

type SCorporation struct {
	Id   ider.Id `json:"id"`
	Name string  `json:"name"`
}

func (this SCorporation) GetName() string {
	return this.Name
}

func (_ SCorporation) GetKind() Kind {
	return Kinds.SCorporation
}
