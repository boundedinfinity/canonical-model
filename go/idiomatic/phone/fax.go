package phone

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/identifier"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
)

type Fax struct {
	Id ider.Id `json:"id,omitempty"`
	// Number NanpNumber `json:"number,omitempty"`
}

var _ identifier.TypeNamer = &Fax{}

func (t Fax) TypeName() string {
	return identifier.TypeNamers.Dotted(Fax{})
}

func (t Fax) String() string {
	return ""
	// return t.Number.String()
}
