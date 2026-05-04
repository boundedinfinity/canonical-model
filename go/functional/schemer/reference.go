package schemer

import "github.com/boundedinfinity/go-commoner/functional/optioner"

var _ Schemer = &Reference{}

type Reference struct {
	Name      optioner.Option[string]
	Reference Schemer
}

func (this Reference) Kind() Kind {
	return Kinds.Object
}
