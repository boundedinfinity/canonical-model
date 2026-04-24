package phone

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/id"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
)

type Fax struct {
	Id     ider.Id    `json:"id,omitempty"`
	Number NanpNumber `json:"number,omitempty"`
}

var _ id.TypeNamer = &Fax{}

func (t Fax) TypeName() string {
	return id.TypeNamers.Dotted(Fax{})
}

func (t Fax) String() string {
	return t.Number.String()
}
