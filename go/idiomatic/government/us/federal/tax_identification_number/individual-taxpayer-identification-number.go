package tax_identification_number

// https://www.irs.gov/tin/taxpayer-identification-numbers-tin#itin

import "github.com/boundedinfinity/canonical-model/go/idiomatic/ider"

var _ TaxIdentificationNumber = &IndividualTaxpayerIdentificationNumber{}

type IndividualTaxpayerIdentificationNumber struct {
	Id   ider.Id `json:"id"`
	Name string  `json:"name"`
}

func (this IndividualTaxpayerIdentificationNumber) GetName() string {
	return this.Name
}

func (_ IndividualTaxpayerIdentificationNumber) GetKind() Kind {
	return Kinds.IndividualTaxpayerIdentificationNumber
}
