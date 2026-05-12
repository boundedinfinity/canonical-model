package payment

type WireTransfer struct {
	Account string `json:"account,omitempty"`
}

var _ Payment = &WireTransfer{}

func (_ WireTransfer) Kind() Kind {
	return Kinds.WireTransfer
}

func (this WireTransfer) String() string {
	return "wire transfer: " + this.Account
}
