package multifactor_secret

import "github.com/boundedinfinity/canonical_model/go/idiomatic/ider"

var _ Secret = &BiometricsFaceModel{}

type BiometricsFaceModel struct {
	Id   ider.Id `json:"id"`
	Name string  `json:"name"`
}

func (this BiometricsFaceModel) GetName() string {
	return this.Name
}

func (_ BiometricsFaceModel) GetKind() Kind {
	return Kinds.BiometricsFace
}
