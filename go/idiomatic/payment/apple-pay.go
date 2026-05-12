package payment

type ApplePay struct {
	Account string `json:"account,omitempty"`
}

var _ Payment = &ApplePay{}

func (_ ApplePay) Kind() Kind {
	return Kinds.ApplePay
}

func (this ApplePay) String() string {
	return "ApplePay: " + this.Account
}
