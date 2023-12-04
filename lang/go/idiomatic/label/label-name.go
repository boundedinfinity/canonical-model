package label

import (
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type LabelName struct {
	Id           id.Id  `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Abbreviation string `json:"abbreviation,omitempty"`
	Description  string `json:"description,omitempty"`
}

var _ id.TypeNamer = &LabelName{}

func (t LabelName) TypeName() string {
	return id.TypeNamers.Dotted(LabelName{})
}
