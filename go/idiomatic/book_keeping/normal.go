package book_keeping

type Kind string

type kinds struct {
	Unknown Kind
	Debit   Kind
	Credit  Kind
}

var Kinds = kinds{
	Unknown: "unknown",
	Debit:   "debit",
	Credit:  "credit",
}
