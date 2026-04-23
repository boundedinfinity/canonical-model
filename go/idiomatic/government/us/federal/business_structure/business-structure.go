package business_structure

// https://www.irs.gov/tin/taxpayer-identification-numbers-tin

type BusinessStructure interface {
	GetKind() Kind
	GetName() string
}
