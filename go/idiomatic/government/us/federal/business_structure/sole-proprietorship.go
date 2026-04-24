package business_structure

// https://www.irs.gov/businesses/small-businesses-self-employed/sole-proprietorships

import "github.com/boundedinfinity/canonical_model/go/idiomatic/ider"

var _ BusinessStructure = &SoleProprietorship{}

type SoleProprietorship struct {
	Id   ider.Id `json:"id"`
	Name string  `json:"name"`
}

func (this SoleProprietorship) GetName() string {
	return this.Name
}

func (_ SoleProprietorship) GetKind() Kind {
	return Kinds.SoleProprietorship
}
