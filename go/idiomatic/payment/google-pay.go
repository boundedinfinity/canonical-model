package payment

type GooglePay struct {
	Account string `json:"account,omitempty"`
}

var _ Payment = &GooglePay{}

func (_ GooglePay) Kind() Kind {
	return Kinds.GooglePay
}

func (this GooglePay) String() string {
	return "google pay: " + this.Account
}
