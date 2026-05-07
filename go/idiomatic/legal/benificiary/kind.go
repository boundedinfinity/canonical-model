package benificiary

type Kind string

var Kinds = kinds{
	IndividualPerson: "legal.benificiary.individual-person",
	MultiplePeople:   "legal.benificiary.multiple-people",
	Business:         "legal.benificiary.business",
	Trust:            "legal.benificiary.trust",
	Charity:          "legal.benificiary.charity",
	NonProfit:        "legal.benificiary.non-profit",
	MinorPerson:      "legal.benificiary.minor-person",
	Contingent:       "legal.benificiary.contingent",
	Descendant:       "legal.benificiary.descendant",
	Pet:              "legal.benificiary.pet",
}

type kinds struct {
	IndividualPerson Kind
	MultiplePeople   Kind
	Business         Kind
	Trust            Kind
	Charity          Kind
	NonProfit        Kind
	MinorPerson      Kind
	Contingent       Kind
	Descendant       Kind
	Pet              Kind
}
