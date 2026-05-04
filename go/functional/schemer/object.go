package schemer

import (
	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

var _ Schemer = &Object{}

type Object struct {
	Name       optioner.Option[string]
	Properties []Property
}

func (this Object) Kind() Kind {
	return Kinds.Object
}

type Property struct {
	Name     optioner.Option[string]
	Optional optioner.Option[bool]
}
