package usa

import "github.com/boundedinfinity/schema/idiomatic/id"

// TIN
// https://www.irs.gov/individuals/individual-taxpayer-identification-number

type TaxIdentificationNumber interface {
	Type() string
}

// ITIN

type IndividualTaxpayerIdenficationNumber struct {
	Id     id.Id  `json:"id,omitempty"`
	Number string `json:"number,omitempty"`
}

func (t IndividualTaxpayerIdenficationNumber) Type() string {
	return "individual-taxpayer-identification-number"
}

var _ TaxIdentificationNumber = &IndividualTaxpayerIdenficationNumber{}

// EIN

type EmployerIdentificationNumber struct {
	Id     id.Id  `json:"id,omitempty"`
	Number string `json:"number,omitempty"`
}

func (t EmployerIdentificationNumber) Type() string {
	return "empoyer-identification-number"
}

var _ TaxIdentificationNumber = &EmployerIdentificationNumber{}
