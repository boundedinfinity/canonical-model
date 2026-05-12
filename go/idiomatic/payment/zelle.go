package payment

type Zelle struct {
	Account string `json:"account,omitempty"`
}

var _ Payment = &Zelle{}

func (_ Zelle) Kind() Kind {
	return Kinds.Zelle
}

func (this Zelle) String() string {
	return "zelle: " + this.Account
}
