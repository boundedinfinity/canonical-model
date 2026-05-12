package payment

type Venmo struct {
	Account string `json:"account,omitempty"`
}

var _ Payment = &Venmo{}

func (_ Venmo) Kind() Kind {
	return Kinds.Venmo
}

func (this Venmo) String() string {
	return "venmo: " + this.Account
}
