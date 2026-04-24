package tax_identification_number

// https://www.irs.gov/tin/taxpayer-identification-numbers-tin#ptin

import "github.com/boundedinfinity/canonical_model/go/idiomatic/ider"

var _ TaxIdentificationNumber = &PreparerTaxpayerIdentificationNumber{}

type PreparerTaxpayerIdentificationNumber struct {
	Id   ider.Id `json:"id"`
	Name string  `json:"name"`
}

func (this PreparerTaxpayerIdentificationNumber) GetName() string {
	return this.Name
}

func (_ PreparerTaxpayerIdentificationNumber) GetKind() Kind {
	return Kinds.PreparerTaxpayerIdentificationNumber
}
