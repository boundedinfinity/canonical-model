package login

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/digital/address/email"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/digital/address/uri"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/phone"
)

type Authenticator struct {
	Id       ider.Id `json:"id"`
	Name     string  `json:"name"`
	Location uri.Url `json:"location"`
}

type TwoFactorMethod struct {
	Id            ider.Id            `json:"id"`
	Authenticator Authenticator      `json:"authenticator"`
	Phone         []phone.PhoneModel `json:"phone"`
	Email         []email.Email      `json:"email"`
}
