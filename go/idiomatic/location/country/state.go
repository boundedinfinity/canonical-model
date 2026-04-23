package country

import "github.com/boundedinfinity/canonical_model/idiomatic/id"

type State struct {
	Id   id.Id  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`
}

var _ id.TypeNamer = &State{}

func (t State) TypeName() string {
	return id.TypeNamers.Dotted(State{})
}

func (t State) String() string {
	return t.Name
}
