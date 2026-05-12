package payment

import "github.com/boundedinfinity/canonical-model/go/idiomatic/banking/account"

type PrivacyCom struct {
	Account account.Account `json:"account,omitempty"`
}

var _ Payment = &PrivacyCom{}

func (_ PrivacyCom) Kind() Kind {
	return Kinds.PrivacyCom
}

func (this PrivacyCom) String() string {
	return "privacy.com: " + this.Account.Last4Digits()
}
