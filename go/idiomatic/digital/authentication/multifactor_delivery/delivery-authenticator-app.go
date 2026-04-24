package multifactor_delivery

import "github.com/boundedinfinity/canonical_model/go/idiomatic/ider"

var _ Delivery = &AuthenticatorAppModel{}

type AuthenticatorAppModel struct {
	Id   ider.Id `json:"id"`
	Name string  `json:"name"`
}

func (this AuthenticatorAppModel) GetName() string {
	return this.Name
}

func (_ AuthenticatorAppModel) GetKind() Kind {
	return Kinds.AuthenticatorApplication
}
