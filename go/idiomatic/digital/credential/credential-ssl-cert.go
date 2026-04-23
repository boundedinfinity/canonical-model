package credential

import "github.com/boundedinfinity/canonical_model/idiomatic/ider"

var _ Credential = &SslCertModel{}

type SslCertModel struct {
	Id         ider.Id `json:"id"`
	Name       string  `json:"name"`
	PublicKey  string  `json:"public-key"`
	PrivateKey string  `json:"private-key"`
	PassPhrase string  `json:"pass-phrase"`
}

func (this SslCertModel) GetName() string {
	return this.Name
}

func (_ SslCertModel) GetKind() Kind {
	return Kinds.SslCert
}
