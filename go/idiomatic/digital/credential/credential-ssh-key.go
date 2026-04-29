package credential

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
)

var _ Credential = &SshKeyModel{}

type SshKeyModel struct {
	Id         ider.Id `json:"id"`
	Name       string  `json:"name"`
	PrivateKey string  `json:"private-key"`
	PassPhrase string  `json:"pass-phrase"`
}

func (this SshKeyModel) GetName() string {
	return this.Name
}

func (_ SshKeyModel) GetKind() Kind {
	return Kinds.SshKey
}
