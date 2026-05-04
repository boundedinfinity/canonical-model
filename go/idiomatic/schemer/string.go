package schemer

var _ Schemer = &String{}

type String struct {
	Name string
}

func (this String) Kind() Kind {
	return Kinds.String
}
