package credential

import "github.com/boundedinfinity/canonical_model/go/idiomatic/ider"

var _ Credential = &WebAuthnModel{}

type WebAuthnModel struct {
	Id   ider.Id `json:"id"`
	Name string  `json:"name"`
}

func (this WebAuthnModel) GetName() string {
	return this.Name
}

func (_ WebAuthnModel) GetKind() Kind {
	return Kinds.Passkey
}
