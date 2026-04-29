package credential

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
)

var _ Credential = &UsernamePasswordModel{}

type UsernamePasswordModel struct {
	Id       ider.Id `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Password string  `json:"password"`
}

func (this UsernamePasswordModel) GetName() string {
	return this.Name
}

func (_ UsernamePasswordModel) GetKind() Kind {
	return Kinds.UsernamePassword
}
