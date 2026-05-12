package payment

type PayPal struct {
	Account string `json:"account,omitempty"`
}

var _ Payment = &PayPal{}

func (_ PayPal) Kind() Kind {
	return Kinds.PayPal
}

func (this PayPal) String() string {
	return "PayPal: " + this.Account
}
