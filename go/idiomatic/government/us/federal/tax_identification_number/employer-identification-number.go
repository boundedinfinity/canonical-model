package tax_identification_number

// https://www.irs.gov/businesses/employer-identification-number

import "github.com/boundedinfinity/canonical_model/idiomatic/ider"

var _ TaxIdentificationNumber = &EmployerIdentificationNumber{}

type EmployerIdentificationNumber struct {
	Id   ider.Id `json:"id"`
	Name string  `json:"name"`
}

func (this EmployerIdentificationNumber) GetName() string {
	return this.Name
}

func (_ EmployerIdentificationNumber) GetKind() Kind {
	return Kinds.EmployerIdentificationNumber
}
