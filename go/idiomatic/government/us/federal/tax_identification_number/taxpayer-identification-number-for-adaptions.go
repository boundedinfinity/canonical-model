package tax_identification_number

// https://www.irs.gov/individuals/adoption-taxpayer-identification-number

import "github.com/boundedinfinity/canonical_model/go/idiomatic/ider"

var _ TaxIdentificationNumber = &TaxpayerIdentificationNumberForAdaptions{}

type TaxpayerIdentificationNumberForAdaptions struct {
	Id   ider.Id `json:"id"`
	Name string  `json:"name"`
}

func (this TaxpayerIdentificationNumberForAdaptions) GetName() string {
	return this.Name
}

func (_ TaxpayerIdentificationNumberForAdaptions) GetKind() Kind {
	return Kinds.TaxpayerIdentificationNumberForAdaptions
}
