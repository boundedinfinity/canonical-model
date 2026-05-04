package schemer

import (
	"regexp"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

var _ Schemer = &String{}

type String struct {
	Name  optioner.Option[string]
	Min   optioner.Option[int]
	Max   optioner.Option[int]
	Regex optioner.Option[*regexp.Regexp]
}

func (this String) Kind() Kind {
	return Kinds.String
}
