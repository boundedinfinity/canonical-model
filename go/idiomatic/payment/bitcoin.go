package payment

type BitCoin struct {
	PublicKey  PublicKey  `json:"public-key,omitempty"`
	PrivateKey PrivateKey `json:"private-key,omitempty"`
}

var _ Payment = &BitCoin{}

func (_ BitCoin) Kind() Kind {
	return Kinds.BitCoin
}

func (this BitCoin) String() string {
	return "BitCoin: " + this.PublicKey.Last4Digits()
}

type PublicKey string

func (this PublicKey) String() string {
	return string(this)
}

func (this PublicKey) Last4Digits() string {
	if len(this) <= 4 {
		return string(this)
	}

	return string(this[len(this)-4:])
}

type PrivateKey string

func (this PrivateKey) String() string {
	return string(this)
}

func (this PrivateKey) Last4Digits() string {
	if len(this) <= 4 {
		return string(this)
	}

	return string(this[len(this)-4:])
}
