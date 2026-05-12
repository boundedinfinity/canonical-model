package payment

type WesternUnion struct {
	Account string `json:"account,omitempty"`
}

var _ Payment = &WesternUnion{}

func (_ WesternUnion) Kind() Kind {
	return Kinds.WesternUnion
}

func (this WesternUnion) String() string {
	return "western union: " + this.Account
}
