package multifactor_secret

import "github.com/boundedinfinity/canonical_model/idiomatic/ider"

var _ Secret = &QuickResponseCodeModel{}

type QuickResponseCodeModel struct {
	Id   ider.Id `json:"id"`
	Name string  `json:"name"`
}

func (this QuickResponseCodeModel) GetName() string {
	return this.Name
}

func (_ QuickResponseCodeModel) GetKind() Kind {
	return Kinds.QuickResponseCode
}
