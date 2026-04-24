package multifactor_secret

import "github.com/boundedinfinity/canonical_model/go/idiomatic/ider"

var _ Secret = &OneTimePasswordModel{}

type OneTimePasswordModel struct {
	Id   ider.Id `json:"id"`
	Name string  `json:"name"`
}

func (this OneTimePasswordModel) GetName() string {
	return this.Name
}

func (_ OneTimePasswordModel) GetKind() Kind {
	return Kinds.OneTimePassword
}
