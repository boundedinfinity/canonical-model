package property

type Kind string

var Kinds = kinds{
	RealEstate: "legal.property.real-estate",
	Money:      "legal.property.money",
	Security:   "legal.property.security",
	Valuables:  "legal.property.valuables",
}

type kinds struct {
	RealEstate Kind
	Money      Kind
	Security   Kind
	Valuables  Kind
}
