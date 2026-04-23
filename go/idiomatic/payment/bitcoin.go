package payment

type BitCoin struct {
}

var _ Payment = &BitCoin{}

func (_ BitCoin) Kind() Kind {
	return Kinds.BitCoin
}
