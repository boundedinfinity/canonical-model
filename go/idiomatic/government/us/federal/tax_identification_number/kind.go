package tax_identification_number

type Kind string

type kinds struct {
	Unknown                                  Kind
	SocialSecurityNumber                     Kind
	EmployerIdentificationNumber             Kind
	IndividualTaxpayerIdentificationNumber   Kind
	TaxpayerIdentificationNumberForAdaptions Kind
	PreparerTaxpayerIdentificationNumber     Kind
}

var Kinds = kinds{
	Unknown:                                  "tax-identification-number.unknown",
	SocialSecurityNumber:                     "tax-identification-number.socical-security-number",
	EmployerIdentificationNumber:             "tax-identification-number.employer-identification-number",
	IndividualTaxpayerIdentificationNumber:   "tax-identification-number.individual-taxpayer-identification-number",
	TaxpayerIdentificationNumberForAdaptions: "tax-identification-number.taxpayer-identification-number-for-adaptions",
	PreparerTaxpayerIdentificationNumber:     "tax-identification-number.preparer-taxpayer-identification-number",
}
