package multifactor_delivery

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/digital/address/email"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
)

var _ Delivery = &EmailModel{}

type EmailModel struct {
	Id    ider.Id     `json:"id"`
	Name  string      `json:"name"`
	Email email.Email `json:"email"`
}

func (this EmailModel) GetName() string {
	return this.Name
}

func (_ EmailModel) GetKind() Kind {
	return Kinds.Email
}
