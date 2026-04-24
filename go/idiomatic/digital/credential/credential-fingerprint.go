package credential

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
)

var _ Credential = &FingerprintModel{}

type FingerprintModel struct {
	Id   ider.Id `json:"id"`
	Name string  `json:"name"`
}

func (this FingerprintModel) GetName() string {
	return this.Name
}

func (_ FingerprintModel) GetKind() Kind {
	return Kinds.SshKey
}
