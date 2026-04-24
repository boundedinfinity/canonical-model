package tax_identification_number

// https://www.irs.gov/tin/taxpayer-identification-numbers-tin

type TaxIdentificationNumber interface {
	GetKind() Kind
	GetName() string
}

// TIN
// https://www.irs.gov/individuals/individual-taxpayer-identification-number

// // ITIN

// type IndividualTaxpayerIdenficationNumber struct {
// 	Id     ider.Id  `json:"id,omitempty"`
// 	Number string `json:"number,omitempty"`
// }

// func (t IndividualTaxpayerIdenficationNumber) Type() string {
// 	return "individual-taxpayer-identification-number"
// }

// var _ TaxIdentificationNumber = &IndividualTaxpayerIdenficationNumber{}

// // EIN

// type EmployerIdentificationNumber struct {
// 	Id     ider.Id  `json:"id,omitempty"`
// 	Number string `json:"number,omitempty"`
// }

// func (t EmployerIdentificationNumber) Type() string {
// 	return "empoyer-identification-number"
// }

// var _ TaxIdentificationNumber = &EmployerIdentificationNumber{}
