package identifier

type Kind string

type kinds struct {
	Unknown            Kind
	Passport           Kind
	DriversLicense     Kind
	MilitaryIdentifier Kind
}

var Kinds = kinds{
	Unknown:            "government.identifier.unknown",
	Passport:           "government.identifier.passport",
	DriversLicense:     "government.identifier.drivers-license",
	MilitaryIdentifier: "government.identifier.military-identifier",
}
