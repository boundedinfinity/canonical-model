package credential

import (
	"github.com/boundedinfinity/canonical_model/idiomatic/ider"
	"github.com/labstack/gommon/email"
)

var _ Credential = &EmailPassword{}

type EmailPassword struct {
	Id    ider.Id     `json:"id"`
	Name  string      `json:"name"`
	Email email.Email `json:"email"`
}

func (this EmailPassword) GetName() string {
	return this.Name
}

func (_ EmailPassword) GetKind() Kind {
	return Kinds.EmailPassword
}
