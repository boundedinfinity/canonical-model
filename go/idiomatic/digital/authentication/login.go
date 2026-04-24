package login

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/digital/address/email"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/digital/address/uri"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/digital/credential"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/digital/phone"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
)

type Login struct {
	Id                ider.Id               `json:"id"`
	Name              string                `json:"name"`
	Credential        credential.Credential `json:"credential"`
	Description       string                `json:"description"`
	SecurityQuestions []SecurityQuestion    `json:"security-questions"`
	LoginLocation     uri.Url               `json:"login-location"`
	AccountEmail      email.Email           `json:"account-email"`
	AccountPhone      phone.PhoneModel      `json:"account-phone"`
	TwoFactorMethod   TwoFactorMethod       `json:"two-factor-method"`
}
