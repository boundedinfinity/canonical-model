package trust

// https://www.legalzoom.com/articles/types-of-trusts

type Kind string

var Kinds = kinds{
	TitleHolding: "legal.trust.title-holding",
	Conservation: "legal.trust.convervation",
	Community:    "legal.trust.community",
	Living:       "legal.trust.living",
}

type kinds struct {
	TitleHolding Kind
	Conservation Kind
	Community    Kind
	Living       Kind
}

type Type string

var Types = types{
	Revocable:   "legal.type.revocable",
	Irrevocable: "legal.type.irrevocable",
}

type types struct {
	Revocable   Type
	Irrevocable Type
}
