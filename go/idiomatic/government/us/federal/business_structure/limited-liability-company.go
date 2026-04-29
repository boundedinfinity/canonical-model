package business_structure

// https://www.irs.gov/businesses/small-businesses-self-employed/limited-liability-company-llc

import "github.com/boundedinfinity/canonical-model/go/idiomatic/ider"

var _ BusinessStructure = &LimitedLiabilityCompany{}

type LimitedLiabilityCompany struct {
	Id   ider.Id `json:"id"`
	Name string  `json:"name"`
}

func (this LimitedLiabilityCompany) GetName() string {
	return this.Name
}

func (_ LimitedLiabilityCompany) GetKind() Kind {
	return Kinds.LimitedLiabilityCompany
}
