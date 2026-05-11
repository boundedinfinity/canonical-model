package business_identifier

type BusinessIdentifer interface {
	Identifier() string
	Kind() Kind
}
