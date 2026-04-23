package amount

type Kind string

type kinds struct {
	Unknown Kind
	Mass    Kind
}

var Kinds = kinds{
	Unknown: "medical.amount.unknown",
	Mass:    "medical.amount.tablet",
}
