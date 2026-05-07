package trustee

type Kind string

var Kinds = kinds{
	IndividualPerson:      "legal.trustee.individual-person",
	MultiplePeople:        "legal.trustee.multiple-people",
	Business:              "legal.trustee.business",
	ProfessionalFiduciary: "legal.trustee.professional-fiduciary",
}

type kinds struct {
	IndividualPerson      Kind
	MultiplePeople        Kind
	Business              Kind
	ProfessionalFiduciary Kind
}
