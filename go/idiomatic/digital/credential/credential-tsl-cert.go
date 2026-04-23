package credential

import "github.com/boundedinfinity/canonical_model/idiomatic/ider"

var _ Credential = &TslCertModel{}

type TslCertModel struct {
	Id         ider.Id `json:"id"`
	Name       string  `json:"name"`
	PublicKey  string  `json:"public-key"`
	PrivateKey string  `json:"private-key"`
	PassPhrase string  `json:"pass-phrase"`
}

func (this TslCertModel) GetName() string {
	return this.Name
}

func (_ TslCertModel) GetKind() Kind {
	return Kinds.TlsCert
}
