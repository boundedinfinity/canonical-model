package country

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/identifier"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
)

type State struct {
	Id   ider.Id `json:"id,omitempty"`
	Name string  `json:"name,omitempty"`
	Code string  `json:"code,omitempty"`
}

var _ identifier.TypeNamer = &State{}

func (t State) TypeName() string {
	return identifier.TypeNamers.Dotted(State{})
}

func (t State) String() string {
	return t.Name
}
