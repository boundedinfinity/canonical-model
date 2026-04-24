package credential

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/value/email"
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
