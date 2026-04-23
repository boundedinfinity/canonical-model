package perscription

type FormFactor string

type formFactors struct {
	Unknown FormFactor
	Tablet  FormFactor
	Capsul  FormFactor
}

var FormFactors = formFactors{
	Unknown: "medical.perscription.unknown",
	Tablet:  "medical.perscription.tablet",
	Capsul:  "medical.perscription.capsul",
}
