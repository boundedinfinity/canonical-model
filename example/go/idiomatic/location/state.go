package location

import "github.com/boundedinfinity/schema/idiomatic/id"

type State struct {
	Id   id.Id     `json:"id,omitempty"`
	Name StateName `json:"name,omitempty"`
	Code StateAnsi `json:"code,omitempty"`
}

var _ id.TypeNamer = &State{}

func (t State) TypeName() string {
	return id.TypeNamers.Dotted(State{})
}

func (t State) String() string {
	return t.Name.String()
}
