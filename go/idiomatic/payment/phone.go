package payment

import "github.com/boundedinfinity/canonical-model/go/idiomatic/phone"

type Phone struct {
	Phone phone.Phone `json:"phone,omitempty"`
}

var _ Payment = &Phone{}

func (_ Phone) Kind() Kind {
	return Kinds.Phone
}

func (this Phone) String() string {
	return "phone: " + this.Phone.Number.Last4Digits()
}
