package business_structure

type Kind string

type kinds struct {
	Unknown                 Kind
	SoleProprietorship      Kind
	Partnership             Kind
	Corporation             Kind
	SCorporation            Kind
	LimitedLiabilityCompany Kind
}

var Kinds = kinds{
	Unknown:                 "business-structure.unknown",
	SoleProprietorship:      "business-structure.sole-proprietorship",
	Partnership:             "business-structure.partnership",
	Corporation:             "business-structure.corporation",
	SCorporation:            "business-structure.s-corporation",
	LimitedLiabilityCompany: "business-structure.limited-liability-company",
}
