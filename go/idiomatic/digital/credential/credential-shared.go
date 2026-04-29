package credential

import "github.com/boundedinfinity/canonical-model/go/idiomatic/ider"

var _ Credential = &Shared{}

type Shared struct {
	Id         ider.Id    `json:"id"`
	Name       string     `json:"name"`
	Credential Credential `json:"credential"`
}

func (this Shared) GetName() string {
	return this.Name
}

func (_ Shared) GetKind() Kind {
	return Kinds.Shared
}
