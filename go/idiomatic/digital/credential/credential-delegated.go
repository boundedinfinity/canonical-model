package credential

import "github.com/boundedinfinity/canonical_model/go/idiomatic/ider"

var _ Credential = &Delegated{}

type Delegated struct {
	Id         ider.Id    `json:"id"`
	Name       string     `json:"name"`
	Credential Credential `json:"credential"`
}

func (this Delegated) GetName() string {
	return this.Name
}

func (_ Delegated) GetKind() Kind {
	return Kinds.Delegated
}
